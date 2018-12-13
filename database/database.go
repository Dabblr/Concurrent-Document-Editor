package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
	"github.com/mxk/go-sqlite/sqlite3"
)

// The Database struct implements the Database interface
type Database struct {
	Path string
}

const createDBFile = "../database/createDb.sql"

// CreateEmptyDb creates an empty database and returns the struct
func CreateEmptyDb(name string) Database {
	os.Remove(name)

	creates, err := ioutil.ReadFile(createDBFile)
	createString := string(creates)
	if err != nil {
		panic(err)
	}

	conn, err := sqlite3.Open(name)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	err = conn.Exec(createString)
	if err != nil {
		panic(err)
	}

	return Database{name}
}

// CreateEmptyFile creates a new file with no contents, gives ownership to userName
// Returns the new file's ID, and the file's latest revision (0) (both -1 in the event  of an error)
func (db *Database) CreateEmptyFile(fileName string, userID string) (int, int, error) {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return -1, -1, err
	}
	err = conn.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return -1, -1, err
	}
	user, err := conn.Query("SELECT id FROM users WHERE username = ?", userID)
	if err != nil {
		return -1, -1, err
	}
	var userNumber int
	user.Scan(&userNumber)
	user.Close()

	statement, err := conn.Prepare("INSERT INTO files (filename, owner) VALUES(?, ?)")
	if err != nil {
		return -1, -1, err
	}
	err = statement.Exec(fileName, userNumber)
	if err != nil {
		return -1, -1, err
	}
	statement.Close()

	// Get file ID
	rows, err := conn.Query("SELECT LAST_INSERT_ROWID();")
	if err != nil {
		return -1, -1, err
	}
	var fileID int
	rows.Scan(&fileID)
	rows.Close()

	if fileID == 0 {
		err = fmt.Errorf("invalid user ID %v", userNumber)
		return -1, -1, err
	}

	// Insert a dummy first change to signify file creation
	err = conn.Exec("INSERT INTO revisions (file, rev_number) VALUES(?, ?)", fileID, 1)
	if err != nil {
		return -1, -1, err
	}

	return fileID, 1, nil
}

// CreateUser creates a user with the given username
// returns the new user's ID
// TODO @nikita: get rid of the int being returned here?
func (db *Database) CreateUser(username string) (int, error) {
	if username == "" {
		return -1, errors.New("empty username")
	}

	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	err = conn.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return -1, err
	}

	statement, err := conn.Prepare("INSERT INTO users (username) VALUES(?)")
	if err != nil {
		return -1, err
	}

	err = statement.Exec(username)
	if err != nil {
		return -1, err
	}

	// Get user ID
	rows, err := conn.Query("SELECT LAST_INSERT_ROWID();")
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return -1, err
	}
	var userID int
	rows.Scan(&userID)

	return userID, nil
}

// GetFileContent return the contents of the latest revision of the given file in the database
// An error is returned if no file with the given id exists.
func (db *Database) GetFileContent(id int) (obj.File, error) {
	var f obj.File

	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return f, err
	}

	rows, err := conn.Query(`
		SELECT filename, data, owner
		FROM files
		WHERE id=?`, id)
	if err != nil {
		return f, err
	}
	var data string
	var user string
	var name string
	var revision int
	rows.Scan(&name, &data, &user)
	rows.Close()

	rows, err = conn.Query(
		`SELECT max(rev_number)
		FROM revisions
		WHERE file = ?`, id)
	rows.Scan(&revision)
	rows.Close()

	f.Content = data
	f.Name = name
	f.User = user
	f.ID = id
	f.RevisionNumber = revision

	return f, nil
}

// GetChangesSinceRevision returns an array of all changes made to the given file after the given revision number.
func (db *Database) GetChangesSinceRevision(id int, revisionNumber int) ([]obj.Change, error) {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	f, err := conn.Query(
		`SELECT id
		FROM files
		WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	var fileCheck int
	f.Scan(&fileCheck)
	if fileCheck == 0 {
		f.Close()
		return nil, errors.New("File not found")
	}
	f.Close()

	var next error
	var changes []obj.Change

	revs, err := conn.Query(
		`SELECT DISTINCT chages.id, changes.file, changes.rev_number, position, character
		FROM revisions
			JOIN changes
		WHERE changes.file = ? AND changes.rev_number > ?`, id, revisionNumber)
	if revs != nil {
		defer revs.Close()
	} else {
		return changes, nil
	}
	if err != nil {
		return nil, err
	}

	for next == nil {
		var fileID int
		var revID int
		var pos int
		var char string
		err = revs.Scan(&fileID, &revID, &pos, &char)
		if err != nil {
			return changes, err
		}
		var c obj.Change
		if char == "" {
			c = obj.NewChange("delete", pos, char)
		} else {
			c = obj.NewChange("insert", pos, char)
		}
		changes = append(changes, c)
		next = revs.Next()
	}

	return changes, nil
}

// InsertChanges inserts an array of changes made to the given file in the database.
func (db *Database) InsertChanges(id int, changes []obj.Change) error {
	if len(changes) == 0 {
		return nil
	}
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return err
	}

	row, err := conn.Query(
		`SELECT MAX(rev_number)
		FROM revisions
		WHERE file = ?`, id)
	defer row.Close()
	if err != nil {
		return err
	}

	var revNum int
	row.Scan(&revNum)
	revNum++

	err = conn.Exec(
		`INSERT INTO revisions (file, rev_number)
			VALUES (?, ?)`, id, revNum)
	if err != nil {
		return err
	}

	statement, err := conn.Prepare(
		`INSERT INTO changes (file, rev_number, position, character)
		VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	for _, c := range changes {
		if c.Type == "delete" {
			c.Value = ""
		}
		err = statement.Exec(id, revNum, c.Position, c.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateFileContent updates the file content for the given file in the database.
func (db *Database) UpdateFileContent(id int, fileContent string) error {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return err
	}

	statement, err := conn.Prepare("UPDATE files SET data = ? WHERE id = ?")
	if err != nil {
		return err
	}
	err = statement.Exec(fileContent, id)
	if err != nil {
		return err
	}

	return nil
}

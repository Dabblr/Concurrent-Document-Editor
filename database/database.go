package database

import (
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

const createDBFile = "createDb.sql"

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
	conn.Exec(createString)

	return Database{name}
}

// CreateEmptyFile creates a new file with no contents, gives ownership to userName
// Returns the new file's ID
func (db *Database) CreateEmptyFile(fileName string, userID int) (int, int, error) {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return -1, -1, err
	}
	err = conn.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return -1, -1, err
	}

	statement, err := conn.Prepare("INSERT INTO files (filename, owner) VALUES(?, ?)")
	defer statement.Close()
	if err != nil {
		return -1, -1, err
	}
	statement.Exec(fileName, userID)

	// Get file ID
	rows, err := conn.Query("SELECT LAST_INSERT_ROWID();")
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return -1, -1, err
	}
	var fileID int
	rows.Scan(&fileID)

	if fileID == 0 {
		err = fmt.Errorf("invalid user ID %v", userID)
		return -1, -1, err
	}

	return fileID, 1, nil
}

// CreateUser creates a user with the given username
// returns the new user's ID
func (db *Database) CreateUser(username string) (int, error) {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return -1, err
	}
	conn.Exec("PRAGMA foreign_keys = ON")

	statement, err := conn.Prepare("INSERT INTO users (username) VALUES(?)")
	if err != nil {
		return -1, err
	}

	statement.Exec(username)

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

	rows, err := conn.Query("SELECT filename, data, owner FROM files WHERE id=?", id)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return f, err
	}
	var data string
	var user string
	var name string
	rows.Scan(&name, &data, &user)

	f.Content = data
	f.Name = name
	f.User = user
	f.ID = id

	return f, nil
}

// GetChangesSinceRevision returns an array of all changes made to the given file after the given revision number.
func (db *Database) GetChangesSinceRevision(id int, revisionNumber int) ([]obj.Change, error) {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	revs, err := conn.Query("SELECT DISTINCT changes.file, rev_number, position, character FROM revisions JOIN changes WHERE changes.file = ? AND number > ?", id, revisionNumber)
	if revs != nil {
		defer revs.Close()
	}
	if err != nil {
		return nil, err
	}

	var next error
	var changes []obj.Change

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
		fmt.Printf("c: %v\n", c)
		changes = append(changes, c)
		next = revs.Next()
	}

	return changes, nil
}

// InsertChanges inserts an array of changes made to the given file in the database.
func (db *Database) InsertChanges(id int, changes []obj.Change) error {
	conn, err := sqlite3.Open(db.Path)
	defer conn.Close()
	if err != nil {
		return err
	}

	row, err := conn.Query("SELECT MAX(rev_number) FROM changes WHERE file = ?", id)
	defer row.Close()
	if err != nil {
		return err
	}

	firstChange := false
	var revNum int
	if !firstChange {
		row.Scan(&revNum)
	}
	revNum++

	conn.Exec("INSERT INTO revisions (file, number) VALUES (?, ?)", id, revNum)

	statement, err := conn.Prepare("INSERT INTO changes (file, rev_number, position, character) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for _, c := range changes {
		statement.Exec(id, revNum, c.Position, c.Value)
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
	statement.Exec(fileContent, id)

	return nil
}

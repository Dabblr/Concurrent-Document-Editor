package main

import (
	"fmt"
	"log"

	"github.com/mxk/go-sqlite/sqlite3"
)

// TODO: tests
// TODO: foreign key constraints seem to be ignored?

type change struct {
	character string
	position  int
}

// Creates a user with the given username
func createUser(username string) {
	// TODO: return user id
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("INSERT INTO users (username) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(username)
}

// Creates an empty file by the given owner with the given file name
func createEmptyFile(fileName string, owner int) {
	// TODO: return file id http://www.sqlite.org/c3ref/last_insert_rowid.html ?
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("INSERT INTO files (filename, owner) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(fileName, owner)
}

// Sets the current file content
func updateFileContents(id int, fileContent string) {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("UPDATE files SET data = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(fileContent, id)

}

// Returns the most up-to-date contents of the given file
func getFileContent(file int) string {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT data FROM files WHERE id=?", file)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data string
	rows.Scan(&data)

	return data
}

// Returns a list of chagnes since the given revision
func getChangesSinceRevision(rev int, file int) []change {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	revs, err := db.Query("SELECT DISTINCT changes.file, rev_number, position, character FROM revisions JOIN changes WHERE changes.file = ? AND number > ?", file, rev)
	defer revs.Close()
	if err != nil {
		log.Fatal(err)
	}

	var next error
	var changes []change

	for next == nil {
		var fileID int
		var revID int
		var pos int
		var char string
		err = revs.Scan(&fileID, &revID, &pos, &char)
		if err != nil {
			log.Fatal(err)
		}
		changes = append(changes, change{character: char, position: pos})
		next = revs.Next()
	}

	return changes
}

// Insert a list of changes into db
// Creates an appropriate revision for the chagnes
func insertChanges(id int, changes []change) {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	row, err := db.Query("SELECT MAX(rev_number) FROM changes WHERE file = ?", id)
	defer row.Close()
	if err != nil {
		log.Panic(err)
	}

	firstChange := false
	var revNum int
	if !firstChange {
		row.Scan(&revNum)
	}
	revNum++

	statement, err := db.Prepare("INSERT INTO changes (file, rev_number, position, character) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range changes {
		statement.Exec(id, revNum, c.position, c.character)
	}

	db.Exec("INSERT INTO revisions (file, number) VALUES (?, ?)", id, revNum)
}

func main() {
	createUser("NikitaIsAwesome")
	createEmptyFile("Test347.txt", 1)
	updateFileContents(1, "q8^)")
	data := getFileContent(1)
	fmt.Println(data)

	var changes []change
	changes = append(changes, change{position: 0, character: "a"})
	changes = append(changes, change{position: 1, character: "b"})
	changes = append(changes, change{position: 2, character: "c"})

	insertChanges(1, changes)
	receivedChanges := getChangesSinceRevision(0, 1)
	fmt.Println(receivedChanges)
}

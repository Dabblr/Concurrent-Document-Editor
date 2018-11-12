package main

import (
	"log"

	"github.com/mxk/go-sqlite/sqlite3"
)

// TODO: list of changes in a revision
// TODO: store the latest revision, latest changes
// TODO: tests lol

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
	// TODO: foreign key constraints seem to be ignored?
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

func updateFileContent(id int, fileContent string) {
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
func getFileContent(file int, dest string) string {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT data FROM files WHERE id=?", file)
	if err != nil {
		log.Fatal(err)
	}
	var data string
	rows.Scan(&data)

	return data
}

type change struct {
	character string
	position  int
}

// Returns a list of chagnes
func getChangesSinceRevision(rev int, file int) []change {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	revs, err := db.Query("SELECT DISTINCT changes.file, rev_number, position, character FROM revisions JOIN changes WHERE changes.file = ? AND number > ?", file, rev)
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

func insertChanges(id int, changes []change) {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("INSERT INTO changes (file, position, character) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range changes {
		statement.Exec(id, c.position, c.character) // TODO: change#
	}
}

func main() {
	// var dest string
	// getChangesSinceRevision(1, 1)

	// updateFileContent(1, "dest")

	// createEmptyFile("PLEEASE", 2)

	// dest = getFileContent(1, dest)
	// fmt.Printf("%s\n", dest)

	var cs []change
	cs = append(cs, change{character: "q", position: 1})
	cs = append(cs, change{character: "w", position: 2})
	// cs[1] = change{character: "w", position: 2}
	insertChanges(1, cs)
}

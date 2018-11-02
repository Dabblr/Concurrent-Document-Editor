package main

import (
	"log"

	"github.com/mxk/go-sqlite/sqlite3"
)

func createUser(username string) { //TODO: fixme :(
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
	q1, _ := db.Query("SELECT * FROM files")
	sqlite3.Print(q1)
}

func readFile(file int64, dest []byte) int {
	// TODO: use row instead of fileId
	db, _ := sqlite3.Open("updates.db")
	defer db.Close()
	data, _ := db.BlobIO("updates.db", "files", "data", file, false)
	n, _ := data.Read(dest) // NOTE: I think this is written into dest
	return n
}

func main() {
	// db, _ := sqlite3.Open("updates.db")
	// q1, _ := db.Query("SELECT * FROM users")
	// sqlite3.Print(q1)
	createEmptyFile("testF.txt", 27)
}

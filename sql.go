package main

import (
	"fmt"
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

func readFile(file int64, dest []byte) []byte {
	db, err := sqlite3.Open("updates.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT data FROM files WHERE id=?", file)
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	rows.Scan(&data)

	fmt.Printf("%c\n", data)
	return data
}

func main() {
	var dest []byte
	readFile(1, dest)
	fmt.Printf("%b\n", dest)
}

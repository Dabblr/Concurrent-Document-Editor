package main

import (
	"errors"
	"log"

	obj "github.com/jcgallegdup/Concurrent-Document-Editor/objects"
)

// Keeps track of how many files have been created to generate new ids.
var mockFileCounter int
var mockFileContent string
var mockChanges []obj.Change

// Mocks storing a new file in the database and returning an id for it.
func mockCreateEmptyFile(fileName string, userName string) int {
	mockFileCounter++
	return mockFileCounter
}

// Mocks returning the latest revision file content for the given file id.
// An error is returned if no file with the given id exists.
func mockGetFileContent(id int) (obj.File, error) {
	if id <= 0 || id > mockFileCounter {
		// Invalid id.
		return obj.NewFile("", 0, "", 0, ""), errors.New("invalid file id")
	}
	return obj.NewFile("", id, "fileName", 1, mockFileContent), nil
}

// Mocks returning an array of all changes after the given revision number for the given file id.
// Currently it just returns all changes recorded in the last update.
func mockGetChangesSinceRevision(id int, revisionNumber int) []obj.Change {
	changes := mockChanges
	mockChanges = []obj.Change{}
	return changes
}

// Mocks inserting a new change to a file in the database.
func mockInsertChange(id int, change obj.Change) {
	mockChanges = append(mockChanges, change)
	return
}

// Mocks updating the file content for the given file id in the database.
func mockUpdateFileContent(id int, fileContent string) {
	mockFileContent = fileContent
	log.Println("New file content:", fileContent)
	return
}

package database

import (
	"errors"
	"log"

	obj "github.com/jcgallegdup/Concurrent-Document-Editor/objects"
)

// MockDB mocks the database functions for unit testing.
type MockDB struct {
	// Keeps track of how many files have been created to generate new ids.
	FileCounter int
	// Keeps track of the content of the file.
	FileContent string
	// Keeps track of all changes recorded in the last update.
	Changes []obj.Change
}

// CreateEmptyFile increments the FileCounter and returns it to mock creating a file.
func (m *MockDB) CreateEmptyFile(fileName string, userName string) int {
	m.FileCounter++
	return m.FileCounter
}

// GetFileContent returns a File object containing FileContent.
// An error is returned if id <= 0 or id > FileCounter.
func (m *MockDB) GetFileContent(id int) (obj.File, error) {
	if id <= 0 || id > m.FileCounter {
		// Invalid id.
		return obj.NewFile("", 0, "", 0, ""), errors.New("invalid file id")
	}
	return obj.NewFile("", id, "fileName", 1, m.FileContent), nil
}

// GetChangesSinceRevision returns all changes recorded in the last update.
func (m *MockDB) GetChangesSinceRevision(id int, revisionNumber int) []obj.Change {
	changes := m.Changes
	m.Changes = []obj.Change{}
	return changes
}

// InsertChanges mocks inserting an array of changes to a file in the database.
func (m *MockDB) InsertChanges(id int, changes []obj.Change) {
	m.Changes = append(m.Changes, changes...)
}

// UpdateFileContent mocks updating the file content for the given file in the database.
func (m *MockDB) UpdateFileContent(id int, fileContent string) {
	m.FileContent = fileContent
	log.Println("New file content:", fileContent)
}

package database

import (
	"errors"
	"log"

	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
)

// The "limit" on the number of files that can be created in the mock DB.
const maxFiles = 100

// MockDB mocks the database functions for unit testing.
type MockDB struct {
	// Keeps track of how many files have been created to generate new ids.
	FileCounter int
	// Keeps track of the content of the file.
	FileContent string
	// Keeps track of all changes recorded in the last update.
	Changes []obj.Change
}

// CreateUser returns an error if the given username string is empty, otherwise returns nil.
// TODO: Once Nikita removes the int from his method, remove it here too.
func (m *MockDB) CreateUser(username string) (int, error) {
	if username == "" {
		return -1, errors.New("cannot create a user with an empty username")
	}
	return 1, nil
}

// CreateEmptyFile increments the FileCounter and returns it to mock creating a file.
// Returns an error if we have already created the maximum number of files (arbitrarily imposed limit to test error response)
func (m *MockDB) CreateEmptyFile(fileName string, userName string) (int, int, error) {
	if m.FileCounter >= maxFiles {
		return -1, -1, errors.New("not enough space to create a new file")
	}
	m.FileCounter++
	return m.FileCounter, 1, nil
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
// An error is returned if id <= 0 or id > FileCounter.
func (m *MockDB) GetChangesSinceRevision(id int, revisionNumber int) ([]obj.Change, error) {
	if id <= 0 || id > m.FileCounter {
		// Invalid id.
		return []obj.Change{}, errors.New("invalid file id")
	}
	changes := m.Changes
	m.Changes = []obj.Change{}
	return changes, nil
}

// InsertChanges mocks inserting an array of changes to a file in the database.
// An error is returned if id <= 0 or id > FileCounter.
func (m *MockDB) InsertChanges(id int, changes []obj.Change) error {
	if id <= 0 || id > m.FileCounter {
		// Invalid id.
		return errors.New("invalid file id")
	}
	m.Changes = append(m.Changes, changes...)
	return nil
}

// UpdateFileContent mocks updating the file content for the given file in the database.
// An error is returned if id <= 0 or id > FileCounter.
func (m *MockDB) UpdateFileContent(id int, fileContent string) error {
	if id <= 0 || id > m.FileCounter {
		// Invalid id.
		return errors.New("invalid file id")
	}
	m.FileContent = fileContent
	log.Println("New file content:", fileContent)
	return nil
}

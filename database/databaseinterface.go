package database

import (
	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
)

// Database contains all the functions that interact with the database.
type DatabaseInterface interface {
	// Creates a new file, stores it in the database, and returns the id for it.
	CreateEmptyFile(fileName string, userID string) (int, int, error)

	// Returns the latest revision file content for the given file id.
	// An error is returned if no file with the given id exists.
	GetFileContent(id int) (obj.File, error)

	// Returns an array of all changes made to the given file after the given revision number.
	GetChangesSinceRevision(id int, revisionNumber int) ([]obj.Change, error)

	// Inserts an array of changes made to the given file in the database.
	InsertChanges(id int, changes []obj.Change) error

	// Updates the file content for the given file in the database.
	UpdateFileContent(id int, fileContent string) error
}

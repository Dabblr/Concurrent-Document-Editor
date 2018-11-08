package database

import (
	obj "github.com/jcgallegdup/Concurrent-Document-Editor/objects"
)

// Database contains all the functions that interact with the database.
type Database interface {
	// Creates a new file, stores it in the database, and returns the id for it.
	CreateEmptyFile(fileName string, userName string) int

	// Returns the latest revision file content for the given file id.
	// An error is returned if no file with the given id exists.
	GetFileContent(id int) (obj.File, error)

	// Returns an array of all changes made to the given file after the given revision number.
	GetChangesSinceRevision(id int, revisionNumber int) []obj.Change

	// Inserts an array of changes made to the given file in the database.
	InsertChanges(id int, changes []obj.Change)

	// Updates the file content for the given file in the database.
	UpdateFileContent(id int, fileContent string)
}

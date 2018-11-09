package objects

import "fmt"

// File contains information about each file.
type File struct {
	// The username of the user updating the file.
	User string `json:"user,omitempty"`
	// The ID of the file (generated on server).
	ID int `json:"id"`
	// The name of the file.
	Name string `json:"name"`
	// The revision number of the file.
	RevisionNumber int `json:"revision_number"`
	// The file content.
	Content string `json:"content,omitempty"`
}

// NewFile is a constructor for the File type.
func NewFile(user string, id int, name string, revision int, content string) File {
	return File{user, id, name, revision, content}
}

// Equals defines what makes two File objects equal.
func (file *File) Equals(file2 File) bool {
	if file.Content == file2.Content && file.ID == file2.ID && file.Name == file2.Name && file.RevisionNumber == file2.RevisionNumber && file.User == file2.User {
		return true
	}
	return false
}

// String determines the default string format for the File type.
func (file File) String() string {
	return fmt.Sprintf("File: user=%s id=%d name=%s revisionNumber=%d content=%s", file.User, file.ID, file.Name, file.RevisionNumber, file.Content)
}

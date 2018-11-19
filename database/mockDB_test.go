package database

import (
	"testing"

	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
)

// Tests that CreateEmptyFile returns the value of FileCounter incremented by 1.
func TestCreateEmptyFileReturnsIncrementedCounter(t *testing.T) {
	var m MockDB
	expID := m.FileCounter + 1
	id, _ := m.CreateEmptyFile("fileName", "userName")

	if id != expID {
		t.Errorf("Expected id %d from CreateEmptyFile but got %d.", expID, id)
	}
}

// Tests that CreateEmptyFile returns a revision number = 1.
func TestCreateEmptyFileReturnsRevisionNumberEqualOne(t *testing.T) {
	var m MockDB
	_, revisionNumber := m.CreateEmptyFile("fileName", "userName")

	if revisionNumber != 1 {
		t.Errorf("Expected revisionNumber 1 from CreateEmptyFile but got %d.", revisionNumber)
	}
}

// Tests that GetFileContent returns a File object with the correct ID and FileContent when the input ID is valid.
func TestGetFileContentReturnsFileWhenValidId(t *testing.T) {
	fileContent := "fileContent"
	fileID := 5
	m := MockDB{fileID, fileContent, []obj.Change{}}
	expFile := obj.NewFile("", fileID, "fileName", 1, fileContent)

	file, err := m.GetFileContent(fileID)
	if expFile.Equals(file) == false || err != nil {
		t.Errorf("Expected %v from GetFileContent but got %v.", expFile, file)
	}
}

// Tests that GetFileContent returns an error when the ID is negative.
func TestGetFileContentReturnsErrorWhenNegativeId(t *testing.T) {
	fileID := -1
	m := MockDB{5, "", []obj.Change{}}

	file, err := m.GetFileContent(fileID)
	if err == nil {
		t.Errorf("Expected GetFileContent to produce an error, but it was nil.\n Returned File: %v", file)
	}
}

// Tests that GetFileContent returns an error when the ID is zero.
func TestGetFileContentReturnsErrorWhenZeroId(t *testing.T) {
	fileID := 0
	m := MockDB{5, "", []obj.Change{}}

	file, err := m.GetFileContent(fileID)
	if err == nil {
		t.Errorf("Expected GetFileContent to produce an error, but it was nil.\n Returned File: %v", file)
	}
}

// Tests that GetFileContent returns an error when the file ID does not exist on the server.
func TestGetFileContentReturnsErrorWhenIdNotFound(t *testing.T) {
	fileID := 5
	m := MockDB{fileID - 1, "", []obj.Change{}}

	file, err := m.GetFileContent(fileID)
	if err == nil {
		t.Errorf("Expected GetFileContent to produce an error, but it was nil.\n Returned File: %v", file)
	}
}

// Tests that GetChangesSinceRevision returns the Change array Changes.
func TestGetChangesSinceRevisionReturnsChangeArray(t *testing.T) {
	fileID := 1
	revisionNumber := 1
	expChanges := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("delete", 0, "a")}
	m := MockDB{fileID, "", expChanges}

	changes := m.GetChangesSinceRevision(fileID, revisionNumber)
	for i, change := range changes {
		if change.Equals(expChanges[i]) == false {
			t.Errorf("Expected %v from GetChangesSinceRevision, but got %v.", expChanges, changes)
			break
		}
	}
}

// Tests that GetChangesSinceRevision returns an empty Change array when Changes is empty.
func TestGetChangesSinceRevisionReturnsEmptyArrayIfNoChanges(t *testing.T) {
	fileID := 1
	revisionNumber := 1
	m := MockDB{fileID, "", []obj.Change{}}

	changes := m.GetChangesSinceRevision(fileID, revisionNumber)
	if len(changes) != 0 {
		t.Errorf("Expected an empty Change array from GetChangesSinceRevision, but got %v.", changes)
	}
}

// Tests that InsertChanges updates the Changes field to include the array of changes.
func TestInsertChangesUpdatesChanges(t *testing.T) {
	fileID := 1
	fileContent := ""
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a")}
	newChanges := []obj.Change{obj.NewChange("delete", 0, "a"), obj.NewChange("insert", 0, "x")}
	m := MockDB{fileID, fileContent, prevChanges}
	expChanges := append(prevChanges, newChanges...)

	m.InsertChanges(fileID, newChanges)
	for i, change := range m.Changes {
		if change.Equals(expChanges[i]) == false {
			t.Errorf("Expected InsertChanges to update Changes to %v, but it updated to %v.", expChanges, m.Changes)
			break
		}
	}
}

// Tests that InsertChanges does not update the Changes field when the array of changes is empty.
func TestInsertChangesDoesNotUpdateWhenChangesEmpty(t *testing.T) {
	fileID := 1
	fileContent := ""
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a")}
	newChanges := []obj.Change{}
	m := MockDB{fileID, fileContent, prevChanges}
	expChanges := prevChanges

	m.InsertChanges(fileID, newChanges)
	for i, change := range m.Changes {
		if change.Equals(expChanges[i]) == false {
			t.Errorf("Expected InsertChanges to remain %v, but it updated to %v.", expChanges, m.Changes)
			break
		}
	}
}

// Tests that UpdateFileContent overwrites the FileContent field with the new value.
func TestUpdateFileContentModifiesFileContent(t *testing.T) {
	fileID := 1
	fileContent := "oldContent"
	newContent := "newContent"
	m := MockDB{fileID, fileContent, []obj.Change{}}

	m.UpdateFileContent(fileID, newContent)
	if m.FileContent != newContent {
		t.Errorf("Expected UpdateFileContent to update FileContent to %v, but it got updated to %v.", newContent, m.FileContent)
	}
}

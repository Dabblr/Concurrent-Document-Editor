package database

import (
	"fmt"
	"testing"

	obj "github.com/jcgallegdup/Concurrent-Document-Editor/objects"
)

const path = "../updates.db"
const invalidID = 9001
const ID = 1

var expectedChanges = []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 1, "b")}

func TestCreateUserReturnsIncrementedCounterReal(t *testing.T) {
	CreateEmptyDb(path)
	db := Database{Path: path}
	userID, _ := db.CreateUser("TESTUSER")

	if userID != ID {
		t.Errorf("Expected userID to be 1, instead it was %d", userID)
	}
}

func TestCreateEmptyFileReturnsIncrementedCounterReal(t *testing.T) {
	db := Database{Path: path}
	fileID, _ := db.CreateEmptyFile("TESTFILE.TEST", ID)

	if fileID != ID {
		t.Errorf("Expected fileID to be 1, instead it was %d", fileID)
	}
}

func TestCreateEmptyFileReturnsErrorWhenBadUserReal(t *testing.T) {
	db := Database{Path: path}
	_, err := db.CreateEmptyFile("TESTFILE.TEST", invalidID)

	if err == nil {
		t.Errorf("Expected CreateEmptyFile to produce an error, but it was nil.")
	}
}

func TestUpdateFileContentModifiesFileContentReal(t *testing.T) {
	db := Database{Path: path}

	err := db.UpdateFileContent(ID, "Updated file content")

	if err != nil {
		t.Errorf("UpdateFileContent threw an error %v", err)
	}
}

func TestGetFileContentReturnsMatchingContentReal(t *testing.T) {
	db := Database{Path: path}
	f, err := db.GetFileContent(ID)

	if err != nil {
		t.Errorf("GetFileContent threw and error %v", err)
	}
	if f.Content != "Updated file content" {
		t.Errorf("Incorrect file content, received content: %v", f.Content)
	}
}

func TestGetFileContentReturnsErrorWithBadFileIDReal(t *testing.T) {
	db := Database{Path: path}
	_, err := db.GetFileContent(invalidID)

	if err == nil {
		t.Error("Expected GetFileContent to produce an error, but it was nil.")
	}
}

func TestInsertChangesUpdatesChangesReal(t *testing.T) {
	db := Database{Path: path}

	err := db.InsertChanges(ID, expectedChanges)

	if err != nil {
		t.Errorf("InsertChanges threw an error %s", err)
	}
}

func TestInsertChangesDoesNotUpdateWhenChangesEmptyReal(t *testing.T) {
	db := Database{Path: path}
	var emptyChanges []obj.Change

	err := db.InsertChanges(ID, emptyChanges)

	if err != nil {
		t.Errorf("InsertChanges threw an error on empty array %v", err)
	}
}

func TestGetChangesSinceRevisionReturnsEmptyArrayIfNoChangesReal(t *testing.T) {
	db := Database{Path: path}
	changes, _ := db.GetChangesSinceRevision(ID, 1)

	for i, c := range changes {
		fmt.Print("F")
		fmt.Printf("%v ", i)
		fmt.Printf("%v\n", c)
	}

	if len(changes) != 0 {
		t.Errorf("GetChangesSinceRevision was supposed to return an empty array. Instead we got %+v", changes)
	}
}

func TestGetChangesSinceRevisionReturnsErrorWithBadFileIDReal(t *testing.T) {
	db := Database{Path: path}
	_, err := db.GetChangesSinceRevision(invalidID, 0)

	if err == nil {
		t.Error("Expected GetChangesSinceRevision to produce an error, but it was nil.")
	}
}

func TestGetChangesSinceRevisionReturnsChangeArrayReal(t *testing.T) {
	db := Database{Path: path}
	changes, _ := db.GetChangesSinceRevision(ID, 0)

	if len(changes) != len(expectedChanges) {
		t.Errorf("GetChangesSinceRevision returned an array with the wrong length: %v, expecetd: %v", changes, expectedChanges)
	} else {
		for i, change := range changes {
			if change.Equals(expectedChanges[i]) == false {
				t.Errorf("GetChangesSinceRevision returned a weird change: %v, expecetd: %v", changes, expectedChanges)
			}
		}
	}
}

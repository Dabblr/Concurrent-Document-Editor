package main

import (
	"testing"

	db "github.com/Dabblr/Concurrent-Document-Editor/database"
	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
)

// Tests that the file content is correctly updated when a change inserts at the start.
func TestApplyChangeInsAtStart(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("insert", 0, "x")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "xabcd"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content is correctly updated when a change inserts somewhere in the middle.
func TestApplyChangeInsInMiddle(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("insert", 2, "x")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "abxcd"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content is correctly updated when a change inserts at the end.
func TestApplyChangeInsAtEnd(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("insert", 4, "x")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "abcdx"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content remains unchanged and an error is returned when the Change position is negative.
func TestApplyChangeInsNegativeReturnErr(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("insert", -1, "x")
	newContent, err := ApplyChange(change, prevContent)

	if newContent != prevContent || err == nil {
		t.Errorf("Expected an error from ApplyChange but got nil. Content changed from %v to %v.", prevContent, newContent)
	}
}

// Tests that the file content remains unchanged and an error is returned when the Change position is greater than the length of the string.
func TestApplyChangeInsGreaterThanLengthReturnsErr(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("insert", 5, "x")
	newContent, err := ApplyChange(change, prevContent)

	if newContent != prevContent || err == nil {
		t.Errorf("Expected an error from ApplyChange but got nil. Content changed from %v to %v.", prevContent, newContent)
	}
}

// Tests that the file content is correctly updated when a change deletes at the start.
func TestApplyChangeDelAtStart(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", 0, "a")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "bcd"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content is correctly updated when a change deletes somewhere in the middle.
func TestApplyChangeDelInMiddle(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", 2, "c")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "abd"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content is correctly updated when a change deletes at the end.
func TestApplyChangeDelAtEnd(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", 3, "d")
	newContent, err := ApplyChange(change, prevContent)
	expContent := "abc"

	if newContent != expContent || err != nil {
		t.Errorf("Expected %v from ApplyChange but got %v.", expContent, newContent)
	}
}

// Tests that the file content remains unchanged and an error is returned when the Change position is negative.
func TestApplyChangeDelNegativeReturnErr(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", -1, "x")
	newContent, err := ApplyChange(change, prevContent)

	if newContent != prevContent || err == nil {
		t.Errorf("Expected an error from ApplyChange but got nil. Content changed from %v to %v.", prevContent, newContent)
	}
}

// Tests that the file content remains unchanged and an error is returned when the Change position is greater than the length of the string for deletion.
func TestApplyChangeDelGreaterThanLengthReturnsErr(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", 5, "x")
	newContent, err := ApplyChange(change, prevContent)

	if newContent != prevContent || err == nil {
		t.Errorf("Expected an error from ApplyChange but got nil. Content changed from %v to %v.", prevContent, newContent)
	}
}

// Tests that the file content remains unchanged and an error is returned when the Change position is equal to the length of the string for deletion.
func TestApplyChangeDelEqualToLengthReturnsErr(t *testing.T) {
	prevContent := "abcd"
	change := obj.NewChange("delete", 4, "x")
	newContent, err := ApplyChange(change, prevContent)

	if newContent != prevContent || err == nil {
		t.Errorf("Expected an error from ApplyChange but got nil. Content changed from %v to %v.", prevContent, newContent)
	}
}

// Note: Need to find a way to efficiently write tests for testing TransformChange on many cases of arrays.
// Tests that TransformChange correctly transforms a dependent insert on top of an array of previous inserts.
func TestTransformChangeDependentInsOnIns(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 1, "b")}
	newChange := obj.NewChange("insert", 0, "c")
	expChange := obj.NewChange("insert", 2, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an indepedent insertion on top of an array of previous insertions.
func TestTransformChangeIndependentInsOnIns(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 1, "a"), obj.NewChange("insert", 2, "b")}
	newChange := obj.NewChange("insert", 0, "c")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange correctly transforms a dependent deletions on top of an array of previous deletions.
func TestTransformChangeDependentDelOnDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("delete", 0, "a"), obj.NewChange("delete", 0, "b")}
	newChange := obj.NewChange("delete", 2, "c")
	expChange := obj.NewChange("delete", 0, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an independent deletion on top of an array of previous deletions.
func TestTransformChangeIndependentDelOnDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("delete", 1, "a"), obj.NewChange("delete", 1, "b")}
	newChange := obj.NewChange("delete", 0, "c")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange correctly transforms a dependent insertion on top of an array of previous deletions.
func TestTransformChangeDependentInsOnDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("delete", 0, "a"), obj.NewChange("delete", 0, "b")}
	newChange := obj.NewChange("insert", 1, "c")
	expChange := obj.NewChange("insert", 0, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an independent insertion on top of an array of previous deletions.
func TestTransformChangeIndependentInsOnDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("delete", 1, "a"), obj.NewChange("delete", 1, "b")}
	newChange := obj.NewChange("insert", 1, "c")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange correctly transforms a dependent deletion on top of an array of previous insertions.
func TestTransformChangeDependentDelOnIns(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 1, "b")}
	newChange := obj.NewChange("delete", 1, "c")
	expChange := obj.NewChange("delete", 3, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an independent deletion on top of an array of previous insertions.
func TestTransformChangeInependentDelOnIns(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 1, "a"), obj.NewChange("insert", 2, "b")}
	newChange := obj.NewChange("delete", 0, "c")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange correctly transforms a dependent insert on top of an array of previous inserts and deletes.
func TestTransformChangeDependentInsOnInsAndDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("delete", 1, "b")}
	newChange := obj.NewChange("insert", 0, "c")
	expChange := obj.NewChange("insert", 1, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an independent insert on top of an array of previous inserts and deletes.
func TestTransformChangeIndependentInsOnInsAndDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 1, "a"), obj.NewChange("delete", 1, "a")}
	newChange := obj.NewChange("insert", 0, "b")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange correctly transforms a dependent deletion on top of an array of previous insertions and deletions.
func TestTransformChangeDependentDelOnInsAndDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("delete", 3, "b")}
	newChange := obj.NewChange("delete", 1, "c")
	expChange := obj.NewChange("delete", 2, "c")
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that TransformChange does not transform an independent deletion on top of an array of previous insertions and deletions.
func TestTransformChangeIndependentDelOnInsAndDel(t *testing.T) {
	prevChanges := []obj.Change{obj.NewChange("insert", 1, "a"), obj.NewChange("delete", 2, "b")}
	newChange := obj.NewChange("delete", 0, "c")
	expChange := newChange
	transformedChange, err := TransformChange(newChange, prevChanges)

	if transformedChange.Equals(expChange) == false || err != nil {
		t.Errorf("Expected %v from TransformChange but got %v.", expChange, transformedChange)
	}
}

// Tests that ApplyUpdate does not return an error when the update is successfully applied.
func TestApplyUpdateReturnsNilWhenSuccess(t *testing.T) {
	changes := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 1, "b")}
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "x"), obj.NewChange("insert", 1, "y")}
	revision := obj.NewRevision("user1", 1, 1, changes)
	file := obj.NewFile("user1", 1, "fileName", 1, "xy")
	mockDB := db.MockDB{FileCounter: 1, FileContent: "xy", Changes: prevChanges}

	err := ApplyUpdate(revision, file, &mockDB)
	if err != nil {
		t.Errorf("Expected ApplyUpdate to be successful, but it returned an error: %v", err)
	}
}

// Tests that ApplyUpdate returns an error when a Change in the update is invalid.
func TestApplyUpdateReturnsErrWhenInvalidChange(t *testing.T) {
	changes := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 1, "bc")}
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "x"), obj.NewChange("insert", 1, "y")}
	revision := obj.NewRevision("user1", 1, 1, changes)
	file := obj.NewFile("user1", 1, "fileName", 1, "xy")
	mockDB := db.MockDB{FileCounter: 1, FileContent: "xy", Changes: prevChanges}

	err := ApplyUpdate(revision, file, &mockDB)
	if err == nil {
		t.Errorf("Expected ApplyUpdate to return an error but it returned nil.")
	}
}

// Tests that ApplyUpdate returns an error when a Change in the update has an out of range position.
func TestApplyUpdateReturnsErrWhenChangeIndexOutOfRange(t *testing.T) {
	changes := []obj.Change{obj.NewChange("insert", 0, "a"), obj.NewChange("insert", 5, "b")}
	prevChanges := []obj.Change{obj.NewChange("insert", 0, "x"), obj.NewChange("insert", 1, "y")}
	revision := obj.NewRevision("user1", 1, 1, changes)
	file := obj.NewFile("user1", 1, "fileName", 1, "xy")
	mockDB := db.MockDB{FileCounter: 1, FileContent: "xy", Changes: prevChanges}

	err := ApplyUpdate(revision, file, &mockDB)
	if err == nil {
		t.Errorf("Expected ApplyUpdate to return an error but it returned nil.")
	}
}

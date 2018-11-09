package objects

import (
	"testing"
)

// Tests that two File objects are equal if they have the same values for all fields.
func TestFileEquivIfEqualFields(t *testing.T) {
	file1 := NewFile("userName", 1, "fileName", 1, "fileContent")
	file2 := NewFile("userName", 1, "fileName", 1, "fileContent")

	if file1.Equals(file2) == false {
		t.Errorf("File %v and File %v should be considered equal.", file1, file2)
	}
}

// Tests that the same instance of a File object is equal.
func TestFileEquivIfSameInstance(t *testing.T) {
	file := NewFile("userName", 1, "fileName", 1, "fileContent")
	if file.Equals(file) == false {
		t.Errorf("File %v should be considered equal to itself.", file)
	}
}

// Tests that two File objects are not equal if they have a differing field.
func TestFileNotEquivIfDiffField(t *testing.T) {
	file1 := NewFile("userName", 1, "fileName", 1, "fileContent")
	file2 := NewFile("userName", 2, "fileName", 1, "fileContent")

	if file1.Equals(file2) == true {
		t.Errorf("File %v and File %v should not be considered equal.", file1, file2)
	}
}

// Tests that two File objects are not equal if they have multiple differing fields.
func TestFileNotEquivIfDiffFields(t *testing.T) {
	file1 := NewFile("userName", 1, "fileName", 1, "fileContent")
	file2 := NewFile("userName2", 2, "fileName", 1, "fileContent")

	if file1.Equals(file2) == true {
		t.Errorf("File %v and File %v should NOT be considered equal.", file1, file2)
	}
}

// Tests that the File Constructor creates a File object with the correct fields.
func TestFileConstructor(t *testing.T) {
	file := NewFile("userName", 1, "fileName", 1, "fileContent")
	expFile := File{"userName", 1, "fileName", 1, "fileContent"}

	if file.Equals(expFile) == false {
		t.Errorf("Expected %v from File constructor but got %v.", expFile, file)
	}
}

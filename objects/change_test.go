package objects

import (
	"testing"

	ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"
)

// Tests that two Change objects are equal if they have the same fields.
func TestChangeEquivIfSameFields(t *testing.T) {
	change1 := NewChange("insert", 0, "a")
	change2 := NewChange("insert", 0, "a")

	if change1.Equals(change2) == false {
		t.Errorf("Change %v and Change %v should be considered equal.", change1, change2)
	}
}

// Tests that an instance of a Change object is equal to itself.
func TestChangeEquivIfSameInstance(t *testing.T) {
	change := NewChange("insert", 0, "a")

	if change.Equals(change) == false {
		t.Errorf("Change %v should be considered equal to itself.", change)
	}
}

// Tests that two Change objects are not equal if they have a differing field.
func TestChangeNotEquivIfDiffField(t *testing.T) {
	change1 := NewChange("insert", 0, "a")
	change2 := NewChange("insert", 1, "a")

	if change1.Equals(change2) == true {
		t.Errorf("Change %v and Change %v should NOT be considered equal.", change1, change2)
	}
}

// Tests that two Change objects are not equal if they have multiple differing fields.
func TestChangeNotEquivIfDiffFields(t *testing.T) {
	change1 := NewChange("insert", 0, "a")
	change2 := NewChange("insert", 1, "z")

	if change1.Equals(change2) == true {
		t.Errorf("Change %v and Change %v should NOT be considered equal.", change1, change2)
	}
}

// Tests that the Change constructor creates a Change object with the correct fields.
func TestChangeConstructor(t *testing.T) {
	change := NewChange("delete", 0, "b")
	expChange := Change{"delete", 0, "b"}

	if change.Equals(expChange) == false {
		t.Errorf("Expected %v from Change constructor but got %v.", expChange, change)
	}
}

// Tests that the IsValid function returns true when the Change is a valid insert.
func TestChangeIsValidForInsert(t *testing.T) {
	validChange := NewChange("insert", 0, "a")

	if validChange.IsValid() == false {
		t.Errorf("Change %v should be valid.", validChange)
	}
}

// Tests that the IsValid function returns true when the Change is a valid delete.
func TestChangeIsValidForDelete(t *testing.T) {
	validChange := NewChange("delete", 0, "a")

	if validChange.IsValid() == false {
		t.Errorf("Change %v should be valid.", validChange)
	}
}

// Tests that the IsValid function returns false when the Change contains a Type that is not "insert" or "delete".
func TestChangeIsNotValidIfInvalidType(t *testing.T) {
	invalidChange := NewChange("append", 0, "a")

	if invalidChange.IsValid() == true {
		t.Errorf("Change %v should be invalid.", invalidChange)
	}
}

// Tests that the IsValid function returns false when the Change contains a Value that is longer than one character.
func TestChangeIsNotValidIfInvalidValue(t *testing.T) {
	invalidChange := NewChange("insert", 0, "ab")

	if invalidChange.IsValid() == true {
		t.Errorf("Change %v should be invalid.", invalidChange)
	}
}

// Tests that the ChangeToIns function correctly converts the Change to an Insertion.
func TestChangeToIns(t *testing.T) {
	change := NewChange("insert", 0, "a")
	ins := change.ChangeToIns()
	expIns := ops.NewInsertion(0, 'a')

	if ins.Equals(expIns) == false {
		t.Errorf("Expected ChangeToIns to convert %v to %v, but got %v", change, expIns, ins)
	}
}

// Tests that the ChangeToDel function correctly converts the Change to a Deletion.
func TestChangeToDel(t *testing.T) {
	change := NewChange("delete", 0, "a")
	del := change.ChangeToDel()
	expDel := ops.NewDeletion(0)

	if del.Equals(expDel) == false {
		t.Errorf("Expected ChangeToDel to convert %v to %v, but got %v.", change, expDel, del)
	}
}

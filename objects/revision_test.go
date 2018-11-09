package objects

import (
	"testing"
)

// Tests that two Revision objects are equal if they have the same fields.
func TestRevisionEquivIfSameFields(t *testing.T) {
	changes := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	rev1 := NewRevision("user1", 1, 1, changes)
	rev2 := NewRevision("user1", 1, 1, changes)

	if rev1.Equals(rev2) == false {
		t.Errorf("Revision %v and Revision %v should be considered equal.", rev1, rev2)
	}
}

// Tests that an instance of a Revision object is equal to itself.
func TestRevisionEquivIfSameInstance(t *testing.T) {
	changes := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	rev := NewRevision("user1", 5, 2, changes)

	if rev.Equals(rev) == false {
		t.Errorf("Revision %v should be considered equal to itself.", rev)
	}
}

// Tests that two Revision objects are not equal if they have a differing field.
func TestRevisionNotEquivIfDiffField(t *testing.T) {
	changes := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	rev1 := NewRevision("user1", 1, 1, changes)
	rev2 := NewRevision("user1", 2, 1, changes)

	if rev1.Equals(rev2) == true {
		t.Errorf("Revision %v and Revision %v should NOT be considered equal.", rev1, rev2)
	}
}

// Tests that two Revision objects are not equal if they have multiple differing fields.
func TestRevisionNotEquivIfDiffFields(t *testing.T) {
	changes := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	rev1 := NewRevision("user1", 1, 1, changes)
	rev2 := NewRevision("user2", 2, 1, changes)

	if rev1.Equals(rev2) == true {
		t.Errorf("Revision %v and Revision %v should NOT be considered equal.", rev1, rev2)
	}
}

// Tests that two Revision objects are not equal if their Change arrays are different sizes.
func TestRevisionNotEquivIfDiffLengthChanges(t *testing.T) {
	changes1 := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	changes2 := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b"), NewChange("delete", 0, "a")}
	rev1 := NewRevision("user1", 1, 1, changes1)
	rev2 := NewRevision("user1", 2, 1, changes2)

	if rev1.Equals(rev2) == true {
		t.Errorf("Revision %v and Revision %v should NOT be considered equal.", rev1, rev2)
	}
}

// Tests that two Revision objects are not equal if their Change arrays are the same length but contain different Changes.
func TestRevisionNotEquivIfDiffChanges(t *testing.T) {
	changes1 := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	changes2 := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "c")}
	rev1 := NewRevision("user1", 1, 1, changes1)
	rev2 := NewRevision("user1", 2, 1, changes2)

	if rev1.Equals(rev2) == true {
		t.Errorf("Revision %v and Revision %v should NOT be considered equal.", rev1, rev2)
	}
}

// Tests that the Revision constructor returns a Revision object with the correct fields.
func TestRevisionConstructor(t *testing.T) {
	changes := []Change{NewChange("insert", 0, "a"), NewChange("insert", 1, "b")}
	rev := NewRevision("user1", 1, 2, changes)
	expRev := Revision{"user1", 1, 2, changes}

	if rev.Equals(expRev) == false {
		t.Errorf("Expected %v from Revision constructor but got %v.", expRev, rev)
	}
}

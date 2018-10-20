package operations

import "testing"

func TestDeletionEquivIfEqualFields(t *testing.T) {
	del1 := NewDeletion(0)
	del2 := NewDeletion(0)

	if del1.Equals(del2) == false {
		t.Errorf("Deletions %v and %v should be considered equal.", del1, del2)
	}
}

func TestDeletionEquivIfSameInstance(t *testing.T) {
	del1 := NewDeletion(0)

	if del1.Equals(del1) == false {
		t.Errorf("Deletion %v should be considered equal to itself.", del1)
	}
}

func TestDeletionNotEquivIfDiffField(t *testing.T) {
	del1 := NewDeletion(0)
	del2 := NewDeletion(1)

	if del1.Equals(del2) == true {
		t.Errorf("Deletions %v and %v should NOT be considered equal.", del1, del2)
	}
}

func TestDeletionNotEquivIfDiffFields(t *testing.T) {
	del1 := NewDeletion(0)
	del2 := NewDeletion(1)

	if del1.Equals(del2) == true {
		t.Errorf("Deletions %v and %v should NOT be considered equal.", del1, del2)
	}
}

// NOTE: this test will be worth more when we introduce more logic into the Deletion constructor
func TestDeletionConstructor(t *testing.T) {
	del := NewDeletion(123)
	expDel := Deletion{123}

	if del.Pos != 123 {
		t.Errorf("Expected %v from Deletion constructor but got %v", expDel, del)
	}
}

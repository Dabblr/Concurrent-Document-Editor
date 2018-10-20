package operations

import "testing"

func TestEquivIfEqualFields(t *testing.T) {
	ins1 := NewInsertion(0, 'a')
	ins2 := NewInsertion(0, 'a')

	if AreEqual(ins1, ins2) == false {
		t.Errorf("Insertions %v and %v should be considered equal.", ins1, ins2)
	}
}

func TestEquivIfSameInstance(t *testing.T) {
	ins1 := NewInsertion(0, 'a')

	if AreEqual(ins1, ins1) == false {
		t.Errorf("Insertion %v should be considered equal to itself.", ins1)
	}
}

func TestNotEquivIfDiffField(t *testing.T) {
	ins1 := NewInsertion(0, 'a')
	ins2 := NewInsertion(1, 'a')

	if AreEqual(ins1, ins2) == true {
		t.Errorf("Insertions %v and %v should NOT be considered equal.", ins1, ins2)
	}
}

func TestNotEquivIfDiffFields(t *testing.T) {
	ins1 := NewInsertion(0, 'a')
	ins2 := NewInsertion(1, 'b')

	if AreEqual(ins1, ins2) == true {
		t.Errorf("Insertions %v and %v should NOT be considered equal.", ins1, ins2)
	}
}

// NOTE: this test will be worth more when we introduce more logic into the Insertion constructor
func TestInsertionConstructor(t *testing.T) {
	ins := NewInsertion(123, 'x')
	expIns := Insertion{123, 'x'}

	if ins.Pos != 123 || ins.Val != 'x' {
		t.Errorf("Expected %v from Insertion constructor but got %v", expIns, ins)
	}
}

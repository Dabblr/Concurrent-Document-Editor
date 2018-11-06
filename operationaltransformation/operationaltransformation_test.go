package operationaltransformation

import (
	"testing"

	ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"
)

// ------------------
// Insertion onto Insertion
// ------------------
// Tests that a new insertion is changed when transformed onto another insertion that occurs at an earlier position
func TestIndirectlyDependentInsertion(t *testing.T) {
	newIns := ops.NewInsertion(5, 'b')
	oldIns := ops.NewInsertion(0, 'a')
	expectedTransformedIns := ops.NewInsertion(6, 'b')

	t.Logf("Transforming %v onto %v => expecting %v", newIns, oldIns, expectedTransformedIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if expectedTransformedIns.Equals(transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedIns, transformedIns)
	}
}

// Tests that a new insertion is NOT changed when transformed onto another insertion that occurs at a later position
func TestIndependentInsertion(t *testing.T) {
	newIns := ops.NewInsertion(0, 'a')
	oldIns := ops.NewInsertion(5, 'b')

	t.Logf("Transforming %v onto %v => expecting %v (no change)", newIns, oldIns, newIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if newIns.Equals(transformedIns) == false {
		t.Errorf("Transformation failed to leave the operation unchanged.\nExpected: %v\nFound: %v", newIns, transformedIns)
	}
}

// Tests that a new insertion is changed when transformed onto another insertion that occurs at the same position
func TestDirectlyDependentInsertion(t *testing.T) {
	newIns := ops.NewInsertion(0, 'b')
	oldIns := ops.NewInsertion(0, 'a')
	expectedTransformedIns := ops.NewInsertion(1, 'b')

	t.Logf("Transforming %v onto %v => expecting %v", newIns, oldIns, expectedTransformedIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if expectedTransformedIns.Equals(transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedIns, transformedIns)
	}
}

// ------------------
// Deletion onto Deletion
// ------------------
// Tests that a new deletion is changed when transformed onto another deletion that occurs at an earlier position
func TestDeletionTransformed(t *testing.T) {
	newDel := ops.NewDeletion(5)
	oldDel := ops.NewDeletion(0)
	expectedTransformedDel := ops.NewDeletion(4)

	t.Logf("Transforming %v onto %v => expecting %v", newDel, oldDel, expectedTransformedDel)
	transformedDel, err := TransformDeletions(newDel, oldDel)

	if expectedTransformedDel.Equals(transformedDel) == false || err != nil {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedDel, transformedDel)
	}
}

// Tests that a new deletion is NOT changed when transformed onto another deletion that occurs at a later position
func TestDeletionNotTransformed(t *testing.T) {
	newDel := ops.NewDeletion(0)
	oldDel := ops.NewDeletion(5)

	t.Logf("Transforming %v onto %v => expecting %v (no change)", newDel, oldDel, newDel)
	transformedDel, err := TransformDeletions(newDel, oldDel)

	if newDel.Equals(transformedDel) == false || err != nil {
		t.Errorf("Transformation failed to leave the operation unchanged.\nExpected: %v\nFound: %v", newDel, transformedDel)
	}
}

// Tests that a new deletion is transformed onto an older deletion if they occur at the same position
func TestDuplicateDeletionCausesError(t *testing.T) {
	newDel := ops.NewDeletion(1)

	t.Logf("Transforming %v onto %v => expecting error.", newDel, newDel)
	transformedDel, err := TransformDeletions(newDel, newDel)

	if err == nil {
		t.Errorf("Expected transformation to produce error, but it was <nil>.\nTranformed Operation: %v", transformedDel)
	}
}

// ------------------
// Deletion onto Insertion
// ------------------
// Tests that a new deletion is changed when transformed onto an insertion that occurs at an earlier position
func TestDeletionTransformedOntoInsertion(t *testing.T) {
	del := ops.NewDeletion(5)
	ins := ops.NewInsertion(0, ' ')
	expectedTransformedDel := ops.NewDeletion(6)

	t.Logf("Transforming %v onto %v => expecting %v", del, ins, expectedTransformedDel)
	transformedDel := TransformDelOnIns(del, ins)

	if expectedTransformedDel.Equals(transformedDel) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedDel, transformedDel)
	}
}

// Tests that a new deletion is NOT changed when transformed onto an insertion that occurs at a later position
func TestDeletionNotTransformedOntoInsertion(t *testing.T) {
	del := ops.NewDeletion(0)
	ins := ops.NewInsertion(5, ' ')

	t.Logf("Transforming %v onto %v => expecting %v (no change)", del, ins, del)
	transformedDel := TransformDelOnIns(del, ins)

	if del.Equals(transformedDel) == false {
		t.Errorf("Transformation failed to leave the operation unchanged.\nExpected: %v\nFound: %v", del, transformedDel)
	}
}

// Tests that a new deletion is changed when transformed onto an insertion that occurs at the same position
func TestDeletionTransformedOntoInsertionAtSamePos(t *testing.T) {
	del := ops.NewDeletion(0)
	ins := ops.NewInsertion(0, ' ')
	expectedTransformedDel := ops.NewDeletion(1)

	t.Logf("Transforming %v onto %v => expecting %v", del, ins, expectedTransformedDel)
	transformedDel := TransformDelOnIns(del, ins)

	if expectedTransformedDel.Equals(transformedDel) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedDel, transformedDel)
	}
}

// ------------------
// Insertion onto Deletion
// ------------------
// Tests that a new insertion is decremented when transformed onto a deletion that occurs at an earlier position
func TestInsertionTransformedOntoDeletion(t *testing.T) {
	ins := ops.NewInsertion(5, ' ')
	del := ops.NewDeletion(0)
	expectedTransformedIns := ops.NewInsertion(4, ' ')

	t.Logf("Transforming %v onto %v => expecting %v", ins, del, expectedTransformedIns)
	transformedIns := TransformInsOnDel(ins, del)

	if expectedTransformedIns.Equals(transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedIns, transformedIns)
	}
}

// Tests that a new insertion is NOT changed when transformed onto a deletion that occurs at a later position
func TestInsertionNotTransformedOntoDeletion(t *testing.T) {
	ins := ops.NewInsertion(0, ' ')
	del := ops.NewDeletion(5)

	t.Logf("Transforming %v onto %v => expecting %v (no change)", ins, del, ins)
	transformedIns := TransformInsOnDel(ins, del)

	if ins.Equals(transformedIns) == false {
		t.Errorf("Transformation failed to leave the operation unchanged.\nExpected: %v\nFound: %v", ins, transformedIns)
	}
}

// Tests that a new insertion is NOT changed when transformed onto a deletion that occurs at the same position
func TestInsertionTransformedOntoDeletionAtSamePos(t *testing.T) {
	ins := ops.NewInsertion(1, ' ')
	del := ops.NewDeletion(1)

	t.Logf("Transforming %v onto %v => expecting %v", ins, del, ins)
	transformedIns := TransformInsOnDel(ins, del)

	if ins.Equals(transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", ins, transformedIns)
	}
}

// Tests that a new insertion with Pos=0 is NOT changed
func TestInsertionAtPos0NotTransformed(t *testing.T) {
	ins := ops.NewInsertion(0, ' ')
	del := ops.NewDeletion(0)

	t.Logf("Transforming %v onto %v => expecting %v.", ins, del, ins)
	transformedIns := TransformInsOnDel(ins, del)

	if ins.Equals(transformedIns) == false {
		t.Errorf("Transformation failed to leave the operation unchanged.\nExpected: %v\nFound: %v", ins, transformedIns)
	}
}

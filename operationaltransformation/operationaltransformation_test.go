package operationaltransformation

import (
	"testing"

	ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"
)

// ------------------
// Insertion
// ------------------
// Tests that a new insertion is transformed onto an older insertion if the latter occurs positionally before the former
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

// Tests that a new insertion is NOT transformed onto an older insertion if the latter occurs positionally after the former
func TestIndependentInsertion(t *testing.T) {
	newIns := ops.NewInsertion(0, 'a')
	oldIns := ops.NewInsertion(5, 'b')

	t.Logf("Transforming %v onto %v => expecting %v (no change)", newIns, oldIns, newIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if newIns.Equals(transformedIns) == false {
		t.Errorf("Transformation failed to leave the operation %v unchanged.\nExpected: %v\nFound: %v", newIns, newIns, transformedIns)
	}
}

// Tests that a new insertion is transformed onto an older insertion if they occur at the same position
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

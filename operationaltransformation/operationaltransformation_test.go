package operationaltransformation

import (
	"testing"
)

// Tests that a new insertion is transformed onto an older insertion if the latter occurs positionally before the former
func TestIndirectlyDependentInsertion(t *testing.T) {
	newIns := NewInsertion(5, 'b')
	oldIns := NewInsertion(0, 'a')
	expectedTransformedIns := NewInsertion(6, 'b')

	t.Logf("Transforming %v onto %v => expecting %v", newIns, oldIns, expectedTransformedIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if AreEqual(expectedTransformedIns, transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedIns, transformedIns)
	}
}

// Tests that a new insertion is NOT transformed onto an older insertion if the latter occurs positionally after the former
func TestIndependentInsertion(t *testing.T) {
	newIns := NewInsertion(0, 'a')
	oldIns := NewInsertion(5, 'b')

	t.Logf("Transforming %v onto %v => expecting %v (no change)", newIns, oldIns, newIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if AreEqual(newIns, transformedIns) == false {
		t.Errorf("Transformation failed to leave the operation %v unchanged.\nExpected: %v\nFound: %v", newIns, newIns, transformedIns)
	}
}

// Tests that a new insertion is transformed onto an older insertion if they occur at the same position
func TestDirectlyDependentInsertion(t *testing.T) {
	newIns := NewInsertion(0, 'b')
	oldIns := NewInsertion(0, 'a')
	expectedTransformedIns := NewInsertion(1, 'b')

	t.Logf("Transforming %v onto %v => expecting %v", newIns, oldIns, expectedTransformedIns)
	transformedIns := TransformInsertions(newIns, oldIns)

	if AreEqual(expectedTransformedIns, transformedIns) == false {
		t.Errorf("Transformation did not result in expected operation.\nExpected: %v\nFound: %v", expectedTransformedIns, transformedIns)
	}
}

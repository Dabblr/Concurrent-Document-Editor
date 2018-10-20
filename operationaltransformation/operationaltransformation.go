package operationaltransformation

import ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"

// TransformInsertions transforms the a new insertion operation on top of a previous insertion operation
// e.g. (Ins{3, 'a'}, Ins[1, 'b']) => Ins{4, 'a'}
func TransformInsertions(newIns, prev ops.Insertion) ops.Insertion {
	if prev.Pos <= newIns.Pos {
		return ops.NewInsertion(newIns.Pos+1, newIns.Val)
	}
	return newIns
}

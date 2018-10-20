package operationaltransformation

import (
	"errors"

	ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"
)

// TransformInsertions transforms the a new insertion operation on top of a previous insertion operation
// e.g. (Ins{3, 'a'}, Ins{1, 'b'}) => Ins{4, 'a'}
func TransformInsertions(newIns, prev ops.Insertion) ops.Insertion {
	if prev.Pos <= newIns.Pos {
		return ops.NewInsertion(newIns.Pos+1, newIns.Val)
	}
	return newIns
}

// TransformDeletions transforms the a new deletion operation on top of a previous deletion operation
// e.g. (Del{3, 'a'}, Del{1, 'b'}) => Del{2, 'a'}
// An error is returned if the two deletions have the same position, since
// newDel's target character has already been deleted.
func TransformDeletions(newDel, prev ops.Deletion) (ops.Deletion, error) {
	if prev.Pos < newDel.Pos {
		return ops.NewDeletion(newDel.Pos - 1), nil

	} else if prev.Pos == newDel.Pos {
		return newDel, errors.New("Prob")
	}
	return newDel, nil
}

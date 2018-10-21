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
// e.g. (Del{3}, Del{1}) => Del{2}
// An error is returned if the two deletions have the same position, since
// newDel's target character has already been deleted.
func TransformDeletions(newDel, prev ops.Deletion) (ops.Deletion, error) {
	if prev.Pos < newDel.Pos {
		return ops.NewDeletion(newDel.Pos - 1), nil

	} else if prev.Pos == newDel.Pos {
		return newDel, errors.New("cannot transform deletion onto duplicate deletion")
	}
	return newDel, nil
}

// TransformDelOnIns transforms the a new deletion operation on top of a previous insertion operation
// e.g. (Del{3}, Ins{1, 'b'}) => Del{4}
func TransformDelOnIns(newDel ops.Deletion, ins ops.Insertion) ops.Deletion {
	if ins.Pos <= newDel.Pos {
		return ops.NewDeletion(newDel.Pos + 1)
	}
	return newDel
}

// TransformInsOnDel transforms the a new insertion operation on top of a previous deletion operation
// e.g. (Ins{3, 'b'}, Del{1}) => Ins{2, 'b'}
func TransformInsOnDel(newIns ops.Insertion, del ops.Deletion) ops.Insertion {
	// Cannot decrement position if already at left-most position
	if del.Pos <= newIns.Pos && newIns.Pos != 0 {
		return ops.NewInsertion(newIns.Pos-1, newIns.Val)
	}
	return newIns
}

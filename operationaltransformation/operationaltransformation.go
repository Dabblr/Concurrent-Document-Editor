package operationaltransformation

// TransformInsertions transforms the a new insertion operation on top of a previous insertion operation
// e.g. (Ins{3, 'a'}, Ins[1, 'b']) => Ins{4, 'a'}
func TransformInsertions(newIns, prev Insertion) Insertion {
	if prev.pos <= newIns.pos {
		return NewInsertion(newIns.pos+1, newIns.val)
	}
	return newIns
}

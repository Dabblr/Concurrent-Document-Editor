package operationaltransformation

import "fmt"

// Operation interface defines

// Insertion represents the addition of a character to a corresponding string
// e.g. Insertion={0, 'z'} for string "oo" => "zoo"
// e.g. Insertion={1, 'z'} for string "oo" => "ozo"
// e.g. Insertion={2, 'z'} for string "oo" => "ooz"
type Insertion struct {
	pos int
	val rune
}

// NewInsertion is a constructor for the Insertion type
func NewInsertion(pos int, val rune) Insertion {
	return Insertion{pos, val}
}

// AreEqual defines what makes two Insertions equal
func AreEqual(ins1, ins2 Insertion) bool {
	if ins1.pos == ins2.pos && ins1.val == ins2.val {
		return true
	}
	return false
}

// String implements the Stringer interface
func (ins *Insertion) String() string {
	return fmt.Sprintf("Ins: pos=%v val=%v", ins.pos, ins.val)
}

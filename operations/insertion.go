package operations

import "fmt"

// Operation interface defines

// Insertion represents the addition of a character to a corresponding string
// e.g. Insertion={0, 'z'} for string "oo" => "zoo"
// e.g. Insertion={1, 'z'} for string "oo" => "ozo"
// e.g. Insertion={2, 'z'} for string "oo" => "ooz"
type Insertion struct {
	Pos int
	Val rune
}

// NewInsertion is a constructor for the Insertion type
func NewInsertion(pos int, val rune) Insertion {
	return Insertion{pos, val}
}

// AreEqual defines what makes two Insertions equal
func AreEqual(ins1, ins2 Insertion) bool {
	if ins1.Pos == ins2.Pos && ins1.Val == ins2.Val {
		return true
	}
	return false
}

// String implements the Stringer interface
func (ins *Insertion) String() string {
	return fmt.Sprintf("Ins: pos=%v val=%v", ins.Pos, ins.Val)
}

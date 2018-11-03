package operations

import "fmt"

// Deletion represents the addition of a character to a corresponding string
// e.g. Deletion={0} for string "abc" => "bc"
// e.g. Deletion={1} for string "abc" => "ac"
// e.g. Deletion={2} for string "abc" => "ab"
type Deletion struct {
	Pos int
}

// NewDeletion is a constructor for the Deletion type
func NewDeletion(pos int) Deletion {
	return Deletion{pos}
}

// Equals defines what makes two Deletions equal
func (del *Deletion) Equals(del2 Deletion) bool {
	if del.Pos == del2.Pos {
		return true
	}
	return false
}

// String implements the Stringer interface
func (del *Deletion) String() string {
	return fmt.Sprintf("Del: pos=%v", del.Pos)
}

package objects

import (
	"fmt"

	ops "github.com/jcgallegdup/Concurrent-Document-Editor/operations"
)

// Change represents each individual change (insertion/deletion).
type Change struct {
	// The type of the change (insert or delete).
	Type string `json:"type"`
	// The index the change is being made at.
	Position int `json:"position"`
	// The value of the character being inserted or deleted.
	Value string `json:"value"`
}

// NewChange is a constructor for the Change type.
func NewChange(changeType string, position int, value string) Change {
	return Change{changeType, position, value}
}

// Equals defines what makes two Change objects equal.
func (change *Change) Equals(change2 Change) bool {
	if change.Type == change2.Type && change.Position == change2.Position && change.Value == change2.Value {
		return true
	}
	return false
}

// IsValid checks that only a single character is being inserted/deleted and the type is insert or delete.
func (change *Change) IsValid() bool {
	if len(change.Value) > 1 {
		return false
	}
	if change.Type != "insert" && change.Type != "delete" {
		return false
	}
	return true
}

// ChangeToIns converts the Change object to an equivalent Insertion.
func (change *Change) ChangeToIns() ops.Insertion {
	return ops.NewInsertion(change.Position, rune(change.Value[0]))
}

// ChangeToDel converts the Change object to an equivalent Deletion.
func (change *Change) ChangeToDel() ops.Deletion {
	return ops.NewDeletion(change.Position)
}

// String determines the default string format for the Change type.
func (change Change) String() string {
	switch change.Type {
	case "insert":
		return fmt.Sprintf("Ins: pos=%d val=%s", change.Position, change.Value)
	case "delete":
		return fmt.Sprintf("Del: pos=%d", change.Position)
	default:
		return ""
	}
}

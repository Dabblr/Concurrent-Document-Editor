package objects

import "fmt"

// Revision represents all changes made in a single update.
type Revision struct {
	// The username of the user making the revision.
	User string `json:"user"`
	// TODO rename this to file ID
	// The ID of the file.
	ID int `json:"id"`
	// The revision number.
	RevisionNumber int `json:"revision_number"`
	// The array of changes contained in the revision.
	Changes []Change `json:"changes"`
}

// NewRevision is a constructor for the Revision Type.
func NewRevision(user string, id int, revision int, changes []Change) Revision {
	return Revision{user, id, revision, changes}
}

// Equals defines what makes two Revision objects equal.
func (rev *Revision) Equals(rev2 Revision) bool {
	if rev.User != rev2.User || rev.ID != rev2.ID || rev.RevisionNumber != rev2.RevisionNumber || len(rev.Changes) != len(rev2.Changes) {
		return false
	}
	for i, change := range rev.Changes {
		if change.Equals(rev2.Changes[i]) == false {
			return false
		}
	}
	return true
}

// String determines the default string format for the Revision type.
func (rev Revision) String() string {
	return fmt.Sprintf("Revision: user=%s id=%d revisionNumber=%d changes=%v", rev.User, rev.ID, rev.RevisionNumber, rev.Changes)
}

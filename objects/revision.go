package objects

// Revision represents all changes made in a single update.
type Revision struct {
	// The username of the user making the revision.
	User string `json:"user"`
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

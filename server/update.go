package main

import (
	"errors"
	"log"

	obj "github.com/jcgallegdup/Concurrent-Document-Editor/objects"
	opTransform "github.com/jcgallegdup/Concurrent-Document-Editor/operationaltransformation"
)

// ApplyUpdate applies all the changes contained in a revision to a file.
// Returns an error if a change in the revision is invalid. It will apply all changes up until the invalid change.
func ApplyUpdate(revision obj.Revision, file obj.File) error {
	var err error
	prevChanges := mockGetChangesSinceRevision(revision.ID, revision.RevisionNumber)
	fileContent := file.Content
	for _, change := range revision.Changes {
		if !change.IsValid() {
			err = errors.New("invalid change:" + change.String())
			break
		}

		log.Println("Original change:", change)
		transformedChange, transformErr := TransformChange(change, prevChanges)
		if transformErr != nil {
			// Change was not applied (duplicate), move on to the next one.
			log.Println("Change was not applied.")
			continue
		}

		log.Println("Transformed change:", transformedChange)
		fileContent, err = ApplyChange(transformedChange, fileContent)
		if err != nil {
			// Index was out of range.
			break
		}
		mockInsertChange(revision.ID, transformedChange)
	}
	mockUpdateFileContent(revision.ID, fileContent)
	return err
}

// TransformChange transforms a new change so it can be applied on top of the changes that have already occurred.
// Returns an error if the change should not be applied (ex: deletion in the same position as a previous deletion)
func TransformChange(newChange obj.Change, prevChanges []obj.Change) (obj.Change, error) {
	for _, change := range prevChanges {
		switch {
		case change.Type == "insert" && newChange.Type == "insert":
			// insertion on insertion
			newIns := opTransform.TransformInsertions(newChange.ChangeToIns(), change.ChangeToIns())
			newChange.Position = newIns.Pos
		case change.Type == "insert" && newChange.Type == "delete":
			// deletion on insertion
			newDel := opTransform.TransformDelOnIns(newChange.ChangeToDel(), change.ChangeToIns())
			newChange.Position = newDel.Pos
		case change.Type == "delete" && newChange.Type == "insert":
			// insertion on deletion
			newIns := opTransform.TransformInsOnDel(newChange.ChangeToIns(), change.ChangeToDel())
			newChange.Position = newIns.Pos
		case change.Type == "delete" && newChange.Type == "delete":
			// deletion on deletion
			newDel, err := opTransform.TransformDeletions(newChange.ChangeToDel(), change.ChangeToDel())
			if err != nil {
				return newChange, err
			}
			newChange.Position = newDel.Pos
		}
	}
	return newChange, nil
}

// ApplyChange updates the file content to reflect the new change.
// Returns an error if the position of the change is out of range.
func ApplyChange(change obj.Change, fileContent string) (string, error) {
	switch change.Type {
	case "insert":
		if change.Position < 0 || change.Position > len(fileContent) {
			// Index out of range.
			return fileContent, errors.New("index out of range")
		}
		return (fileContent[:change.Position] + change.Value + fileContent[change.Position:]), nil
	case "delete":
		if change.Position < 0 || change.Position >= len(fileContent) {
			// Index out of range.
			return fileContent, errors.New("index out of range")
		}
		return (fileContent[:change.Position] + fileContent[change.Position+1:]), nil
	default:
		return fileContent, errors.New("invalid change type")
	}
}

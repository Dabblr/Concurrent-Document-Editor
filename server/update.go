package main

import (
	"errors"
	"fmt"
	"log"

	db "github.com/Dabblr/Concurrent-Document-Editor/database"
	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
	opTransform "github.com/Dabblr/Concurrent-Document-Editor/operationaltransformation"
)

// ApplyUpdate applies all the changes contained in a revision to a file.
// Returns an error if a change in the revision is invalid.
func ApplyUpdate(revision obj.Revision, file obj.File, database db.Interface) error {
	var err error
	var changesToApply []obj.Change
	prevChanges, err := database.GetChangesSinceRevision(revision.ID, revision.RevisionNumber)
	if err != nil {
		return err
	}

	fileContent := file.Content
	for _, change := range revision.Changes {
		if !change.IsValid() {
			return errors.New("invalid change:" + change.String())
		}

		log.Println("Original change:", change)
		transformedChange, transformErr := TransformChange(change, prevChanges)
		if transformErr != nil {
			// A change had an invalid type.
			return transformErr
		}

		log.Println("Transformed change:", transformedChange)
		fileContent, err = ApplyChange(transformedChange, fileContent)
		if err != nil {
			// Index was out of range.
			return err
		}
		changesToApply = append(changesToApply, transformedChange)
	}

	err = database.InsertChanges(revision.ID, changesToApply)
	if err != nil {
		return err
	}

	err = database.UpdateFileContent(revision.ID, fileContent)
	if err != nil {
		return err
	}

	return nil
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
			if err == nil {
				// only update the position if no duplicate deletion error was returned
				newChange.Position = newDel.Pos
			}
		default:
			return newChange, errors.New("invalid change type: " + newChange.Type)
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
			return fileContent, fmt.Errorf("index %d out of range", change.Position)
		}
		return (fileContent[:change.Position] + change.Value + fileContent[change.Position:]), nil
	case "delete":
		if change.Position < 0 || change.Position >= len(fileContent) {
			// Index out of range.
			return fileContent, fmt.Errorf("index %d out of range", change.Position)
		}
		return (fileContent[:change.Position] + fileContent[change.Position+1:]), nil
	default:
		return fileContent, errors.New("invalid change type")
	}
}

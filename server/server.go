package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// File contains information about each file.
type File struct {
	// The username of the user updating the file.
	User string `json:"user,omitempty"`
	// The ID of the file (generated on server).
	ID int `json:"id"`
	// The name of the file.
	Name string `json:"name"`
	// The revision number of the file.
	RevisionNumber int `json:"revision_number"`
	// The file content.
	Content string `json:"content,omitempty"`
}

// Change contains information about each individual change (insertion/deletion).
type Change struct {
	// The type of the change (insert or delete).
	Type string `json:"type"`
	// The index the change is being made at.
	Position int `json:"position"`
	// The value of the character being inserted or deleted.
	Value string `json:"value"`
}

// Revision contains information about all changes made in a single update.
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

// Keeps track of how many files have been created to generate new ids.
var mockFileCounter int

// Mocks storing a new file in the database and returning an id for it.
func mockCreateEmptyFile(fileName string, userName string) int {
	mockFileCounter++
	return mockFileCounter
}

// Mocks returning the latest revision file content for the given file id.
func mockGetFileContent(id int) (File, error) {
	if id <= 0 || id > mockFileCounter {
		// Invalid id.
		return File{"", 0, "", 0, ""}, errors.New("invalid file id")
	}
	return File{"", id, "fileName", 1, "This is the file content."}, nil
}

// Mocks returning an array of all changes after the given revision number for the given file id.
func mockGetChangesSinceRevision(id int, revisionNumber int) []Change {
	changes := []Change{
		{"insert", 0, "a"},
		{"insert", 1, "b"},
	}
	return changes
}

// Mocks inserting a new change to a file in the database.
func mockInsertChange(id int, change Change) {
	return
}

// Mocks updating the file content for the given file id in the database.
func mockUpdateFileContent(id int, fileContent string) {
	return
}

func transformChange(newChange Change, databaseChanges []Change) Change {
	for _, change := range databaseChanges {
		switch {
		case change.Type == "insert" && newChange.Type == "insert":
			// insertion on insertion
		case change.Type == "insert" && newChange.Type == "delete":
			// deletion on insertion
		case change.Type == "delete" && newChange.Type == "insert":
			// insertion on deletion
		case change.Type == "delete" && newChange.Type == "delete":
			// deletion on deletion
		}
	}
	return newChange
}

func applyChangeToFileContent(change Change, fileContent string) string {
	switch change.Type {
	case "insert":
		return (fileContent[:change.Position] + change.Value + fileContent[change.Position:])
	case "delete":
		return (fileContent[:change.Position] + fileContent[change.Position+1:])
	default:
		return fileContent
	}
}

func makeChanges(revision Revision, file File) {
	databaseChanges := mockGetChangesSinceRevision(revision.ID, revision.RevisionNumber)
	fileContent := file.Content
	for _, change := range revision.Changes {
		log.Println("Original change:", change)
		transformedChange := transformChange(change, databaseChanges)
		log.Println("Transformed change:", change)
		log.Println("Original file content:", fileContent)
		fileContent = applyChangeToFileContent(transformedChange, fileContent)
		log.Println("New file content:", fileContent, "\n")
		mockInsertChange(revision.ID, transformedChange)
	}
	mockUpdateFileContent(revision.ID, fileContent)
}

// CreateFile creates a new empty file and returns the associated file object.
func CreateFile(w http.ResponseWriter, r *http.Request) {
	var file File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil || file.Name == "" || file.User == "" {
		// Request was missing required fields or poorly formed.
		log.Println("POST request to /file was missing required field(s) or poorly formed.")
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file.ID = mockCreateEmptyFile(file.Name, file.User)
	file.RevisionNumber = 1
	log.Printf("\"%s\" created a new file called \"%s\" with ID %d.\n", file.User, file.Name, file.ID)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(file)
}

// GetFile returns the latest revision file content for the given file id.
func GetFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Invalid ID.
		log.Printf("GET request to /file/%s contained an invalid file ID.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := mockGetFileContent(id)
	if err != nil {
		// Invalid ID.
		log.Printf("GET request to /file/%s contained an invalid file ID.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(file)
}

// PostUpdates adds the updates to the file.
func PostUpdates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var revision Revision
	err := json.NewDecoder(r.Body).Decode(&revision)
	if err != nil || revision.User == "" || revision.ID == 0 || revision.RevisionNumber == 0 {
		// Missing required fields or poorly formed request.
		log.Printf("POST request to /file/%s was missing required field(s) or poorly formed.\n", params["id"])
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := mockGetFileContent(revision.ID)
	if err != nil {
		// Invalid ID.
		log.Printf("POST request to /file/%s contained an invalid file ID.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("\"%s\" made updates to File ID %d\n", revision.User, revision.ID)
	makeChanges(revision, file)
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/file", CreateFile).Methods("POST")
	router.HandleFunc("/file/{id}", GetFile).Methods("GET")
	router.HandleFunc("/file/{id}", PostUpdates).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	db "github.com/Dabblr/Concurrent-Document-Editor/database"
	obj "github.com/Dabblr/Concurrent-Document-Editor/objects"
	"github.com/gorilla/mux"
)

// Change the type of this based on environment.
var database db.Interface

// DBPATH is the path to the database file
const DBPATH = "../updates.db"

// CreateUser creates a new user and adds it to the user database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user obj.File
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.User == "" {
		// Request was missing required fields or poorly formed.
		log.Println("POST request to /users was missing required field(s) or poorly formed.")
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = database.CreateUser(user.User)
	if err != nil {
		// Unable to create a user with the given username.
		log.Println("POST request to /users contained an invalid username, error generated:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("User with username %s created.\n", user.User)
	w.WriteHeader(http.StatusCreated)
}

// CreateFile creates a new empty file and returns the associated file object.
func CreateFile(w http.ResponseWriter, r *http.Request) {
	var file obj.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil || file.Name == "" || file.User == "" {
		// Request was missing required fields or poorly formed.
		log.Println("POST request to /files was missing required field(s) or poorly formed.")
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file.ID, file.RevisionNumber, err = database.CreateEmptyFile(file.Name, file.User)
	if err != nil {
		log.Println("POST request to /files failed, unable to create new file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("File created.", file)
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
		log.Printf("GET request to /files/%s did not contain an integer file ID.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := database.GetFileContent(id)
	if err != nil {
		// Invalid ID.
		log.Printf("Could not retrieve file for given ID: %s.\n", params["id"])
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
	id, conversionErr := strconv.Atoi(params["id"])
	if conversionErr != nil {
		// Post request contained an invalid id.
		log.Printf("POST request to /files/%s did not contain an integer file ID.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var revision obj.Revision
	err := json.NewDecoder(r.Body).Decode(&revision)
	if err != nil || revision.User == "" || revision.ID == 0 || revision.RevisionNumber == 0 || revision.ID != id {
		// Missing required fields or poorly formed request.
		log.Printf("POST request to /files/%s was missing required field(s) or poorly formed.\n", params["id"])
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := database.GetFileContent(revision.ID)
	if err != nil {
		// Invalid ID.
		log.Printf("Could not retrieve file for given ID: %s.\n", params["id"])
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("\"%s\" made updates to File ID %d\n", revision.User, revision.ID)
	err = ApplyUpdate(revision, file, database)
	if err != nil {
		// Updates were not applied.
		log.Println("Updates were not applied:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	var dev *bool
	dev = flag.Bool("dev", false, "set this flag to use the mock database")
	flag.Parse()
	if *dev {
		fmt.Println("Starting mock database")
		database = &db.MockDB{}
	} else {
		fmt.Println("Starting real database")
		dbtemp := db.CreateEmptyDb(DBPATH)
		database = &dbtemp
	}
	// dbtemp := db.CreateEmptyDb(DBPATH)
	// dbtemp.CreateUser("Nikita")
	// dbtemp.CreateEmptyFile("file27.txt", "Nikita")
	// err := dbtemp.UpdateFileContent(1, "UPDATED!!!")
	// fmt.Printf("'Error' %v\n", err)

	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/files", CreateFile).Methods("POST")
	router.HandleFunc("/files/{id}", GetFile).Methods("GET")
	router.HandleFunc("/files/{id}", PostUpdates).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

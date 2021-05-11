
// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Document - Our struct for all Documents
type Document struct {
	Id      string    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Documents []Document

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllDocs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Documents)
}

func returnSingleDocs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, Document := range Documents {
		if Document.Id == key {
			json.NewEncoder(w).Encode(Document)
		}
	}
}


func createNewDocs(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Document struct
	// append this to our Documents array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var Document Document
	json.Unmarshal(reqBody, &Document)
	// update our global Documents array to include
	// our new Document
	Documents = append(Documents, Document)

	json.NewEncoder(w).Encode(Document)
}

func deleteDocs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, Document := range Documents {
		if Document.Id == id {
			Documents = append(Documents[:index], Documents[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/Documents", returnAllDocs)
	myRouter.HandleFunc("/Document", createNewDocs).Methods("POST")
	myRouter.HandleFunc("/Document/{id}", deleteDocs).Methods("DELETE")
	myRouter.HandleFunc("/Document/{id}", returnSingleDocs)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Documents = []Document{
		Document{Id: "1", Title: "Hello", Desc: "Document Description", Content: "Document Content"},
		Document{Id: "2", Title: "Hello 2", Desc: "Document Description", Content: "Document Content"},
	}
	handleRequests()
}
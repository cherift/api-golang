package controllers

import(
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	model "models"
)

// Gets list of documents
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Vade Secure Test\n\n")
	json.NewEncoder(w).Encode(model.Documents)
}

// Creates a new document
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	id, converr := strconv.Atoi(params.Get("id"))
	name := params.Get("name")
	desc := params.Get("desc")

	if converr == nil && name != "" && desc != "" {
		model.Documents[id] = model.Document{ID: id, Name: name, Description: desc}
	}

	fmt.Fprintf(w, "A new document has been added.\n{ID : %d, Name: %s, Description: %s}", id, name, desc)
}

// Removes document from the list
func RemoveDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	if document, isPresent := model.Documents[id]; isPresent {
		delete(model.Documents, document.ID)
		fmt.Fprintf(w, "The document %d has been deleted \n", document.ID)
		return
	}

	fmt.Fprintf(w, "No existing document corresponding to the ID %d \n", id)
}
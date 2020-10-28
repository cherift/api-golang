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

	id, _ := strconv.Atoi(params.Get("id"))
	name := params.Get("name")
	desc := params.Get("desc")

	// check if the document is present or a parameter is missing
	if _, isPresent := model.Documents[id]; !isPresent {
		if name != "" && desc != "" {
			model.Documents[id] = model.Document{ID: id, Name: name, Description: desc}
			fmt.Fprintf(w, "A new document has been added.\n{ID : %d, Name: %s, Description: %s}", id, name, desc)
		} else {
			fmt.Fprint(w, "One or more arguments missing (name, desc)")
		}
		return
	}

	// print a message if the document exists
	fmt.Fprintf(w, "The document %d allready exists\n", id)
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

// Gets a document
func GetDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	if document, isPresent := model.Documents[id]; isPresent {
		fmt.Fprintf(w, "Document Founded !\nID : %d \nName : %s \nDescription : %s \n", 
			document.ID, document.Name, document.Description)
		return
	}

	fmt.Fprintf(w, "No existing document corresponding to the ID %d \n", id)
}
package controllers

import(
	"fmt"
	"net/http"
	"strconv"
	model "models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Vade Secure Test")
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
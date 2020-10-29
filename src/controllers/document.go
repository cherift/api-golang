package controllers

import(
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	model "models"
)

type Result struct {
	Message string
	Code	int
	Result 	[]model.Document
}  

// Gets list of documents
func HomePage(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Message :"Vade Secure Test",
		Code    : 200, 
		Result  : func (datas map[int]model.Document) []model.Document {
			list := []model.Document{}
			for _, value := range datas {
				list = append(list, value)
			}
			return list
		}(model.Documents),
	}

	json.NewEncoder(w).Encode(result)
}

// Creates a new document
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	var result Result

	params := r.URL.Query()

	id, _ := strconv.Atoi(params.Get("id"))
	name := params.Get("name")
	desc := params.Get("desc")

	// check if the document is present or a parameter is missing
	if _, isPresent := model.Documents[id]; !isPresent {
		if name != "" && desc != "" {
			model.Documents[id] = model.Document{ID: id, Name: name, Description: desc}

			result.Message = "New document added"
			result.Code    = 200
			result.Result  = []model.Document{model.Documents[id],}
		} else {
			result.Message = "One or more arguments missing (name, desc)"
		}
	} else {
		result.Message = "The document allready exists"
		result.Result  = []model.Document{model.Documents[id],}
	}

	json.NewEncoder(w).Encode(result)
}

// Removes document from the list
func RemoveDocument(w http.ResponseWriter, r *http.Request) {
	var result Result

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	if document, isPresent := model.Documents[id]; isPresent {
		result.Message = "The document has been deleted"
		result.Code    = 200
		result.Result  = []model.Document{document,}
		delete(model.Documents, document.ID)
	} else {
		result.Message = "No corresponding document"
	}

	json.NewEncoder(w).Encode(result)
}

// Gets a document
func GetDocument(w http.ResponseWriter, r *http.Request) {
	var result Result

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	if document, isPresent := model.Documents[id]; isPresent {
		result.Message = "The document has been founded"
		result.Code    = 200
		result.Result  = []model.Document{document,}
	} else {
		result.Message = "No corresponding document"
	}

	json.NewEncoder(w).Encode(result)
}
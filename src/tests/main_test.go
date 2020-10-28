package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	model "models"
	control "controllers"
)

// Tests the creating document funciton
func TestCreateDocument(t *testing.T) {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=12&name=name12&desc=description12", nil)
	rec := httptest.NewRecorder()
	
	assert.Equal(t, 0, len(model.Documents), "No document has been added yet")
	control.CreateDocument(rec, req)
	assert.Equal(t, 1, len(model.Documents), "document has been added")	
}

// Tests the creating of existing document document 
func TestCreateExistingDocument(t *testing.T) {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=12&name=newname12&desc=newdescription1", nil)
	rec := httptest.NewRecorder()
	
	size := len(model.Documents)
	assert.Equal(t, size, len(model.Documents), "No document has been added yet")
	control.CreateDocument(rec, req)
	// The size is always the same
	assert.NotEqual(t, size-1, len(model.Documents), "document has been added")
	assert.Equal(t, size, len(model.Documents), "document has been added")	
}

// Tests getting document
func TestGetDocument(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/document/{id}", control.GetDocument).Methods("GET")
	
	const result string = "Document Founded !\nID : 12 \nName : name12 \nDescription : description12 \n"

	req, _ := http.NewRequest("GET", "/document/12", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, result, rec.Body.String(), "Get document information")
}

// Tests getting non existing document
func TestGetNonExistingDocument(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/document/{id}", control.GetDocument).Methods("GET")
	
	const result string = "No existing document corresponding to the ID 444444 \n"

	req, _ := http.NewRequest("GET", "/document/444444", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, result, rec.Body.String(), "Get document information")
}

// Tests removig document funciton
func TestRemoveDocument(t *testing.T) {
	size := len(model.Documents)

	assert.Equal(t, size, len(model.Documents), "size of the list of documents not changed yet")

	router := mux.NewRouter()
	router.HandleFunc("/remove/{id}", control.RemoveDocument).Methods("DELETE")
	
	req, _ := http.NewRequest("DELETE", "/remove/12", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// size has been decreased
	assert.Equal(t, size-1, len(model.Documents), "The only one object has been removed")
}

// Tests removig non existing projet
func TestRemoveNotExistingDocument(t *testing.T) {
	size := len(model.Documents)

	assert.Equal(t, size, len(model.Documents), "size of the list of documents not changed yet")

	router := mux.NewRouter()
	router.HandleFunc("/remove/{id}", control.RemoveDocument).Methods("DELETE")
	
	req, _ := http.NewRequest("DELETE", "/remove/500000", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	// size has is the same
	assert.Equal(t, size, len(model.Documents), "The only one object has been removed")
}
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
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=12&name=name1&desc=description1", nil)
	rec := httptest.NewRecorder()
	
	assert.Equal(t, 0, len(model.Documents), "No document has been added yet")
	control.CreateDocument(rec, req)
	assert.Equal(t, 1, len(model.Documents), "document has been added")	
}


// Tests the creating of existing document document 
func TestCreateExistingDocument(t *testing.T) {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=12&name=name1&desc=description1", nil)
	rec := httptest.NewRecorder()
	
	size := len(model.Documents)
	assert.Equal(t, size, len(model.Documents), "No document has been added yet")
	control.CreateDocument(rec, req)
	// The size is always the same
	assert.NotEqual(t, size-1, len(model.Documents), "document has been added")
	assert.Equal(t, size, len(model.Documents), "document has been added")	
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
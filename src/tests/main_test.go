package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	model "models"
	control "controllers"
)

//Tests home page
func TestHomePAge(t *testing.T)  {
	req, _ := http.NewRequest("POST", "localhost:8080", nil)
	rec := httptest.NewRecorder()

	control.HomePage(rec, req)

	mock := control.Result{
		Message : "Vade Secure Test",
		Code	: 200,
		Result 	: []model.Document{},
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}

}

// Tests the creating document funciton
func TestCreateDocument(t *testing.T) {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=12&name=name12&desc=description12", nil)
	rec := httptest.NewRecorder()
	
	size := len(model.Documents)
	assert.Equal(t, size, len(model.Documents), "No document has been added yet")
	control.CreateDocument(rec, req)
	assert.Equal(t, size+1 , len(model.Documents), "document has been added")	
}

//Tests the response after the creating of document
func TestCreateDocumentResponse(t *testing.T) {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=1&name=name1&desc=description1", nil)
	rec := httptest.NewRecorder()

	control.CreateDocument(rec, req)

	mock := control.Result{
		Message : "New document added",
		Code	: 200,
		Result 	: []model.Document{
			model.Document{ID: 1, Name: "name1", Description: "description1"},
		},
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}
} 

// Tests the creating of existing document 
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

// Tests the creating of existing document response
func TestCrateExtingDocumentReponse(t *testing.T)  {
	req, _ := http.NewRequest("POST", "localhost:8080/create?id=1&name=name1&desc=description1", nil)
	rec := httptest.NewRecorder()

	control.CreateDocument(rec, req)

	mock := control.Result{
		Message : "The document allready exists",
		Code	: 0,
		Result 	: []model.Document{
			model.Document{ID: 1, Name: "name1", Description: "description1"},
		},
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}
}

// Tests getting document
func TestGetDocument(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/document/{id}", control.GetDocument).Methods("GET")
	
	req, _ := http.NewRequest("GET", "/document/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	mock := control.Result{
		Message : "The document has been founded",
		Code	: 200,
		Result 	: []model.Document{
			model.Document{ID: 1, Name: "name1", Description: "description1"},
		},
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}
}

// Tests getting non existing document
func TestGetNonExistingDocument(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/document/{id}", control.GetDocument).Methods("GET")

	req, _ := http.NewRequest("GET", "/document/444444", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	mock := control.Result{
		Message : "No corresponding document",
		Code	: 0,
		Result 	: nil,
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}	
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

// Tests removing document response
func TestRemovingDocumentResponse(t *testing.T)  {
	router := mux.NewRouter()
	router.HandleFunc("/remove/{id}", control.RemoveDocument).Methods("DELETE")
	
	req, _ := http.NewRequest("DELETE", "/remove/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	mock := control.Result{
		Message : "The document has been deleted",
		Code	: 200,
		Result 	: []model.Document{
			model.Document{ID: 1, Name: "name1", Description: "description1"},
		},
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}
}

// Tests removig non existing document
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

// Tests removing non existing document response
func TestRemoveNotExistingDocumentResponse(t *testing.T)  {
	router := mux.NewRouter()
	router.HandleFunc("/remove/{id}", control.RemoveDocument).Methods("DELETE")
	
	req, _ := http.NewRequest("DELETE", "/remove/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	mock := control.Result{
		Message : "No corresponding document",
		Code	: 0,
		Result 	: nil,
	}

	body := rec.Body.String()
	var result control.Result
	err := json.Unmarshal([]byte(body), &result)

	if err != nil {
		assert.Equal(t, mock, result, "result is corresponding to the mock value")
	}
}
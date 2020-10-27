package models

// a documment object
type Document struct {
	ID          int 	`json:"ID"`
	Name        string	`json:"Name"`
	Description string	`json:"Description"`
}

// a dictionnary of documents
var Documents map[int]Document

// Initilizes the dictionnary of documents
func init() {
	Documents = map[int]Document{}
} 
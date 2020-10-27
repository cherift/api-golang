package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	control "controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", control.HomePage)
	router.HandleFunc("/create", control.CreateDocument)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

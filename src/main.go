package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	control "controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", control.HomePage).Methods("GET")
	router.HandleFunc("/create", control.CreateDocument).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

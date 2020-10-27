package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Vade Secure Test")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomePage)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

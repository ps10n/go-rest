package main

import (
	"log"
	"net/http"

	h "./handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/msgs", h.GetMessages).Methods("GET")
	r.HandleFunc("/msgs/{id}", h.GetMessage).Methods("GET")
	r.HandleFunc("/msgs", h.CreateMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

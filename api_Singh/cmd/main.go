package main

import (
	"log"
	"net/http"

	"github.com/RSGuelfi/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)

	// Here we'll define our api endpoints

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
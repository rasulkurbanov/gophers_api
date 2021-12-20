package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rasulkurbanov/gophers_api/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	//Get all books from mock db
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)

	//Create a new book
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)

	//Get a single book by ID
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)

	//Update a book by ID
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)

	//Delete a book by ID
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("Server is running on PORT")
	http.ListenAndServe(":4000", router)

}
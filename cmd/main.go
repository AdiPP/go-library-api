package main

import (
	"log"
	"net/http"

	"github.com/AdiPP/go/library-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main () {
	router := mux.NewRouter()

	// Ping Routes
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// Business Routes
	// Fetch All Books
	router.HandleFunc("/books", handlers.FetchAllBooks).Methods(http.MethodGet)

	// Add Book
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)

	// Get Book
	router.HandleFunc("/books/{book}", handlers.GetBook).Methods(http.MethodGet)

	// Update Book
	router.HandleFunc("/books/{book}", handlers.UpdateBook).Methods(http.MethodPut)

	// Delete Book
	router.HandleFunc("/books/{book}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8000", router)
}
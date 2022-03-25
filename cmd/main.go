package main

import (
	"log"
	"net/http"

	"github.com/AdiPP/go/library-api/pkg/db"
	"github.com/AdiPP/go/library-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main () {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	// Ping Routes
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// Business Routes
	// Fetch All Books
	router.HandleFunc("/books", h.FetchAllBooks).Methods(http.MethodGet)

	// Add Book
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)

	// Get Book
	router.HandleFunc("/books/{book}", h.GetBook).Methods(http.MethodGet)

	// Update Book
	router.HandleFunc("/books/{book}", h.UpdateBook).Methods(http.MethodPut)

	// Delete Book
	router.HandleFunc("/books/{book}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8000", router)
}
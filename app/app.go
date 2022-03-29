package app

import (
	"net/http"

	"github.com/AdiPP/go/library-api/app/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB database.LibraryDB
}

func New() *App {
	a := &App {
		Router: mux.NewRouter(),
	}

	a.initRoutes()

	return a;
}

func (a *App) initRoutes() {
	router := a.Router

	// Ping Routes
	router.HandleFunc("/v1/ping", a.PingHandler()).Methods(http.MethodGet)

	// Business Routes
	// Get Books
	router.HandleFunc("/v1/books", a.GetBooksHandler()).Methods(http.MethodGet)

	// Add Book
	router.HandleFunc("/v1/books", a.CreateBookHandler()).Methods(http.MethodPost)

	// // Get Book
	// router.HandleFunc("/books/{book}", h.GetBook).Methods(http.MethodGet)

	// // Update Book
	// router.HandleFunc("/books/{book}", h.UpdateBook).Methods(http.MethodPut)

	// // Delete Book
	// router.HandleFunc("/books/{book}", h.DeleteBook).Methods(http.MethodDelete)
}
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AdiPP/go/library-api/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to Book mocks
	// book.Id = rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)

	// Append to the books table
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}
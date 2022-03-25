package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/AdiPP/go/library-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book"])

	// Read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var UpdateBook models.Book
	json.Unmarshal(body, &UpdateBook)

	// Iterate oer all mock Books
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		// Update and send a response when book Id mathces dynamic Id
	// 		book.Title = UpdateBook.Title
	// 		book.Author = UpdateBook.Author
	// 		book.Desc = UpdateBook.Desc

	// 		mocks.Books[index] = book
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(book)
	// 		break
	// 	}
	// }

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = UpdateBook.Title
	book.Author = UpdateBook.Author
	book.Desc = UpdateBook.Desc

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
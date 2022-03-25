package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AdiPP/go/library-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book"])

	// Iterate over all the mock books
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		// Delete book and send a response if the book Id matches dynamic Id
	// 		mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
			
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(book)
	// 		break
	// 	}
	// }

	// Find the book by Id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
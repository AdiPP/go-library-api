package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AdiPP/go/library-api/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book"])

	// Iterate over all the mock books
	// for _, book := range mocks.Books {
	// 	if book.Id == id {
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
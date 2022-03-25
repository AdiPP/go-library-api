package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdiPP/go/library-api/pkg/models"
)

func (h handler) FetchAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AdiPP/go/library-api/app/models"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Cannot format json. err%v \n", err)
	}
}

func mapPingToJSON(p *models.Ping) models.JsonPing {
	return models.JsonPing{
		Status: p.Status,
	}
}

func mapBookToJSON(b *models.Book) models.JsonBook {
	return models.JsonBook{
		Id: b.Id,
		Title: b.Title,
		Author: b.Author,
		Desc: b.Desc,
	}
}
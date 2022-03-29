package app

import (
	"log"
	"net/http"

	"github.com/AdiPP/go/library-api/app/mocks"
	"github.com/AdiPP/go/library-api/app/models"
)

func (a *App) PingHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		p := &mocks.Ping

		resp := mapPingToJSON(p)

		sendResponse(w, r, resp, http.StatusAccepted)
	}
}

func (a *App) CreateBookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.BookRequest{}

		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		b := &models.Book{
			Id: 0,
			Title: req.Title,
			Author: req.Author,
			Desc: req.Desc,
		}

		err = a.DB.CreateBook(b)

		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapBookToJSON(b)

		sendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *App) GetBooksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := a.DB.GetBooks()

		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonBook, len(books))

		for idx, book := range books {
			resp[idx] = mapBookToJSON(book)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
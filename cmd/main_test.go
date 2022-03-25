package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AdiPP/go/library-api/pkg/db"
	"github.com/AdiPP/go/library-api/pkg/mocks"
	"github.com/AdiPP/go/library-api/pkg/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	DB := db.Init()
	return Routes(DB)
}

func TestPing(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	var result models.Ping

	json.Unmarshal([]byte(response.Body.Bytes()), &result)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, result.Status, "pong")
}

func TestAddBook(t *testing.T) {
	book := mocks.Book
	jsonValue, _ := json.Marshal(book)

	request, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	var result models.Book

	json.Unmarshal([]byte(response.Body.Bytes()), &result)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, result.Title, book.Title)
}

func TestGetBook(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/books", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}
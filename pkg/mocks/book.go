package mocks

import (
	"github.com/AdiPP/go/library-api/pkg/models"
)

var Book = models.Book {
	Id: 1,
	Title: "Golang",
	Author: "Gopher",
	Desc: "A book for Go",
}

var Books = []models.Book{
	{
		Id: 1,
		Title: "Golang",
		Author: "Gopher",
		Desc: "A book for Go",
	},
}
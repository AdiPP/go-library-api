package models

type Book struct {
	Id int64 `db:"id"`
	Title string `db:"title"`
	Author string `db:"author"`
	Desc string `db:"desc"`
}

type JsonBook struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}

type BookRequest struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}
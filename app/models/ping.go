package models

type Ping struct {
	Status string
}

type JsonPing struct {
	Status string `json:"status"`
}
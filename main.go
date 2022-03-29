package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AdiPP/go/library-api/app"
	"github.com/AdiPP/go/library-api/app/database"
)

func main () {
	app := app.New()
	app.DB = &database.DB{}
	databaseErr := app.DB.Open()

	check(databaseErr)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("API is running!")

	serveErr := http.ListenAndServe(":8000", nil)

	check(serveErr)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
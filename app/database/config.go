package database

import "fmt"

var (
	dbUsername = "pg"
	dbPassword = "pass"
	dbHost = "localhost"
	dbTable = "library"
	dbPort = "5432"
	pgConnStr = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword,
	)
)
package database

const createSchema = `
	CREATE TABLE IF NOT EXISTS books
	(
		id SERIAL PRIMARY KEY,
		title VARCHAR,
		author VARCHAR,
		"desc" VARCHAR
	)
`

var insertBookSchema = `
	INSERT INTO books(title, author, "desc") values($1, $2, $3) RETURNING id
`
package database

import "github.com/AdiPP/go/library-api/app/models"

func (d *DB) CreateBook(b *models.Book) error {
	res, err := d.db.Exec(insertBookSchema, b.Title, b.Author, b.Desc)

	if err != nil {
		return err
	}

	res.LastInsertId()

	return err
}

func (d *DB) GetBooks() ([]*models.Book, error) {
	var books []*models.Book

	err := d.db.Select(&books, "SELECT * FROM books")

	if err != nil {
		return books, err
	}

	return books, err
}
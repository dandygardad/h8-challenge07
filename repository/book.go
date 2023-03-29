package repository

import (
	"challenge07/model/entity"
	"errors"
)

type BookRepository interface {
	Create(book entity.Book) error
	GetAll() ([]entity.Book, error)
	GetOne(id int) (entity.Book, error)
	UpdateOne(id int, book entity.Book) (int, error)
	DeleteOne(id int) (int, error)
}

func (r Repo) Create(book entity.Book) error {
	// query sql insert data
	command := `INSERT INTO books (title, author, "desc") VALUES($1, $2, $3) RETURNING *`
	result, err := r.db.Exec(command, book.Title, book.Author, book.Desc)
	if err != nil {
		return err
	}
	count, errCreated := result.RowsAffected()
	if errCreated != nil {
		return errCreated
	}

	if count < 1 {
		return errors.New("no data created")
	}
	return nil
}
func (r Repo) GetAll() ([]entity.Book, error) {
	command := `SELECT id, title, author, "desc" FROM books`
	rows, err := r.db.Query(command)
	if err != nil {
		return []entity.Book{}, err
	}
	defer rows.Close()

	// get all books
	var allBooks []entity.Book
	for rows.Next() {
		tempBook := entity.Book{}
		err := rows.Scan(&tempBook.ID, &tempBook.Title, &tempBook.Author, &tempBook.Desc)
		if err != nil {
			return []entity.Book{}, err
		}
		allBooks = append(allBooks, tempBook)
	}
	return allBooks, nil
}

func (r Repo) GetOne(id int) (entity.Book, error) {
	command := `SELECT id, title, author, "desc" FROM books WHERE id=$1`
	result := r.db.QueryRow(command, id)

	var book entity.Book
	err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (r Repo) UpdateOne(id int, book entity.Book) (int, error) {
	command := `UPDATE books SET title=$1, author=$2, "desc"=$3 WHERE id=$4`
	result, err := r.db.Exec(command, book.Title, book.Author, book.Desc, id)
	if err != nil {
		return 0, err
	}

	count, errUpdated := result.RowsAffected()
	if errUpdated != nil {
		return 0, errUpdated
	}

	return int(count), nil
}

func (r Repo) DeleteOne(id int) (int, error) {
	command := `DELETE FROM books WHERE id=$1`
	result, err := r.db.Exec(command, id)
	if err != nil {
		return 0, err
	}

	count, errDeleted := result.RowsAffected()
	if errDeleted != nil {
		return 0, errDeleted
	}

	return int(count), nil
}

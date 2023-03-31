package repository

import (
	"challenge07/model/entity"
	"time"
)

type BookRepository interface {
	Create(book entity.Book) (entity.Book, error)
	GetAll() ([]entity.Book, error)
	GetOne(id int) (entity.Book, error)
	UpdateOne(id int, book entity.Book) (entity.Book, error)
	DeleteOne(id int) (int, error)
}

func (r Repo) Create(book entity.Book) (entity.Book, error) {
	// query sql insert data
	var resultBook entity.Book
	command := `INSERT INTO books (name_book, author) VALUES($1, $2) RETURNING *`
	err := r.db.QueryRow(command, book.NameBook, book.Author).Scan(&resultBook.ID, &resultBook.NameBook, &resultBook.Author, &resultBook.CreatedAt, &resultBook.UpdatedAt)
	if err != nil {
		return entity.Book{}, err
	}

	return resultBook, nil
}

func (r Repo) GetAll() ([]entity.Book, error) {
	command := `SELECT id, name_book, author, created_at, updated_at FROM books`
	rows, err := r.db.Query(command)
	if err != nil {
		return []entity.Book{}, err
	}
	defer rows.Close()

	// get all books
	var allBooks []entity.Book
	for rows.Next() {
		tempBook := entity.Book{}
		err := rows.Scan(&tempBook.ID, &tempBook.NameBook, &tempBook.Author, &tempBook.CreatedAt, &tempBook.UpdatedAt)
		if err != nil {
			return []entity.Book{}, err
		}
		allBooks = append(allBooks, tempBook)
	}
	return allBooks, nil
}

func (r Repo) GetOne(id int) (entity.Book, error) {
	command := `SELECT id, name_book, author, created_at, updated_at FROM books WHERE id=$1`
	result := r.db.QueryRow(command, id)

	var book entity.Book
	err := result.Scan(&book.ID, &book.NameBook, &book.Author, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return entity.Book{}, err
	}
	return book, nil
}

func (r Repo) UpdateOne(id int, book entity.Book) (entity.Book, error) {
	var updatedBook entity.Book
	currentTime := time.Now()
	command := `UPDATE books SET name_book=$1, author=$2, updated_at=$3 WHERE id=$4 RETURNING *`
	err := r.db.QueryRow(command, book.NameBook, book.Author, currentTime, id).Scan(&updatedBook.ID, &updatedBook.NameBook, &updatedBook.Author, &updatedBook.CreatedAt, &updatedBook.UpdatedAt)
	if err != nil {
		return entity.Book{}, err
	}

	return updatedBook, nil
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

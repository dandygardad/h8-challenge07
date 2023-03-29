package service

import (
	"challenge07/model/entity"
	"errors"
)

type BookService interface {
	CreateBook(book entity.Book) error
	GetAllBooks() ([]entity.Book, error)
	GetBook(id int) (entity.Book, error)
	UpdateBook(id int, book entity.Book) error
	DeleteBook(id int) error
}

func (s Service) CreateBook(book entity.Book) error {
	err := s.repo.Create(book)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetAllBooks() ([]entity.Book, error) {
	all, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// jika tidak ada data buku, maka kembalikan error tidak ada buku
	if len(all) == 0 {
		return nil, errors.New("no books")
	}

	return all, nil
}

func (s Service) GetBook(id int) (entity.Book, error) {
	result, err := s.repo.GetOne(id)
	if err != nil {
		return entity.Book{}, err
	}
	return result, nil
}

func (s Service) UpdateBook(id int, book entity.Book) error {
	count, err := s.repo.UpdateOne(id, book)
	if err != nil {
		return err
	}

	// Jika tidak ada data maka kembalikan error
	if count < 1 {
		return errors.New("no data updated")
	}
	return nil
}

func (s Service) DeleteBook(id int) error {
	count, err := s.repo.DeleteOne(id)
	if err != nil {
		return err
	}

	// Jika tidak ada data maka kembalikan error
	if count < 1 {
		return errors.New("no data deleted")
	}
	return nil
}

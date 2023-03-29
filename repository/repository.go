// Digunakan untuk memasukkan repository ke main

package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

type RepositoryInterface interface {
	BookRepository
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

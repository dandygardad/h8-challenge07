package entity

import "time"

type Book struct {
	ID        int       `json:"id"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

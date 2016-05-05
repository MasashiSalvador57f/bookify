package model

import "time"

// Book is a struct for book
type Book struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	UserID    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

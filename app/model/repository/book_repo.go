package repository

import (
	"time"

	"github.com/MasashiSalvador57f/bookify/app/model"
	"github.com/MasashiSalvador57f/bookify/vendor/gopkg.in/gorp.v1"
)

// BookRepository is an interface for bookRepo
type BookRepository interface {
	FindByID(int64) (*model.Book, error)
	FindByUserID(int64) ([]*model.Book, error)
}

// BookRepo is repository accessor for book Book
type BookRepo struct {
	BookRepository
	dbmap *gorp.DbMap
	now   time.Time
}

// NewBookRepo Creates new Book repository instance
func NewBookRepo(dbm *gorp.DbMap, now time.Time) BookRepo {
	return BookRepo{
		dbmap: dbm,
		now:   now,
	}
}

// FindByID Find Record by its unique id
func (br *BookRepo) FindByID(id int64) (*model.Book, error) {
	b := new(model.Book)
	err := br.dbmap.SelectOne(b, "SELECT * FROM book WHERE id = ?", id)
	return b, err
}

// FindByUserID find book record by its creator's id
func (br *BookRepo) FindByUserID(uID int64) ([]model.Book, error) {
	var books []model.Book
	_, err := br.dbmap.Select(&books, "SELECT * FROM book WHERE user_id = ?", uID)
	return books, err
}

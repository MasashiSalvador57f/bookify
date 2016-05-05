package bookservice

import (
	"time"

	"github.com/MasashiSalvador57f/bookify/app/model"
	"github.com/MasashiSalvador57f/bookify/app/model/repository"
	"github.com/MasashiSalvador57f/bookify/vendor/github.com/gin-gonic/gin"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/gorp.v1"
)

const okFindBook int = 0
const errorFindError int = 1

type BookFindService struct {
	ctx *gin.Context
}

// NewBookFindService creates book find service by context (gin)
func NewBookFindService(ctx *gin.Context) BookFindService {
	return BookFindService{ctx: ctx}
}

func (s *BookFindService) FindBooksByUserID(uID int64) ([]model.Book, int) {
	ctx := s.ctx
	now := ctx.MustGet("now").(time.Time)
	dbm := ctx.MustGet("dbmap").(*gorp.DbMap)

	br := repository.NewBookRepo(dbm, now)
	books, err := br.FindByUserID(uID)
	if err != nil {
		spew.Dump(err)
		return []model.Book{}, errorFindError
	}

	return books, okFindBook
}

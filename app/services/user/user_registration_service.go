package userservice

import (
	"time"

	"fmt"

	"github.com/MasashiSalvador57f/bookify/app/model"
	"github.com/MasashiSalvador57f/bookify/app/model/repository"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"log"
)

const (
	Complete = 0 + iota
	CantBeginTransactionOnInitialize
	FailToCreateUserByFacebookIDAndEmail
)

type UserRegistrationService struct {
	ctx *gin.Context
}

func NewUserService(ctx *gin.Context) UserRegistrationService {
	return UserRegistrationService{ctx: ctx}
}

func (urs *UserRegistrationService) InitializeUserByFbIDAndEmail(facebookID string, email string) (*model.User, int) {
	ctx := urs.ctx
	now := ctx.MustGet("now").(time.Time)
	dbm := ctx.MustGet("dbmap").(*gorp.DbMap)

	tx, err := dbm.Begin()
	if err != nil {
		fmt.Println(err)
		return &model.User{}, CantBeginTransactionOnInitialize
	}

	uRepo := repository.NewUserRepo(dbm, now)

	u, err := uRepo.CreateByFacebookIDAndEmail(facebookID, email)
	if err != nil {
		log.Fatalln(err)
		tx.Rollback()
		return &model.User{}, FailToCreateUserByFacebookIDAndEmail
	}

	tx.Commit()
	return u, Complete
}

func (urs *UserRegistrationService) FindByFacebookID(facebookID string) (*model.User, int) {
	ctx := urs.ctx
	now := ctx.MustGet("now").(time.Time)
	dbm := ctx.MustGet("dbmap").(*gorp.DbMap)

	uRepo := repository.NewUserRepo(dbm, now)

	u, _ := uRepo.FindByFacebookID(facebookID)
	return u, Complete
}

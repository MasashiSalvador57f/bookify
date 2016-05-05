package userservice

import (
	"time"

	"github.com/MasashiSalvador57f/bookify/app/model/repository"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
)

type UserLoginService struct {
	ctx *gin.Context
}

var findUserByAccessToken int = 0
var notFoundUserByAccessToken int = 1

func NewUserLoginService(c *gin.Context) UserLoginService {
	return UserLoginService{ctx: c}
}
func (s *UserLoginService) setUserInContext(accessToken string) int {
	now := s.ctx.MustGet("now").(time.Time)
	dbm := s.ctx.MustGet("dbmap").(*gorp.DbMap)
	ur := repository.NewUserRepo(dbm, now)

	u, err := ur.FindByAccessToken(accessToken)

	if err != nil {
		return notFoundUserByAccessToken
	}
	s.ctx.Set("user", u)

	return findUserByAccessToken
}

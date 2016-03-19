package userservice
import (
	"github.com/gin-gonic/gin"
	"github.com/MasashiSalvador57f/bookify/app/model/repository"
)

type LoginUserService struct {
	ctx *gin.Context
}

func (s *LoginUserService) setUserInContext(accessToken string) int {
	dbm := ctx.M
	ur := repository.NewUserRepo()
}
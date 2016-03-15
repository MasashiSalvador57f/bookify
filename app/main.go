package main

import (
	"time"

	"github.com/MasashiSalvador57f/bookify/app/controller/auth"
	"github.com/MasashiSalvador57f/bookify/app/controller/hoge"
	"github.com/MasashiSalvador57f/bookify/app/controller/user"
	"github.com/MasashiSalvador57f/bookify/app/shared/database_accessor"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(da.InitDB)
	r.Use(initializeTime)

	authCtr := authcontroller.AuthController{}
	r.GET("/auth/index", authCtr.Index)
	r.GET("/auth/callback", authCtr.AuthCallback)
	r.GET("/auth/show", authCtr.Show)

	userCtr := usercontroller.UserController{}
	r.GET("/", userCtr.Index)

	hogeCtr := hogecontroller.HogeController{}
	r.GET("/hoge", hogeCtr.Hoge)
	r.Run(":8080")
}

func initializeTime(c *gin.Context) {
	now := time.Now()
	c.Set("now", now)
}

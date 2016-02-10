package main

import (
	"github.com/MasashiSalvador57f/bookify/app/controller"
	"github.com/MasashiSalvador57f/bookify/app/shared/database_accessor"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(da.InitDB)

	authCtr := controller.AuthController{}
	r.GET("/auth/index", authCtr.Index)
	r.GET("/auth/callback", authCtr.AuthCallback)
	r.GET("/auth/show", authCtr.Show)

	userCtr := controller.UserController{}
	r.GET("/", userCtr.Index)
	r.Run(":8080")
}

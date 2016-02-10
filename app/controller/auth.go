package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MasashiSalvador57f/bookify/app/shared/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type FbUser struct {
	MailAddress string `json:"email"`
	ID          string `json:"id"`
}

// AuthController is controller for authentication
type AuthController struct {
}

const fbGraphMeURL = "https://graph.facebook.com/me?fields=email,id"

func (ctr *AuthController) Index(c *gin.Context) {
	fbConfig := auth.Facebook
	url := fbConfig.AuthCodeURL("")
	fmt.Println(url)
	c.Redirect(302, url)
	return
}

func (ctr *AuthController) AuthCallback(c *gin.Context) {
	code := c.Query("code")

	fbConfig := auth.Facebook
	token, err := fbConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		log.Fatal(err)
		c.Redirect(500, "/")
		return
	}

	client := fbConfig.Client(oauth2.NoContext, token)
	result, err := client.Get(fbGraphMeURL)

	if err != nil {
		log.Fatal("access token is not valid")
		c.Redirect(302, "/")
		return
	}
	defer result.Body.Close()

	fbuser := FbUser{}
	err = json.NewDecoder(result.Body).Decode(&fbuser)

	if err != nil {
		c.Redirect(301, "/")
		return
	}

	redirectURL := fmt.Sprintf("/auth/show?email=%s", fbuser.MailAddress)
	c.Redirect(301, redirectURL)
	return
}

func (ctr *AuthController) Show(c *gin.Context) {
	email := c.Query("email")
	c.JSON(200, gin.H{
		"status": "posted",
		"email":  email,
	})
}

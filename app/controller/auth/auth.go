package authcontroller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MasashiSalvador57f/bookify/app/services/user"
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

// TODO: Access Control
func (ctr *AuthController) Index(c *gin.Context) {
	fbConfig := auth.Facebook
	url := fbConfig.AuthCodeURL("")
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

	us := userservice.NewUserService(c)

	u, errint := us.FindByFacebookID(fbuser.ID)
	if errint != 0 {
		redirectURL := "/"
		c.Redirect(301, redirectURL)
		return
	}
	if !u.IsEmptyUser() {
		redirectURL := fmt.Sprintf("/auth/show?email=%s&ddd=1", u.Email)
		c.Redirect(301, redirectURL)
		return
	}

	u, errint = us.InitializeUserByFbIDAndEmail(
		fbuser.ID,
		fbuser.MailAddress,
	)

	if errint != userservice.Complete {
		c.Redirect(301, "/")
		return
	}
	redirectURL := fmt.Sprintf("/auth/show?email=%s", u.Email)
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

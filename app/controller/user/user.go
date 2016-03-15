package usercontroller

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (ctr *UserController) Index(c *gin.Context) {
	tpl, err := pongo2.FromFile("templates/user/index.html")
	if err != nil {
		c.String(500, "can't load template")
	}
	err = tpl.ExecuteWriter(
		pongo2.Context{"hoge": "hogehoge"},
		c.Writer,
	)
	if err != nil {
		c.String(500, "can't write to template")
	}
}

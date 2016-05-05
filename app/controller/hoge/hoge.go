package hogecontroller

import (
	"github.com/MasashiSalvador57f/bookify/app/services/book"
	"github.com/MasashiSalvador57f/bookify/vendor/github.com/gin-gonic/gin"
	"github.com/davecgh/go-spew/spew"
)

type HogeController struct {
}

// Hoge is ...
func (h *HogeController) Hoge(c *gin.Context) {
	bs := bookservice.NewBookFindService(c)
	books, errint := bs.FindBooksByUserID(3)
	spew.Dump(errint)
	spew.Dump(books)
}

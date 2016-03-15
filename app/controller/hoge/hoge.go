package hogecontroller

import (
	"fmt"

	"github.com/MasashiSalvador57f/bookify/app/shared/crypto"
	"github.com/gin-gonic/gin"
)

type HogeController struct {
}

// Hoge is ...
func (h *HogeController) Hoge(c *gin.Context) {
	k := accesstoken.ConvertToAccessToken("hoge")
	fmt.Println(k)
}

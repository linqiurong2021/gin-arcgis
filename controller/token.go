package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/services"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// GetToken 获取Token
func GetToken(c *gin.Context) {
	token := &services.Token{
		Name: "Test",
		ID:   1,
	}
	tokenStr, err := services.CreateToken(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
	}
	c.JSON(http.StatusOK, utils.Success("get success", tokenStr))
}

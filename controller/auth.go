package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/logic"
	"github.com/linqiurong2021/gin-book-frontend/utils"
)

// Auth 新增授权
func Auth(c *gin.Context) {
	ok, err := logic.Auth(c)
	if !ok && err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
}

// UnAuth 取消授权
func UnAuth(c *gin.Context) {
	ok, err := logic.UnAuth(c)
	if !ok && err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
}

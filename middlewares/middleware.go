package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/services"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"gorm.io/gorm"
)

// DomainCheck  校验
func DomainCheck() gin.HandlerFunc {
	// 判断是否存在用户
	return func(c *gin.Context) {
		// 1 获取参数 判断是否有权限
		ClientIP := c.ClientIP()
		// 通过IP查询
		_, err := services.GetDomainByName(ClientIP)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, utils.Unauthorized("unauthorized", ""))
				c.Abort()
				return
			}
			// 非空错误
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			c.Abort()
			return
		}
		c.Next()
	}
}

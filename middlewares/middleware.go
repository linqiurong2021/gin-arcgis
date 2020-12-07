package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/redis"
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

// TokenCheck token 校验
func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		// 如果开启RedisToken校验
		if config.Conf.RedisCheck {
			// 校验token
			redisToken, ok, err := redis.Get(c, "token")
			fmt.Printf("\ntoken:%t,%s\n", ok, redisToken)
			if err != nil {
				c.JSON(http.StatusInternalServerError, utils.ServerError(err.Error(), ""))
				c.Abort()
				return
			} else if !ok {
				c.JSON(http.StatusInternalServerError, utils.ServerError("get token failure", ""))
				c.Abort()
				return
			} else if token != redisToken {
				c.JSON(http.StatusBadRequest, utils.BadRequest("token invalidate", ""))
				c.Abort()
				return
			}
		}

		//
		if token == "" {
			c.JSON(http.StatusBadRequest, utils.BadRequest("token must", ""))
			c.Abort()
			return
		}
		jwtToken, err := services.Parse(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			c.Abort()
			return
		}
		myCliams, ok := services.Check(jwtToken)
		if ok && jwtToken.Valid {
			// 存储当前用户信息
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, myCliams)
			c.Abort()
			return
		}
	}
}

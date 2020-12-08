package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/config"
)

// Claims 自定义
type Claims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// Token 模型
type Token struct {
	Name string
	ID   uint
}

// var ctx context.Context

// CreateToken 创建
func CreateToken(c *gin.Context, token *Token) (string, error) {
	//
	mySigningKey := []byte(config.Conf.JWTSignKey)
	//
	now := time.Now().Unix()
	delayTime := config.Conf.TokenExpireMinutes * 60
	expiresAt := now + delayTime

	claims := Claims{
		token.ID,
		token.Name,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "gin-arcgis",
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singString, err := newToken.SignedString(mySigningKey)

	return singString, err
}

// Parse 解析Token
func Parse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JWTSignKey), nil
	})
	return token, err
}

// Check Token校验
func Check(jwtToken *jwt.Token) (*Claims, bool) {
	claims, ok := jwtToken.Claims.(*Claims)
	return claims, ok
}

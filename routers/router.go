package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/controller"
	"github.com/linqiurong2021/gin-arcgis/middlewares"
)

// InitRouters 初始化
func InitRouters(router *gin.Engine) {
	//
	v1Group := router.Group("/v1")

	arcgisGroup := v1Group.Group("/arcgis")
	// 需要授权后才可以使用 如果未授权则不能使用
	arcgisGroup.Use(middlewares.TokenCheck(), middlewares.DomainCheck())
	// 参数校验
	arcgisGroup.GET("", controller.QueryFeatures)
	arcgisGroup.GET("/all", controller.QueryAllLayerFeatures)
	arcgisGroup.PUT("", controller.UpdateFeatures)
	arcgisGroup.POST("", controller.AddFeatures)
	arcgisGroup.DELETE("", controller.DeleteFeatures)
	// 
	authDomainGroup := v1Group.Group("/auth")
	authDomainGroup.Use(middlewares.TokenCheck(), middlewares.DomainCheck())
	authDomainGroup.POST("", controller.Auth)
	authDomainGroup.DELETE("", controller.UnAuth)
	// Token
	tokenGroup := v1Group.Group("/token")
	tokenGroup.GET("", controller.GetToken)
}

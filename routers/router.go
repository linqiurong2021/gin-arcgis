package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/arcgis"
)

// InitRouters 初始化
func InitRouters(router *gin.Engine) {
	//
	v1Group := router.Group("/v1")
	arcgisGroup := v1Group.Group("/arcgis")
	// 参数校验
	arcgisGroup.GET("", arcgis.QueryFeatures)
	arcgisGroup.GET("/all", arcgis.QueryAllLayerFeatures)
	arcgisGroup.PUT("", arcgis.UpdateFeatures)
	arcgisGroup.POST("", arcgis.AddFeatures)
	arcgisGroup.DELETE("", arcgis.DeleteFeatures)
}

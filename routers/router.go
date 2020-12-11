package routers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/client"
	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/controller"
	"github.com/linqiurong2021/gin-arcgis/middlewares"
)

// Router Router
type Router struct {
}

// Service 注册服务结构体
type Service struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Host  string `json:"host"`
	Port  uint   `json:"port"`
	URL   string `json:"url"`
	Alive bool   `json:"alive"`
}

// InitRouters 初始化
func (r *Router) InitRouters(router *gin.Engine) {
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

// RegisterToGateway 注册到网关
func (r *Router) RegisterToGateway() {
	// client.HTTPClient.Post() 是用接口请求的形式还是用包引用的形式
	// 如果有 Post就可以不需要引用model 否则需要引包
	registerServerURL := fmt.Sprintf("%s:%d/%s", config.Conf.RegisterServer.Host, config.Conf.RegisterServer.Port, config.Conf.RegisterServer.URL)
	serviceData := &Service{
		Name:  "Arcgis请求空间数据接口",
		Host:  config.Conf.Host,
		Port:  config.Conf.Port,
		URL:   "/v1/arcgis",
		Alive: true,
	}
	data, err := json.Marshal(serviceData)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}
	body := bytes.NewReader(data)
	//
	_, err = client.HTTPClient.Post(registerServerURL, "application/json", body)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}
	// fmt.Printf("%#v", resp)
	// 第二个服务
	serviceData.ID = serviceData.ID + 1
	serviceData.URL = "/v1/auth"
	data, err = json.Marshal(serviceData)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}
	body = bytes.NewReader(data)
	//
	_, err = client.HTTPClient.Post(registerServerURL, "application/json", body)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}

	// 第二个服务
	serviceData.ID = serviceData.ID + 1
	serviceData.URL = "/v1/token"
	data, err = json.Marshal(serviceData)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}
	body = bytes.NewReader(data)
	//
	_, err = client.HTTPClient.Post(registerServerURL, "application/json", body)
	if err != nil {
		fmt.Printf("json invalidate: %s\n", err.Error())
		return
	}
	// fmt.Printf("%#v", resp)

}

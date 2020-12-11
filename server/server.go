package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/routers"
)

// Server Server
type Server struct{}

// Start 启动服务
func (s *Server) Start() {
	engine := gin.Default()
	var router = new(routers.Router)
	// 注册到网关
	router.RegisterToGateway()

	router.InitRouters(engine)
	addr := fmt.Sprintf(":%d", config.Conf.Port)
	fmt.Printf("Server Start At: %s%s\n", config.Conf.Host, addr)
	engine.Run(addr)
}

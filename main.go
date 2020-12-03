package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/routers"
)

func main() {

	engine := gin.Default()
	routers.InitRouters(engine)
	engine.Run(":8080")
}

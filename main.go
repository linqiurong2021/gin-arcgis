package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/models"
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"github.com/linqiurong2021/gin-arcgis/routers"
)

func main() {
	// 加载配置文件(这里可以使用默认的配置文件)
	if len(os.Args) > 1 {
		if err := config.Init(os.Args[1]); err != nil {
			fmt.Printf("load config from file falure !, err:%v\n", err)
			return
		}
	} else {
		if err := config.Init("./config/config.ini"); err != nil {
			fmt.Printf("load config from file falure !, err:%v\n", err)
			return
		}
		fmt.Printf("\n\n#### load config from config/config.ini ! ####\n\n")
	}

	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	engine := gin.Default()
	routers.InitRouters(engine)
	// 模型绑定
	// AutoMigrate 模型绑定
	mysql.DB.AutoMigrate(&models.Domain{}, &models.URL{}, &models.DomainURL{})
	engine.Run(fmt.Sprintf(":%#v", config.Conf.Port))
}

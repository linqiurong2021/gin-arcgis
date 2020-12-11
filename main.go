package main

import (
	"fmt"

	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"github.com/linqiurong2021/gin-arcgis/server"
)

func main() {
	// 加载配置文件(这里可以使用默认的配置文件)
	if err := config.Init("./config/config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	fmt.Printf("%#v", config.Conf.RegisterServer)
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	server := new(server.Server)
	server.Start()
}

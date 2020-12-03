package test

import (
	"fmt"

	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/models"
	"github.com/linqiurong2021/gin-arcgis/mysql"
)

func init() {
	if err := config.Init("config.ini"); err != nil {
		fmt.Printf("load config from file falure !, err:%v\n", err)
		return
	}
	fmt.Printf("\n\n#### load config from ./config.ini ! ####\n\n")
	if err := mysql.InitMySQL(config.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	mysql.DB.AutoMigrate(&models.Domain{}, &models.Path{})

}

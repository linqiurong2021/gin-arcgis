package test

import (
	"fmt"

	"github.com/linqiurong2021/gin-arcgis/config"
	"github.com/linqiurong2021/gin-arcgis/models"
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"github.com/linqiurong2021/gin-arcgis/redis"
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
	// redis 初始化
	redis.InitRedisSession(config.Conf.RedisConfig)

	// AutoMigrate 模型绑定
	mysql.DB.AutoMigrate(&models.Domain{}, &models.URL{}, &models.DomainURL{})
	// mysql.DB.AutoMigrate(&models.Domain{}, &models.URL{}, &models.DomainURL{})
	// mysql.DB.AutoMigrate(&models.Person{}, &models.Address{}, &models.PersonAddress{})
}

package config

import (
	"gopkg.in/ini.v1"
)

// AppConfig App配置项
type AppConfig struct {
	Release            bool   `ini:"release"`
	Port               int    `ini:"port"`
	JWTSignKey         string `ini:"JWTSingKey"`
	TokenExpireMinutes int64  `ini:"tokenExpireMinutes"`
	*MySQLConfig       `ini:"mysql"`
	*RedisConfig       `ini:"redis"`
	*TokenCheckConfig  `ini:"token"`
}

// MySQLConfig 数据库配置项
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Charset  string `ini:"charset"`
}

// RedisConfig redis配置文件
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// TokenCheckConfig 校验
type TokenCheckConfig struct {
	RedisCheck bool `int:"redisCheck"`
}

// Conf 配置
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	return ini.MapTo(Conf, file)
}

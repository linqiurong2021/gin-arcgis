package test

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
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
}

// TestPathCreate 测试
func TestPathCreate(t *testing.T) {
	//

	var path = new(models.Path)
	path.URL = "http://dev.eginsosft.cn:6080/"
	outPath, err := models.CreatePath(path)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outPath.ID, 0)
}

func TestPathUpdate(t *testing.T) {
	//

	var path = new(models.Path)
	path.ID = 3
	path.URL = "http://dev.eginsosft.cn:6081/"
	outPath, err := models.UpdatePath(path)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outPath.ID, 0)
}

func TestPathGetByID(t *testing.T) {
	//

	outPath, err := models.GetPathByID(3)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outPath.ID, 0)
}

func TestPathDeleteByID(t *testing.T) {
	//
	ok, _ := models.DeletePathByID(3)

	if !ok {
		assert.Equal(t, 1, 0)
	} else {
		assert.Equal(t, 1, 1)
	}
}

package test

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/linqiurong2021/gin-arcgis/models"
)

// TestURLCreate 测试
func TestURLCreate(t *testing.T) {
	//

	var URL = new(models.URL)
	URL.URL = "http://dev.eginsosft.cn:6080/"
	outURL, err := models.CreateURL(URL)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outURL.ID, 0)

	URL = new(models.URL)
	URL.URL = "http://dev.eginsosft.cn:6081/"
	outURL, err = models.CreateURL(URL)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outURL.ID, 0)
}

func TestURLUpdate(t *testing.T) {
	//
	URL, err := models.GetURLByID(1)
	if err != nil {
		assert.Equal(t, 1, 0)
	}
	URL.URL = "http://dev.eginsosft.cn:6081/"
	outURL, err := models.SaveURL(URL)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outURL.ID, 0)
}

func TestURLGetByID(t *testing.T) {
	//

	outURL, err := models.GetURLByID(1)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outURL.ID, 0)
}

func TestURLDeleteByID(t *testing.T) {
	//
	ok, _ := models.DeleteURLByID(1)

	if !ok {
		assert.Equal(t, 1, 0)
	} else {
		assert.Equal(t, 1, 1)
	}
}

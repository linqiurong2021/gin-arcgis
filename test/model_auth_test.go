package test

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/linqiurong2021/gin-arcgis/models"
)

func TestAuthPath(t *testing.T) {

	path := models.Path{
		URL: "http://dev.eginsoft.cn:6081",
		Domains: []models.Domain{
			{Name: "127.0.0.1", Note: "本地"},
		},
	}
	//
	ok, err := models.AuthPath(&path)
	if err != nil {
		fmt.Printf("\nError:\n%s\n", err.Error())
		return
	}
	if !ok {
		assert.Equal(t, 1, 0)
	} else {
		assert.Equal(t, 1, 1)
	}

}

func TestAuthDomain(t *testing.T) {

	domain := models.Domain{
		Name: "127.0.0.1",
		Note: "本地",
		Paths: []*models.Path{
			{URL: "http://dev.eginsoft.cn:6080"},
		},
	}
	//
	ok, err := models.AuthDomain(&domain)
	if err != nil {
		fmt.Printf("\nError:\n%s\n", err.Error())
		return
	}
	if !ok {
		assert.Equal(t, 1, 0)
	} else {
		assert.Equal(t, 1, 1)
	}

}

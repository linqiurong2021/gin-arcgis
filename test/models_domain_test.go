package test

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/linqiurong2021/gin-arcgis/models"
)

// TestDomainCreate 测试
func TestDomainCreate(t *testing.T) {
	//

	var domain = new(models.Domain)
	domain.Name = "localhost2"
	domain.Note = "本地"
	outDomain, err := models.CreateDomain(domain)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outDomain.ID, 0)
}

func TestDomainUpdate(t *testing.T) {
	//
	domain, err := models.GetDomainByID(1)
	if err != nil {
		assert.Equal(t, 1, 0)
	}
	domain.Note = "本地222"
	domain.UserID = 1
	outDomain, err := models.SaveDomain(domain)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outDomain.ID, 0)
}

func TestDomainGetByID(t *testing.T) {
	//
	outDomain, err := models.GetDomainByID(1)
	if err != nil {
		fmt.Printf("Error:%s\n", err.Error())
	}
	assert.NotEqual(t, outDomain.ID, 0)
}

func TestDomainDeleteByID(t *testing.T) {
	//
	ok, _ := models.DeleteDomainByID(1)
	if !ok {
		assert.Equal(t, 1, 0)
	} else {
		assert.Equal(t, 1, 1)
	}
}

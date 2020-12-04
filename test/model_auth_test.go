package test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/linqiurong2021/gin-arcgis/models"
	"github.com/linqiurong2021/gin-arcgis/services"
)

func TestAuth(t *testing.T) {
	domainURL := []*models.DomainURL{
		{URLID: 2, DomainID: 2},
		{URLID: 1, DomainID: 2},
	}
	// domainUrl = []*models.DomainURL
	result, err := services.Auth(domainURL)
	if err != nil {
		assert.Equal(t, 1, 0)
	}
	if len(result) > 0 {
		assert.Equal(t, 1, 1)
	} else {

		assert.Equal(t, 1, 0)
	}
}

func TestUnAuthByURLID(t *testing.T) {
	ok, _ := services.UnAuthByDURLID(2)
	if !ok {
		assert.Equal(t, 1, 0)
	}
	assert.Equal(t, 1, 1)
}

func TestUnAuthByDomainID(t *testing.T) {
	ok, _ := services.UnAuthByDomainID(1)
	if !ok {
		assert.Equal(t, 1, 0)
	}
	assert.Equal(t, 1, 1)
}

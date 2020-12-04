package services

import "github.com/linqiurong2021/gin-arcgis/models"

// CreateDomain 创建Domain
func CreateDomain(inDomain *models.Domain) (outDomain *models.Domain, err error) {

	return models.CreateDomain(inDomain)
}

// UpdateDomain 更新数据
func UpdateDomain(domain *models.Domain) (outDomain *models.Domain, err error) {

	return models.UpdateDomain(domain)
}

// GetDomainByID 获取购物车
func GetDomainByID(DomainID uint) (outDomain *models.Domain, err error) {
	return models.GetDomainByID(DomainID)
}

// GetDomainByName 获取购物车
func GetDomainByName(name string) (outDomain *models.Domain, err error) {
	return models.GetDomainByName(name)
}

package services

import "github.com/linqiurong2021/gin-arcgis/models"

// Auth 创建请求路径
func Auth(inDomainURLs []*models.DomainURL) (outDomainURL []*models.DomainURL, err error) {
	return models.CreateDomainURL(inDomainURLs)
}

// UnAuthByDomainID 取消授权
func UnAuthByDomainID(domainID uint) (bool, error) {
	return models.DeleteDomainURLByDomainID(domainID)
}

// UnAuthByDURLID 取消授权
func UnAuthByDURLID(URLID uint) (bool, error) {
	return models.DeleteDomainURLByURLID(URLID)
}

// GetDomainURLByDomainIDAndURLD 通过URLID与DomainID获取权限
func GetDomainURLByDomainIDAndURLD(domainID uint, URLID uint) (*models.DomainURL, error) {
	return models.GetDomainURLByDomainIDAndURLD(domainID, URLID)
}

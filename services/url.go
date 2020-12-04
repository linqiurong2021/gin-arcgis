package services

import "github.com/linqiurong2021/gin-arcgis/models"

// CreateURL 创建URL
func CreateURL(inURL *models.URL) (outURL *models.URL, err error) {

	return models.CreateURL(inURL)
}

// UpdateURL 更新数据
func UpdateURL(URL *models.URL) (outURL *models.URL, err error) {

	return models.UpdateURL(URL)
}

// GetURLByID 获取购物车
func GetURLByID(URLID uint) (outURL *models.URL, err error) {
	return models.GetURLByID(URLID)
}

// GetURLByURL 获取购物车
func GetURLByURL(URL string) (outURL *models.URL, err error) {
	return models.GetURLByURL(URL)
}

// DeleteURLByIDs 通过ID删除URL
func DeleteURLByIDs(URLID uint) (ok bool, err error) {
	return models.DeleteURLByID(URLID)
}

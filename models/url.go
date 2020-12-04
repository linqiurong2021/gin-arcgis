package models

import (
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"gorm.io/gorm"
)

// URL 请求的URL
type URL struct {
	gorm.Model
	URL string `json:"url" gorm:"url;unique" binding:"required" label:"请求的链接"`
}

// CreateURL 创建请求URL
func CreateURL(inURL *URL) (outURL *URL, err error) {

	if err := mysql.DB.Create(&inURL).Error; err != nil {
		return nil, err
	}
	outURL = inURL
	return
}

// GetURLByID 通过ID获取请求URL信息
func GetURLByID(URLID uint) (outURL *URL, err error) {
	var URL = new(URL)
	if err := mysql.DB.Where("id = ?", URLID).First(&URL).Error; err != nil {
		return nil, err
	}
	return URL, nil
}

// GetURLByURL 通过ID获取请求URL信息
func GetURLByURL(url string) (outURL *URL, err error) {
	var URL = new(URL)
	if err := mysql.DB.Where("url = ?", url).First(&URL).Error; err != nil {
		return nil, err
	}
	return URL, nil
}

// UpdateURL 更新数据
func UpdateURL(inURL *URL) (outURL *URL, err error) {
	if err := mysql.DB.Save(inURL).Error; err != nil {
		return nil, err
	}
	outURL = inURL
	return
}

// SaveURL 保存数据
func SaveURL(inURL *URL) (outURL *URL, err error) {
	if err := mysql.DB.Save(inURL).Error; err != nil {
		return nil, err
	}
	outURL = inURL
	return
}

// DeleteURLByID 通过ID删除多个
func DeleteURLByID(URLID uint) (ok bool, err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := mysql.DB.Where("id = ?", URLID).Delete(&URL{}).Error; err != nil {
			return err
		}
		//
		if ok, err := DeleteDomainURLByURLID(URLID); !ok {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	//
	return true, nil
}

// GetListURLByPage 获取列表 分页
func GetListURLByPage(page int, pageSize int) (outURLList []*URL, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&outURLList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&URL{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return outURLList, count, nil
}

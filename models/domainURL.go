package models

import (
	"time"

	"github.com/linqiurong2021/gin-arcgis/mysql"
	"gorm.io/gorm"
)

// DomainURL DomainURL与URL中间表
type DomainURL struct {
	ID        uint           `json:"ID" gorm:"id;primaryKey"`
	DomainID  uint           `json:"domain_id" gorm:"domain_id"`
	URLID     uint           `json:"url_id" gorm:"url_id"`
	CreatedAt time.Time      `json:"create_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}

// BeforeCreate 创建前
func (domainURL *DomainURL) BeforeCreate(db *gorm.DB) error {

	return nil
}

// CreateDomainURL 创建域名
func CreateDomainURL(inDomainURL []*DomainURL) (outDomainURL []*DomainURL, err error) {
	if err := mysql.DB.Create(&inDomainURL).Error; err != nil {
		return nil, err
	}
	outDomainURL = inDomainURL
	return
}

// GetDomainURLByDomainID 通过ID获取域名信息
func GetDomainURLByDomainID(domainID uint) (outDomainURL *DomainURL, err error) {
	var domainURL = new(DomainURL)
	if err := mysql.DB.Where("domain_id = ?", domainID).First(&domainURL).Error; err != nil {
		return nil, err
	}
	return domainURL, nil
}

// GetDomainURLByDomainIDAndURLD判断是否存在
func GetDomainURLByDomainIDAndURLD(domainID uint, URLID uint) (outDomainURL *DomainURL, err error) {
	var domainURL = new(DomainURL)
	if err := mysql.DB.Where("domain_id = ?", domainID).Where("url_id = ?", URLID).First(&domainURL).Error; err != nil {
		return nil, err
	}
	return domainURL, nil
}

// GetDomainURLByURLID 通过ID获取域名信息
func GetDomainURLByURLID(URLID uint) (outDomainURL *DomainURL, err error) {
	var DomainURL = new(DomainURL)
	if err := mysql.DB.Where("url_id = ?", URLID).First(&DomainURL).Error; err != nil {
		return nil, err
	}
	return DomainURL, nil
}

// DeleteDomainURLByDomainID 通过ID删除多个
func DeleteDomainURLByDomainID(domainID uint) (ok bool, err error) {
	//
	if err := mysql.DB.Where("domain_id = ?", domainID).Delete(&DomainURL{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// DeleteDomainURLByURLID 通过ID删除多个
func DeleteDomainURLByURLID(URLID uint) (ok bool, err error) {
	//
	if err := mysql.DB.Where("url_id = ?", URLID).Delete(&DomainURL{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

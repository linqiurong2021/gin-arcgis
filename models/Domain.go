package models

import (
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"gorm.io/gorm"
)

// 逻辑
// RequestURL 与 Domain 是多对多关系  一个请求URL可对应多个域名  一个域名也可请求多个URL
// WhiteURL 授权域名可访问的URL

// Domain 白名单
type Domain struct {
	gorm.Model
	Name  string  `json:"domain" gorm:"domain" bind:"required" label:"域名或ip"`
	Note  string  `json:"note" gorm:"note" bind:"required" label:"备注"`
	Paths []*Path `json:"paths" gorm:"many2many:domain_url;ForeignKey:id;References:id"`
}

// CreateDomain 创建域名
func CreateDomain(inDomain *Domain) (outDomain *Domain, err error) {
	if err := mysql.DB.Create(&inDomain).Error; err != nil {
		return nil, err
	}
	outDomain = inDomain
	return
}

// GetDomainByID 通过ID获取域名信息
func GetDomainByID(DomainID uint) (outDomain *Domain, err error) {
	var Domain = new(Domain)
	if err := mysql.DB.Where("id = ?", DomainID).First(&Domain).Error; err != nil {
		return nil, err
	}
	return Domain, nil
}

// UpdateDomain 更新数据
func UpdateDomain(inDomain *Domain) (outDomain *Domain, err error) {
	if err := mysql.DB.Where("id = ?", inDomain.ID).Save(inDomain).Error; err != nil {
		return nil, err
	}
	outDomain = inDomain
	return
}

// DeleteDomainByID 通过ID删除多个
func DeleteDomainByID(domainID uint) (ok bool, err error) {
	//
	if err := mysql.DB.Where("id = ?", domainID).Delete(&Domain{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetListDomainByPage 获取列表 分页
func GetListDomainByPage(page int, pageSize int) (outDomainList []*Domain, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&outDomainList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&Domain{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return outDomainList, count, nil
}

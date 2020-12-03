package models

import (
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"gorm.io/gorm"
)

// Path 请求的URL
type Path struct {
	gorm.Model
	URL     string   `json:"url" gorm:"url;unique" label:"请求的链接"`
	Domains []Domain `json:"domains" gorm:"many2many:domain_url;ForeignKey:id;References:id"`
}

// CreatePath 创建请求URL
func CreatePath(inPath *Path) (outPath *Path, err error) {

	if err := mysql.DB.Create(&inPath).Error; err != nil {
		return nil, err
	}
	outPath = inPath
	return
}

// GetPathByID 通过ID获取请求URL信息
func GetPathByID(PathID uint) (outPath *Path, err error) {
	var Path = new(Path)
	if err := mysql.DB.Where("id = ?", PathID).First(&Path).Error; err != nil {
		return nil, err
	}
	return Path, nil
}

// UpdatePath 更新数据
func UpdatePath(inPath *Path) (outPath *Path, err error) {
	if err := mysql.DB.Where("id = ?", inPath.ID).Save(inPath).Error; err != nil {
		return nil, err
	}
	outPath = inPath
	return
}

// SavePath 保存数据
func SavePath(inPath *Path) (outPath *Path, err error) {
	if err := mysql.DB.Save(inPath).Error; err != nil {
		return nil, err
	}
	outPath = inPath
	return
}

// DeletePathByID 通过ID删除多个
func DeletePathByID(pathID uint) (ok bool, err error) {
	//
	if err := mysql.DB.Where("id = ?", pathID).Delete(&Path{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// GetListPathByPage 获取列表 分页
func GetListPathByPage(page int, pageSize int) (outPathList []*Path, count int64, err error) {
	if err := mysql.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&outPathList).Error; err != nil {
		return nil, 0, err
	}
	if err := mysql.DB.Find(&Path{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return outPathList, count, nil
}

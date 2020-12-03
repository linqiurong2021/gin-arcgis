package models

import (
	"github.com/linqiurong2021/gin-arcgis/mysql"
	"gorm.io/gorm"
)

// AuthPath 授权
func AuthPath(inPath *Path) (ok bool, err error) {

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := mysql.DB.Create(&inPath).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//
		if err := mysql.DB.Save(&inPath).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		// 返回 nil 提交事务
		return false, nil
	}
	return true, nil
}

// AuthDomain 授权
func AuthDomain(inDomain *Domain) (ok bool, err error) {

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := mysql.DB.Create(&inDomain).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//
		if err := mysql.DB.Save(&inDomain).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		// 返回 nil 提交事务
		return false, nil
	}
	return true, nil
}

// UnAuthPath 取消授权
// func UnAuthPath(inPath *Path) (ok bool, err error) {
// 	//
// 	if err = mysql.DB.Association("Paths").Delete(inPath).Error; err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

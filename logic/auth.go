package logic

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/dao"
	"github.com/linqiurong2021/gin-arcgis/models"
	"github.com/linqiurong2021/gin-arcgis/services"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"
	"gorm.io/gorm"
)

// Auth 创建请求路径
func Auth(c *gin.Context) (ok bool, err error) {

	var authParams dao.AuthParams
	err = c.ShouldBindJSON(&authParams) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, err
	}
	// 判断申请的URL是否存在
	URL, err := existsApplyURL(authParams.URL)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, nil
	}
	if URL == nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(authParams.URL+" not exists ", ""))
		return false, nil
	}
	//
	domain, err := existsApplyDomain(authParams.Domain)
	// 域名不存在则需要新增 如果存在则直接用ID
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	// 不存在域名
	if domain == nil {
		domain, err = createDomain(&authParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
			return false, nil
		}
	}
	// 还需要判断是否已存在 如果已存在则不需要操作数据库
	hasDomainURL, err := services.GetDomainURLByDomainIDAndURLD(domain.ID, URL.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return false, nil
	}
	// 已存在
	if hasDomainURL != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest("has auth", ""))
		return true, nil
	} //
	domainURLs := []*models.DomainURL{
		{DomainID: domain.ID, URLID: URL.ID},
	}
	outDomainURLs, err := services.Auth(domainURLs)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	if len(outDomainURLs) > 0 {
		c.JSON(http.StatusOK, utils.Success("auth success", ""))
		return true, nil
	}
	return false, nil
}

// UnAuth 取消授权
func UnAuth(c *gin.Context) (ok bool, err error) {
	var authParams dao.AuthParams
	err = c.ShouldBindJSON(&authParams) // 绑定并校验
	// 参数校验判断
	ok = validator.Validate(c, err)
	if !ok {
		return false, nil
	}
	//
	domain, err := existsApplyDomain(authParams.Domain)
	// 域名不存在则需要新增 如果存在则直接用ID
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	//
	if domain == nil {
		//
		c.JSON(http.StatusBadRequest, utils.BadRequest(authParams.Domain+" not exists", ""))
		return false, nil
	}
	//
	ok, err = services.UnAuthByDomainID(domain.ID)
	if err != nil {
		return false, err
	}
	c.JSON(http.StatusOK, utils.Success(" unauth success", ""))
	return true, nil
}

// ExistsApplyURL 申请的URL是否存在
func existsApplyURL(url string) (URL *models.URL, err error) {
	URL, err = services.GetURLByURL(url)
	if err != nil {
		return nil, err
	}
	return URL, nil
}

// ExistsApplyDomain 是否存在Domain
func existsApplyDomain(name string) (domain *models.Domain, err error) {
	//
	return services.GetDomainByName(name)
}

// CreateDomain 创建域名
func createDomain(authParams *dao.AuthParams) (outDomain *models.Domain, err error) {
	//
	inDomain := &models.Domain{
		Name: authParams.Domain,
		Note: " ",
	}

	return services.CreateDomain(inDomain)
}

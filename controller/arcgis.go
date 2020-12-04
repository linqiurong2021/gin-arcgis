package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linqiurong2021/gin-arcgis/client"
	"github.com/linqiurong2021/gin-arcgis/dao"
	"github.com/linqiurong2021/gin-arcgis/response"
	"github.com/linqiurong2021/gin-book-frontend/utils"
	"github.com/linqiurong2021/gin-book-frontend/validator"
)

// AddFeatures 新增
func AddFeatures(c *gin.Context) {
	//
	var postData dao.PostParams
	err := c.BindJSON(&postData)
	// // 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return
	}
	// 请求返回结果
	result := client.AddFeature(postData.URL, postData.Features)
	var response response.AddResponseResult
	err = json.Unmarshal(result, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, utils.Success("do add success", response))
}

// UpdateFeatures 更新
func UpdateFeatures(c *gin.Context) {
	//
	var postData dao.PostParams
	err := c.BindJSON(&postData)
	// // 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return
	}
	// 请求返回结果
	result := client.UpdateFeature(postData.URL, postData.Features)
	var response response.UpdateResponseResult
	err = json.Unmarshal(result, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))

		return
	}
	c.JSON(http.StatusOK, utils.Success("do update success", response))
}

// DeleteFeatures 删除
func DeleteFeatures(c *gin.Context) {

	var postData dao.DeleteParams
	err := c.BindJSON(&postData)
	// // 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return
	}
	// 请求返回结果
	result := client.DeleteFeature(postData.URL, postData.ObjectIDs, postData.Where)
	var response response.DeleteResponseResult
	err = json.Unmarshal(result, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Success("do delete success", response))
}

// QueryFeatures 获取某个图层数据
func QueryFeatures(c *gin.Context) {

	var postData dao.QueryParams
	err := c.BindJSON(&postData)
	// // 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return
	}
	// 请求返回结果
	result := client.QueryFeature(postData.URL, postData.ObjectIDs, postData.Where, postData.OutFields, postData.ReturnGeometry, postData.OrderByFields)
	var response response.QueryResponseResult
	err = json.Unmarshal(result, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Success("do query success", response))
}

// QueryAllLayerFeatures 获取所有图层数据
func QueryAllLayerFeatures(c *gin.Context) {
	//
	var postData dao.QueryAllParams
	err := c.BindJSON(&postData)
	// // 参数校验判断
	ok := validator.Validate(c, err)
	if !ok {
		return
	}
	layerDefs, err := json.Marshal(postData.LayerDefs)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}
	// 请求返回结果
	result := client.QueryAllFeature(postData.URL, string(layerDefs), postData.ReturnGeometry, postData.GeometryPrecision)
	var response response.QueryAllResponseResult
	err = json.Unmarshal(result, &response)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequest(err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.Success("do query all success", response))

}

package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient 客户端
var HTTPClient *http.Client

// 初始化Client
func init() {
	HTTPClient = NewClient()
}

// NewClient 新建客户端
func NewClient() *http.Client {
	// 超时时间：5秒
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

// PostForm  表单提交
func PostForm(URL string, formData url.Values) []byte {
	resp, err := HTTPClient.PostForm(URL, formData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

// Get Get
func Get(URL string) []byte {
	resp, err := HTTPClient.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

// AddFeature POST请求
func AddFeature(URL string, data interface{}) []byte {

	// 超时时间：5秒
	jsonStr, _ := json.Marshal(data)
	// fmt.Printf("REQUEST BODY: , %s\n, %s\n", jsonStr, contentType)
	formData := url.Values{
		"f":        {"json"},
		"features": {string(jsonStr)},
	}
	return PostForm(URL, formData)
}

// UpdateFeature POST请求
func UpdateFeature(URL string, data interface{}) []byte {

	// 超时时间：5秒
	jsonStr, _ := json.Marshal(data)
	// fmt.Printf("REQUEST BODY: , %s\n, %s\n", jsonStr, contentType)
	formData := url.Values{
		"f":        {"json"},
		"features": {string(jsonStr)},
	}
	return PostForm(URL, formData)
}

// DeleteFeature POST请求
func DeleteFeature(URL string, ObjectIDs string, Where string) []byte {

	// 超时时间：5秒
	formData := url.Values{
		"f":         {"json"},
		"objectIds": {ObjectIDs},
		"where":     {Where},
	}
	return PostForm(URL, formData)
}

// QueryFeature POST请求
func QueryFeature(URL string, ObjectIDs string, Where string, outFields string, returnGeometry string, orderByFields string) []byte {

	// 超时时间：5秒
	formData := url.Values{
		"f":              {"json"},
		"objectIds":      {ObjectIDs},
		"where":          {Where},
		"outFields":      {outFields},
		"returnGeometry": {returnGeometry},
		"orderByFields":  {orderByFields},
	}
	return PostForm(URL, formData)
}

// QueryAllFeature POST请求
func QueryAllFeature(URL string, layerDefs string, returnGeometry string, geometryPrecision uint) []byte {

	// 超时时间：5秒
	formData := url.Values{
		"f":                 {"json"},
		"layerDefs":         {layerDefs},
		"returnGeometry":    {returnGeometry},
		"geometryPrecision": {string(rune(geometryPrecision))},
	}
	// url := URL + "?" + url.Values.Encode(formData)
	// fmt.Println(url, "url")
	// return Get(url)
	return PostForm(URL, formData)
}

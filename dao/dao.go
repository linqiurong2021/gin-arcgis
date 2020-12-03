package dao

// CommonURL URL
type CommonURL struct {
	URL string `json:"url" binding:"required" label:"请求链接"`
}

// CommonWhere Where
type CommonWhere struct {
	Where string `json:"where" label:"条件"`
}

// CommonOutFields OutFields
type CommonOutFields struct {
	OutFields string `json:"outFields" label:"输出字段"`
}

// PostParams 参数
type PostParams struct {
	CommonURL
	Features Features `json:"features" binding:"required" label:"空间数据"`
}

// DeleteParams 参数
type DeleteParams struct {
	CommonURL
	CommonWhere

	ObjectIDs string `json:"objectIds" binding:"required" label:"OBJECTID"`
}

// QueryParams 参数
type QueryParams struct {
	CommonURL
	CommonWhere
	CommonOutFields

	ObjectIDs      string `json:"objectIds" label:"OBJECTID"`
	ReturnGeometry string `json:"returnGeometry" label:"返回空间数据"`
	OrderByFields  string `json:"orderByFields" label:"排序"`
}

// QueryAllParams 搜索图层数据
type QueryAllParams struct {
	CommonURL
	LayerDefs         LayerDefs `json:"layerDefs"` // 例[{"layerId" : 0, "where" : "OBJECTID<100", "outFields" : "*"}, {"layerId" : 1, "where" : "OBJECTID<323", "outFields" : "OBJECTID,CREATOR"}]
	ReturnGeometry    string    `json:"returnGeometry" label:"返回空间数据"`
	GeometryPrecision uint      `json:"geometryPrecision" label:"空间数据精度"`
}

// LayerDefs 图层定义条件集
type LayerDefs []LayerDef

// LayerDef 图层定义条件
type LayerDef struct {
	LayerID uint `json:"layerId"`
	CommonWhere
	CommonOutFields
}

// Features 空间数组
type Features []Feature

// Feature 空间数据
type Feature struct {
	Geometry   map[string]interface{} `json:"geometry" bind:"required" label:"空间数据"`
	Atttibutes map[string]interface{} `json:"attributes" bind:"required" label:"属性"`
}

package response

// QueryResponseResult 搜索结果集
type QueryResponseResult struct {
	ObjectIDFieldName string           `json:"objectIdFieldName"`
	GlobalIDFieldName string           `json:"globalIdFieldName"`
	GeometryType      string           `json:"geometryType"`
	SpatialReference  SpatialReference `json:"spatialReference"`
	Fields            FieldItems       `json:"fields"`
	Features          FeatureItems     `json:"features"`
}

// QueryAllResponseResult 服务搜索结果集
type QueryAllResponseResult struct {
	Layers Layers `json:"layers"`
}

// Layers 图层集合
type Layers []Layer

// Layer 图层
type Layer struct {
	ID uint `json:"id"`
	QueryResponseResult
}

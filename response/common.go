package response

// FeatureResults 返回结果
type FeatureResults []FeatureResultItem

// FeatureResultItem 新增结果项
type FeatureResultItem struct {
	ObjectID int  `json:"objectId"`
	Success  bool `json:"success"`
}


// SpatialReference 坐标系统
type SpatialReference struct {
	Wkid uint `json:"wkid"`
}

// FieldItems 字段集
type FieldItems []FieldItem

// FieldItem 字段项
type FieldItem struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Alias  string `json:"alias"`
	Length uint   `json:"length"`
}

// FeatureItems 空间数据集
type FeatureItems []FeatureItem

// FeatureItem 空间数据项
type FeatureItem struct {
	Geometry   map[string]interface{} `json:"geometry" bind:"required" label:"空间数据"`
	Atttibutes map[string]interface{} `json:"attributes" bind:"required" label:"属性"`
}

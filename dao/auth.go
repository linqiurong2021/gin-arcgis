package dao

// DomainParams 域名
type DomainParams struct {
	Domain string `json:"domain" binding:"required" label:"授权域名"`
}

// AuthParams 授权请求参数
type AuthParams struct {
	DomainParams
	URL string `json:"url" binding:"required,url" label:"授权URL"`
}

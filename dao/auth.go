package dao

// DomainParams 域名
type DomainParams struct {
	Domain string `json:"domain" bind:"required"`
}

// AuthParams 授权请求参数
type AuthParams struct {
	DomainParams
	URL string `json:"url" bind:"required"`
}

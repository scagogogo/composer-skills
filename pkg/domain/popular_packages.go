// Package domain 提供了Composer API响应的数据结构定义
package domain

// PopularPackagesResponse 表示流行包列表响应的顶层结构
// 对应API: https://packagist.org/explore/popular.json
type PopularPackagesResponse struct {
	// 包信息列表
	Packages []PopularPackage `json:"packages"`
	// 总包数
	Total int `json:"total"`
	// 下一页的URL，如果没有下一页则为空
	Next string `json:"next,omitempty"`
}

// PopularPackage 表示流行包的基本信息
type PopularPackage struct {
	// 包名
	Name string `json:"name"`
	// 包描述
	Description string `json:"description"`
	// 包URL
	URL string `json:"url"`
	// 下载次数
	Downloads int `json:"downloads"`
	// 收藏次数
	Favers int `json:"favers"`
}

// Package domain 提供了Composer API响应的数据结构定义
package domain

// SearchResponse 表示搜索结果响应的顶层结构
// 对应API: https://packagist.org/search.json
type SearchResponse struct {
	// 搜索结果列表
	Results []SearchResult `json:"results"`
	// 总结果数
	Total int `json:"total"`
	// 下一页的URL，如果没有下一页则为空
	Next string `json:"next,omitempty"`
}

// SearchResult 表示单个搜索结果
type SearchResult struct {
	// 包名
	Name string `json:"name"`
	// 包描述
	Description string `json:"description"`
	// 包URL
	URL string `json:"url"`
	// 仓库URL
	Repository string `json:"repository"`
	// 下载次数
	Downloads int `json:"downloads"`
	// 收藏次数
	Favers int `json:"favers"`
}

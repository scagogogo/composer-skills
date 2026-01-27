// Package domain 提供了Composer API响应的数据结构定义
// 主要包含从Composer API返回的各类响应数据结构
package domain

// StatisticsResponse 表示统计数据响应的顶层结构
// 包含了Composer仓库的整体统计信息
// 示例: {"totals": {"downloads": 10000000000, "packages": 300000, "versions": 2500000}}
type StatisticsResponse struct {
	// 总计统计数据
	Totals Totals `json:"totals"`
}

// Totals 包含了Composer仓库的总体统计信息
// 包括下载次数、包数量和版本数量
// 示例: {"downloads": 10000000000, "packages": 300000, "versions": 2500000}
type Totals struct {
	// 所有包的总下载次数
	// 示例: 10000000000
	Downloads int64 `json:"downloads"`

	// 仓库中的包总数
	// 示例: 300000
	Packages int `json:"packages"`

	// 所有包的版本总数
	// 示例: 2500000
	Versions int `json:"versions"`
}

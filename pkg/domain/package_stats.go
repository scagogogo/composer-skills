// Package domain 提供了Composer API响应的数据结构定义
package domain

// PackageStatsResponse 表示包下载统计响应的顶层结构
// 对应API: https://packagist.org/packages/[vendor]/[package]/stats.json
type PackageStatsResponse struct {
	// 下载统计信息
	Downloads PackageDownloads `json:"downloads"`
	// 可用版本列表
	Versions []string `json:"versions"`
	// 统计开始日期
	Date string `json:"date"`
}

// ChangeTrackingResponse 表示变更跟踪响应
// 对应API: https://packagist.org/metadata/changes.json
type ChangeTrackingResponse struct {
	// 如果缺少或无效的since参数，则返回错误信息
	Error string `json:"error,omitempty"`
	// 当前时间戳
	Timestamp int64 `json:"timestamp"`
	// 变更操作列表
	Actions []ChangeAction `json:"actions,omitempty"`
}

// ChangeAction 表示单个变更操作
type ChangeAction struct {
	// 操作类型：update或delete
	Type string `json:"type"`
	// 包名
	Package string `json:"package"`
	// 操作时间的Unix时间戳
	Time int64 `json:"time"`
}

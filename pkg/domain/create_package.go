// Package domain 提供了Composer API响应的数据结构定义
package domain

// PackageCreateRequest 表示创建包的请求结构
type PackageCreateRequest struct {
	// 仓库URL
	Repository string `json:"repository"`
}

// PackageCreateResponse 表示创建包的响应结构
type PackageCreateResponse struct {
	// 操作状态
	Status string `json:"status"`
}

// PackageEditRequest 表示编辑包的请求结构
type PackageEditRequest struct {
	// 仓库URL
	Repository string `json:"repository"`
}

// PackageEditResponse 表示编辑包的响应结构
type PackageEditResponse struct {
	// 操作状态
	Status string `json:"status"`
}

// PackageUpdateRequest 表示更新包的请求结构
type PackageUpdateRequest struct {
	// 仓库URL
	Repository string `json:"repository"`
}

// PackageUpdateResponse 表示更新包的响应结构
type PackageUpdateResponse struct {
	// 操作状态
	Status string `json:"status"`
	// 触发的作业ID列表
	Jobs []string `json:"jobs,omitempty"`
}

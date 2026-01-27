// Package domain 提供了Composer API响应的数据结构定义
package domain

// PackageListResponse 表示包列表响应的顶层结构
// 对应API: https://packagist.org/packages/list.json
type PackageListResponse struct {
	// 包名列表
	PackageNames []string `json:"packageNames"`
}

// PackageListWithDataResponse 表示带附加数据的包列表响应
// 对应API: https://packagist.org/packages/list.json?fields[]=repository&fields[]=type
type PackageListWithDataResponse struct {
	// 包信息映射，键为包名，值为包的附加数据
	Packages map[string]PackageData `json:"package"`
}

// PackageData 表示包的附加数据
type PackageData struct {
	// 包类型
	Type string `json:"type,omitempty"`
	// 仓库URL
	Repository string `json:"repository,omitempty"`
	// 是否被废弃
	Abandoned interface{} `json:"abandoned,omitempty"` // 可以是布尔值(false/true)或字符串(推荐替代包)
}

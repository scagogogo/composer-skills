// Package domain 提供了Composer包信息的数据模型定义
package domain

import "time"

// PackageDownloads 表示包的下载统计信息
type PackageDownloads struct {
	// 历史上总共被下载过多少次
	Total int `json:"total" bson:"total"`
	// 本月被下载过多少次
	Monthly int `json:"monthly" bson:"monthly"`
	// 今日被下载过多少次
	Daily int `json:"daily" bson:"daily"`
}

// PackageInfo 表示包的基本信息
type PackageInfo struct {
	// 包的名字
	Name string `json:"name" bson:"name"`

	// 包描述信息
	Description string `json:"description" bson:"description"`

	// 包被创建的时间
	Time time.Time `json:"time" bson:"time"`

	// 维护者的信息
	Maintainers []*Maintainer `json:"maintainers" bson:"maintainers"`

	// 这个包的所有版本相关信息
	Versions map[string]*Version `json:"versions" bson:"versions"`

	// 包的类型
	Type string `json:"type" bson:"type"`
	// 包所对应的仓库
	Repository string `json:"repository" bson:"repository"`
	// 仓库的star有多少
	GithubStars int `json:"github_stars" bson:"github_stars"`
	// 仓库的watch有多少
	GithubWatchers int `json:"github_watchers" bson:"github_watchers"`
	// 仓库的fork有多少
	GithubForks int `json:"github_forks" bson:"github_forks"`
	// 仓库的open issue
	GithubOpenIssues int `json:"github_open_issues" bson:"github_open_issues"`
	// 仓库是啥语言？难不成还会有别的语言？
	Language string `json:"language" bson:"language"`
	// 被多少个其它包所依赖
	Dependents int `json:"dependents" bson:"dependents"`
	// 建议者数量
	Suggesters int `json:"suggesters" bson:"suggesters"`
	// 下载信息
	Downloads PackageDownloads `json:"downloads" bson:"downloads"`
	// 收藏数
	Favers int `json:"favers" bson:"favers"`
}

// ComposerPackageInfo 表示一个Composer包的完整信息
// 包含包的基本信息、版本信息、下载统计和时间戳等
type ComposerPackageInfo struct {
	// 包的名字，用来做唯一主键
	PackageName string `json:"package_name" bson:"package_name"`
	// 小写的包名，用于不区分大小写的查询
	PackageNameLowercase string `json:"package_name_lowercase" bson:"package_name_lowercase"`

	// 包的详细信息
	Package PackageInfo `json:"package" bson:"package"`

	// 包相关信息的md5，用来识别信息是否较之前发生了变化
	PackageInfoMd5 string `json:"package_info_md5" bson:"package_info_md5"`

	// 三个时间戳
	CreateTime *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime *time.Time `json:"update_time" bson:"update_time"`
	ChangeTime *time.Time `json:"change_time" bson:"change_time"`
}

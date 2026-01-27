// Package domain 提供了Composer包信息的数据模型定义
package domain

import "time"

// Version 表示包的一个版本信息
type Version struct {
	// 版本的名字
	Name string `json:"name"`
	// 版本描述
	Description string `json:"description"`
	// 版本关键词
	Keywords []string `json:"keywords"`
	// 首页
	Homepage string `json:"homepage"`
	// 版本号
	Version string `json:"version"`
	// 版本号的格式统一化
	VersionNormalized string `json:"version_normalized"`
	// 这个版本所使用的license
	License []string `json:"license" bson:"license"`
	// 作者信息
	Authors []struct {
		Name     string `json:"name" bson:"name"`
		Email    string `json:"email" bson:"email"`
		Homepage string `json:"homepage" bson:"homepage"`
	} `json:"authors" bson:"authors"`
	// 源代码信息
	Source struct {
		URL       string `json:"url" bson:"url"`
		Type      string `json:"type" bson:"type"`
		Reference string `json:"reference" bson:"reference"`
	} `json:"source" bson:"source"`
	// 分发信息
	Dist struct {
		URL       string `json:"url" bson:"url"`
		Type      string `json:"type" bson:"type"`
		Shasum    string `json:"shasum" bson:"shasum"`
		Reference string `json:"reference" bson:"reference"`
	} `json:"dist" bson:"dist"`
	// 包类型
	Type string `json:"type" bson:"type"`
	// 支持信息
	Support struct {
		Source string `json:"source" bson:"source"`
	} `json:"support" bson:"support"`
	// 资助信息
	Funding []struct {
		URL  string `json:"url" bson:"url"`
		Type string `json:"type" bson:"type"`
	} `json:"funding" bson:"funding"`

	// 可以看做是版本的发布时间
	Time time.Time `json:"time" bson:"time"`

	// 这个版本所依赖的其它包的其它版本
	Require map[string]string `json:"require" bson:"require"`

	// 自动加载相关配置
	Autoload interface{} `json:"autoload" bson:"autoload"`
	// 额外配置
	Extra interface{} `json:"extra" bson:"extra"`
	// 建议安装的包
	Suggest interface{} `json:"suggest" bson:"suggest"`
	// 提供的功能
	Provide interface{} `json:"provide" bson:"provide"`
}

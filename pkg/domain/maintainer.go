// Package domain 提供了Composer包信息的数据模型定义
package domain

// Maintainer 表示一个包维护者相关的信息
type Maintainer struct {
	// 维护者的名字
	Name string `json:"name" bson:"name"`
	// 维护者的头像
	AvatarURL string `json:"avatar_url" bson:"avatar_url"`
}

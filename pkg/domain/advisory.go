// Package domain 提供了Composer API响应的数据结构定义
// 主要包含从Composer安全公告API返回的数据结构
package domain

// AdvisoriesResponse 表示安全公告响应的顶层结构
// 包含了按照包名分组的安全公告列表
// 示例: {"advisories": {"symfony/http-foundation": [{"advisoryId": "PKSA-38s9-s9dj",...}]}}
type AdvisoriesResponse struct {
	// 安全公告映射，键为包名，值为该包的安全公告列表
	Advisories map[string][]*Advisory `json:"advisories"`
}

// Advisory 表示单个安全公告的详细信息
// 包含了漏洞信息、影响范围、报告时间等关键安全数据
// 示例: {"advisoryId": "PKSA-38s9-s9dj", "packageName": "symfony/http-foundation", ...}
type Advisory struct {
	// 安全公告唯一标识符
	// 示例: "PKSA-38s9-s9dj"
	AdvisoryID string `json:"advisoryId"`

	// 受影响的Composer包名称
	// 示例: "symfony/http-foundation"
	PackageName string `json:"packageName"`

	// 远程系统中的ID
	// 示例: "CVE-2022-24894"
	RemoteID string `json:"remoteId"`

	// 安全公告的标题
	// 示例: "HTTP Request Smuggling in Symfony HttpFoundation"
	Title string `json:"title"`

	// 安全公告的详细链接
	// 示例: "https://github.com/advisories/GHSA-rc93-5vf2-xh7q"
	Link string `json:"link"`

	// CVE编号(常见漏洞和暴露标识)
	// 示例: "CVE-2022-24894"
	Cve string `json:"cve"`

	// 受影响的版本范围，使用Composer版本约束语法
	// 示例: ">=5.4.0,<5.4.19|>=6.0.0,<6.0.4"
	AffectedVersions string `json:"affectedVersions"`

	// 安全公告的来源平台
	// 示例: "GitHub"
	Source string `json:"source"`

	// 公告报告时间，ISO 8601格式
	// 示例: "2022-03-10T12:00:00Z"
	ReportedAt string `json:"reportedAt"`

	// 相关的Composer仓库信息
	// 示例: "packagist"
	ComposerRepository string `json:"composerRepository"`

	// 多个来源信息列表
	// 示例: [{"name": "GitHub", "remoteId": "GHSA-rc93-5vf2-xh7q"}]
	Sources []*Source `json:"sources"`
}

// Source 表示安全公告的来源信息
// 可能包含来自不同安全数据库的引用
// 示例: {"name": "GitHub", "remoteId": "GHSA-rc93-5vf2-xh7q"}
type Source struct {
	// 来源名称，如GitHub、NVD等
	// 示例: "GitHub"
	Name string `json:"name"`

	// 来源平台中的远程ID
	// 示例: "GHSA-rc93-5vf2-xh7q"
	RemoteID string `json:"remoteId"`
}

package composer

import (
	"encoding/json"
	"strings"
)

// AuditResult 表示安全审计结果
type AuditResult struct {
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Found           int             `json:"found"`
	Advisory        string          `json:"advisory,omitempty"`
	WithoutDev      bool            `json:"without-dev,omitempty"`
}

// Vulnerability 表示一个安全漏洞
type Vulnerability struct {
	Package     string   `json:"package"`
	Version     string   `json:"version"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	CVE         []string `json:"cve,omitempty"`
	Advisory    string   `json:"advisory"`
	Abandoned   bool     `json:"abandoned,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Source      string   `json:"source,omitempty"`
	Affectedver string   `json:"affectedver,omitempty"`
}

// Audit 执行安全审计
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法对项目的依赖进行安全审计，检查已知的安全漏洞。
//	相当于执行`composer audit`命令，并返回标准格式的输出。
//
// 用法示例：
//
//	output, err := comp.Audit()
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Println("安全审计结果:", output)
func (c *Composer) Audit() (string, error) {
	return c.Run("audit")
}

// AuditWithJSON 执行安全审计并返回 JSON 格式结果
//
// 返回值：
//   - *AuditResult: 解析后的安全审计结果结构体
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法对项目的依赖进行安全审计，并返回结构化的JSON格式结果。
//	相当于执行`composer audit --format=json`命令，并将输出解析为结构体。
//
// 用法示例：
//
//	result, err := comp.AuditWithJSON()
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Printf("发现 %d 个漏洞\n", result.Found)
//	for _, vuln := range result.Vulnerabilities {
//	    fmt.Printf("漏洞: %s %s\n", vuln.Package, vuln.Title)
//	    fmt.Printf("严重性: %s\n", vuln.Severity)
//	    fmt.Printf("详情: %s\n\n", vuln.Link)
//	}
func (c *Composer) AuditWithJSON() (*AuditResult, error) {
	output, err := c.Run("audit", "--format=json")
	if err != nil {
		return nil, err
	}

	var result AuditResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AuditWithoutDev 执行安全审计，不包含开发依赖
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法对项目的非开发依赖进行安全审计，忽略开发环境依赖。
//	相当于执行`composer audit --no-dev`命令。
//
// 用法示例：
//
//	output, err := comp.AuditWithoutDev()
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Println("生产环境依赖安全审计结果:", output)
func (c *Composer) AuditWithoutDev() (string, error) {
	return c.Run("audit", "--no-dev")
}

// AuditWithFormat 使用指定格式执行安全审计
//
// 参数：
//   - format: 输出格式，例如"json"、"table"、"plain"
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法对项目的依赖进行安全审计，并以指定的格式输出结果。
//	相当于执行`composer audit --format=FORMAT`命令。
//
// 用法示例：
//
//	// 以表格形式输出
//	output, err := comp.AuditWithFormat("table")
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Println(output)
//
//	// 以纯文本形式输出
//	output, err = comp.AuditWithFormat("plain")
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) AuditWithFormat(format string) (string, error) {
	return c.Run("audit", "--format="+format)
}

// HasVulnerabilities 检查当前项目是否有漏洞
//
// 返回值：
//   - bool: 如果项目中存在漏洞则返回true，否则返回false
//   - error: 如果检查漏洞过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查项目是否存在任何安全漏洞，如果存在则返回true。
//	这个方法会自动解析审计命令的输出，即使命令因发现漏洞而返回非零退出码也能正确处理。
//
// 用法示例：
//
//	hasVulns, err := comp.HasVulnerabilities()
//	if err != nil {
//	    log.Fatalf("检查漏洞失败: %v", err)
//	}
//
//	if hasVulns {
//	    fmt.Println("警告: 项目中存在安全漏洞!")
//	} else {
//	    fmt.Println("项目中未发现安全漏洞。")
//	}
func (c *Composer) HasVulnerabilities() (bool, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		// 如果命令失败且出错信息包含漏洞，也认为有漏洞
		if strings.Contains(err.Error(), "Found") && strings.Contains(err.Error(), "vulnerabilities") {
			return true, nil
		}
		return false, err
	}

	return result.Found > 0, nil
}

// GetHighSeverityVulnerabilities 获取高严重性的漏洞
//
// 返回值：
//   - []Vulnerability: 高严重性漏洞列表
//   - error: 如果获取漏洞信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行安全审计并筛选出严重性为"high"或"critical"的漏洞。
//	这对于优先修复最危险的安全问题非常有用。
//
// 用法示例：
//
//	highVulns, err := comp.GetHighSeverityVulnerabilities()
//	if err != nil {
//	    log.Fatalf("获取高危漏洞失败: %v", err)
//	}
//
//	if len(highVulns) > 0 {
//	    fmt.Printf("发现 %d 个高危漏洞:\n", len(highVulns))
//	    for _, vuln := range highVulns {
//	        fmt.Printf("包: %s 版本: %s\n", vuln.Package, vuln.Version)
//	        fmt.Printf("漏洞: %s\n", vuln.Title)
//	        fmt.Printf("详情: %s\n\n", vuln.Link)
//	    }
//	} else {
//	    fmt.Println("未发现高危漏洞。")
//	}
func (c *Composer) GetHighSeverityVulnerabilities() ([]Vulnerability, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		return nil, err
	}

	var highSeverity []Vulnerability
	for _, vuln := range result.Vulnerabilities {
		if vuln.Severity == "high" || vuln.Severity == "critical" {
			highSeverity = append(highSeverity, vuln)
		}
	}

	return highSeverity, nil
}

// AuditWithOptions 使用自定义选项执行安全审计
//
// 参数：
//   - options: 审计选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项对项目的依赖进行安全审计，提供最大的灵活性。
//
// 用法示例：
//
//	// 使用多个选项执行审计
//	options := map[string]string{
//	    "no-dev": "",
//	    "format": "json",
//	    "locked": "",
//	}
//	output, err := comp.AuditWithOptions(options)
//	if err != nil {
//	    log.Fatalf("执行安全审计失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) AuditWithOptions(options map[string]string) (string, error) {
	args := []string{"audit"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// AuditLock 审计 composer.lock 文件
//
// 参数：
//   - lockFilePath: composer.lock文件的路径，为空时使用当前目录的lock文件
//
// 返回值：
//   - string: 安全审计的输出结果
//   - error: 如果执行安全审计过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法对指定的composer.lock文件进行安全审计。
//	可以用于审计尚未安装依赖的项目或其他项目的lock文件。
//
// 用法示例：
//
//	// 审计当前项目的lock文件
//	output, err := comp.AuditLock("")
//	if err != nil {
//	    log.Fatalf("审计lock文件失败: %v", err)
//	}
//	fmt.Println(output)
//
//	// 审计其他项目的lock文件
//	output, err = comp.AuditLock("/path/to/other/project/composer.lock")
//	if err != nil {
//	    log.Fatalf("审计lock文件失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) AuditLock(lockFilePath string) (string, error) {
	if lockFilePath == "" {
		return c.Run("audit")
	}
	return c.Run("audit", lockFilePath)
}

// GetAbandonedPackages 获取已放弃的包
//
// 返回值：
//   - []Vulnerability: 已放弃的包列表
//   - error: 如果获取包信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行安全审计并筛选出已被标记为"abandoned"（已放弃维护）的包。
//	使用已放弃的包可能存在安全风险，应考虑替换它们。
//
// 用法示例：
//
//	abandoned, err := comp.GetAbandonedPackages()
//	if err != nil {
//	    log.Fatalf("获取已放弃的包失败: %v", err)
//	}
//
//	if len(abandoned) > 0 {
//	    fmt.Printf("发现 %d 个已放弃维护的包:\n", len(abandoned))
//	    for _, pkg := range abandoned {
//	        fmt.Printf("包: %s 版本: %s\n", pkg.Package, pkg.Version)
//	        fmt.Printf("详情: %s\n\n", pkg.Link)
//	    }
//	    fmt.Println("建议替换这些包以避免潜在的安全风险。")
//	} else {
//	    fmt.Println("未发现已放弃维护的包。")
//	}
func (c *Composer) GetAbandonedPackages() ([]Vulnerability, error) {
	result, err := c.AuditWithJSON()
	if err != nil {
		return nil, err
	}

	var abandoned []Vulnerability
	for _, vuln := range result.Vulnerabilities {
		if vuln.Abandoned {
			abandoned = append(abandoned, vuln)
		}
	}

	return abandoned, nil
}

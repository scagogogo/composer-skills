package composer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/scagogogo/composer-skills/pkg/installer"
)

// ==================== 状态和检查命令 ====================

// StatusResult 表示 `composer status` 命令的结果
type StatusResult struct {
	// Modified 是否有修改的文件
	Modified bool `json:"modified"`
	// Files 修改的文件列表
	Files []string `json:"files,omitempty"`
	// Output 原始输出
	Output string `json:"output,omitempty"`
}

// StatusStructured 检查依赖是否有本地修改（结构化结果）
//
// 返回值：
//   - *StatusResult: 状态检查结果
//   - error: 执行错误
//
// 功能说明：
//
//	该方法检查已安装的依赖是否有本地修改，并返回结构化结果。
//	基于diagnosis.go中已有的Status方法。
//
// 用法示例：
//
//	result, err := comp.StatusStructured()
//	if err != nil {
//	    log.Fatalf("检查状态失败: %v", err)
//	}
//	if result.Modified {
//	    fmt.Printf("发现 %d 个修改的文件\n", len(result.Files))
//	}
func (c *Composer) StatusStructured() (*StatusResult, error) {
	output, err := c.Run("status")
	if err != nil {
		return nil, err
	}
	return ParseStatusOutput(output), nil
}

// ParseStatusOutput 解析 `composer status` 的输出
func ParseStatusOutput(output string) *StatusResult {
	result := &StatusResult{Output: output}
	if strings.TrimSpace(output) == "" {
		result.Modified = false
		return result
	}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result.Modified = true
			result.Files = append(result.Files, line)
		}
	}
	return result
}

// CheckResult 表示 `composer check` 命令的结果
type CheckResult struct {
	// Valid composer.json和composer.lock是否同步
	Valid bool `json:"valid"`
	// Messages 检查消息
	Messages []string `json:"messages,omitempty"`
	// Warnings 警告信息
	Warnings []string `json:"warnings,omitempty"`
	// Errors 错误信息
	Errors []string `json:"errors,omitempty"`
}

// CheckStructured 检查composer.json和composer.lock是否同步（结构化结果）
//
// 返回值：
//   - *CheckResult: 检查结果
//   - error: 执行错误
//
// 功能说明：
//
//	基于diagnosis.go中已有的Check方法，返回结构化结果。
func (c *Composer) CheckStructured() (*CheckResult, error) {
	output, err := c.Run("check")
	result := ParseCheckOutput(output)
	if err != nil {
		result.Valid = false
	}
	return result, nil
}

// ParseCheckOutput 解析 `composer check` 的输出
func ParseCheckOutput(output string) *CheckResult {
	result := &CheckResult{Valid: true}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "error") || strings.Contains(line, "Error") || strings.Contains(line, "FAIL") {
			result.Valid = false
			result.Errors = append(result.Errors, line)
		} else if strings.Contains(line, "warning") || strings.Contains(line, "Warning") || strings.Contains(line, "WARN") {
			result.Warnings = append(result.Warnings, line)
		} else {
			result.Messages = append(result.Messages, line)
		}
	}
	return result
}

// ==================== 批量操作 ====================

// BatchRequireResult 表示批量添加包的结果
type BatchRequireResult struct {
	// Results 各包的添加结果
	Results []RequireResult `json:"results,omitempty"`
	// SuccessCount 成功数量
	SuccessCount int `json:"success_count"`
	// FailCount 失败数量
	FailCount int `json:"fail_count"`
	// TotalCount 总数量
	TotalCount int `json:"total_count"`
}

// BatchRequire 批量添加多个包
//
// 参数：
//   - packages: 包名到版本约束的映射
//   - dev: 是否作为开发依赖
//   - continueOnError: 遇到错误是否继续
//
// 返回值：
//   - *BatchRequireResult: 批量添加结果
//   - error: 第一个遇到的错误（如果continueOnError为false）
//
// 功能说明：
//
//	该方法批量添加多个包。如果continueOnError为true，即使某些包添加失败，
//	也会继续添加其他包。如果为false，遇到第一个错误就停止。
//
// 用法示例：
//
//	packages := map[string]string{
//	    "symfony/console": "^5.4",
//	    "monolog/monolog": "^2.0",
//	    "psr/log": "^1.1",
//	}
//	result, err := comp.BatchRequire(packages, false, true)
//	if err != nil {
//	    log.Fatalf("批量添加失败: %v", err)
//	}
//	fmt.Printf("成功: %d, 失败: %d\n", result.SuccessCount, result.FailCount)
func (c *Composer) BatchRequire(packages map[string]string, dev bool, continueOnError bool) (*BatchRequireResult, error) {
	result := &BatchRequireResult{
		TotalCount: len(packages),
	}

	for pkg, version := range packages {
		err := c.RequirePackage(pkg, version, dev)
		requireResult := RequireResult{
			PackageName: pkg,
			Version:     version,
		}

		if err != nil {
			requireResult.Warnings = append(requireResult.Warnings, err.Error())
			result.FailCount++
			if !continueOnError {
				result.Results = append(result.Results, requireResult)
				return result, err
			}
		} else {
			result.SuccessCount++
		}
		result.Results = append(result.Results, requireResult)
	}

	return result, nil
}

// BatchRemoveResult 表示批量移除包的结果
type BatchRemoveResult struct {
	// Results 各包的移除结果
	Results []RemoveResult `json:"results,omitempty"`
	// SuccessCount 成功数量
	SuccessCount int `json:"success_count"`
	// FailCount 失败数量
	FailCount int `json:"fail_count"`
	// TotalCount 总数量
	TotalCount int `json:"total_count"`
}

// BatchRemove 批量移除多个包
//
// 参数：
//   - packages: 要移除的包名列表
//   - dev: 是否从开发依赖中移除
//   - continueOnError: 遇到错误是否继续
//
// 返回值：
//   - *BatchRemoveResult: 批量移除结果
//   - error: 第一个遇到的错误（如果continueOnError为false）
func (c *Composer) BatchRemove(packages []string, dev bool, continueOnError bool) (*BatchRemoveResult, error) {
	result := &BatchRemoveResult{
		TotalCount: len(packages),
	}

	for _, pkg := range packages {
		err := c.Remove(pkg, dev)
		removeResult := RemoveResult{
			PackageName: pkg,
		}

		if err != nil {
			removeResult.Warnings = append(removeResult.Warnings, err.Error())
			result.FailCount++
			if !continueOnError {
				result.Results = append(result.Results, removeResult)
				return result, err
			}
		} else {
			result.SuccessCount++
		}
		result.Results = append(result.Results, removeResult)
	}

	return result, nil
}

// ==================== 项目健康检查 ====================

// HealthStatus 表示项目健康状态
type HealthStatus struct {
	// ComposerInstalled Composer是否已安装
	ComposerInstalled bool `json:"composer_installed"`
	// ComposerVersion Composer版本
	ComposerVersion string `json:"composer_version,omitempty"`
	// PHPAvailable PHP是否可用
	PHPAvailable bool `json:"php_available"`
	// PHPVersion PHP版本
	PHPVersion string `json:"php_version,omitempty"`
	// HasComposerJson 是否有composer.json
	HasComposerJson bool `json:"has_composer_json"`
	// HasComposerLock 是否有composer.lock
	HasComposerLock bool `json:"has_composer_lock"`
	// HasVendorDir 是否有vendor目录
	HasVendorDir bool `json:"has_vendor_dir"`
	// Valid composer.json是否有效
	Valid bool `json:"valid,omitempty"`
	// OutdatedCount 过时包数量
	OutdatedCount int `json:"outdated_count,omitempty"`
	// VulnerabilityCount 安全漏洞数量
	VulnerabilityCount int `json:"vulnerability_count,omitempty"`
	// AbandonedCount 废弃包数量
	AbandonedCount int `json:"abandoned_count,omitempty"`
	// OverallStatus 总体状态: "healthy", "warning", "critical"
	OverallStatus string `json:"overall_status"`
	// Issues 发现的问题列表
	Issues []string `json:"issues,omitempty"`
}

// HealthCheck 执行项目健康检查
//
// 返回值：
//   - *HealthStatus: 健康状态
//   - error: 错误信息
//
// 功能说明：
//
//	该方法执行全面的项目健康检查，包括：
//	- Composer和PHP环境检查
//	- 项目文件检查（composer.json, composer.lock, vendor）
//	- 依赖有效性检查
//	- 过时包检查
//	- 安全漏洞检查
//	- 废弃包检查
//
//	根据检查结果，OverallStatus可能是：
//	- "healthy": 所有检查通过
//	- "warning": 存在警告（过时包、废弃包等）
//	- "critical": 存在严重问题（安全漏洞、无效配置等）
//
// 用法示例：
//
//	health, err := comp.HealthCheck()
//	if err != nil {
//	    log.Fatalf("健康检查失败: %v", err)
//	}
//	fmt.Printf("总体状态: %s\n", health.OverallStatus)
//	if len(health.Issues) > 0 {
//	    for _, issue := range health.Issues {
//	        fmt.Printf("- %s\n", issue)
//	    }
//	}
func (c *Composer) HealthCheck() (*HealthStatus, error) {
	status := &HealthStatus{
		OverallStatus: "healthy",
	}

	// 1. 检查Composer安装
	status.ComposerInstalled = c.IsInstalled()
	if !status.ComposerInstalled {
		status.OverallStatus = "critical"
		status.Issues = append(status.Issues, "Composer未安装")
	} else {
		if version, err := c.GetVersion(); err == nil {
			status.ComposerVersion = version
		}
	}

	// 2. 检查PHP
	status.PHPAvailable = installer.HasPHP()
	if !status.PHPAvailable {
		status.OverallStatus = "critical"
		status.Issues = append(status.Issues, "PHP未安装")
	} else {
		if phpVer, err := installer.GetPHPVersion(); err == nil {
			status.PHPVersion = phpVer
		}
	}

	// 3. 检查项目文件
	status.HasComposerJson = c.IsProject()
	if !status.HasComposerJson {
		status.OverallStatus = "critical"
		status.Issues = append(status.Issues, "缺少composer.json")
	}

	status.HasComposerLock = c.HasComposerLock()
	if !status.HasComposerLock && status.HasComposerJson {
		if status.OverallStatus == "healthy" {
			status.OverallStatus = "warning"
		}
		status.Issues = append(status.Issues, "缺少composer.lock（建议运行composer install）")
	}

	status.HasVendorDir = c.HasVendorDir()
	if !status.HasVendorDir && status.HasComposerLock {
		if status.OverallStatus == "healthy" {
			status.OverallStatus = "warning"
		}
		status.Issues = append(status.Issues, "缺少vendor目录（建议运行composer install）")
	}

	// 4. 验证composer.json
	if status.HasComposerJson {
		if validateResult, err := c.ValidateStructured(); err == nil {
			status.Valid = validateResult.Valid
			if !validateResult.Valid {
				status.OverallStatus = "critical"
				status.Issues = append(status.Issues, "composer.json验证失败")
				for _, e := range validateResult.Errors {
					status.Issues = append(status.Issues, fmt.Sprintf("验证错误: %s", e))
				}
			}
		}
	}

	// 5. 检查过时包
	if outdated, err := c.GetOutdatedInfo(); err == nil {
		status.OutdatedCount = outdated.Count
		if outdated.Count > 0 {
			if status.OverallStatus == "healthy" {
				status.OverallStatus = "warning"
			}
			status.Issues = append(status.Issues, fmt.Sprintf("有 %d 个过时的包", outdated.Count))
		}
	}

	// 6. 检查安全漏洞
	if audit, err := c.GetAuditInfo(); err == nil {
		status.VulnerabilityCount = audit.Count
		if audit.Count > 0 {
			status.OverallStatus = "critical"
			status.Issues = append(status.Issues, fmt.Sprintf("发现 %d 个安全漏洞", audit.Count))
		}
	}

	// 7. 检查废弃包
	if abandoned, err := c.GetAbandonedPackagesFromLock(); err == nil {
		status.AbandonedCount = len(abandoned)
		if len(abandoned) > 0 {
			if status.OverallStatus == "healthy" {
				status.OverallStatus = "warning"
			}
			status.Issues = append(status.Issues, fmt.Sprintf("有 %d 个废弃的包", len(abandoned)))
		}
	}

	return status, nil
}

// ==================== JSON格式化输出 ====================

// GetInfoAsJSON 获取项目信息并格式化为JSON
//
// 返回值：
//   - string: JSON格式的项目信息
//   - error: 错误信息
//
// 功能说明：
//
//	该方法综合获取项目信息并格式化为JSON字符串。
//	适合在API或CLI工具中使用。
func (c *Composer) GetInfoAsJSON() (string, error) {
	summary, err := c.GetProjectSummary()
	if err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// GetHealthAsJSON 获取健康状态并格式化为JSON
//
// 返回值：
//   - string: JSON格式的健康状态
//   - error: 错误信息
func (c *Composer) GetHealthAsJSON() (string, error) {
	health, err := c.HealthCheck()
	if err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(health, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

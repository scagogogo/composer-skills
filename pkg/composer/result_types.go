package composer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ==================== 结构化输出类型 ====================

// VersionInfo 表示 Composer 版本信息
type VersionInfo struct {
	Version     string    `json:"version"`
	FullOutput  string    `json:"full_output"`
	Major       int       `json:"major"`
	Minor       int       `json:"minor"`
	Patch       int       `json:"patch"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
}

// ParseVersionOutput 解析 `composer --version` 的输出
//
// 参数：
//   - output: composer --version 命令的原始输出
//
// 返回值：
//   - *VersionInfo: 解析后的版本信息
//   - error: 解析错误
//
// 功能说明：
//
//	解析类似 "Composer version 2.6.6 2024-02-22 15:37:50" 的输出，
//	提取版本号和发布日期。
func ParseVersionOutput(output string) (*VersionInfo, error) {
	output = strings.TrimSpace(output)
	info := &VersionInfo{FullOutput: output}

	// 尝试解析 "Composer version X.Y.Z YYYY-MM-DD HH:MM:SS" 格式
	parts := strings.Fields(output)
	if len(parts) < 3 {
		return nil, fmt.Errorf("无法解析版本输出: %s", output)
	}

	// 找到 "version" 关键字后的版本号
	versionStr := ""
	for i, p := range parts {
		if p == "version" && i+1 < len(parts) {
			versionStr = parts[i+1]
			break
		}
	}

	if versionStr == "" {
		return nil, fmt.Errorf("无法从输出中提取版本号: %s", output)
	}

	info.Version = versionStr

	// 解析 major.minor.patch
	var major, minor, patch int
	n, err := fmt.Sscanf(versionStr, "%d.%d.%d", &major, &minor, &patch)
	if err == nil && n >= 2 {
		info.Major = major
		info.Minor = minor
		info.Patch = patch
	}

	// 尝试解析日期 (格式: 2024-02-22 15:37:50)
	for i, p := range parts {
		if len(p) == 10 && strings.Contains(p, "-") {
			dateStr := p
			if i+1 < len(parts) && strings.Contains(parts[i+1], ":") {
				dateStr = p + " " + parts[i+1]
			}
			if t, err := time.Parse("2006-01-02 15:04:05", dateStr); err == nil {
				info.ReleaseDate = t
			}
		}
	}

	return info, nil
}

// GetVersionInfo 获取 Composer 版本信息（结构化）
//
// 返回值：
//   - *VersionInfo: Composer 版本信息
//   - error: 执行或解析错误
//
// 功能说明：
//
//	该方法执行 `composer --version` 并将输出解析为结构化的版本信息，
//	包括主版本号、次版本号、补丁版本号和发布日期。
//
// 用法示例：
//
//	info, err := comp.GetVersionInfo()
//	if err != nil {
//	    log.Fatalf("获取版本信息失败: %v", err)
//	}
//	fmt.Printf("Composer %d.%d.%d\n", info.Major, info.Minor, info.Patch)
func (c *Composer) GetVersionInfo() (*VersionInfo, error) {
	output, err := c.Run("--version")
	if err != nil {
		return nil, err
	}
	return ParseVersionOutput(output)
}

// PackageInfo 表示一个 Composer 包的信息
type PackageInfo struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description,omitempty"`
	Type        string            `json:"type,omitempty"`
	Keywords    []string          `json:"keywords,omitempty"`
	Homepage    string            `json:"homepage,omitempty"`
	License     []string          `json:"license,omitempty"`
	Authors     []PackageAuthor   `json:"authors,omitempty"`
	Support     map[string]string `json:"support,omitempty"`
	Require     map[string]string `json:"require,omitempty"`
	RequireDev  map[string]string `json:"require_dev,omitempty"`
	Autoload    map[string]interface{} `json:"autoload,omitempty"`
	Source      PackageSource     `json:"source,omitempty"`
	Dist        PackageDist       `json:"dist,omitempty"`
	Abandoned   interface{}       `json:"abandoned,omitempty"` // bool or string
	Time        string            `json:"time,omitempty"`
}

// PackageAuthor 表示包的作者信息
type PackageAuthor struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Role     string `json:"role,omitempty"`
}

// PackageSource 表示包的源代码信息
type PackageSource struct {
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Reference string `json:"reference,omitempty"`
}

// PackageDist 表示包的分发信息
type PackageDist struct {
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	Reference string `json:"reference,omitempty"`
	Shasum    string `json:"shasum,omitempty"`
}

// ShowPackageInfo 获取包的详细信息（结构化）
//
// 参数：
//   - packageName: 包名
//
// 返回值：
//   - *PackageInfo: 包的详细信息
//   - error: 执行或解析错误
//
// 功能说明：
//
//	该方法执行 `composer show packageName --format=json` 并将输出解析为结构化的包信息。
//
// 用法示例：
//
//	info, err := comp.ShowPackageInfo("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取包信息失败: %v", err)
//	}
//	fmt.Printf("包 %s 版本 %s\n", info.Name, info.Version)
func (c *Composer) ShowPackageInfo(packageName string) (*PackageInfo, error) {
	output, err := c.Run("show", packageName, "--format", "json")
	if err != nil {
		return nil, err
	}
	return ParsePackageInfo(output)
}

// ParsePackageInfo 解析 `composer show --format=json` 的输出
func ParsePackageInfo(output string) (*PackageInfo, error) {
	output = strings.TrimSpace(output)
	var info PackageInfo
	if err := json.Unmarshal([]byte(output), &info); err != nil {
		return nil, fmt.Errorf("解析包信息失败: %w, 原始输出: %s", err, output)
	}
	return &info, nil
}

// OutdatedPackage 表示一个过时的包
type OutdatedPackage struct {
	Name         string `json:"name"`
	Latest       string `json:"latest"`
	Installed    string `json:"version"`
	LatestStatus string `json:"latest_status"` // "semver-safe-update", "update-possible", "up-to-date"
	Abandoned    interface{} `json:"abandoned,omitempty"`
}

// OutdatedResult 表示过时包检查的结果
type OutdatedResult struct {
	Installed []OutdatedPackage `json:"installed"`
	Count     int               `json:"count,omitempty"`
}

// GetOutdatedInfo 获取过时包信息（结构化）
//
// 返回值：
//   - *OutdatedResult: 过时包信息
//   - error: 执行或解析错误
//
// 功能说明：
//
//	该方法执行 `composer outdated --format=json` 并将输出解析为结构化的过时包信息。
//
// 用法示例：
//
//	result, err := comp.GetOutdatedInfo()
//	if err != nil {
//	    log.Fatalf("获取过时包信息失败: %v", err)
//	}
//	for _, pkg := range result.Installed {
//	    fmt.Printf("%s: %s -> %s\n", pkg.Name, pkg.Installed, pkg.Latest)
//	}
func (c *Composer) GetOutdatedInfo() (*OutdatedResult, error) {
	output, err := c.Run("outdated", "--format", "json")
	if err != nil {
		// composer outdated 在没有过时包时可能返回非零退出码
		// 但输出仍然包含有效的JSON
		if output == "" {
			return &OutdatedResult{Installed: []OutdatedPackage{}}, nil
		}
	}
	return ParseOutdatedResult(output)
}

// GetOutdatedInfoWithOptions 使用选项获取过时包信息（结构化）
//
// 参数：
//   - options: 额外选项
//
// 返回值：
//   - *OutdatedResult: 过时包信息
//   - error: 执行或解析错误
func (c *Composer) GetOutdatedInfoWithOptions(options map[string]string) (*OutdatedResult, error) {
	args := []string{"outdated", "--format", "json"}
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}
	output, err := c.Run(args...)
	if err != nil {
		if output == "" {
			return &OutdatedResult{Installed: []OutdatedPackage{}}, nil
		}
	}
	return ParseOutdatedResult(output)
}

// ParseOutdatedResult 解析 `composer outdated --format=json` 的输出
func ParseOutdatedResult(output string) (*OutdatedResult, error) {
	output = strings.TrimSpace(output)
	var result OutdatedResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("解析过时包信息失败: %w", err)
	}
	result.Count = len(result.Installed)
	return &result, nil
}

// AuditAdvisoryInfo 表示一个安全漏洞公告（用于GetAuditInfo的结构化输出）
// 注意：AuditResult已定义在audit.go中，此处的AuditAdvisoryInfo用于更细粒度的审计结果
type AuditAdvisoryInfo struct {
	PackageName string `json:"package"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Severity    string `json:"severity"` // "critical", "high", "medium", "low"
	CVE         string `json:"cve,omitempty"`
	Link        string `json:"link,omitempty"`
	ReportedAt  string `json:"reportedAt,omitempty"`
}

// AuditInfoResult 表示安全审计的详细结果（使用AuditAdvisoryInfo）
type AuditInfoResult struct {
	Advisories []AuditAdvisoryInfo `json:"advisories"`
	Count      int                 `json:"count,omitempty"`
}

// GetAuditInfo 执行安全审计并返回结构化结果
//
// 返回值：
//   - *AuditResult: 安全审计结果
//   - error: 执行或解析错误
//
// 功能说明：
//
//	该方法执行 `composer audit --format=json` 并将输出解析为结构化的安全审计结果。
//
// 用法示例：
//
//	result, err := comp.GetAuditInfo()
//	if err != nil {
//	    log.Fatalf("安全审计失败: %v", err)
//	}
//	for _, adv := range result.Advisories {
//	    fmt.Printf("漏洞: %s (%s) - %s\n", adv.PackageName, adv.Severity, adv.Title)
//	}
func (c *Composer) GetAuditInfo() (*AuditInfoResult, error) {
	output, err := c.Run("audit", "--format", "json")
	if err != nil {
		// 审计发现漏洞时返回非零退出码，但输出仍包含有效JSON
		if output == "" {
			return &AuditInfoResult{Advisories: []AuditAdvisoryInfo{}}, nil
		}
	}
	return ParseAuditInfoResult(output)
}

// GetAuditInfoWithOptions 使用选项执行安全审计并返回结构化结果
func (c *Composer) GetAuditInfoWithOptions(options map[string]string) (*AuditInfoResult, error) {
	args := []string{"audit", "--format", "json"}
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}
	output, err := c.Run(args...)
	if err != nil {
		if output == "" {
			return &AuditInfoResult{Advisories: []AuditAdvisoryInfo{}}, nil
		}
	}
	return ParseAuditInfoResult(output)
}

// ParseAuditInfoResult 解析 `composer audit --format=json` 的输出为AuditInfoResult
func ParseAuditInfoResult(output string) (*AuditInfoResult, error) {
	output = strings.TrimSpace(output)
	var result AuditInfoResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("解析审计结果失败: %w", err)
	}
	result.Count = len(result.Advisories)
	return &result, nil
}

// SearchResultItem 表示搜索结果中的一个包
type SearchResultItem struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	Repository  string `json:"repository,omitempty"`
}

// SearchResult 表示搜索结果
type SearchResult struct {
	Results []SearchResultItem `json:"results"`
	Total   int                `json:"total,omitempty"`
}

// SearchInfo 搜索包并返回结构化结果
//
// 参数：
//   - query: 搜索关键词
//
// 返回值：
//   - *SearchResult: 搜索结果
//   - error: 执行或解析错误
//
// 功能说明：
//
//	该方法执行 `composer search query --format=json` 并将输出解析为结构化的搜索结果。
//
// 用法示例：
//
//	result, err := comp.SearchInfo("logger")
//	if err != nil {
//	    log.Fatalf("搜索失败: %v", err)
//	}
//	for _, pkg := range result.Results {
//	    fmt.Printf("%s: %s\n", pkg.Name, pkg.Description)
//	}
func (c *Composer) SearchInfo(query string) (*SearchResult, error) {
	output, err := c.Run("search", query, "--format", "json")
	if err != nil {
		return nil, err
	}
	return ParseSearchResult(output)
}

// ParseSearchResult 解析 `composer search --format=json` 的输出
func ParseSearchResult(output string) (*SearchResult, error) {
	output = strings.TrimSpace(output)
	var result SearchResult
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		return nil, fmt.Errorf("解析搜索结果失败: %w", err)
	}
	return &result, nil
}

// DependencyNode 表示依赖树中的一个节点
type DependencyNode struct {
	Name     string           `json:"name"`
	Version  string           `json:"version,omitempty"`
	Children []DependencyNode `json:"children,omitempty"`
}

// ValidateResult 表示验证结果
type ValidateResult struct {
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ParseValidateOutput 解析 `composer validate` 的输出
func ParseValidateOutput(output string) *ValidateResult {
	result := &ValidateResult{Valid: true}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "error") || strings.Contains(line, "Error") {
			result.Valid = false
			result.Errors = append(result.Errors, line)
		} else if strings.Contains(line, "warning") || strings.Contains(line, "Warning") {
			result.Warnings = append(result.Warnings, line)
		}
	}
	return result
}

// ValidateStructured 执行验证并返回结构化结果
//
// 返回值：
//   - *ValidateResult: 验证结果
//   - error: 执行错误
//
// 功能说明：
//
//	该方法执行 `composer validate` 并将输出解析为结构化的验证结果。
//
// 用法示例：
//
//	result, err := comp.ValidateStructured()
//	if err != nil {
//	    log.Fatalf("验证失败: %v", err)
//	}
//	if !result.Valid {
//	    for _, e := range result.Errors {
//	        fmt.Println("错误:", e)
//	    }
//	}
func (c *Composer) ValidateStructured() (*ValidateResult, error) {
	output, err := c.Run("validate")
	// validate 在发现问题时返回非零退出码
	// 但我们仍然需要解析输出
	result := ParseValidateOutput(output)
	if err != nil {
		result.Valid = false
	}
	return result, nil
}

// PlatformRequirement 表示平台需求检查结果
type PlatformRequirement struct {
	Package  string `json:"package"`
	Version  string `json:"version,omitempty"`
	Status   string `json:"status"` // "ok", "missing", "mismatch"
	Required string `json:"required,omitempty"`
}

// PlatformCheckResult 表示平台需求检查结果
type PlatformCheckResult struct {
	Requirements []PlatformRequirement `json:"requirements"`
	OK           bool                  `json:"ok"`
}

// CheckPlatformReqsStructured 检查平台需求并返回结构化结果
//
// 返回值：
//   - *PlatformCheckResult: 平台需求检查结果
//   - error: 执行或解析错误
func (c *Composer) CheckPlatformReqsStructured() (*PlatformCheckResult, error) {
	output, err := c.Run("check-platform-reqs", "--format", "json")
	if err != nil {
		if output == "" {
			return nil, err
		}
	}
	return ParsePlatformCheckResult(output)
}

// ParsePlatformCheckResult 解析 `composer check-platform-reqs --format=json` 的输出
func ParsePlatformCheckResult(output string) (*PlatformCheckResult, error) {
	output = strings.TrimSpace(output)
	var reqs []PlatformRequirement
	if err := json.Unmarshal([]byte(output), &reqs); err != nil {
		return nil, fmt.Errorf("解析平台需求检查结果失败: %w", err)
	}
	result := &PlatformCheckResult{Requirements: reqs, OK: true}
	for _, req := range reqs {
		if req.Status != "ok" && req.Status != "success" {
			result.OK = false
			break
		}
	}
	return result, nil
}

// FundInfo 表示包的资金信息
type FundInfo struct {
	Package string `json:"package"`
	Type    string `json:"type,omitempty"` // "github", "patreon", "tidelift", etc.
	URL     string `json:"url,omitempty"`
}

// FundResult 表示资金信息查询结果
type FundResult struct {
	Funds []FundInfo `json:"funds,omitempty"`
}

// LicenseInfo 表示许可证信息
type LicenseInfo struct {
	Package  string   `json:"package"`
	Version  string   `json:"version,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
}

// LicensesResult 表示许可证查询结果
type LicensesResult struct {
	Licenses []LicenseInfo `json:"licenses,omitempty"`
}

// GetLicensesInfo 获取许可证信息（结构化）
//
// 返回值：
//   - *LicensesResult: 许可证信息
//   - error: 执行或解析错误
func (c *Composer) GetLicensesInfo() (*LicensesResult, error) {
	output, err := c.Run("licenses", "--format", "json")
	if err != nil {
		return nil, err
	}
	return ParseLicensesResult(output)
}

// ParseLicensesResult 解析 `composer licenses --format=json` 的输出
func ParseLicensesResult(output string) (*LicensesResult, error) {
	output = strings.TrimSpace(output)
	// composer licenses --format=json 输出格式: {"vendor/package":["MIT"],...}
	var raw map[string][]string
	if err := json.Unmarshal([]byte(output), &raw); err != nil {
		return nil, fmt.Errorf("解析许可证信息失败: %w", err)
	}
	result := &LicensesResult{}
	for pkg, licenses := range raw {
		result.Licenses = append(result.Licenses, LicenseInfo{
			Package:  pkg,
			Licenses: licenses,
		})
	}
	return result, nil
}

// ScriptInfo 表示脚本信息
type ScriptInfo struct {
	Name    string   `json:"name"`
	Command string   `json:"command,omitempty"`
	Scripts []string `json:"scripts,omitempty"`
}

// DiagnoseResult 表示诊断结果
type DiagnoseResult struct {
	Checks []DiagnoseCheck `json:"checks,omitempty"`
}

// DiagnoseCheck 表示单个诊断检查项
type DiagnoseCheck struct {
	Name   string `json:"name"`
	Status string `json:"status"` // "ok", "warning", "error"
	Detail string `json:"detail,omitempty"`
}

// ParseDiagnoseOutput 解析 `composer diagnose` 的输出
func ParseDiagnoseOutput(output string) *DiagnoseResult {
	result := &DiagnoseResult{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		check := DiagnoseCheck{Detail: line}
		if strings.HasPrefix(line, "[OK]") || strings.HasPrefix(line, "✓") {
			check.Status = "ok"
			check.Name = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "[OK]"), "✓"))
		} else if strings.HasPrefix(line, "[WARNING]") || strings.HasPrefix(line, "⚠") {
			check.Status = "warning"
			check.Name = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "[WARNING]"), "⚠"))
		} else if strings.HasPrefix(line, "[ERROR]") || strings.HasPrefix(line, "✗") {
			check.Status = "error"
			check.Name = strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "[ERROR]"), "✗"))
		} else {
			check.Status = "info"
			check.Name = line
		}
		result.Checks = append(result.Checks, check)
	}
	return result
}

// DiagnoseStructured 执行诊断并返回结构化结果
//
// 返回值：
//   - *DiagnoseResult: 诊断结果
//   - error: 执行错误
func (c *Composer) DiagnoseStructured() (*DiagnoseResult, error) {
	output, err := c.Run("diagnose")
	result := ParseDiagnoseOutput(output)
	return result, err
}

// ConfigItem 表示配置项
type ConfigItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Source string `json:"source,omitempty"`
}

// ConfigResult 表示配置查询结果
type ConfigResult struct {
	Items []ConfigItem `json:"items,omitempty"`
}

// GetConfigStructured 获取配置信息（结构化）
//
// 返回值：
//   - *ConfigResult: 配置信息
//   - error: 执行或解析错误
func (c *Composer) GetConfigStructured() (*ConfigResult, error) {
	output, err := c.Run("config", "--list")
	if err != nil {
		return nil, err
	}
	return ParseConfigList(output), nil
}

// ParseConfigList 解析 `composer config --list` 的输出
func ParseConfigList(output string) *ConfigResult {
	result := &ConfigResult{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts) >= 1 {
			item := ConfigItem{Key: parts[0]}
			if len(parts) >= 2 {
				item.Value = strings.TrimSpace(parts[1])
			}
			result.Items = append(result.Items, item)
		}
	}
	return result
}

// BumpResult 表示包升级结果
type BumpResult struct {
	Bumped []string `json:"bumped,omitempty"`
	Output string   `json:"output,omitempty"`
}

// InstallResult 表示安装结果
type InstallResult struct {
	PackagesInstalled int      `json:"packages_installed"`
	PackagesUpdated   int      `json:"packages_updated"`
	PackagesRemoved   int      `json:"packages_removed"`
	Output            string   `json:"output"`
	Warnings          []string `json:"warnings,omitempty"`
}

// UpdateResult 表示更新结果
type UpdateResult struct {
	PackagesUpdated int      `json:"packages_updated"`
	Output          string   `json:"output"`
	Warnings        []string `json:"warnings,omitempty"`
}

// RequireResult 表示添加包的结果
type RequireResult struct {
	PackageName string   `json:"package_name"`
	Version     string   `json:"version,omitempty"`
	Output      string   `json:"output"`
	Warnings    []string `json:"warnings,omitempty"`
}

// RemoveResult 表示移除包的结果
type RemoveResult struct {
	PackageName string   `json:"package_name"`
	Output      string   `json:"output"`
	Warnings    []string `json:"warnings,omitempty"`
}

package composer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/scagogogo/composer-skills/pkg/installer"
)

// ==================== 便捷方法 ====================

// IsProject 检查当前工作目录是否是一个Composer项目
//
// 返回值：
//   - bool: 是否是Composer项目
//
// 功能说明：
//
//	检查当前工作目录中是否存在composer.json文件。
//	这是判断一个目录是否为PHP/Composer项目的基本方法。
//
// 用法示例：
//
//	if comp.IsProject() {
//	    fmt.Println("这是一个Composer项目")
//	} else {
//	    fmt.Println("这不是一个Composer项目")
//	}
func (c *Composer) IsProject() bool {
	dir := c.workingDir
	if dir == "" {
		dir = "."
	}
	_, err := os.Stat(filepath.Join(dir, "composer.json"))
	return err == nil
}

// IsProjectIn 检查指定目录是否是一个Composer项目
//
// 参数：
//   - dir: 要检查的目录路径
//
// 返回值：
//   - bool: 是否是Composer项目
func (c *Composer) IsProjectIn(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "composer.json"))
	return err == nil
}

// HasComposerLock 检查是否存在composer.lock文件
//
// 返回值：
//   - bool: 是否存在composer.lock
func (c *Composer) HasComposerLock() bool {
	dir := c.workingDir
	if dir == "" {
		dir = "."
	}
	_, err := os.Stat(filepath.Join(dir, "composer.lock"))
	return err == nil
}

// HasVendorDir 检查是否存在vendor目录
//
// 返回值：
//   - bool: 是否存在vendor目录
func (c *Composer) HasVendorDir() bool {
	dir := c.workingDir
	if dir == "" {
		dir = "."
	}
	info, err := os.Stat(filepath.Join(dir, "vendor"))
	return err == nil && info.IsDir()
}

// ComposerJsonData 表示composer.json文件的数据
type ComposerJsonData struct {
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Type            string                 `json:"type,omitempty"`
	Keywords        []string               `json:"keywords,omitempty"`
	Homepage        string                 `json:"homepage,omitempty"`
	License         interface{}            `json:"license,omitempty"` // string or []string
	Authors         []map[string]string    `json:"authors,omitempty"`
	Support         map[string]string      `json:"support,omitempty"`
	Require         map[string]string      `json:"require,omitempty"`
	RequireDev      map[string]string      `json:"require-dev,omitempty"`
	Autoload        map[string]interface{} `json:"autoload,omitempty"`
	AutoloadDev     map[string]interface{} `json:"autoload-dev,omitempty"`
	MinimumStability string               `json:"minimum-stability,omitempty"`
	PreferStable    bool                   `json:"prefer-stable,omitempty"`
	Repositories    []interface{}          `json:"repositories,omitempty"`
	Scripts         map[string]interface{} `json:"scripts,omitempty"`
	Extra           map[string]interface{} `json:"extra,omitempty"`
	Config          map[string]interface{} `json:"config,omitempty"`
	Provide         map[string]string      `json:"provide,omitempty"`
	Suggest         map[string]string      `json:"suggest,omitempty"`
	Replace         map[string]string      `json:"replace,omitempty"`
	Conflict        map[string]string      `json:"conflict,omitempty"`
	Bin             []string               `json:"bin,omitempty"`
}

// ComposerLockData 表示composer.lock文件的数据
type ComposerLockData struct {
	ContentHash  string                 `json:"content-hash,omitempty"`
	Packages     []LockPackageData      `json:"packages,omitempty"`
	PackagesDev  []LockPackageData      `json:"packages-dev,omitempty"`
	Platform     map[string]string      `json:"platform,omitempty"`
	PlatformDev  map[string]string      `json:"platform-dev,omitempty"`
	PluginApiVersion string             `json:"plugin-api-version,omitempty"`
}

// LockPackageData 表示lock文件中的一个包
type LockPackageData struct {
	Name       string            `json:"name"`
	Version    string            `json:"version"`
	Source     map[string]string `json:"source,omitempty"`
	Dist       map[string]string `json:"dist,omitempty"`
	Require    map[string]string `json:"require,omitempty"`
	RequireDev map[string]string `json:"require-dev,omitempty"`
	Type       string            `json:"type,omitempty"`
	License    []string          `json:"license,omitempty"`
	Autoload   map[string]interface{} `json:"autoload,omitempty"`
	Time       string            `json:"time,omitempty"`
	Abandoned  interface{}       `json:"abandoned,omitempty"`
}

// ReadComposerJson 读取并解析composer.json文件
//
// 返回值：
//   - *ComposerJsonData: composer.json数据
//   - error: 读取或解析错误
//
// 功能说明：
//
//	读取当前工作目录中的composer.json文件，并将其解析为结构化的Go数据。
//	不需要执行Composer命令，直接读取文件。
//
// 用法示例：
//
//	data, err := comp.ReadComposerJson()
//	if err != nil {
//	    log.Fatalf("读取composer.json失败: %v", err)
//	}
//	fmt.Printf("项目: %s\n", data.Name)
//	fmt.Printf("依赖数量: %d\n", len(data.Require))
func (c *Composer) ReadComposerJson() (*ComposerJsonData, error) {
	dir := c.workingDir
	if dir == "" {
		dir = "."
	}
	return ReadComposerJsonFile(filepath.Join(dir, "composer.json"))
}

// ReadComposerJsonFile 从指定文件路径读取并解析composer.json
//
// 参数：
//   - filePath: composer.json文件路径
//
// 返回值：
//   - *ComposerJsonData: composer.json数据
//   - error: 读取或解析错误
func ReadComposerJsonFile(filePath string) (*ComposerJsonData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取composer.json失败: %w", err)
	}

	var json_data ComposerJsonData
	if err := json.Unmarshal(data, &json_data); err != nil {
		return nil, fmt.Errorf("解析composer.json失败: %w", err)
	}

	return &json_data, nil
}

// ReadComposerLock 读取并解析composer.lock文件
//
// 返回值：
//   - *ComposerLockData: composer.lock数据
//   - error: 读取或解析错误
//
// 功能说明：
//
//	读取当前工作目录中的composer.lock文件，并将其解析为结构化的Go数据。
//
// 用法示例：
//
//	data, err := comp.ReadComposerLock()
//	if err != nil {
//	    log.Fatalf("读取composer.lock失败: %v", err)
//	}
//	fmt.Printf("已安装包数量: %d\n", len(data.Packages))
func (c *Composer) ReadComposerLock() (*ComposerLockData, error) {
	dir := c.workingDir
	if dir == "" {
		dir = "."
	}
	return ReadComposerLockFile(filepath.Join(dir, "composer.lock"))
}

// ReadComposerLockFile 从指定文件路径读取并解析composer.lock
//
// 参数：
//   - filePath: composer.lock文件路径
//
// 返回值：
//   - *ComposerLockData: composer.lock数据
//   - error: 读取或解析错误
func ReadComposerLockFile(filePath string) (*ComposerLockData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取composer.lock失败: %w", err)
	}

	var lock_data ComposerLockData
	if err := json.Unmarshal(data, &lock_data); err != nil {
		return nil, fmt.Errorf("解析composer.lock失败: %w", err)
	}

	return &lock_data, nil
}

// GetInstalledPackageNames 获取所有已安装包的名称列表
//
// 返回值：
//   - []string: 包名列表
//   - error: 执行错误
//
// 功能说明：
//
//	该方法执行 `composer show` 并解析输出，返回所有已安装包的名称。
//	比获取完整输出更高效，适合只需要包名的场景。
//
// 用法示例：
//
//	packages, err := comp.GetInstalledPackageNames()
//	if err != nil {
//	    log.Fatalf("获取包列表失败: %v", err)
//	}
//	for _, name := range packages {
//	    fmt.Println(name)
//	}
func (c *Composer) GetInstalledPackageNames() ([]string, error) {
	output, err := c.Run("show")
	if err != nil {
		return nil, err
	}
	return ParsePackageList(output), nil
}

// GetDirectDependencyNames 获取直接依赖的包名列表
//
// 返回值：
//   - []string: 直接依赖包名列表
//   - error: 执行错误
func (c *Composer) GetDirectDependencyNames() ([]string, error) {
	output, err := c.Run("show", "--direct")
	if err != nil {
		return nil, err
	}
	return ParsePackageList(output), nil
}

// ParsePackageList 解析 `composer show` 的文本输出，提取包名列表
//
// 参数：
//   - output: composer show 命令的输出
//
// 返回值：
//   - []string: 包名列表
//
// 功能说明：
//
//	解析类似以下格式的输出：
//	  symfony/console    v5.4.0  Eases the creation of beautiful and testable command line interfaces
//	  monolog/monolog    2.5.0   Sends your logs to files, sockets, inboxes, databases and various web services
func ParsePackageList(output string) []string {
	var packages []string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// 格式: "package/name  version  description"
		fields := strings.Fields(line)
		if len(fields) >= 1 && strings.Contains(fields[0], "/") {
			packages = append(packages, fields[0])
		}
	}
	return packages
}

// GetPackageVersionsList 获取指定包的所有可用版本（结构化列表）
//
// 参数：
//   - packageName: 包名
//
// 返回值：
//   - []string: 版本列表
//   - error: 执行错误
//
// 功能说明：
//
//	该方法执行 `composer show packageName --all --format=json` 并解析版本信息。
//	返回该包所有可用版本的列表。
//	与version_constraints.go中的GetPackageVersions不同，此方法返回结构化的版本列表。
//
// 用法示例：
//
//	versions, err := comp.GetPackageVersionsList("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取版本列表失败: %v", err)
//	}
//	for _, v := range versions {
//	    fmt.Println(v)
//	}
func (c *Composer) GetPackageVersionsList(packageName string) ([]string, error) {
	output, err := c.Run("show", packageName, "--all", "--format", "json")
	if err != nil {
		return nil, err
	}
	return ParsePackageVersions(output)
}

// ParsePackageVersions 从 `composer show --all --format=json` 的输出中解析版本列表
func ParsePackageVersions(output string) ([]string, error) {
	output = strings.TrimSpace(output)
	var data struct {
		Versions []string `json:"versions"`
	}
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, fmt.Errorf("解析版本信息失败: %w", err)
	}
	return data.Versions, nil
}

// GetProjectDependencies 获取项目依赖摘要
//
// 返回值：
//   - *ProjectDependencies: 依赖摘要
//   - error: 执行错误
//
// 功能说明：
//
//	该方法综合读取composer.json和composer.lock，返回项目依赖的摘要信息。
//	包括直接依赖数量、开发依赖数量、总安装包数量等。
//
// 用法示例：
//
//	deps, err := comp.GetProjectDependencies()
//	if err != nil {
//	    log.Fatalf("获取依赖摘要失败: %v", err)
//	}
//	fmt.Printf("直接依赖: %d, 开发依赖: %d, 总安装: %d\n",
//	    deps.DirectCount, deps.DevCount, deps.TotalInstalled)
func (c *Composer) GetProjectDependencies() (*ProjectDependencies, error) {
	result := &ProjectDependencies{}

	// 从composer.json获取直接依赖
	jsonData, err := c.ReadComposerJson()
	if err == nil && jsonData != nil {
		result.DirectCount = len(jsonData.Require)
		result.DevCount = len(jsonData.RequireDev)
		result.DirectPackages = make([]string, 0, len(jsonData.Require))
		for name := range jsonData.Require {
			// 排除php和ext-*这样的平台要求
			if name != "php" && !strings.HasPrefix(name, "ext-") && !strings.HasPrefix(name, "lib-") {
				result.DirectPackages = append(result.DirectPackages, name)
			}
		}
		result.DevPackages = make([]string, 0, len(jsonData.RequireDev))
		for name := range jsonData.RequireDev {
			if name != "php" && !strings.HasPrefix(name, "ext-") && !strings.HasPrefix(name, "lib-") {
				result.DevPackages = append(result.DevPackages, name)
			}
		}
	}

	// 从composer.lock获取安装的包
	lockData, err := c.ReadComposerLock()
	if err == nil && lockData != nil {
		result.TotalInstalled = len(lockData.Packages) + len(lockData.PackagesDev)
		result.InstalledPackages = make([]string, 0, result.TotalInstalled)
		for _, pkg := range lockData.Packages {
			result.InstalledPackages = append(result.InstalledPackages, pkg.Name)
		}
		for _, pkg := range lockData.PackagesDev {
			result.InstalledPackages = append(result.InstalledPackages, pkg.Name)
		}
	}

	return result, nil
}

// ProjectDependencies 表示项目依赖摘要
type ProjectDependencies struct {
	// DirectCount 直接依赖数量
	DirectCount int `json:"direct_count"`
	// DevCount 开发依赖数量
	DevCount int `json:"dev_count"`
	// TotalInstalled 总安装包数量（包括间接依赖）
	TotalInstalled int `json:"total_installed"`
	// DirectPackages 直接依赖包名列表
	DirectPackages []string `json:"direct_packages,omitempty"`
	// DevPackages 开发依赖包名列表
	DevPackages []string `json:"dev_packages,omitempty"`
	// InstalledPackages 所有已安装的包名列表
	InstalledPackages []string `json:"installed_packages,omitempty"`
}

// GetProjectSummary 获取项目摘要信息
//
// 返回值：
//   - *ProjectSummary: 项目摘要
//   - error: 错误信息
//
// 功能说明：
//
//	该方法综合多个信息源，返回项目的完整摘要，包括：
//	- composer.json基本信息
//	- 依赖摘要
//	- 过时包数量
//	- 安全漏洞数量
//	- Composer版本
//	- PHP版本
//
// 用法示例：
//
//	summary, err := comp.GetProjectSummary()
//	if err != nil {
//	    log.Fatalf("获取项目摘要失败: %v", err)
//	}
//	fmt.Printf("项目: %s\n", summary.Name)
//	fmt.Printf("过时包: %d, 安全漏洞: %d\n", summary.OutdatedCount, summary.VulnerabilityCount)
func (c *Composer) GetProjectSummary() (*ProjectSummary, error) {
	summary := &ProjectSummary{}

	// 获取项目基本信息
	jsonData, err := c.ReadComposerJson()
	if err == nil && jsonData != nil {
		summary.Name = jsonData.Name
		summary.Description = jsonData.Description
		summary.Type = jsonData.Type
		summary.License = fmt.Sprintf("%v", jsonData.License)
	}

	// 获取依赖信息
	deps, err := c.GetProjectDependencies()
	if err == nil && deps != nil {
		summary.DirectDependencyCount = deps.DirectCount
		summary.DevDependencyCount = deps.DevCount
		summary.TotalInstalledCount = deps.TotalInstalled
	}

	// 获取过时包数量
	outdated, err := c.GetOutdatedInfo()
	if err == nil && outdated != nil {
		summary.OutdatedCount = outdated.Count
	}

	// 获取安全漏洞数量
	auditResult, err := c.GetAuditInfo()
	if err == nil && auditResult != nil {
		summary.VulnerabilityCount = auditResult.Count
	}

	// 获取Composer版本
	version, err := c.GetVersion()
	if err == nil {
		summary.ComposerVersion = version
	}

	// 获取PHP版本
	if installer.HasPHP() {
		if phpVer, err := installer.GetPHPVersion(); err == nil {
			summary.PHPVersion = phpVer
		}
	}

	return summary, nil
}

// ProjectSummary 表示项目摘要信息
type ProjectSummary struct {
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	Type                 string `json:"type,omitempty"`
	License              string `json:"license,omitempty"`
	DirectDependencyCount int   `json:"direct_dependency_count"`
	DevDependencyCount   int    `json:"dev_dependency_count"`
	TotalInstalledCount  int    `json:"total_installed_count"`
	OutdatedCount        int    `json:"outdated_count"`
	VulnerabilityCount   int    `json:"vulnerability_count"`
	ComposerVersion      string `json:"composer_version,omitempty"`
	PHPVersion           string `json:"php_version,omitempty"`
}

// GetRequireWithVersion 获取指定包的已安装版本
//
// 参数：
//   - packageName: 包名
//
// 返回值：
//   - string: 已安装版本号
//   - error: 执行错误
//
// 功能说明：
//
//	获取指定包当前安装的版本号。如果包未安装，返回错误。
//
// 用法示例：
//
//	version, err := comp.GetRequireWithVersion("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取版本失败: %v", err)
//	}
//	fmt.Printf("symfony/console 版本: %s\n", version)
func (c *Composer) GetRequireWithVersion(packageName string) (string, error) {
	output, err := c.Run("show", packageName, "--format", "json")
	if err != nil {
		return "", err
	}

	var data struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return "", fmt.Errorf("解析版本信息失败: %w", err)
	}

	return data.Version, nil
}

// IsPackageInstalled 检查指定包是否已安装
//
// 参数：
//   - packageName: 包名
//
// 返回值：
//   - bool: 是否已安装
func (c *Composer) IsPackageInstalled(packageName string) bool {
	_, err := c.Run("show", packageName)
	return err == nil
}

// IsPackageDev 检查指定包是否为开发依赖
//
// 参数：
//   - packageName: 包名
//
// 返回值：
//   - bool: 是否为开发依赖
//   - error: 错误信息
func (c *Composer) IsPackageDev(packageName string) (bool, error) {
	// 检查composer.json中的require-dev
	jsonData, err := c.ReadComposerJson()
	if err != nil {
		return false, err
	}
	if jsonData.RequireDev != nil {
		if _, ok := jsonData.RequireDev[packageName]; ok {
			return true, nil
		}
	}
	return false, nil
}

// GetPackagesByType 按类型获取已安装的包
//
// 参数：
//   - packageType: 包类型，如"library", "composer-plugin", "project"等
//
// 返回值：
//   - []string: 匹配的包名列表
//   - error: 执行错误
func (c *Composer) GetPackagesByType(packageType string) ([]string, error) {
	// 读取composer.lock获取包类型
	lockData, err := c.ReadComposerLock()
	if err != nil {
		return nil, err
	}

	var packages []string
	for _, pkg := range lockData.Packages {
		if pkg.Type == packageType {
			packages = append(packages, pkg.Name)
		}
	}
	for _, pkg := range lockData.PackagesDev {
		if pkg.Type == packageType {
			packages = append(packages, pkg.Name)
		}
	}

	return packages, nil
}

// GetAbandonedPackagesFromLock 从composer.lock获取已废弃的包列表
//
// 返回值：
//   - []string: 已废弃的包名列表
//   - error: 执行错误
//
// 功能说明：
//
//	与audit.go中的GetAbandonedPackages不同，此方法直接从composer.lock文件读取
//	abandoned标记，而不需要执行composer命令。
func (c *Composer) GetAbandonedPackagesFromLock() ([]string, error) {
	// 读取composer.lock检查abandoned标记
	lockData, err := c.ReadComposerLock()
	if err != nil {
		return nil, err
	}

	var abandoned []string
	for _, pkg := range lockData.Packages {
		if pkg.Abandoned != nil {
			abandoned = append(abandoned, pkg.Name)
		}
	}
	for _, pkg := range lockData.PackagesDev {
		if pkg.Abandoned != nil {
			abandoned = append(abandoned, pkg.Name)
		}
	}

	return abandoned, nil
}

// GetNamespaceMap 获取自动加载命名空间映射
//
// 返回值：
//   - map[string]string: 命名空间到目录的映射
//   - error: 错误信息
//
// 功能说明：
//
//	从composer.json的autoload.psr-4中提取命名空间到目录的映射。
func (c *Composer) GetNamespaceMap() (map[string]string, error) {
	jsonData, err := c.ReadComposerJson()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	if jsonData.Autoload != nil {
		if psr4, ok := jsonData.Autoload["psr-4"]; ok {
			if psr4Map, ok := psr4.(map[string]interface{}); ok {
				for ns, dir := range psr4Map {
					if dirStr, ok := dir.(string); ok {
						result[ns] = dirStr
					} else if dirSlice, ok := dir.([]interface{}); ok && len(dirSlice) > 0 {
						if first, ok := dirSlice[0].(string); ok {
							result[ns] = first
						}
					}
				}
			}
		}
	}

	return result, nil
}

// GetScripts 获取composer.json中定义的脚本
//
// 返回值：
//   - map[string]interface{}: 脚本映射
//   - error: 错误信息
func (c *Composer) GetScripts() (map[string]interface{}, error) {
	jsonData, err := c.ReadComposerJson()
	if err != nil {
		return nil, err
	}
	return jsonData.Scripts, nil
}

// GetComposerHomeDir 获取Composer主目录
//
// 返回值：
//   - string: Composer主目录路径
//   - error: 执行错误
func (c *Composer) GetComposerHomeDir() (string, error) {
	output, err := c.Run("config", "home")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetCacheDir 获取Composer缓存目录
//
// 返回值：
//   - string: 缓存目录路径
//   - error: 执行错误
func (c *Composer) GetCacheDir() (string, error) {
	output, err := c.Run("config", "cache-dir")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// GetVendorDir 获取vendor目录路径
//
// 返回值：
//   - string: vendor目录路径
//   - error: 执行错误
func (c *Composer) GetVendorDir() (string, error) {
	output, err := c.Run("config", "vendor-dir")
	if err != nil {
		// 如果没有自定义vendor-dir，返回默认值
		dir := c.workingDir
		if dir == "" {
			dir = "."
		}
		return filepath.Join(dir, "vendor"), nil
	}
	return strings.TrimSpace(output), nil
}

// GetBinDir 获取bin目录路径
//
// 返回值：
//   - string: bin目录路径
//   - error: 执行错误
func (c *Composer) GetBinDir() (string, error) {
	output, err := c.Run("config", "bin-dir")
	if err != nil {
		dir := c.workingDir
		if dir == "" {
			dir = "."
		}
		return filepath.Join(dir, "vendor", "bin"), nil
	}
	return strings.TrimSpace(output), nil
}

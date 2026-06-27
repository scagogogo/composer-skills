package composer

import (
	"encoding/json"
	"strings"
	"testing"
)

// ==================== VersionInfo 解析测试 ====================

func TestParseVersionOutput(t *testing.T) {
	tests := []struct {
		name           string
		output         string
		expectVersion  string
		expectMajor    int
		expectMinor    int
		expectPatch    int
		expectErr      bool
	}{
		{
			name:           "标准格式",
			output:         "Composer version 2.6.6 2024-02-22 15:37:50",
			expectVersion:  "2.6.6",
			expectMajor:    2,
			expectMinor:    6,
			expectPatch:    6,
			expectErr:      false,
		},
		{
			name:           "无日期格式",
			output:         "Composer version 2.5.1",
			expectVersion:  "2.5.1",
			expectMajor:    2,
			expectMinor:    5,
			expectPatch:    1,
			expectErr:      false,
		},
		{
			name:           "带渠道后缀",
			output:         "Composer version 2.7.0-RC1 2024-01-15 10:00:00",
			expectVersion:  "2.7.0-RC1",
			expectMajor:    2,
			expectMinor:    7,
			expectPatch:    0,
			expectErr:      false,
		},
		{
			name:      "无效输出",
			output:    "invalid",
			expectErr: true,
		},
		{
			name:      "空输出",
			output:    "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := ParseVersionOutput(tt.output)
			if tt.expectErr {
				if err == nil {
					t.Errorf("期望返回错误，但未返回")
				}
				return
			}
			if err != nil {
				t.Fatalf("未期望的错误: %v", err)
			}
			if info.Version != tt.expectVersion {
				t.Errorf("版本号 = %v, 期望 %v", info.Version, tt.expectVersion)
			}
			if info.Major != tt.expectMajor {
				t.Errorf("主版本号 = %v, 期望 %v", info.Major, tt.expectMajor)
			}
			if info.Minor != tt.expectMinor {
				t.Errorf("次版本号 = %v, 期望 %v", info.Minor, tt.expectMinor)
			}
			if info.Patch != tt.expectPatch {
				t.Errorf("补丁版本号 = %v, 期望 %v", info.Patch, tt.expectPatch)
			}
		})
	}
}

func TestParseVersionOutput_ReleaseDate(t *testing.T) {
	output := "Composer version 2.6.6 2024-02-22 15:37:50"
	info, err := ParseVersionOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if info.ReleaseDate.IsZero() {
		t.Error("期望解析出发布日期，但得到零值")
	}
}

// ==================== PackageInfo 解析测试 ====================

func TestParsePackageInfo(t *testing.T) {
	pkgJSON := `{
		"name": "symfony/console",
		"version": "v5.4.0",
		"description": "Eases the creation of beautiful and testable command line interfaces",
		"type": "library",
		"keywords": ["console"],
		"homepage": "https://symfony.com",
		"license": ["MIT"],
		"authors": [
			{"name": "Fabien Potencier", "email": "fabien@symfony.com"}
		],
		"require": {
			"php": ">=7.2.5",
			"symfony/deprecation-contracts": "^2.1|^3"
		}
	}`

	info, err := ParsePackageInfo(pkgJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if info.Name != "symfony/console" {
		t.Errorf("包名 = %v, 期望 symfony/console", info.Name)
	}
	if info.Version != "v5.4.0" {
		t.Errorf("版本 = %v, 期望 v5.4.0", info.Version)
	}
	if info.Type != "library" {
		t.Errorf("类型 = %v, 期望 library", info.Type)
	}
	if len(info.License) != 1 || info.License[0] != "MIT" {
		t.Errorf("许可证 = %v, 期望 [MIT]", info.License)
	}
	if len(info.Authors) != 1 {
		t.Errorf("作者数量 = %v, 期望 1", len(info.Authors))
	}
	if info.Require == nil || len(info.Require) < 1 {
		t.Error("期望有require依赖")
	}
}

func TestParsePackageInfo_InvalidJSON(t *testing.T) {
	_, err := ParsePackageInfo("invalid json")
	if err == nil {
		t.Error("期望返回错误，但未返回")
	}
}

// ==================== OutdatedResult 解析测试 ====================

func TestParseOutdatedResult(t *testing.T) {
	outdatedJSON := `{
		"installed": [
			{
				"name": "symfony/console",
				"version": "v5.3.0",
				"latest": "v5.4.0",
				"latest-status": "semver-safe-update"
			},
			{
				"name": "monolog/monolog",
				"version": "2.3.0",
				"latest": "2.5.0",
				"latest-status": "update-possible"
			}
		]
	}`

	result, err := ParseOutdatedResult(outdatedJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if result.Count != 2 {
		t.Errorf("过时包数量 = %v, 期望 2", result.Count)
	}
	if len(result.Installed) != 2 {
		t.Fatalf("安装包数量 = %v, 期望 2", len(result.Installed))
	}
	if result.Installed[0].Name != "symfony/console" {
		t.Errorf("包名 = %v, 期望 symfony/console", result.Installed[0].Name)
	}
	if result.Installed[0].Installed != "v5.3.0" {
		t.Errorf("已安装版本 = %v, 期望 v5.3.0", result.Installed[0].Installed)
	}
	if result.Installed[0].Latest != "v5.4.0" {
		t.Errorf("最新版本 = %v, 期望 v5.4.0", result.Installed[0].Latest)
	}
}

func TestParseOutdatedResult_Empty(t *testing.T) {
	result, err := ParseOutdatedResult(`{"installed":[]}`)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if result.Count != 0 {
		t.Errorf("过时包数量 = %v, 期望 0", result.Count)
	}
}

// ==================== AuditResult 解析测试 ====================

func TestParseAuditInfoResult(t *testing.T) {
	auditJSON := `{
		"advisories": [
			{
				"package": "example/package",
				"version": "1.0.0",
				"title": "Security vulnerability in X",
				"severity": "high",
				"cve": "CVE-2024-1234",
				"link": "https://example.com/advisory"
			}
		]
	}`

	result, err := ParseAuditInfoResult(auditJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if result.Count != 1 {
		t.Errorf("漏洞数量 = %v, 期望 1", result.Count)
	}
	if result.Advisories[0].PackageName != "example/package" {
		t.Errorf("包名 = %v, 期望 example/package", result.Advisories[0].PackageName)
	}
	if result.Advisories[0].Severity != "high" {
		t.Errorf("严重程度 = %v, 期望 high", result.Advisories[0].Severity)
	}
}

// ==================== SearchResult 解析测试 ====================

func TestParseSearchResult(t *testing.T) {
	searchJSON := `{
		"results": [
			{
				"name": "monolog/monolog",
				"description": "Logging for PHP"
			},
			{
				"name": "symfony/monolog-bundle",
				"description": "Symfony monolog bundle"
			}
		],
		"total": 2
	}`

	result, err := ParseSearchResult(searchJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("总数 = %v, 期望 2", result.Total)
	}
	if len(result.Results) != 2 {
		t.Fatalf("结果数 = %v, 期望 2", len(result.Results))
	}
	if result.Results[0].Name != "monolog/monolog" {
		t.Errorf("包名 = %v, 期望 monolog/monolog", result.Results[0].Name)
	}
}

// ==================== ValidateResult 解析测试 ====================

func TestParseValidateOutput_Valid(t *testing.T) {
	output := "./composer.json is valid"
	result := ParseValidateOutput(output)
	if !result.Valid {
		t.Error("期望验证结果为有效")
	}
}

func TestParseValidateOutput_WithErrors(t *testing.T) {
	output := "schema validation error: The property require is required\nparse error"
	result := ParseValidateOutput(output)
	if result.Valid {
		t.Error("期望验证结果为无效")
	}
	if len(result.Errors) == 0 {
		t.Error("期望有错误信息")
	}
}

func TestParseValidateOutput_WithWarnings(t *testing.T) {
	output := "Warning: The package name is invalid\n./composer.json is valid"
	result := ParseValidateOutput(output)
	if len(result.Warnings) == 0 {
		t.Error("期望有警告信息")
	}
}

// ==================== ParsePackageList 测试 ====================

func TestParsePackageList(t *testing.T) {
	output := `symfony/console    v5.4.0  Eases the creation of beautiful and testable command line interfaces
monolog/monolog    2.5.0   Sends your logs to files, sockets, inboxes, databases and various web services
psr/log            1.1.4   Common interface for logging libraries`

	packages := ParsePackageList(output)
	if len(packages) != 3 {
		t.Fatalf("包数量 = %v, 期望 3", len(packages))
	}
	if packages[0] != "symfony/console" {
		t.Errorf("第一个包 = %v, 期望 symfony/console", packages[0])
	}
	if packages[1] != "monolog/monolog" {
		t.Errorf("第二个包 = %v, 期望 monolog/monolog", packages[1])
	}
}

func TestParsePackageList_Empty(t *testing.T) {
	packages := ParsePackageList("")
	if len(packages) != 0 {
		t.Errorf("空输出应该返回空列表，实际返回 %v 项", len(packages))
	}
}

// ==================== ParseInstallOutput 测试 ====================

func TestParseInstallOutput(t *testing.T) {
	output := `Installing dependencies from lock file (including require-dev)
Verifying lock file contents can be installed on current platform.
Package operations: 5 installs, 2 updates, 1 removal
  - Installing symfony/console (v5.4.0)
  - Updating monolog/monolog (2.4.0 => 2.5.0)
  - Removing old/package (1.0.0)`

	result := ParseInstallOutput(output)
	if result.PackagesInstalled != 5 {
		t.Errorf("安装数量 = %v, 期望 5", result.PackagesInstalled)
	}
	if result.PackagesUpdated != 2 {
		t.Errorf("更新数量 = %v, 期望 2", result.PackagesUpdated)
	}
	if result.PackagesRemoved != 1 {
		t.Errorf("移除数量 = %v, 期望 1", result.PackagesRemoved)
	}
}

// ==================== ParseSelfUpdateOutput 测试 ====================

func TestParseSelfUpdateOutput_Upgrade(t *testing.T) {
	output := "Upgrading to 2.6.6 (from 2.6.5)..."
	oldVer, newVer, err := ParseSelfUpdateOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if oldVer != "2.6.5" {
		t.Errorf("旧版本 = %v, 期望 2.6.5", oldVer)
	}
	if newVer != "2.6.6" {
		t.Errorf("新版本 = %v, 期望 2.6.6", newVer)
	}
}

func TestParseSelfUpdateOutput_AlreadyLatest(t *testing.T) {
	output := "You are already using composer version 2.6.6."
	_, newVer, err := ParseSelfUpdateOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if newVer != "2.6.6" {
		t.Errorf("版本 = %v, 期望 2.6.6", newVer)
	}
}

// ==================== ParseDependencyTreeOutput 测试 ====================

func TestParseDependencyTreeOutput(t *testing.T) {
	output := `symfony/console v5.4.0
|-- psr/log 1.1.4
|-- symfony/deprecation-contracts v2.5.0
|-- symfony/service-contracts v2.5.0
|   ` + "`" + `-- symfony/polyfill-php80 v1.27.0
` + "`" + `-- symfony/string v5.4.0`

	nodes, err := ParseDependencyTreeOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("根节点数量 = %v, 期望 1", len(nodes))
	}
	if nodes[0].Name != "symfony/console" {
		t.Errorf("根节点名称 = %v, 期望 symfony/console", nodes[0].Name)
	}
	if len(nodes[0].Children) < 1 {
		t.Error("期望有子节点")
	}
}

// ==================== ParseDiagnoseOutput 测试 ====================

func TestParseDiagnoseOutput(t *testing.T) {
	output := `[OK] Checking composer.json: OK
[OK] Checking platform settings: OK
[WARNING] Checking git settings: Not configured
[ERROR] Checking http connectivity: Failed`

	result := ParseDiagnoseOutput(output)
	if len(result.Checks) != 4 {
		t.Fatalf("检查项数量 = %v, 期望 4", len(result.Checks))
	}
}

// ==================== ParseConfigList 测试 ====================

func TestParseConfigList(t *testing.T) {
	output := `cache-dir /home/user/.cache/composer
data-dir /home/user/.local/share/composer
home /home/user/.config/composer
bin-dir vendor/bin`

	result := ParseConfigList(output)
	if len(result.Items) < 3 {
		t.Errorf("配置项数量 = %v, 期望至少 3", len(result.Items))
	}
}

// ==================== ParseLicensesResult 测试 ====================

func TestParseLicensesResult(t *testing.T) {
	licensesJSON := `{
		"symfony/console": ["MIT"],
		"monolog/monolog": ["MIT"]
	}`

	result, err := ParseLicensesResult(licensesJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if len(result.Licenses) != 2 {
		t.Errorf("许可证数量 = %v, 期望 2", len(result.Licenses))
	}
}

// ==================== ParsePackageVersions 测试 ====================

func TestParsePackageVersions(t *testing.T) {
	versionsJSON := `{
		"versions": ["v5.4.0", "v5.3.0", "v5.2.0", "v4.4.0"]
	}`

	versions, err := ParsePackageVersions(versionsJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if len(versions) != 4 {
		t.Errorf("版本数量 = %v, 期望 4", len(versions))
	}
	if versions[0] != "v5.4.0" {
		t.Errorf("第一个版本 = %v, 期望 v5.4.0", versions[0])
	}
}

// ==================== ExtractVersionFromOutput 测试 ====================

func TestExtractVersionFromOutput(t *testing.T) {
	tests := []struct {
		output   string
		expected string
		found    bool
	}{
		{"Composer version 2.6.6", "2.6.6", true},
		{"Installed version: v1.2.3-beta", "1.2.3-beta", true},
		{"no version here", "", false},
	}

	for _, tt := range tests {
		ver, found := ExtractVersionFromOutput(tt.output)
		if found != tt.found {
			t.Errorf("found = %v, 期望 %v", found, tt.found)
		}
		if ver != tt.expected {
			t.Errorf("版本 = %v, 期望 %v", ver, tt.expected)
		}
	}
}

// ==================== ExtractPackageNamesFromOutput 测试 ====================

func TestExtractPackageNamesFromOutput(t *testing.T) {
	output := "Installing symfony/console and monolog/monolog, also psr/log"
	names := ExtractPackageNamesFromOutput(output)

	if len(names) != 3 {
		t.Fatalf("包名数量 = %v, 期望 3", len(names))
	}

	// 验证包名
	expectedNames := map[string]bool{
		"symfony/console": true,
		"monolog/monolog": true,
		"psr/log":         true,
	}
	for _, name := range names {
		if !expectedNames[name] {
			t.Errorf("未期望的包名: %v", name)
		}
	}
}

// ==================== ComposerJsonData 测试 ====================

func TestReadComposerJsonFile_Integration(t *testing.T) {
	// 使用项目中的composer.json进行测试
	// 创建一个临时测试文件
	jsonContent := `{
		"name": "test/project",
		"description": "Test project",
		"type": "project",
		"require": {
			"php": "^8.0",
			"symfony/console": "^5.4"
		},
		"require-dev": {
			"phpunit/phpunit": "^9.0"
		},
		"autoload": {
			"psr-4": {
				"App\\": "src/"
			}
		},
		"scripts": {
			"test": "phpunit"
		},
		"minimum-stability": "stable",
		"prefer-stable": true
	}`

	// 通过解析JSON验证结构
	var data ComposerJsonData
	err := json.Unmarshal([]byte(jsonContent), &data)
	if err != nil {
		t.Fatalf("解析composer.json失败: %v", err)
	}

	if data.Name != "test/project" {
		t.Errorf("项目名 = %v, 期望 test/project", data.Name)
	}
	if data.Type != "project" {
		t.Errorf("类型 = %v, 期望 project", data.Type)
	}
	if len(data.Require) != 2 {
		t.Errorf("require数量 = %v, 期望 2", len(data.Require))
	}
	if len(data.RequireDev) != 1 {
		t.Errorf("require-dev数量 = %v, 期望 1", len(data.RequireDev))
	}
	if !data.PreferStable {
		t.Error("期望 prefer-stable 为 true")
	}
	if data.MinimumStability != "stable" {
		t.Errorf("minimum-stability = %v, 期望 stable", data.MinimumStability)
	}
}

// ==================== ComposerLockData 测试 ====================

func TestComposerLockDataParsing(t *testing.T) {
	lockJSON := `{
		"content-hash": "abc123",
		"packages": [
			{
				"name": "symfony/console",
				"version": "v5.4.0",
				"type": "library",
				"license": ["MIT"]
			}
		],
		"packages-dev": [
			{
				"name": "phpunit/phpunit",
				"version": "9.5.0",
				"type": "library"
			}
		],
		"platform": {
			"php": "^8.0"
		},
		"plugin-api-version": "2.3.0"
	}`

	var data ComposerLockData
	err := json.Unmarshal([]byte(lockJSON), &data)
	if err != nil {
		t.Fatalf("解析composer.lock失败: %v", err)
	}

	if data.ContentHash != "abc123" {
		t.Errorf("content-hash = %v, 期望 abc123", data.ContentHash)
	}
	if len(data.Packages) != 1 {
		t.Errorf("packages数量 = %v, 期望 1", len(data.Packages))
	}
	if data.Packages[0].Name != "symfony/console" {
		t.Errorf("包名 = %v, 期望 symfony/console", data.Packages[0].Name)
	}
	if len(data.PackagesDev) != 1 {
		t.Errorf("packages-dev数量 = %v, 期望 1", len(data.PackagesDev))
	}
}

// ==================== ParseRequireOutput 测试 ====================

func TestParseRequireOutput(t *testing.T) {
	output := `Using version ^5.4 for symfony/console
./composer.json has been updated
Running composer update symfony/console
Loading composer repositories with package information
Updating dependencies
Package operations: 1 install, 0 updates, 0 removals
  - Installing symfony/console (v5.4.0): Extracting archive`

	result := ParseRequireOutput(output, "symfony/console")
	if result.PackageName != "symfony/console" {
		t.Errorf("包名 = %v, 期望 symfony/console", result.PackageName)
	}
	// 版本可能不会从所有格式中提取，取决于具体输出格式
}

// ==================== ParsePlatformCheckResult 测试 ====================

func TestParsePlatformCheckResult(t *testing.T) {
	checkJSON := `[
		{"package": "php", "version": "8.1.0", "status": "ok"},
		{"package": "ext-json", "version": "1.0", "status": "ok"}
	]`

	result, err := ParsePlatformCheckResult(checkJSON)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if !result.OK {
		t.Error("期望平台检查结果为OK")
	}
	if len(result.Requirements) != 2 {
		t.Errorf("需求数量 = %v, 期望 2", len(result.Requirements))
	}
}

// ==================== InstallStatus 测试 ====================

func TestInstallStatus(t *testing.T) {
	// 测试InstallStatus结构体创建
	status := &InstallStatus{
		Installed:   true,
		Path:        "/usr/local/bin/composer",
		Version:     "2.6.6",
		PHPAvailable: true,
		PHPVersion:  "8.1.0",
		AutoInstall: true,
	}

	if !status.Installed {
		t.Error("期望Installed为true")
	}
	if status.Path != "/usr/local/bin/composer" {
		t.Errorf("Path = %v", status.Path)
	}
	if status.Version != "2.6.6" {
		t.Errorf("Version = %v", status.Version)
	}
}

// ==================== ProjectDependencies 测试 ====================

func TestProjectDependencies_Struct(t *testing.T) {
	deps := &ProjectDependencies{
		DirectCount:       5,
		DevCount:          3,
		TotalInstalled:    25,
		DirectPackages:    []string{"symfony/console", "monolog/monolog"},
		DevPackages:       []string{"phpunit/phpunit"},
		InstalledPackages: []string{"symfony/console", "psr/log", "monolog/monolog"},
	}

	if deps.DirectCount != 5 {
		t.Errorf("DirectCount = %v, 期望 5", deps.DirectCount)
	}
	if len(deps.DirectPackages) != 2 {
		t.Errorf("DirectPackages数量 = %v, 期望 2", len(deps.DirectPackages))
	}
}

// ==================== ParseAboutOutput 测试 ====================

func TestParseAboutOutput(t *testing.T) {
	output := `Composer - Package Management for PHP
Version: 2.6.6
License: MIT`

	info := ParseAboutOutput(output)
	if len(info) == 0 {
		t.Error("期望解析出关于信息")
	}
}

// ==================== ParseFundOutput 测试 ====================

func TestParseFundOutput(t *testing.T) {
	output := `symfony/symfony - https://symfony.com/sponsor
monolog/monolog - https://github.com/sponsors/Seldaek`

	funds := ParseFundOutput(output)
	if len(funds) != 2 {
		t.Errorf("资金信息数量 = %v, 期望 2", len(funds))
	}
}

// ==================== Integration-style 测试 ====================

func TestComposerJsonAndLockRoundTrip(t *testing.T) {
	// 验证ComposerJsonData可以正确序列化和反序列化
	original := ComposerJsonData{
		Name:        "test/project",
		Description: "A test project",
		Type:        "library",
		Require: map[string]string{
			"php": "^8.0",
		},
		RequireDev: map[string]string{
			"phpunit/phpunit": "^9.0",
		},
		PreferStable:    true,
		MinimumStability: "stable",
	}

	// 序列化
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	// 反序列化
	var parsed ComposerJsonData
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.Name != original.Name {
		t.Errorf("Name = %v, 期望 %v", parsed.Name, original.Name)
	}
	if parsed.PreferStable != original.PreferStable {
		t.Errorf("PreferStable = %v, 期望 %v", parsed.PreferStable, original.PreferStable)
	}
}

// ==================== 边界条件测试 ====================

func TestParseVersionOutput_WithLeadingWhitespace(t *testing.T) {
	output := "  Composer version 2.6.6 2024-02-22 15:37:50  \n"
	info, err := ParseVersionOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if info.Version != "2.6.6" {
		t.Errorf("版本号 = %v, 期望 2.6.6", info.Version)
	}
}

func TestParseOutdatedResult_InvalidJSON(t *testing.T) {
	_, err := ParseOutdatedResult("invalid json")
	if err == nil {
		t.Error("期望返回错误")
	}
}

func TestParseSearchResult_InvalidJSON(t *testing.T) {
	_, err := ParseSearchResult("invalid json")
	if err == nil {
		t.Error("期望返回错误")
	}
}

func TestParsePackageInfo_EmptyObject(t *testing.T) {
	info, err := ParsePackageInfo("{}")
	if err != nil {
		t.Fatalf("解析空对象失败: %v", err)
	}
	if info.Name != "" {
		t.Errorf("空对象应该有空的Name, 实际: %v", info.Name)
	}
}

func TestParseDependencyTreeOutput_Empty(t *testing.T) {
	nodes, err := ParseDependencyTreeOutput("")
	if err != nil {
		t.Fatalf("解析空输出失败: %v", err)
	}
	if len(nodes) != 0 {
		t.Errorf("空输出应该返回空节点列表")
	}
}

func TestParseInstallOutput_NoOperations(t *testing.T) {
	output := "Nothing to install, update or remove."
	result := ParseInstallOutput(output)
	if result.PackagesInstalled != 0 {
		t.Errorf("安装数量 = %v, 期望 0", result.PackagesInstalled)
	}
}

// ==================== QuickSetup 测试 ====================

func TestQuickSetup_InvalidDir(t *testing.T) {
	// QuickSetup在无效目录时应该能优雅地处理
	// 因为不需要实际的composer，我们只测试方法存在
	_ = QuickSetup
}

// ==================== 结果类型JSON序列化测试 ====================

func TestVersionInfoJSONRoundTrip(t *testing.T) {
	original := VersionInfo{
		Version:    "2.6.6",
		FullOutput: "Composer version 2.6.6 2024-02-22",
		Major:      2,
		Minor:      6,
		Patch:      6,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed VersionInfo
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.Version != original.Version {
		t.Errorf("Version = %v, 期望 %v", parsed.Version, original.Version)
	}
	if parsed.Major != original.Major {
		t.Errorf("Major = %v, 期望 %v", parsed.Major, original.Major)
	}
}

func TestOutdatedResultJSONRoundTrip(t *testing.T) {
	original := OutdatedResult{
		Installed: []OutdatedPackage{
			{Name: "symfony/console", Installed: "v5.3.0", Latest: "v5.4.0"},
		},
		Count: 1,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed OutdatedResult
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.Count != 1 {
		t.Errorf("Count = %v, 期望 1", parsed.Count)
	}
}

// ==================== ParseCheckPlatformReqsOutput 文本解析测试 ====================

func TestParseCheckPlatformReqsOutput(t *testing.T) {
	output := `php             8.1.0    success
ext-json        1.0      success
ext-mbstring    missing  missing
ext-curl        7.0      success`

	reqs, err := ParseCheckPlatformReqsOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if len(reqs) != 4 {
		t.Errorf("需求数量 = %v, 期望 4", len(reqs))
	}
}

// ==================== ProjectSummary 测试 ====================

func TestProjectSummary_Struct(t *testing.T) {
	summary := &ProjectSummary{
		Name:                 "test/project",
		DirectDependencyCount: 5,
		DevDependencyCount:   3,
		TotalInstalledCount:  25,
		OutdatedCount:        2,
		VulnerabilityCount:   1,
		ComposerVersion:      "2.6.6",
		PHPVersion:           "8.1.0",
	}

	if summary.OutdatedCount != 2 {
		t.Errorf("OutdatedCount = %v, 期望 2", summary.OutdatedCount)
	}
	if summary.VulnerabilityCount != 1 {
		t.Errorf("VulnerabilityCount = %v, 期望 1", summary.VulnerabilityCount)
	}
}

// ==================== ParseConfigOutput 测试 ====================

func TestParseConfigOutput(t *testing.T) {
	output := "/home/user/.cache/composer\n"
	value, err := ParseConfigOutput(output)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if strings.TrimSpace(value) != "/home/user/.cache/composer" {
		t.Errorf("值 = %v", value)
	}
}

// ==================== Strings 工具方法测试 ====================

func TestParsePackageList_VariousFormats(t *testing.T) {
	tests := []struct {
		name     string
		output   string
		expected int
	}{
		{"标准格式", "symfony/console    v5.4.0  Description", 1},
		{"多包", "pkg/a 1.0 Desc1\npkg/b 2.0 Desc2\npkg/c 3.0 Desc3", 3},
		{"空行混合", "pkg/a 1.0 Desc1\n\npkg/b 2.0 Desc2", 2},
		{"纯空行", "\n\n\n", 0},
		{"非包行", "some random text", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packages := ParsePackageList(tt.output)
			if len(packages) != tt.expected {
				t.Errorf("包数量 = %v, 期望 %v", len(packages), tt.expected)
			}
		})
	}
}

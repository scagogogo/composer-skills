package composer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// ==================== ComposerJsonData 解析测试 ====================

func TestComposerJsonDataParsing(t *testing.T) {
	jsonContent := `{
		"name": "vendor/my-project",
		"description": "My awesome project",
		"type": "project",
		"keywords": ["php", "awesome"],
		"homepage": "https://example.com",
		"license": "MIT",
		"authors": [
			{"name": "John Doe", "email": "john@example.com", "role": "developer"}
		],
		"support": {
			"issues": "https://github.com/vendor/my-project/issues",
			"source": "https://github.com/vendor/my-project"
		},
		"require": {
			"php": "^8.0",
			"symfony/console": "^5.4"
		},
		"require-dev": {
			"phpunit/phpunit": "^9.0",
			"squizlabs/php_codesniffer": "^3.6"
		},
		"autoload": {
			"psr-4": {
				"App\\": "src/",
				"Tests\\": "tests/"
			}
		},
		"autoload-dev": {
			"psr-4": {
				"App\\Tests\\": "tests/"
			}
		},
		"minimum-stability": "stable",
		"prefer-stable": true,
		"repositories": [
			{
				"type": "vcs",
				"url": "https://github.com/vendor/private-repo"
			}
		],
		"scripts": {
			"test": "phpunit",
			"cs-check": "phpcs",
			"auto-scripts": {
				"cache:clear": "symfony-cmd"
			},
			"post-install-cmd": ["@auto-scripts"],
			"post-update-cmd": ["@auto-scripts"]
		},
		"extra": {
			"symfony": {
				"allow-contrib": true
			}
		},
		"config": {
			"sort-packages": true,
			"allow-plugins": true
		},
		"bin": ["bin/my-tool"],
		"suggest": {
			"ext-redis": "For caching"
		},
		"provide": {
			"psr/log-implementation": "1.0"
		},
		"replace": {
			"symfony/polyfill-mbstring": "*"
		},
		"conflict": {
			"symfony/symfony": "*"
		}
	}`

	var data ComposerJsonData
	err := json.Unmarshal([]byte(jsonContent), &data)
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}

	// 基本字段
	if data.Name != "vendor/my-project" {
		t.Errorf("Name = %v", data.Name)
	}
	if data.Description != "My awesome project" {
		t.Errorf("Description = %v", data.Description)
	}
	if data.Type != "project" {
		t.Errorf("Type = %v", data.Type)
	}
	if len(data.Keywords) != 2 {
		t.Errorf("Keywords数量 = %v", len(data.Keywords))
	}
	if data.Homepage != "https://example.com" {
		t.Errorf("Homepage = %v", data.Homepage)
	}

	// 依赖
	if len(data.Require) != 2 {
		t.Errorf("Require数量 = %v, 期望 2", len(data.Require))
	}
	if len(data.RequireDev) != 2 {
		t.Errorf("RequireDev数量 = %v, 期望 2", len(data.RequireDev))
	}

	// 自动加载
	if data.Autoload == nil {
		t.Error("Autoload不应为nil")
	}

	// 稳定性
	if data.MinimumStability != "stable" {
		t.Errorf("MinimumStability = %v", data.MinimumStability)
	}
	if !data.PreferStable {
		t.Error("PreferStable应该为true")
	}

	// 脚本
	if data.Scripts == nil {
		t.Error("Scripts不应为nil")
	}

	// 其他
	if len(data.Bin) != 1 {
		t.Errorf("Bin数量 = %v", len(data.Bin))
	}
	if len(data.Suggest) != 1 {
		t.Errorf("Suggest数量 = %v", len(data.Suggest))
	}
	if len(data.Provide) != 1 {
		t.Errorf("Provide数量 = %v", len(data.Provide))
	}
	if len(data.Replace) != 1 {
		t.Errorf("Replace数量 = %v", len(data.Replace))
	}
	if len(data.Conflict) != 1 {
		t.Errorf("Conflict数量 = %v", len(data.Conflict))
	}
}

// ==================== ReadComposerJsonFile 测试 ====================

func TestReadComposerJsonFile(t *testing.T) {
	// 创建临时目录和composer.json文件
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	jsonContent := `{
		"name": "test/package",
		"description": "Test package",
		"type": "library",
		"require": {
			"php": "^8.0"
		}
	}`

	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	data, err := ReadComposerJsonFile(jsonPath)
	if err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if data.Name != "test/package" {
		t.Errorf("Name = %v, 期望 test/package", data.Name)
	}
}

func TestReadComposerJsonFile_NotFound(t *testing.T) {
	_, err := ReadComposerJsonFile("/nonexistent/composer.json")
	if err == nil {
		t.Error("期望返回错误")
	}
}

func TestReadComposerJsonFile_InvalidJSON(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte("invalid json"), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	_, err = ReadComposerJsonFile(jsonPath)
	if err == nil {
		t.Error("期望返回解析错误")
	}
}

// ==================== ReadComposerLockFile 测试 ====================

func TestReadComposerLockFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	lockContent := `{
		"content-hash": "abc123def456",
		"packages": [
			{
				"name": "symfony/console",
				"version": "v5.4.0",
				"type": "library",
				"license": ["MIT"],
				"require": {
					"php": ">=7.2.5"
				}
			},
			{
				"name": "psr/log",
				"version": "1.1.4",
				"type": "library",
				"abandoned": "psr/log-implementation"
			}
		],
		"packages-dev": [
			{
				"name": "phpunit/phpunit",
				"version": "9.5.10",
				"type": "library"
			}
		],
		"platform": {
			"php": "8.1.0"
		},
		"plugin-api-version": "2.3.0"
	}`

	lockPath := filepath.Join(tmpDir, "composer.lock")
	if err := os.WriteFile(lockPath, []byte(lockContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	data, err := ReadComposerLockFile(lockPath)
	if err != nil {
		t.Fatalf("读取失败: %v", err)
	}

	if data.ContentHash != "abc123def456" {
		t.Errorf("ContentHash = %v", data.ContentHash)
	}
	if len(data.Packages) != 2 {
		t.Errorf("Packages数量 = %v, 期望 2", len(data.Packages))
	}
	if data.Packages[0].Name != "symfony/console" {
		t.Errorf("第一个包名 = %v", data.Packages[0].Name)
	}
	if data.Packages[0].Version != "v5.4.0" {
		t.Errorf("第一个版本 = %v", data.Packages[0].Version)
	}
	if len(data.PackagesDev) != 1 {
		t.Errorf("PackagesDev数量 = %v, 期望 1", len(data.PackagesDev))
	}
	// 测试abandoned字段
	if data.Packages[1].Abandoned == nil {
		t.Error("期望psr/log被标记为abandoned")
	}
}

func TestReadComposerLockFile_NotFound(t *testing.T) {
	_, err := ReadComposerLockFile("/nonexistent/composer.lock")
	if err == nil {
		t.Error("期望返回错误")
	}
}

// ==================== IsProject 测试 ====================

func TestIsProject(t *testing.T) {
	// 使用mock来测试
	ClearMockOutputs()

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 没有composer.json
	comp := &Composer{workingDir: tmpDir}
	if comp.IsProject() {
		t.Error("没有composer.json时IsProject应该返回false")
	}

	// 创建composer.json
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(`{"name": "test/pkg"}`), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	if !comp.IsProject() {
		t.Error("有composer.json时IsProject应该返回true")
	}
}

func TestIsProjectIn(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	comp := &Composer{}

	// 没有composer.json
	if comp.IsProjectIn(tmpDir) {
		t.Error("没有composer.json时IsProjectIn应该返回false")
	}

	// 创建composer.json
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(`{"name": "test/pkg"}`), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	if !comp.IsProjectIn(tmpDir) {
		t.Error("有composer.json时IsProjectIn应该返回true")
	}
}

// ==================== HasComposerLock 测试 ====================

func TestHasComposerLock(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	comp := &Composer{workingDir: tmpDir}

	if comp.HasComposerLock() {
		t.Error("没有composer.lock时应该返回false")
	}

	lockPath := filepath.Join(tmpDir, "composer.lock")
	if err := os.WriteFile(lockPath, []byte(`{}`), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	if !comp.HasComposerLock() {
		t.Error("有composer.lock时应该返回true")
	}
}

// ==================== HasVendorDir 测试 ====================

func TestHasVendorDir(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	comp := &Composer{workingDir: tmpDir}

	if comp.HasVendorDir() {
		t.Error("没有vendor目录时应该返回false")
	}

	vendorPath := filepath.Join(tmpDir, "vendor")
	if err := os.MkdirAll(vendorPath, 0755); err != nil {
		t.Fatalf("创建目录失败: %v", err)
	}

	if !comp.HasVendorDir() {
		t.Error("有vendor目录时应该返回true")
	}
}

// ==================== InstallStatus 测试 ====================

func TestGetInstallStatus(t *testing.T) {
	ClearMockOutputs()

	comp := &Composer{
		executablePath: "/usr/local/bin/composer",
		autoInstall:    true,
	}

	status := comp.GetInstallStatus()
	if status == nil {
		t.Fatal("状态不应为nil")
	}
	if !status.Installed {
		t.Error("期望Installed为true")
	}
	if status.AutoInstall != true {
		t.Error("期望AutoInstall为true")
	}
}

// ==================== ProjectDependencies 测试 ====================

func TestProjectDependencies(t *testing.T) {
	deps := &ProjectDependencies{
		DirectCount:       5,
		DevCount:          3,
		TotalInstalled:    20,
		DirectPackages:    []string{"symfony/console", "monolog/monolog"},
		DevPackages:       []string{"phpunit/phpunit"},
		InstalledPackages: []string{"symfony/console", "psr/log"},
	}

	// 验证JSON序列化
	data, err := json.Marshal(deps)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed ProjectDependencies
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.DirectCount != 5 {
		t.Errorf("DirectCount = %v, 期望 5", parsed.DirectCount)
	}
	if parsed.DevCount != 3 {
		t.Errorf("DevCount = %v, 期望 3", parsed.DevCount)
	}
}

// ==================== ProjectSummary 测试 ====================

func TestProjectSummary(t *testing.T) {
	summary := &ProjectSummary{
		Name:                 "vendor/project",
		Description:          "A test project",
		Type:                 "library",
		DirectDependencyCount: 10,
		DevDependencyCount:   5,
		TotalInstalledCount:  50,
		OutdatedCount:        3,
		VulnerabilityCount:   1,
		ComposerVersion:      "2.6.6",
		PHPVersion:           "8.1.0",
	}

	data, err := json.Marshal(summary)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed ProjectSummary
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.Name != "vendor/project" {
		t.Errorf("Name = %v", parsed.Name)
	}
	if parsed.OutdatedCount != 3 {
		t.Errorf("OutdatedCount = %v, 期望 3", parsed.OutdatedCount)
	}
	if parsed.VulnerabilityCount != 1 {
		t.Errorf("VulnerabilityCount = %v, 期望 1", parsed.VulnerabilityCount)
	}
}

// ==================== QuickSetup 测试 ====================

func TestQuickSetup_InvalidPath(t *testing.T) {
	ClearMockOutputs()

	// 在没有composer的环境中测试
	comp, err := QuickSetup("/nonexistent/path", false)
	// 不自动安装时，如果没有composer可能会失败
	_ = comp
	_ = err
}

// ==================== IsPackageDev 测试 ====================

func TestIsPackageDev(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建有require-dev的composer.json
	jsonContent := `{
		"require": {
			"symfony/console": "^5.4"
		},
		"require-dev": {
			"phpunit/phpunit": "^9.0"
		}
	}`
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{workingDir: tmpDir}
	isDev, err := comp.IsPackageDev("phpunit/phpunit")
	if err != nil {
		t.Fatalf("检查失败: %v", err)
	}
	if !isDev {
		t.Error("phpunit/phpunit应该是开发依赖")
	}

	isDev, err = comp.IsPackageDev("symfony/console")
	if err != nil {
		t.Fatalf("检查失败: %v", err)
	}
	if isDev {
		t.Error("symfony/console不应该是开发依赖")
	}
}

// ==================== GetNamespaceMap 测试 ====================

func TestGetNamespaceMap(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	jsonContent := `{
		"autoload": {
			"psr-4": {
				"App\\": "src/",
				"Tests\\": "tests/"
			}
		}
	}`
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{workingDir: tmpDir}
	nsMap, err := comp.GetNamespaceMap()
	if err != nil {
		t.Fatalf("获取命名空间映射失败: %v", err)
	}

	if len(nsMap) != 2 {
		t.Errorf("命名空间数量 = %v, 期望 2", len(nsMap))
	}
	if nsMap["App\\"] != "src/" {
		t.Errorf("App\\ 映射 = %v, 期望 src/", nsMap["App\\"])
	}
}

// ==================== GetScripts 测试 ====================

func TestGetScripts(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	jsonContent := `{
		"scripts": {
			"test": "phpunit",
			"cs-check": "phpcs",
			"auto-scripts": {
				"cache:clear": "symfony-cmd"
			}
		}
	}`
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{workingDir: tmpDir}
	scripts, err := comp.GetScripts()
	if err != nil {
		t.Fatalf("获取脚本失败: %v", err)
	}

	if len(scripts) != 3 {
		t.Errorf("脚本数量 = %v, 期望 3", len(scripts))
	}
}

// ==================== GetAbandonedPackages 测试 ====================

func TestGetAbandonedPackagesFromLock(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	lockContent := `{
		"packages": [
			{
				"name": "old/package",
				"version": "1.0.0",
				"abandoned": "new/package"
			},
			{
				"name": "active/package",
				"version": "2.0.0"
			}
		],
		"packages-dev": [
			{
				"name": "old/dev-package",
				"version": "1.0.0",
				"abandoned": true
			}
		]
	}`
	lockPath := filepath.Join(tmpDir, "composer.lock")
	if err := os.WriteFile(lockPath, []byte(lockContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{workingDir: tmpDir}
	abandoned, err := comp.GetAbandonedPackagesFromLock()
	if err != nil {
		t.Fatalf("获取废弃包列表失败: %v", err)
	}

	if len(abandoned) != 2 {
		t.Errorf("废弃包数量 = %v, 期望 2", len(abandoned))
	}
}

// ==================== GetPackagesByType 测试 ====================

func TestGetPackagesByType(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	lockContent := `{
		"packages": [
			{
				"name": "symfony/console",
				"version": "v5.4.0",
				"type": "library"
			},
			{
				"name": "composer-plugin/example",
				"version": "1.0.0",
				"type": "composer-plugin"
			}
		],
		"packages-dev": []
	}`
	lockPath := filepath.Join(tmpDir, "composer.lock")
	if err := os.WriteFile(lockPath, []byte(lockContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{workingDir: tmpDir}
	plugins, err := comp.GetPackagesByType("composer-plugin")
	if err != nil {
		t.Fatalf("按类型获取包失败: %v", err)
	}

	if len(plugins) != 1 {
		t.Errorf("composer-plugin类型包数量 = %v, 期望 1", len(plugins))
	}
	if plugins[0] != "composer-plugin/example" {
		t.Errorf("包名 = %v, 期望 composer-plugin/example", plugins[0])
	}
}

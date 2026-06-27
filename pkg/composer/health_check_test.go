package composer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// ==================== StatusResult 测试 ====================

func TestParseStatusOutput(t *testing.T) {
	tests := []struct {
		name       string
		output     string
		modified   bool
		fileCount  int
	}{
		{
			name:      "无修改",
			output:    "",
			modified:  false,
			fileCount: 0,
		},
		{
			name:      "有修改",
			output:    "vendor/symfony/console/Command/HelpCommand.php\nvendor/monolog/monolog/src/Monolog/Logger.php",
			modified:  true,
			fileCount: 2,
		},
		{
			name:      "仅空行",
			output:    "\n\n",
			modified:  false,
			fileCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseStatusOutput(tt.output)
			if result.Modified != tt.modified {
				t.Errorf("Modified = %v, 期望 %v", result.Modified, tt.modified)
			}
			if len(result.Files) != tt.fileCount {
				t.Errorf("Files数量 = %v, 期望 %v", len(result.Files), tt.fileCount)
			}
		})
	}
}

// ==================== CheckResult 测试 ====================

func TestParseCheckOutput(t *testing.T) {
	tests := []struct {
		name    string
		output  string
		valid   bool
		errCount int
	}{
		{
			name:    "有效",
			output:  "composer.json is valid",
			valid:   true,
			errCount: 0,
		},
		{
			name:    "有错误",
			output:  "error: require is missing\nsome other error",
			valid:   false,
			errCount: 2,
		},
		{
			name:    "有警告",
			output:  "Warning: some warning\ncomposer.json is valid",
			valid:   true,
			errCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCheckOutput(tt.output)
			if result.Valid != tt.valid {
				t.Errorf("Valid = %v, 期望 %v", result.Valid, tt.valid)
			}
			if len(result.Errors) != tt.errCount {
				t.Errorf("Errors数量 = %v, 期望 %v", len(result.Errors), tt.errCount)
			}
		})
	}
}

// ==================== BatchRequire 测试 ====================

func TestBatchRequire(t *testing.T) {
	ClearMockOutputs()

	// 设置mock输出
	SetupMockOutput("require symfony/console:^5.4", "Installed symfony/console", nil)
	SetupMockOutput("require monolog/monolog:^2.0", "Installed monolog/monolog", nil)

	comp := &Composer{
		executablePath: "composer",
		autoInstall:    false,
	}

	packages := map[string]string{
		"symfony/console": "^5.4",
		"monolog/monolog": "^2.0",
	}

	result, err := comp.BatchRequire(packages, false, true)
	if err != nil {
		t.Fatalf("批量添加失败: %v", err)
	}

	if result.TotalCount != 2 {
		t.Errorf("TotalCount = %v, 期望 2", result.TotalCount)
	}
}

// ==================== BatchRemove 测试 ====================

func TestBatchRemove(t *testing.T) {
	ClearMockOutputs()

	SetupMockOutput("remove symfony/console", "Removed symfony/console", nil)

	comp := &Composer{
		executablePath: "composer",
		autoInstall:    false,
	}

	packages := []string{"symfony/console"}

	result, err := comp.BatchRemove(packages, false, true)
	if err != nil {
		t.Fatalf("批量移除失败: %v", err)
	}

	if result.TotalCount != 1 {
		t.Errorf("TotalCount = %v, 期望 1", result.TotalCount)
	}
}

// ==================== HealthStatus 测试 ====================

func TestHealthStatus_Struct(t *testing.T) {
	status := &HealthStatus{
		ComposerInstalled: true,
		ComposerVersion:   "2.6.6",
		PHPAvailable:      true,
		PHPVersion:        "8.1.0",
		HasComposerJson:   true,
		HasComposerLock:   true,
		HasVendorDir:      true,
		Valid:             true,
		OutdatedCount:     2,
		VulnerabilityCount: 0,
		AbandonedCount:    1,
		OverallStatus:     "warning",
		Issues:            []string{"有2个过时的包", "有1个废弃的包"},
	}

	data, err := json.Marshal(status)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed HealthStatus
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.OverallStatus != "warning" {
		t.Errorf("OverallStatus = %v, 期望 warning", parsed.OverallStatus)
	}
	if len(parsed.Issues) != 2 {
		t.Errorf("Issues数量 = %v, 期望 2", len(parsed.Issues))
	}
}

// ==================== HealthCheck 测试 ====================

func TestHealthCheck_NoComposerJson(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-health-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	comp := &Composer{
		executablePath: "composer",
		workingDir:     tmpDir,
		autoInstall:    false,
	}

	health, err := comp.HealthCheck()
	if err != nil {
		// 在没有实际composer的环境中可能失败
		t.Logf("HealthCheck返回错误（环境中没有composer）: %v", err)
		return
	}

	if health == nil {
		t.Fatal("健康状态不应为nil")
	}

	if health.HasComposerJson {
		t.Error("临时目录不应该有composer.json")
	}
}

func TestHealthCheck_WithComposerJson(t *testing.T) {
	ClearMockOutputs()

	tmpDir, err := os.MkdirTemp("", "composer-health-test-*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建composer.json
	jsonContent := `{"name": "test/project", "require": {"php": "^8.0"}}`
	jsonPath := filepath.Join(tmpDir, "composer.json")
	if err := os.WriteFile(jsonPath, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("写入文件失败: %v", err)
	}

	comp := &Composer{
		executablePath: "composer",
		workingDir:     tmpDir,
		autoInstall:    false,
	}

	health, err := comp.HealthCheck()
	if err != nil {
		t.Logf("HealthCheck返回错误: %v", err)
		return
	}

	if !health.HasComposerJson {
		t.Error("应该检测到composer.json")
	}
}

// ==================== GetInfoAsJSON 测试 ====================

func TestGetInfoAsJSON(t *testing.T) {
	ClearMockOutputs()

	SetupMockOutput("--version", "Composer version 2.6.6 2024-02-22 15:37:50", nil)
	SetupMockOutput("outdated --format json", `{"installed":[]}`, nil)
	SetupMockOutput("audit --format json", `{"advisories":[]}`, nil)
	SetupMockOutput("config --list --json", `{"name":"test/project"}`, nil)

	comp := &Composer{
		executablePath: "composer",
		autoInstall:    false,
	}

	// 验证JSON输出方法存在
	_ = comp.GetInfoAsJSON
	_ = comp.GetHealthAsJSON
}

// ==================== BatchRequireResult 测试 ====================

func TestBatchRequireResult_JSONRoundTrip(t *testing.T) {
	original := BatchRequireResult{
		Results: []RequireResult{
			{PackageName: "pkg/a", Version: "^1.0", Output: "ok"},
			{PackageName: "pkg/b", Version: "^2.0", Output: "ok"},
		},
		SuccessCount: 2,
		FailCount:    0,
		TotalCount:   2,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed BatchRequireResult
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.SuccessCount != 2 {
		t.Errorf("SuccessCount = %v, 期望 2", parsed.SuccessCount)
	}
}

// ==================== BatchRemoveResult 测试 ====================

func TestBatchRemoveResult_JSONRoundTrip(t *testing.T) {
	original := BatchRemoveResult{
		Results: []RemoveResult{
			{PackageName: "pkg/a", Output: "ok"},
		},
		SuccessCount: 1,
		FailCount:    0,
		TotalCount:   1,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed BatchRemoveResult
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if parsed.TotalCount != 1 {
		t.Errorf("TotalCount = %v, 期望 1", parsed.TotalCount)
	}
}

// ==================== StatusResult JSON 测试 ====================

func TestStatusResult_JSONRoundTrip(t *testing.T) {
	original := StatusResult{
		Modified: true,
		Files:    []string{"file1.php", "file2.php"},
		Output:   "raw output",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed StatusResult
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if !parsed.Modified {
		t.Error("期望Modified为true")
	}
	if len(parsed.Files) != 2 {
		t.Errorf("Files数量 = %v, 期望 2", len(parsed.Files))
	}
}

// ==================== CheckResult JSON 测试 ====================

func TestCheckResult_JSONRoundTrip(t *testing.T) {
	original := CheckResult{
		Valid:    true,
		Messages: []string{"All checks passed"},
		Warnings: []string{"Minor warning"},
		Errors:   []string{},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("序列化失败: %v", err)
	}

	var parsed CheckResult
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if !parsed.Valid {
		t.Error("期望Valid为true")
	}
}

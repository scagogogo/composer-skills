package composer

import (
	"errors"
	"testing"
)

func TestLicenses(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command
	SetupMockOutput("licenses", "Name: vendor/package1\nVersion: 1.0.0\nLicenses: MIT\nAuthors: John Doe <john@example.com>\n\nName: vendor/package2\nVersion: 2.0.0\nLicenses: Apache-2.0\nAuthors: Jane Smith <jane@example.com>", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Licenses()
	if err != nil {
		t.Errorf("Licenses执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}

	if !contains(output, "MIT") {
		t.Errorf("输出应包含许可证信息，实际为\"%s\"", output)
	}

	if !contains(output, "Apache-2.0") {
		t.Errorf("输出应包含许可证信息，实际为\"%s\"", output)
	}
}

func TestLicensesWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("licenses", "", errors.New("licenses command failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.Licenses()
	if err == nil {
		t.Error("Licenses命令失败时应该返回错误")
	}
}

func TestLicensesWithFormat(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command with JSON format
	SetupMockOutput("licenses --format=json", `[{"name":"vendor/package1","version":"1.0.0","license":["MIT"],"authors":[{"name":"John Doe","email":"john@example.com"}]},{"name":"vendor/package2","version":"2.0.0","license":["Apache-2.0"],"authors":[{"name":"Jane Smith","email":"jane@example.com"}]}]`, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.LicensesWithFormat("json")
	if err != nil {
		t.Errorf("LicensesWithFormat执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("JSON输出应包含包名，实际为\"%s\"", output)
	}

	if !contains(output, "MIT") {
		t.Errorf("JSON输出应包含许可证信息，实际为\"%s\"", output)
	}
}

func TestLicensesWithFormatAndEmptyFormat(t *testing.T) {
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.LicensesWithFormat("")
	if err == nil {
		t.Error("空格式应该返回错误")
	}
}

func TestLicensesWithInvalidFormat(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error for invalid format
	SetupMockOutput("licenses --format=invalid", "", errors.New("invalid format"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	_, err = composer.LicensesWithFormat("invalid")
	if err == nil {
		t.Error("无效格式应该返回错误")
	}
}

func TestLicensesWithNoDevDependencies(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command without dev dependencies
	SetupMockOutput("licenses --no-dev", "Name: vendor/package1\nVersion: 1.0.0\nLicenses: MIT\nAuthors: John Doe <john@example.com>", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"no-dev": "",
	}

	output, err := composer.LicensesWithOptions(options)
	if err != nil {
		t.Errorf("LicensesWithOptions执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}

	if !contains(output, "MIT") {
		t.Errorf("输出应包含许可证信息，实际为\"%s\"", output)
	}
}

func TestLicensesWithOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command with options
	SetupMockOutput("licenses --format=json", `[{"name":"vendor/package1","version":"1.0.0","license":["MIT"]}]`, nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"format": "json",
	}

	output, err := composer.LicensesWithOptions(options)
	if err != nil {
		t.Errorf("LicensesWithOptions执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestLicensesWithNilOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command without options
	SetupMockOutput("licenses", "Name: vendor/package1\nVersion: 1.0.0\nLicenses: MIT\nAuthors: John Doe <john@example.com>", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.LicensesWithOptions(nil)
	if err != nil {
		t.Errorf("LicensesWithOptions（nil选项）执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestLicensesWithEmptyOptions(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command without options
	SetupMockOutput("licenses", "Name: vendor/package1\nVersion: 1.0.0\nLicenses: MIT\nAuthors: John Doe <john@example.com>", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.LicensesWithOptions(map[string]string{})
	if err != nil {
		t.Errorf("LicensesWithOptions（空选项）执行失败: %v", err)
	}

	if !contains(output, "vendor/package1") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestLicensesWithEmptyOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with empty result
	SetupMockOutput("licenses", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Licenses()
	if err != nil {
		t.Errorf("Licenses执行失败: %v", err)
	}

	if output != "" {
		t.Errorf("空输出应该返回空字符串，实际为\"%s\"", output)
	}
}

func TestLicensesWithSummary(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for licenses command with summary
	SetupMockOutput("licenses --summary", "MIT: 5 packages\nApache-2.0: 3 packages\nBSD-3-Clause: 2 packages\nGPL-2.0+: 1 package", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	options := map[string]string{
		"summary": "",
	}
	output, err := composer.LicensesWithOptions(options)
	if err != nil {
		t.Errorf("LicensesWithOptions执行失败: %v", err)
	}

	if !contains(output, "MIT") {
		t.Errorf("输出应包含MIT，实际为\"%s\"", output)
	}
}

func TestCheckLicenses(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("licenses --check", "Licenses check passed", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.CheckLicenses()

	if err != nil {
		t.Errorf("CheckLicenses failed: %v", err)
	}
	if !contains(output, "check") {
		t.Errorf("Expected 'check' in output, got '%s'", output)
	}
}

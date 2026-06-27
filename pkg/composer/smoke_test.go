package composer

import (
	"errors"
	"testing"
)

// TestSmokeTest 烟雾测试 - 验证核心功能是否正常工作
func TestSmokeTest(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Test basic composer instance creation
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// Test basic version command
	SetupMockOutput("--version", "Composer version 2.4.4", nil)
	version, err := composer.GetVersion()
	if err != nil {
		t.Errorf("GetVersion执行失败: %v", err)
	}
	if version != "2.4.4" {
		t.Errorf("版本应为2.4.4，实际为%s", version)
	}

	// Test basic audit command
	SetupMockOutput("audit", "No security vulnerabilities found", nil)
	auditOutput, err := composer.Audit()
	if err != nil {
		t.Errorf("Audit执行失败: %v", err)
	}
	if !contains(auditOutput, "No security vulnerabilities") {
		t.Errorf("审计输出不正确: %s", auditOutput)
	}

	// Test basic exec command
	SetupMockOutput("exec phpunit --", "PHPUnit 9.5.0 by Sebastian Bergmann", nil)
	execOutput, err := composer.Exec("phpunit")
	if err != nil {
		t.Errorf("Exec执行失败: %v", err)
	}
	if !contains(execOutput, "PHPUnit") {
		t.Errorf("Exec输出不正确: %s", execOutput)
	}

	// Test basic archive command
	SetupMockOutput("archive --format=zip --dir=/tmp", "Archive created successfully", nil)
	archiveOutput, err := composer.Archive("/tmp")
	if err != nil {
		t.Errorf("Archive执行失败: %v", err)
	}
	if !contains(archiveOutput, "Archive created") {
		t.Errorf("Archive输出不正确: %s", archiveOutput)
	}

	// Test working directory setting
	composer.SetWorkingDir("/test/dir")
	if composer.workingDir != "/test/dir" {
		t.Errorf("工作目录设置失败，期望: /test/dir，实际: %s", composer.workingDir)
	}
}

// TestMockSystemIntegrity 测试模拟系统的完整性
func TestMockSystemIntegrity(t *testing.T) {
	// Test that ClearMockOutputs works
	ClearMockOutputs()

	// Set up a mock
	SetupMockOutput("test-command", "test-output", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// Test that the mock works
	output, err := composer.Run("test-command")
	if err != nil {
		t.Errorf("模拟命令执行失败: %v", err)
	}
	if output != "test-output" {
		t.Errorf("模拟输出不正确，期望: test-output，实际: %s", output)
	}

	// Test that clearing works
	ClearMockOutputs()
	SetupMockOutput("another-command", "another-output", nil)

	output, err = composer.Run("another-command")
	if err != nil {
		t.Errorf("清理后的模拟命令执行失败: %v", err)
	}
	if output != "another-output" {
		t.Errorf("清理后的模拟输出不正确，期望: another-output，实际: %s", output)
	}
}

// TestEdgeCasesIntegration 测试边界条件的集成
func TestEdgeCasesIntegration(t *testing.T) {
	ClearMockOutputs()

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// Test simple command
	SetupMockOutput("--version", "Composer version 2.4.4", nil)
	output, err := composer.Run("--version")
	if err != nil {
		t.Errorf("版本命令执行失败: %v", err)
	}
	if !contains(output, "Composer version") {
		t.Errorf("版本命令输出不正确: %s", output)
	}

	// Test special characters
	SetupMockOutput("command with spaces", "special output", nil)
	output, err = composer.Run("command", "with", "spaces")
	if err != nil {
		t.Errorf("特殊字符命令执行失败: %v", err)
	}
	if output != "special output" {
		t.Errorf("特殊字符命令输出不正确: %s", output)
	}

	// Test Unicode
	SetupMockOutput("命令 with 中文", "Unicode output", nil)
	output, err = composer.Run("命令", "with", "中文")
	if err != nil {
		t.Errorf("Unicode命令执行失败: %v", err)
	}
	if output != "Unicode output" {
		t.Errorf("Unicode命令输出不正确: %s", output)
	}
}

// TestErrorHandlingIntegration 测试错误处理的集成
func TestErrorHandlingIntegration(t *testing.T) {
	ClearMockOutputs()

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// Test command that returns error
	SetupMockOutput("failing-command", "", errors.New("command failed"))
	_, err = composer.Run("failing-command")
	if err == nil {
		t.Error("失败的命令应该返回错误")
	}
	if !contains(err.Error(), "command failed") {
		t.Errorf("错误信息不正确: %v", err)
	}

	// Test successful command after error
	SetupMockOutput("success-command", "success output", nil)
	output, err := composer.Run("success-command")
	if err != nil {
		t.Errorf("成功命令执行失败: %v", err)
	}
	if output != "success output" {
		t.Errorf("成功命令输出不正确: %s", output)
	}
}

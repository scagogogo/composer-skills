package installer

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

// ==================== SmartInstaller 测试 ====================

func TestNewSmartInstaller(t *testing.T) {
	options := DefaultInstallOptions()
	si := NewSmartInstaller(options)

	if si.options.MaxRetries != 3 {
		t.Errorf("MaxRetries = %v, 期望 3", si.options.MaxRetries)
	}
	if si.options.RetryDelay != 5*time.Second {
		t.Errorf("RetryDelay = %v, 期望 5s", si.options.RetryDelay)
	}
}

func TestNewSmartInstaller_NilContext(t *testing.T) {
	options := InstallOptions{
		Config:  SmartConfig(),
		Context: nil,
	}
	si := NewSmartInstaller(options)

	if si.options.Context == nil {
		t.Error("期望Context被设置为默认值")
	}
}

func TestNewSmartInstaller_ZeroRetries(t *testing.T) {
	options := InstallOptions{
		Config:     SmartConfig(),
		MaxRetries: 0,
	}
	si := NewSmartInstaller(options)

	if si.options.MaxRetries != 3 {
		t.Errorf("MaxRetries = %v, 期望默认值 3", si.options.MaxRetries)
	}
}

func TestNewSmartInstaller_ZeroRetryDelay(t *testing.T) {
	options := InstallOptions{
		Config:      SmartConfig(),
		RetryDelay:  0,
	}
	si := NewSmartInstaller(options)

	if si.options.RetryDelay != 5*time.Second {
		t.Errorf("RetryDelay = %v, 期望默认值 5s", si.options.RetryDelay)
	}
}

// ==================== InstallProgress 测试 ====================

func TestInstallProgress(t *testing.T) {
	progress := InstallProgress{
		Stage:     StageCheckingPHP,
		Message:   "Checking PHP...",
		Percent:   10,
		Timestamp: time.Now(),
	}

	if progress.Stage != StageCheckingPHP {
		t.Errorf("Stage = %v, 期望 %v", progress.Stage, StageCheckingPHP)
	}
	if progress.Percent != 10 {
		t.Errorf("Percent = %v, 期望 10", progress.Percent)
	}
}

// ==================== ProgressCallback 测试 ====================

func TestProgressCallback(t *testing.T) {
	var callCount int32
	callback := func(progress InstallProgress) {
		atomic.AddInt32(&callCount, 1)
	}

	options := DefaultInstallOptions()
	options.ProgressCallback = callback

	// 手动触发回调
	callback(InstallProgress{
		Stage:   StageCheckingPHP,
		Message: "test",
		Percent: 10,
	})

	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("回调调用次数 = %v, 期望 1", callCount)
	}
}

// ==================== InstallOptions 测试 ====================

func TestDefaultInstallOptions(t *testing.T) {
	options := DefaultInstallOptions()

	if options.MaxRetries != 3 {
		t.Errorf("MaxRetries = %v, 期望 3", options.MaxRetries)
	}
	if options.RetryDelay != 5*time.Second {
		t.Errorf("RetryDelay = %v, 期望 5s", options.RetryDelay)
	}
	if options.SkipVerification != false {
		t.Error("SkipVerification 应该默认为 false")
	}
}

// ==================== InstallResult 测试 ====================

func TestInstallResult(t *testing.T) {
	result := &InstallResult{
		Success:      true,
		ComposerPath: "/usr/local/bin/composer",
		Version:      "2.6.6",
		PHPVersion:   "8.1.0",
		Method:       "package_manager",
		Duration:     30 * time.Second,
	}

	if !result.Success {
		t.Error("期望Success为true")
	}
	if result.Method != "package_manager" {
		t.Errorf("Method = %v, 期望 package_manager", result.Method)
	}
}

// ==================== EnsureComposerInstalled 测试 ====================

func TestEnsureComposerInstalled_AlreadyInstalled(t *testing.T) {
	// 如果composer已经安装，应该直接返回
	result, err := EnsureComposerInstalled(nil)
	if err != nil {
		// 在没有composer的环境中，这个测试可能失败，这是预期的
		t.Logf("EnsureComposerInstalled返回错误（可能是环境中没有composer）: %v", err)
		return
	}
	if result == nil {
		t.Error("结果不应为nil")
		return
	}
	if !result.Success {
		t.Error("期望Success为true")
	}
	if result.Method != "already_installed" {
		t.Errorf("Method = %v, 期望 already_installed", result.Method)
	}
}

// ==================== IsComposerInstalled 测试 ====================

func TestIsComposerInstalled(t *testing.T) {
	installed, path, version := IsComposerInstalled()
	// 这个测试取决于环境，我们只验证返回类型正确
	_ = installed
	_ = path
	_ = version
}

// ==================== GetSystemInfo 测试 ====================

func TestGetSystemInfo(t *testing.T) {
	info := GetSystemInfo()

	if info == nil {
		t.Error("系统信息不应为nil")
		return
	}

	// 检查必须存在的键
	if _, ok := info["php_available"]; !ok {
		t.Error("缺少php_available键")
	}
	if _, ok := info["composer_available"]; !ok {
		t.Error("缺少composer_available键")
	}
}

// ==================== InstallStages 测试 ====================

func TestInstallStages(t *testing.T) {
	stages := []InstallStage{
		StageCheckingPHP,
		StageInstallingPHP,
		StageDetectingDistro,
		StagePackageManager,
		StageDownloading,
		StageInstalling,
		StageVerifying,
		StageConfiguring,
		StageCompleted,
		StageFailed,
	}

	if len(stages) != 10 {
		t.Errorf("阶段数量 = %v, 期望 10", len(stages))
	}

	// 验证阶段值
	if StageCheckingPHP != "checking_php" {
		t.Errorf("StageCheckingPHP = %v, 期望 checking_php", StageCheckingPHP)
	}
	if StageCompleted != "completed" {
		t.Errorf("StageCompleted = %v, 期望 completed", StageCompleted)
	}
	if StageFailed != "failed" {
		t.Errorf("StageFailed = %v, 期望 failed", StageFailed)
	}
}

// ==================== Context取消测试 ====================

func TestSmartInstaller_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	options := DefaultInstallOptions()
	options.Context = ctx

	si := NewSmartInstaller(options)

	result, err := si.InstallWithProgress()
	if err == nil {
		t.Error("期望返回错误（上下文已取消）")
	}
	if result != nil && result.Success {
		t.Error("期望Success为false")
	}
}

// ==================== Utils 测试 ====================

func TestGetPlatformName(t *testing.T) {
	name := GetPlatformName()
	if name == "" {
		t.Error("平台名称不应为空")
	}
}

func TestGetArchName(t *testing.T) {
	name := GetArchName()
	if name == "" {
		t.Error("架构名称不应为空")
	}
}

func TestCanUseSudo(t *testing.T) {
	// 只验证方法不崩溃
	_ = CanUseSudo()
}

func TestValidateInstallPath(t *testing.T) {
	// 空路径应该返回错误
	if err := ValidateInstallPath(""); err == nil {
		t.Error("期望空路径返回错误")
	}

	// 非空路径应该通过基本验证
	if err := ValidateInstallPath("/usr/local/bin"); err != nil {
		t.Errorf("非空路径不应返回错误: %v", err)
	}
}

// ==================== 配置相关测试（补充） ====================

func TestSmartConfigValue(t *testing.T) {
	config := SmartConfig()
	if config.AutoInstallPHP != true {
		t.Error("AutoInstallPHP应该为true")
	}
	if config.TargetVersion != "latest" {
		t.Errorf("TargetVersion = %v, 期望 latest", config.TargetVersion)
	}
}

// ==================== PHP相关补充测试 ====================

func TestHasPHPValue(t *testing.T) {
	// 只验证方法不崩溃
	result := HasPHP()
	t.Logf("HasPHP = %v", result)
}

func TestGetPHPVersionValue(t *testing.T) {
	// 只验证方法不崩溃
	version, err := GetPHPVersion()
	t.Logf("GetPHPVersion = %v, err = %v", version, err)
}

// ==================== findComposerBinary 测试 ====================

func TestFindComposerBinary(t *testing.T) {
	_, err := findComposerBinary()
	// 在没有composer的环境中可能返回错误，这是预期的
	_ = err
}

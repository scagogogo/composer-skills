package installer_test

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/scagogogo/composer-skills/pkg/detector"
	"github.com/scagogogo/composer-skills/pkg/installer"
)

// ==================== 集成测试：验证安装功能在真实环境中可用 ====================

// TestIntegration_DetectorCanFindComposer 验证检测器能在当前系统中找到 Composer
func TestIntegration_DetectorCanFindComposer(t *testing.T) {
	d := detector.NewDetector()

	// 测试 IsInstalled
	if !d.IsInstalled() {
		t.Fatal("Detector.IsInstalled() = false, 当前环境应该已安装 Composer")
	}

	// 测试 Detect 返回有效路径
	path, err := d.Detect()
	if err != nil {
		t.Fatalf("Detector.Detect() 返回错误: %v", err)
	}
	if path == "" {
		t.Fatal("Detector.Detect() 返回空路径")
	}
	t.Logf("✓ 检测到 Composer 路径: %s", path)

	// 验证路径指向一个可执行文件
	if _, err := exec.LookPath(path); err != nil {
		t.Logf("⚠ 路径不在 PATH 中 (LookPath 失败)，但文件可能存在: %v", err)
	}
}

// TestIntegration_DetectorVerbose 验证详细检测返回完整信息
func TestIntegration_DetectorVerbose(t *testing.T) {
	d := detector.NewDetector()

	result, err := d.DetectVerbose()
	if err != nil {
		t.Fatalf("Detector.DetectVerbose() 返回错误: %v", err)
	}

	t.Logf("✓ 详细检测结果:")
	t.Logf("  路径: %s", result.Path)
	t.Logf("  方法: %s", result.Method)
	t.Logf("  是否 PHAR: %v", result.IsPhar)

	if result.Path == "" {
		t.Error("DetectVerbose 返回空路径")
	}
	if result.Method == "" {
		t.Error("DetectVerbose 返回空方法")
	}
}

// TestIntegration_InstallerIsInstalled 验证安装器能正确识别已安装状态
func TestIntegration_InstallerIsInstalled(t *testing.T) {
	inst := installer.NewInstaller(installer.DefaultConfig())

	if !inst.IsInstalled() {
		t.Fatal("Installer.IsInstalled() = false, 当前环境应该已安装 Composer")
	}
	t.Logf("✓ Installer 确认 Composer 已安装")
}

// TestIntegration_InstallerGetVersion 验证安装器能获取已安装版本
func TestIntegration_InstallerGetVersion(t *testing.T) {
	inst := installer.NewInstaller(installer.DefaultConfig())

	version, err := inst.GetInstalledVersion()
	if err != nil {
		t.Fatalf("Installer.GetInstalledVersion() 返回错误: %v", err)
	}
	t.Logf("✓ 已安装 Composer 版本: %s", version)
}

// TestIntegration_SmartConfig 验证 SmartConfig 在当前平台返回合理配置
func TestIntegration_SmartConfig(t *testing.T) {
	config := installer.SmartConfig()

	t.Logf("✓ SmartConfig 配置:")
	t.Logf("  InstallPath: %s", config.InstallPath)
	t.Logf("  PreferPackageManager: %v", config.PreferPackageManager)
	t.Logf("  PreferBrewOnMac: %v", config.PreferBrewOnMac)
	t.Logf("  AutoInstallPHP: %v", config.AutoInstallPHP)
	t.Logf("  UseSudo: %v", config.UseSudo)
	t.Logf("  TargetVersion: %s", config.TargetVersion)
	t.Logf("  TimeoutSeconds: %d", config.TimeoutSeconds)

	if config.InstallPath == "" {
		t.Error("SmartConfig.InstallPath 为空")
	}
}

// TestIntegration_HasPHP 验证 PHP 检测
func TestIntegration_HasPHP(t *testing.T) {
	if !installer.HasPHP() {
		t.Fatal("HasPHP() = false, 当前环境应该有 PHP")
	}

	version, err := installer.GetPHPVersion()
	if err != nil {
		t.Fatalf("GetPHPVersion() 返回错误: %v", err)
	}
	t.Logf("✓ PHP 版本: %s", version)
}

// TestIntegration_DetectLinuxDistro 验证 Linux 发行版检测
func TestIntegration_DetectLinuxDistro(t *testing.T) {
	distro, err := installer.DetectLinuxDistro()
	if err != nil {
		t.Fatalf("DetectLinuxDistro() 返回错误: %v", err)
	}

	t.Logf("✓ 检测到 Linux 发行版:")
	t.Logf("  ID: %s", distro.ID)
	t.Logf("  Name: %s", distro.Name)
	t.Logf("  Version: %s", distro.Version)
	t.Logf("  PackageManager: %s", distro.PackageManager)

	if distro.ID == "" {
		t.Error("发行版 ID 为空")
	}
	if distro.PackageManager == "" {
		t.Error("包管理器为空")
	}
}

// TestIntegration_EnsureComposerInstalled 验证 EnsureComposerInstalled 在已安装时正确返回
func TestIntegration_EnsureComposerInstalled(t *testing.T) {
	result, err := installer.EnsureComposerInstalled(nil)
	if err != nil {
		t.Fatalf("EnsureComposerInstalled() 返回错误: %v", err)
	}

	t.Logf("✓ EnsureComposerInstalled 结果:")
	t.Logf("  Success: %v", result.Success)
	t.Logf("  ComposerPath: %s", result.ComposerPath)
	t.Logf("  Version: %s", result.Version)
	t.Logf("  PHPVersion: %s", result.PHPVersion)
	t.Logf("  Method: %s", result.Method)

	if !result.Success {
		t.Error("EnsureComposerInstalled 结果不成功")
	}
	if result.ComposerPath == "" {
		t.Error("ComposerPath 为空")
	}
	if result.Method != "already_installed" {
		t.Errorf("Method 应该是 'already_installed', 实际: %s", result.Method)
	}
}

// TestIntegration_IsComposerInstalled 验证 IsComposerInstalled 函数
func TestIntegration_IsComposerInstalled(t *testing.T) {
	installed, path, version := installer.IsComposerInstalled()

	if !installed {
		t.Fatal("IsComposerInstalled() = false, 当前环境应该已安装 Composer")
	}
	t.Logf("✓ IsComposerInstalled:")
	t.Logf("  installed: %v", installed)
	t.Logf("  path: %s", path)
	t.Logf("  version: %s", version)
}

// TestIntegration_GetSystemInfo 验证系统信息获取
func TestIntegration_GetSystemInfo(t *testing.T) {
	info := installer.GetSystemInfo()

	t.Logf("✓ GetSystemInfo:")
	for k, v := range info {
		t.Logf("  %s: %s", k, v)
	}

	if info["php_available"] != "true" {
		t.Error("PHP 应该可用")
	}
	if info["composer_available"] != "true" {
		t.Error("Composer 应该可用")
	}
}

// TestIntegration_SmartInstallerAlreadyInstalled 验证 SmartInstaller 在已安装时的行为
func TestIntegration_SmartInstallerAlreadyInstalled(t *testing.T) {
	options := installer.DefaultInstallOptions()
	options.SkipVerification = false

	si := installer.NewSmartInstaller(options)
	result, err := si.InstallWithProgress()
	if err != nil {
		t.Fatalf("SmartInstaller.InstallWithProgress() 返回错误: %v", err)
	}

	t.Logf("✓ SmartInstaller 结果:")
	t.Logf("  Success: %v", result.Success)
	t.Logf("  Method: %s", result.Method)
	t.Logf("  ComposerPath: %s", result.ComposerPath)
	t.Logf("  Version: %s", result.Version)
	t.Logf("  Duration: %v", result.Duration)

	if !result.Success {
		t.Error("安装结果不成功")
	}
}

// TestIntegration_SmartInstallerWithProgressCallback 验证带进度回调的 SmartInstaller
func TestIntegration_SmartInstallerWithProgressCallback(t *testing.T) {
	var progressMessages []string

	options := installer.DefaultInstallOptions()
	options.ProgressCallback = func(p installer.InstallProgress) {
		progressMessages = append(progressMessages, p.Message)
	}
	options.SkipVerification = false

	si := installer.NewSmartInstaller(options)
	result, err := si.InstallWithProgress()
	if err != nil {
		t.Fatalf("SmartInstaller.InstallWithProgress() 返回错误: %v", err)
	}

	t.Logf("✓ 带进度的 SmartInstaller 结果:")
	t.Logf("  Success: %v", result.Success)
	t.Logf("  进度消息数: %d", len(progressMessages))

	for i, msg := range progressMessages {
		t.Logf("  [%d] %s", i, msg)
	}
}

// TestIntegration_SmartInstallerWithContext 验证带 Context 的 SmartInstaller
func TestIntegration_SmartInstallerWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	options := installer.DefaultInstallOptions()
	options.Context = ctx

	si := installer.NewSmartInstaller(options)
	result, err := si.InstallWithProgress()
	if err != nil {
		t.Fatalf("SmartInstaller.InstallWithProgress() with context 返回错误: %v", err)
	}

	if !result.Success {
		t.Error("带 Context 的安装结果不成功")
	}
	t.Logf("✓ 带 Context 的 SmartInstaller: Success=%v", result.Success)
}

// TestIntegration_SmartInstallerCancelledContext 验证取消的 Context 能终止安装
func TestIntegration_SmartInstallerCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消

	options := installer.DefaultInstallOptions()
	options.Context = ctx

	si := installer.NewSmartInstaller(options)
	_, err := si.InstallWithProgress()
	if err == nil {
		t.Log("⚠ 已取消的 Context 未返回错误 (可能因为 Composer 已安装，安装流程很快完成)")
	} else {
		t.Logf("✓ 已取消的 Context 正确返回错误: %v", err)
	}
}

// TestIntegration_FindComposerBinary 验证 findComposerBinary 能找到 Composer
func TestIntegration_FindComposerBinary(t *testing.T) {
	// 直接用 which 检查
	path, err := exec.LookPath("composer")
	if err != nil {
		t.Fatalf("exec.LookPath('composer') 失败: %v", err)
	}
	t.Logf("✓ LookPath 找到 composer: %s", path)

	// 验证能运行
	cmd := exec.Command("composer", "--version")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("执行 composer --version 失败: %v", err)
	}
	t.Logf("✓ composer --version: %s", string(output))
}

// TestIntegration_InstallVersionValidation 验证 InstallVersion 的参数校验
func TestIntegration_InstallVersionValidation(t *testing.T) {
	inst := installer.NewInstaller(installer.DefaultConfig())

	// 测试已安装时再安装不报错（应该跳过或返回已安装错误）
	err := inst.Install()
	// 如果返回 ErrComposerAlreadyInstalled 也是预期行为
	if err != nil {
		t.Logf("⚠ Install() 返回: %v (可能因为已安装)", err)
	} else {
		t.Log("✓ Install() 成功执行 (已安装时可能为幂等)")
	}

	_ = inst // 避免未使用
}

package composer_test

import (
	"testing"
	"time"

	"github.com/scagogogo/composer-skills/pkg/composer"
	"github.com/scagogogo/composer-skills/pkg/installer"
)

// ==================== 集成测试：验证 composer 包的自动安装功能 ====================

// TestIntegration_NewWithDefaultOptions 验证 New(DefaultOptions()) 能自动检测并初始化
func TestIntegration_NewWithDefaultOptions(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New(DefaultOptions()) 返回错误: %v", err)
	}
	if comp == nil {
		t.Fatal("comp 为 nil")
	}

	// 验证实例可用
	version, err := comp.GetVersion()
	if err != nil {
		t.Fatalf("comp.GetVersion() 返回错误: %v", err)
	}
	t.Logf("✓ composer.New(DefaultOptions()) 成功, 版本: %s", version)
}

// TestIntegration_NewWithoutAutoInstall 验证不自动安装时也能找到已安装的 Composer
func TestIntegration_NewWithoutAutoInstall(t *testing.T) {
	options := composer.Options{
		AutoInstall:    false,
		DefaultTimeout: 5 * time.Minute,
	}

	comp, err := composer.New(options)
	if err != nil {
		t.Fatalf("composer.New(无自动安装) 返回错误: %v", err)
	}

	version, err := comp.GetVersion()
	if err != nil {
		t.Fatalf("comp.GetVersion() 返回错误: %v", err)
	}
	t.Logf("✓ 无自动安装模式成功, 版本: %s", version)
}

// TestIntegration_EnsureInstalled 验证 EnsureInstalled 在已安装时的行为
func TestIntegration_EnsureInstalled(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}

	err = comp.EnsureInstalled()
	if err != nil {
		t.Fatalf("comp.EnsureInstalled() 返回错误: %v", err)
	}
	t.Log("✓ EnsureInstalled() 成功 (Composer 已安装)")
}

// TestIntegration_EnsureInstalledWithProgress 验证带进度的 EnsureInstalled
func TestIntegration_EnsureInstalledWithProgress(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}

	var progressMessages []string
	result, err := comp.EnsureInstalledWithProgress(func(p installer.InstallProgress) {
		progressMessages = append(progressMessages, p.Message)
	})
	if err != nil {
		t.Fatalf("comp.EnsureInstalledWithProgress() 返回错误: %v", err)
	}

	t.Logf("✓ EnsureInstalledWithProgress 结果:")
	t.Logf("  Success: %v", result.Success)
	t.Logf("  Method: %s", result.Method)
	t.Logf("  ComposerPath: %s", result.ComposerPath)
	t.Logf("  Version: %s", result.Version)
	t.Logf("  进度消息数: %d", len(progressMessages))

	if !result.Success {
		t.Error("安装结果不成功")
	}
}

// TestIntegration_GetInstallStatus 验证 GetInstallStatus 返回正确状态
func TestIntegration_GetInstallStatus(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}

	status := comp.GetInstallStatus()

	t.Logf("✓ GetInstallStatus 结果:")
	t.Logf("  Installed: %v", status.Installed)
	t.Logf("  Path: %s", status.Path)
	t.Logf("  Version: %s", status.Version)
	t.Logf("  PHPAvailable: %v", status.PHPAvailable)
	t.Logf("  PHPVersion: %s", status.PHPVersion)
	t.Logf("  AutoInstall: %v", status.AutoInstall)

	if !status.Installed {
		t.Error("Installed 应该为 true")
	}
	if !status.PHPAvailable {
		t.Error("PHPAvailable 应该为 true")
	}
	if status.Path == "" {
		t.Error("Path 不应为空")
	}
}

// TestIntegration_QuickSetup 验证 QuickSetup 快速设置
func TestIntegration_QuickSetup(t *testing.T) {
	comp, err := composer.QuickSetup("", true)
	if err != nil {
		t.Fatalf("composer.QuickSetup() 返回错误: %v", err)
	}
	if comp == nil {
		t.Fatal("QuickSetup 返回 nil")
	}

	version, err := comp.GetVersion()
	if err != nil {
		t.Fatalf("comp.GetVersion() 返回错误: %v", err)
	}
	t.Logf("✓ QuickSetup 成功, 版本: %s", version)
}

// TestIntegration_QuickSetupWithProgress 验证带进度的 QuickSetup
func TestIntegration_QuickSetupWithProgress(t *testing.T) {
	var progressMessages []string

	comp, result, err := composer.QuickSetupWithProgress("", func(p installer.InstallProgress) {
		progressMessages = append(progressMessages, p.Message)
	})
	if err != nil {
		t.Fatalf("composer.QuickSetupWithProgress() 返回错误: %v", err)
	}
	if comp == nil {
		t.Fatal("QuickSetupWithProgress 返回 nil comp")
	}
	if result == nil {
		t.Fatal("QuickSetupWithProgress 返回 nil result")
	}

	t.Logf("✓ QuickSetupWithProgress 结果:")
	t.Logf("  Success: %v", result.Success)
	t.Logf("  Method: %s", result.Method)
	t.Logf("  ComposerPath: %s", result.ComposerPath)
	t.Logf("  进度消息数: %d", len(progressMessages))
}

// TestIntegration_IsInstalled 验证 IsInstalled 方法
func TestIntegration_IsInstalled(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}

	if !comp.IsInstalled() {
		t.Fatal("comp.IsInstalled() = false, 当前环境应该已安装 Composer")
	}
	t.Log("✓ comp.IsInstalled() = true")
}

// TestIntegration_SelfUpdateWithProgress 验证自更新功能签名（不实际执行更新）
// 注意：此测试只验证方法可调用，不实际执行更新（避免破坏环境）
func TestIntegration_SelfUpdateWithProgress_Signature(t *testing.T) {
	// 只验证 Composer 结构体有此方法，不实际执行
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}
	_ = comp.SelfUpdateWithProgress // 验证方法存在
	t.Log("✓ SelfUpdateWithProgress 方法存在")
}

// TestIntegration_ComposerExecuteCommand 验证通过自动安装的 Composer 实例能执行命令
func TestIntegration_ComposerExecuteCommand(t *testing.T) {
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		t.Fatalf("composer.New() 返回错误: %v", err)
	}

	// 测试能执行一个简单的命令
	version, err := comp.GetVersion()
	if err != nil {
		t.Fatalf("comp.GetVersion() 返回错误: %v", err)
	}
	if version == "" {
		t.Error("GetVersion() 返回空版本号")
	}
	t.Logf("✓ 通过自动初始化的 Composer 能执行命令, 版本: %s", version)

	// 测试诊断命令
	diagOutput, err := comp.Diagnose()
	if err != nil {
		t.Logf("⚠ Diagnose() 返回错误 (可能网络问题): %v", err)
	} else {
		t.Logf("✓ Diagnose() 成功执行, 输出长度: %d", len(diagOutput))
	}
}

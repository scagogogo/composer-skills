package composer

import (
	"fmt"
	"time"

	"github.com/scagogogo/composer-skills/pkg/installer"
)

// ==================== 自动安装增强 ====================

// EnsureInstalled 确保Composer已安装，如果未安装则自动安装
//
// 返回值：
//   - error: 如果安装或检测失败，则返回错误
//
// 功能说明：
//
//	该方法检查Composer是否已安装。如果未安装且autoInstall为true，
//	则自动安装Composer。安装完成后会验证安装是否成功。
//	这是最便捷的初始化方式，适合在程序启动时调用。
//
// 用法示例：
//
//	options := composer.DefaultOptions()
//	options.AutoInstall = true
//	comp, err := composer.New(options)
//	if err != nil {
//	    log.Fatalf("初始化失败: %v", err)
//	}
//	// 确保Composer已安装
//	if err := comp.EnsureInstalled(); err != nil {
//	    log.Fatalf("确保安装失败: %v", err)
//	}
func (c *Composer) EnsureInstalled() error {
	if c.IsInstalled() {
		return nil
	}

	// 尝试检测
	execPath, err := c.detector.Detect()
	if err == nil {
		c.executablePath = execPath
		return nil
	}

	// 如果autoInstall为true，尝试安装
	if c.autoInstall {
		if err := c.installer.Install(); err != nil {
			return fmt.Errorf("%w: %v", ErrComposerInstallation, err)
		}

		// 重新检测
		execPath, err = c.detector.Detect()
		if err != nil {
			return fmt.Errorf("%w: 安装后仍无法检测到Composer", ErrComposerNotFound)
		}
		c.executablePath = execPath
		return nil
	}

	return ErrComposerNotFound
}

// EnsureInstalledWithProgress 确保Composer已安装，带进度回调
//
// 参数：
//   - callback: 安装进度回调函数
//
// 返回值：
//   - *installer.InstallResult: 安装结果
//   - error: 安装或检测错误
//
// 功能说明：
//
//	该方法与EnsureInstalled类似，但使用智能安装器，支持进度回调、
//	自动重试和上下文取消。适合需要显示安装进度或需要更精细控制的场景。
//
// 用法示例：
//
//	comp, _ := composer.New(composer.DefaultOptions())
//	result, err := comp.EnsureInstalledWithProgress(func(p installer.InstallProgress) {
//	    fmt.Printf("[%s] %s (%d%%)\n", p.Stage, p.Message, p.Percent)
//	})
//	if err != nil {
//	    log.Fatalf("安装失败: %v", err)
//	}
//	fmt.Printf("Composer %s 已安装\n", result.Version)
func (c *Composer) EnsureInstalledWithProgress(callback installer.ProgressCallback) (*installer.InstallResult, error) {
	// 先检查是否已安装
	if c.IsInstalled() {
		return &installer.InstallResult{
			Success:      true,
			ComposerPath: c.executablePath,
			Method:       "already_installed",
		}, nil
	}

	// 检测是否已安装
	execPath, err := c.detector.Detect()
	if err == nil {
		c.executablePath = execPath
		return &installer.InstallResult{
			Success:      true,
			ComposerPath: execPath,
			Method:       "already_installed",
		}, nil
	}

	// 使用智能安装器安装
	options := installer.DefaultInstallOptions()
	options.ProgressCallback = callback
	options.Config = c.installer.GetConfig()

	smartInstaller := installer.NewSmartInstaller(options)
	result, err := smartInstaller.InstallWithProgress()
	if err != nil {
		return result, err
	}

	// 更新可执行文件路径
	if result.Success {
		execPath, err = c.detector.Detect()
		if err != nil {
			return result, fmt.Errorf("安装成功但无法检测到Composer: %w", err)
		}
		c.executablePath = execPath
	}

	return result, nil
}

// SelfUpdateWithProgress 带进度报告的自更新
//
// 返回值：
//   - string: 更新后的版本号
//   - error: 更新错误
//
// 功能说明：
//
//	该方法执行Composer自更新，并返回更新后的版本号。
//
// 用法示例：
//
//	version, err := comp.SelfUpdateWithProgress()
//	if err != nil {
//	    log.Fatalf("更新失败: %v", err)
//	}
//	fmt.Printf("已更新到版本: %s\n", version)
func (c *Composer) SelfUpdateWithProgress() (string, error) {
	if err := c.SelfUpdate(); err != nil {
		return "", err
	}
	return c.GetVersion()
}

// InstallStatus 表示安装状态
type InstallStatus struct {
	// Installed 是否已安装
	Installed bool `json:"installed"`
	// Path Composer可执行文件路径
	Path string `json:"path,omitempty"`
	// Version Composer版本
	Version string `json:"version,omitempty"`
	// PHPAvailable PHP是否可用
	PHPAvailable bool `json:"php_available"`
	// PHPVersion PHP版本
	PHPVersion string `json:"php_version,omitempty"`
	// AutoInstall 是否启用自动安装
	AutoInstall bool `json:"auto_install"`
}

// GetInstallStatus 获取当前安装状态
//
// 返回值：
//   - *InstallStatus: 安装状态信息
//
// 功能说明：
//
//	该方法返回Composer的当前安装状态，包括是否已安装、版本号、
//	PHP是否可用等信息。适合在程序启动时检查环境状态。
//
// 用法示例：
//
//	status := comp.GetInstallStatus()
//	fmt.Printf("Composer已安装: %v\n", status.Installed)
//	fmt.Printf("版本: %s\n", status.Version)
//	fmt.Printf("PHP可用: %v\n", status.PHPAvailable)
func (c *Composer) GetInstallStatus() *InstallStatus {
	status := &InstallStatus{
		Installed:   c.IsInstalled(),
		Path:        c.executablePath,
		AutoInstall: c.autoInstall,
	}

	// 获取版本
	if status.Installed {
		if version, err := c.GetVersion(); err == nil {
			status.Version = version
		}
	}

	// 检查PHP
	status.PHPAvailable = installer.HasPHP()
	if status.PHPAvailable {
		if phpVer, err := installer.GetPHPVersion(); err == nil {
			status.PHPVersion = phpVer
		}
	}

	return status
}

// QuickSetup 快速设置Composer环境
//
// 参数：
//   - workingDir: 工作目录
//   - autoInstall: 是否自动安装
//
// 返回值：
//   - *Composer: Composer实例
//   - error: 设置错误
//
// 功能说明：
//
//	这是一个便捷方法，一步完成Composer的检测、安装和初始化。
//	适合在程序启动时快速设置Composer环境。
//
// 用法示例：
//
//	comp, err := composer.QuickSetup("/path/to/project", true)
//	if err != nil {
//	    log.Fatalf("快速设置失败: %v", err)
//	}
//	// 现在可以使用comp执行各种Composer操作
func QuickSetup(workingDir string, autoInstall bool) (*Composer, error) {
	options := Options{
		WorkingDir:     workingDir,
		AutoInstall:    autoInstall,
		DefaultTimeout: 10 * time.Minute,
	}

	comp, err := New(options)
	if err != nil {
		return nil, err
	}

	if autoInstall {
		if err := comp.EnsureInstalled(); err != nil {
			return nil, err
		}
	}

	return comp, nil
}

// QuickSetupWithProgress 快速设置Composer环境，带进度回调
//
// 参数：
//   - workingDir: 工作目录
//   - callback: 安装进度回调
//
// 返回值：
//   - *Composer: Composer实例
//   - *installer.InstallResult: 安装结果
//   - error: 设置错误
func QuickSetupWithProgress(workingDir string, callback installer.ProgressCallback) (*Composer, *installer.InstallResult, error) {
	options := Options{
		WorkingDir:     workingDir,
		AutoInstall:    true,
		DefaultTimeout: 10 * time.Minute,
	}

	comp, err := New(options)
	if err != nil {
		// 即使New失败，也尝试安装
		if options.AutoInstall {
			result, installErr := installer.EnsureComposerInstalled(nil)
			if installErr != nil {
				return nil, result, installErr
			}
			// 安装成功后重新创建
			comp, err = New(options)
			if err != nil {
				return nil, result, err
			}
			return comp, result, nil
		}
		return nil, nil, err
	}

	result, err := comp.EnsureInstalledWithProgress(callback)
	if err != nil {
		return nil, result, err
	}

	return comp, result, nil
}

package installer

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// InstallProgress 表示安装进度
type InstallProgress struct {
	// Stage 当前安装阶段
	Stage InstallStage `json:"stage"`
	// Message 进度消息
	Message string `json:"message"`
	// Percent 进度百分比 (0-100)
	Percent int `json:"percent"`
	// Error 安装过程中的错误（如果有）
	Error error `json:"error,omitempty"`
	// Timestamp 进度时间戳
	Timestamp time.Time `json:"timestamp"`
}

// InstallStage 表示安装阶段
type InstallStage string

const (
	StageCheckingPHP      InstallStage = "checking_php"
	StageInstallingPHP    InstallStage = "installing_php"
	StageDetectingDistro  InstallStage = "detecting_distro"
	StagePackageManager   InstallStage = "package_manager"
	StageDownloading      InstallStage = "downloading"
	StageInstalling       InstallStage = "installing"
	StageVerifying        InstallStage = "verifying"
	StageConfiguring      InstallStage = "configuring"
	StageCompleted        InstallStage = "completed"
	StageFailed           InstallStage = "failed"
)

// ProgressCallback 是安装进度回调函数类型
//
// 参数：
//   - progress: 当前进度信息
//
// 功能说明：
//
//	当安装过程中各个阶段发生变化时，会调用此回调函数。
//	可以用于显示进度条、记录日志或通知用户。
type ProgressCallback func(progress InstallProgress)

// InstallOptions 表示安装选项
type InstallOptions struct {
	// Config 安装器配置
	Config Config
	// MaxRetries 最大重试次数
	MaxRetries int
	// RetryDelay 重试间隔
	RetryDelay time.Duration
	// ProgressCallback 进度回调
	ProgressCallback ProgressCallback
	// Context 上下文，用于取消安装
	Context context.Context
	// SkipVerification 是否跳过安装后验证
	SkipVerification bool
	// OnPHPInstalled PHP安装完成后的回调
	OnPHPInstalled func()
}

// DefaultInstallOptions 返回默认安装选项
func DefaultInstallOptions() InstallOptions {
	return InstallOptions{
		Config:            SmartConfig(),
		MaxRetries:        3,
		RetryDelay:        5 * time.Second,
		ProgressCallback:  nil,
		Context:           context.Background(),
		SkipVerification:  false,
	}
}

// InstallResult 表示安装结果
type InstallResult struct {
	// Success 是否安装成功
	Success bool `json:"success"`
	// ComposerPath Composer可执行文件路径
	ComposerPath string `json:"composer_path,omitempty"`
	// Version 安装的Composer版本
	Version string `json:"version,omitempty"`
	// PHPVersion PHP版本
	PHPVersion string `json:"php_version,omitempty"`
	// Method 安装方法
	Method string `json:"method,omitempty"` // "package_manager", "homebrew", "direct_download"
	// Duration 安装耗时
	Duration time.Duration `json:"duration,omitempty"`
	// Error 安装错误
	Error error `json:"error,omitempty"`
	// Stages 各阶段耗时
	Stages map[InstallStage]time.Duration `json:"stages,omitempty"`
}

// SmartInstaller 提供智能安装功能，包括重试、进度回调和上下文取消
type SmartInstaller struct {
	options InstallOptions
	mu      sync.Mutex
}

// NewSmartInstaller 创建智能安装器
//
// 参数：
//   - options: 安装选项
//
// 返回值：
//   - *SmartInstaller: 智能安装器实例
func NewSmartInstaller(options InstallOptions) *SmartInstaller {
	if options.Config.InstallPath == "" {
		options.Config = SmartConfig()
	}
	if options.MaxRetries <= 0 {
		options.MaxRetries = 3
	}
	if options.RetryDelay <= 0 {
		options.RetryDelay = 5 * time.Second
	}
	if options.Context == nil {
		options.Context = context.Background()
	}
	return &SmartInstaller{
		options: options,
	}
}

// reportProgress 报告安装进度
func (si *SmartInstaller) reportProgress(stage InstallStage, message string, percent int) {
	if si.options.ProgressCallback != nil {
		si.options.ProgressCallback(InstallProgress{
			Stage:     stage,
			Message:   message,
			Percent:   percent,
			Timestamp: time.Now(),
		})
	}
}

// reportError 报告安装错误
func (si *SmartInstaller) reportError(stage InstallStage, message string, err error) {
	if si.options.ProgressCallback != nil {
		si.options.ProgressCallback(InstallProgress{
			Stage:     stage,
			Message:   message,
			Percent:   0,
			Error:     err,
			Timestamp: time.Now(),
		})
	}
}

// InstallWithProgress 执行带进度报告的安装
//
// 返回值：
//   - *InstallResult: 安装结果
//   - error: 安装错误
//
// 功能说明：
//
//	该方法执行Composer的智能安装，支持：
//	- 进度回调：通过ProgressCallback获取安装进度
//	- 自动重试：安装失败时自动重试
//	- 上下文取消：通过Context取消安装
//	- PHP自动安装：如果未安装PHP，自动安装
//	- 安装后验证：验证安装是否成功
//
// 用法示例：
//
//	options := installer.DefaultInstallOptions()
//	options.ProgressCallback = func(p installer.InstallProgress) {
//	    fmt.Printf("[%s] %s (%d%%)\n", p.Stage, p.Message, p.Percent)
//	}
//	si := installer.NewSmartInstaller(options)
//	result, err := si.InstallWithProgress()
//	if err != nil {
//	    log.Fatalf("安装失败: %v", err)
//	}
//	fmt.Printf("Composer %s 安装成功，路径: %s\n", result.Version, result.ComposerPath)
func (si *SmartInstaller) InstallWithProgress() (*InstallResult, error) {
	startTime := time.Now()
	result := &InstallResult{
		Stages: make(map[InstallStage]time.Duration),
	}

	// 检查上下文是否已取消
	if err := si.checkContext(); err != nil {
		result.Error = err
		return result, err
	}

	// 阶段1: 检查PHP
	stageStart := time.Now()
	si.reportProgress(StageCheckingPHP, "检查PHP环境...", 5)

	if !HasPHP() {
		if si.options.Config.AutoInstallPHP {
			si.reportProgress(StageInstallingPHP, "自动安装PHP...", 10)
			distro, err := DetectLinuxDistro()
			if err != nil || distro == nil || distro.PackageManager == "unknown" {
				err = fmt.Errorf("无法自动安装PHP: 未知Linux发行版")
				si.reportError(StageInstallingPHP, "PHP自动安装失败", err)
				result.Error = err
				result.Success = false
				return result, err
			}
			if err := InstallPHP(distro, si.options.Config.UseSudo); err != nil {
				si.reportError(StageInstallingPHP, "PHP自动安装失败", err)
				result.Error = fmt.Errorf("%w: 自动安装PHP失败: %v", ErrPHPNotFound, err)
				result.Success = false
				return result, err
			}
			if !HasPHP() {
				err = fmt.Errorf("PHP安装后仍无法检测到")
				si.reportError(StageInstallingPHP, "PHP安装验证失败", err)
				result.Error = err
				result.Success = false
				return result, err
			}
			if si.options.OnPHPInstalled != nil {
				si.options.OnPHPInstalled()
			}
		} else {
			err := ErrPHPNotFound
			si.reportError(StageCheckingPHP, "未找到PHP", err)
			result.Error = err
			result.Success = false
			return result, err
		}
	}

	// 获取PHP版本
	if phpVer, err := GetPHPVersion(); err == nil {
		result.PHPVersion = phpVer
	}
	result.Stages[StageCheckingPHP] = time.Since(stageStart)

	// 阶段2: 尝试安装Composer（带重试）
	var lastErr error
	for attempt := 1; attempt <= si.options.MaxRetries; attempt++ {
		if err := si.checkContext(); err != nil {
			result.Error = err
			return result, err
		}

		si.reportProgress(StageDetectingDistro, fmt.Sprintf("检测系统环境 (尝试 %d/%d)...", attempt, si.options.MaxRetries), 20+attempt*5)

		installMethod, err := si.doInstall(attempt)
		if err == nil {
			result.Method = installMethod
			break
		}

		lastErr = err
		if attempt < si.options.MaxRetries {
			si.reportProgress(StageFailed, fmt.Sprintf("安装失败，%v后重试 (尝试 %d/%d)...", si.options.RetryDelay, attempt, si.options.MaxRetries), 0)

			// 等待重试
			select {
			case <-si.options.Context.Done():
				result.Error = si.options.Context.Err()
				return result, result.Error
			case <-time.After(si.options.RetryDelay):
				// 继续重试
			}
		}
	}

	if lastErr != nil {
		si.reportError(StageFailed, "所有安装尝试均失败", lastErr)
		result.Error = lastErr
		result.Success = false
		result.Duration = time.Since(startTime)
		return result, lastErr
	}

	// 阶段3: 验证安装
	if !si.options.SkipVerification {
		stageStart = time.Now()
		si.reportProgress(StageVerifying, "验证安装...", 90)

		composerPath, err := findComposerBinary()
		if err != nil {
			si.reportError(StageVerifying, "安装验证失败: 无法找到composer", err)
			result.Error = err
			result.Success = false
			result.Duration = time.Since(startTime)
			return result, err
		}
		result.ComposerPath = composerPath

		// 获取版本
		if version, err := CheckComposerVersion(composerPath); err == nil {
			result.Version = version
		}

		result.Stages[StageVerifying] = time.Since(stageStart)
	}

	result.Success = true
	result.Duration = time.Since(startTime)
	si.reportProgress(StageCompleted, "安装完成!", 100)

	return result, nil
}

// doInstall 执行实际的安装操作
func (si *SmartInstaller) doInstall(attempt int) (string, error) {
	// 尝试包管理器安装
	si.reportProgress(StagePackageManager, "尝试通过包管理器安装...", 30)

	installer := NewInstaller(si.options.Config)

	// 在Linux上优先尝试包管理器
	if si.options.Config.PreferPackageManager {
		distro, err := DetectLinuxDistro()
		if err == nil && distro != nil && distro.PackageManager != "unknown" {
			attempted, err := InstallComposerViaPackageManager(distro, si.options.Config.UseSudo)
			if attempted && err == nil {
				if installer.IsInstalled() {
					return "package_manager", nil
				}
			}
		}
	}

	// 尝试Homebrew (macOS)
	if si.options.Config.PreferBrewOnMac {
		si.reportProgress(StagePackageManager, "尝试通过Homebrew安装...", 40)
		if method := tryBrewInstall(); method != "" {
			if installer.IsInstalled() {
				return method, nil
			}
		}
	}

	// 直接下载安装
	si.reportProgress(StageDownloading, "下载Composer安装程序...", 50)
	if err := installer.Install(); err != nil {
		return "", err
	}

	si.reportProgress(StageInstalling, "安装Composer...", 70)
	return "direct_download", nil
}

// tryBrewInstall 尝试通过Homebrew安装
func tryBrewInstall() string {
	// 检查brew是否可用
	if _, err := findBinary("brew"); err != nil {
		return ""
	}
	// 尝试安装
	cmd := findCommand("brew", "install", "composer")
	if output, err := cmd.CombinedOutput(); err != nil {
		_ = output
		return ""
	}
	return "homebrew"
}

// checkContext 检查上下文是否已取消
func (si *SmartInstaller) checkContext() error {
	if si.options.Context == nil {
		return nil
	}
	select {
	case <-si.options.Context.Done():
		return si.options.Context.Err()
	default:
		return nil
	}
}

// EnsureComposerInstalled 确保Composer已安装，如果未安装则自动安装
//
// 参数：
//   - options: 安装选项（可选，为nil则使用默认选项）
//
// 返回值：
//   - *InstallResult: 安装结果
//   - error: 错误信息
//
// 功能说明：
//
//	这是一个便捷方法，首先检查Composer是否已安装。
//	如果已安装，直接返回结果；如果未安装，则执行智能安装。
//
// 用法示例：
//
//	result, err := installer.EnsureComposerInstalled(nil)
//	if err != nil {
//	    log.Fatalf("确保Composer安装失败: %v", err)
//	}
//	fmt.Printf("Composer已就绪: %s\n", result.ComposerPath)
func EnsureComposerInstalled(options *InstallOptions) (*InstallResult, error) {
	// 先检查是否已安装
	if path, err := findComposerBinary(); err == nil {
		result := &InstallResult{
			Success:      true,
			ComposerPath: path,
			Method:       "already_installed",
		}
		if version, err := CheckComposerVersion(path); err == nil {
			result.Version = version
		}
		if phpVer, err := GetPHPVersion(); err == nil {
			result.PHPVersion = phpVer
		}
		return result, nil
	}

	// 未安装，执行安装
	opts := DefaultInstallOptions()
	if options != nil {
		opts = *options
	}

	si := NewSmartInstaller(opts)
	return si.InstallWithProgress()
}

// IsComposerInstalled 检查Composer是否已安装
//
// 返回值：
//   - bool: 是否已安装
//   - string: 安装路径（如果已安装）
//   - string: 版本号（如果已安装）
func IsComposerInstalled() (bool, string, string) {
	path, err := findComposerBinary()
	if err != nil {
		return false, "", ""
	}
	version, _ := CheckComposerVersion(path)
	return true, path, version
}

// GetSystemInfo 获取系统信息（用于诊断安装问题）
//
// 返回值：
//   - map[string]string: 系统信息
func GetSystemInfo() map[string]string {
	info := make(map[string]string)

	// PHP信息
	if HasPHP() {
		info["php_available"] = "true"
		if version, err := GetPHPVersion(); err == nil {
			info["php_version"] = version
		}
	} else {
		info["php_available"] = "false"
	}

	// Composer信息
	if installed, path, version := IsComposerInstalled(); installed {
		info["composer_available"] = "true"
		info["composer_path"] = path
		info["composer_version"] = version
	} else {
		info["composer_available"] = "false"
	}

	// 包管理器信息
	distro, err := DetectLinuxDistro()
	if err == nil && distro != nil {
		info["distro_id"] = distro.ID
		info["distro_name"] = distro.Name
		info["distro_version"] = distro.Version
		info["package_manager"] = distro.PackageManager
	}

	// Homebrew信息
	if _, err := findBinary("brew"); err == nil {
		info["brew_available"] = "true"
	} else {
		info["brew_available"] = "false"
	}

	// curl信息
	if _, err := findBinary("curl"); err == nil {
		info["curl_available"] = "true"
	} else {
		info["curl_available"] = "false"
	}

	return info
}

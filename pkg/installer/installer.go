// Package installer provides the functionality to install and manage PHP Composer.
package installer

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

var (
	// ErrInstallationFailed indicates that the installation process failed
	ErrInstallationFailed = errors.New("安装失败")
	// ErrInsufficientRights indicates insufficient permissions
	ErrInsufficientRights = errors.New("权限不足，请使用管理员/sudo权限")
	// ErrUnsupportedPlatform indicates an unsupported platform
	ErrUnsupportedPlatform = errors.New("不支持的操作系统平台")
	// ErrDownloadFailed indicates a download failure
	ErrDownloadFailed = errors.New("下载失败")
	// ErrPHPNotFound indicates that PHP is not installed
	ErrPHPNotFound = errors.New("未找到PHP，Composer需要PHP才能运行")
	// ErrComposerAlreadyInstalled indicates Composer is already installed
	ErrComposerAlreadyInstalled = errors.New("Composer已经安装")
)

// Installer is responsible for installing Composer
type Installer struct {
	config Config
}

// NewInstaller creates a new installer instance
func NewInstaller(config Config) *Installer {
	return &Installer{
		config: config,
	}
}

// DefaultInstaller creates an installer with default configuration
func DefaultInstaller() *Installer {
	return &Installer{
		config: DefaultConfig(),
	}
}

// GetConfig returns the installer configuration
func (i *Installer) GetConfig() Config {
	return i.config
}

// SetConfig updates the installer configuration
func (i *Installer) SetConfig(config Config) {
	i.config = config
}

// Install installs Composer
func (i *Installer) Install() error {
	// Check if PHP is available first
	if !HasPHP() {
		if i.config.AutoInstallPHP {
			// Try to install PHP automatically
			distro, err := DetectLinuxDistro()
			if err != nil || distro == nil || distro.PackageManager == "unknown" {
				return fmt.Errorf("%w: 无法自动安装PHP，请手动安装PHP后再试", ErrPHPNotFound)
			}
			if err := InstallPHP(distro, i.config.UseSudo); err != nil {
				return fmt.Errorf("%w: 自动安装PHP失败: %v", ErrPHPNotFound, err)
			}
			// Verify PHP is now available
			if !HasPHP() {
				return fmt.Errorf("%w: PHP安装后仍无法检测到", ErrPHPNotFound)
			}
		} else {
			return ErrPHPNotFound
		}
	}

	// On Linux, try package manager first if configured
	if runtime.GOOS == "linux" && i.config.PreferPackageManager {
		distro, err := DetectLinuxDistro()
		if err == nil && distro != nil && distro.PackageManager != "unknown" {
			attempted, err := InstallComposerViaPackageManager(distro, i.config.UseSudo)
			if attempted && err == nil {
				// Verify installation
				if i.verifyInstallation() {
					return nil
				}
			}
			// If package manager install failed, fall through to direct install
		}
	}

	// Get platform-specific installer
	platformInstaller, err := GetPlatformInstaller(i.config)
	if err != nil {
		return err
	}

	// Execute installation
	return platformInstaller.Install()
}

// InstallVersion installs a specific version of Composer
// version can be "1", "2", "latest", "preview", or a specific version like "2.5.1"
func (i *Installer) InstallVersion(version string) error {
	// Check PHP first
	if !HasPHP() {
		return ErrPHPNotFound
	}

	return InstallComposerVersion(version, i.config.InstallPath, i.config.UseSudo)
}

// Uninstall removes Composer from the system
func (i *Installer) Uninstall() error {
	// Find composer path
	composerPath, err := findComposerBinary()
	if err != nil {
		return fmt.Errorf("cannot find composer to uninstall: %w", err)
	}

	return UninstallComposer(composerPath, i.config.UseSudo)
}

// IsInstalled checks if Composer is already installed
func (i *Installer) IsInstalled() bool {
	_, err := findComposerBinary()
	return err == nil
}

// GetInstalledVersion returns the installed Composer version
func (i *Installer) GetInstalledVersion() (string, error) {
	composerPath, err := findComposerBinary()
	if err != nil {
		return "", err
	}
	return CheckComposerVersion(composerPath)
}

// verifyInstallation checks that composer is accessible after install
func (i *Installer) verifyInstallation() bool {
	cmd := exec.Command("composer", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// findComposerBinary tries to locate the composer binary
func findComposerBinary() (string, error) {
	// Try PATH lookup first
	if path, err := exec.LookPath("composer"); err == nil {
		return path, nil
	}

	// Try common locations
	commonPaths := []string{
		"/usr/local/bin/composer",
		"/usr/bin/composer",
		"/opt/homebrew/bin/composer",
	}

	for _, p := range commonPaths {
		if _, err := exec.LookPath(p); err == nil {
			return p, nil
		}
	}

	return "", ErrComposerNotFound
}

var ErrComposerNotFound = errors.New("composer not found")

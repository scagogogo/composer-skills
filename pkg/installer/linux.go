package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composerutils"
)

// LinuxInstaller is the Composer installer for Linux
type LinuxInstaller struct {
	config Config
}

// NewLinuxInstaller creates an installer for Linux
func NewLinuxInstaller(config Config) *LinuxInstaller {
	return &LinuxInstaller{
		config: config,
	}
}

// Install installs Composer on Linux
func (i *LinuxInstaller) Install() error {
	// Try native package manager first if configured
	if i.config.PreferPackageManager {
		distro, err := DetectLinuxDistro()
		if err == nil && distro != nil && distro.PackageManager != "unknown" {
			attempted, err := InstallComposerViaPackageManager(distro, i.config.UseSudo)
			if attempted && err == nil {
				// Verify the installation worked
				if path, err := exec.LookPath("composer"); err == nil {
					_ = path
					return nil
				}
			}
			// If package manager failed, fall through to direct download
		}
	}

	return i.installDirect()
}

// installDirect downloads and installs Composer directly from getcomposer.org
func (i *LinuxInstaller) installDirect() error {
	// Check write permission on install directory
	if err := composerutils.CheckWritePermission(i.config.InstallPath); err != nil {
		if !i.config.UseSudo {
			return fmt.Errorf("%w, 目标路径: %s", ErrInsufficientRights, i.config.InstallPath)
		}
	}

	// Check if a specific version is requested
	if i.config.TargetVersion != "" && i.config.TargetVersion != "latest" {
		return InstallComposerVersion(i.config.TargetVersion, i.config.InstallPath, i.config.UseSudo)
	}

	// Download the installer script
	scriptPath := filepath.Join(os.TempDir(), "composer-setup.php")
	downloadConfig := composerutils.DownloadConfig{
		UseProxy:       i.config.UseProxy,
		ProxyURL:       i.config.ProxyURL,
		TimeoutSeconds: i.config.TimeoutSeconds,
	}
	if err := composerutils.DownloadFile(i.config.DownloadURL, scriptPath, downloadConfig); err != nil {
		return err
	}
	defer os.Remove(scriptPath)

	// Execute PHP installer script
	pharPath := filepath.Join(i.config.InstallPath, "composer.phar")

	var cmd *exec.Cmd
	if i.config.UseSudo {
		cmd = exec.Command("sudo", "php", scriptPath, "--install-dir="+i.config.InstallPath, "--filename=composer.phar")
	} else {
		cmd = exec.Command("php", scriptPath, "--install-dir="+i.config.InstallPath, "--filename=composer.phar")
	}

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w: %s, 错误: %v", ErrInstallationFailed, string(output), err)
	}

	// Create executable wrapper script
	binPath := filepath.Join(i.config.InstallPath, "composer")
	binContent := fmt.Sprintf("#!/bin/sh\nphp \"%s\" \"$@\"", pharPath)

	if i.config.UseSudo {
		// Use echo + sudo tee to create file
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' | sudo tee %s > /dev/null", binContent, binPath))
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("使用sudo创建可执行文件失败: %w", err)
		}

		// Set executable permission
		chmodCmd := exec.Command("sudo", "chmod", "755", binPath)
		if err := chmodCmd.Run(); err != nil {
			return fmt.Errorf("设置可执行权限失败: %w", err)
		}
	} else {
		if err := composerutils.CreateFileWithContent(binPath, []byte(binContent), 0755); err != nil {
			return fmt.Errorf("创建可执行文件失败: %w", err)
		}
	}

	return nil
}

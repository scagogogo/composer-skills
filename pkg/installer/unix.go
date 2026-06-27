package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composerutils"
)

// UnixInstaller 是通用 Unix 平台的 Composer 安装器
type UnixInstaller struct {
	config Config
}

// NewUnixInstaller 创建通用Unix平台的安装器
func NewUnixInstaller(config Config) *UnixInstaller {
	return &UnixInstaller{
		config: config,
	}
}

// Install 在类 Unix 系统上安装 Composer
func (i *UnixInstaller) Install() error {
	// 确保安装目录存在并且可写
	if err := composerutils.CheckWritePermission(i.config.InstallPath); err != nil {
		if !i.config.UseSudo {
			return fmt.Errorf("%w, 目标路径: %s", ErrInsufficientRights, i.config.InstallPath)
		}
	}

	// 下载安装脚本
	scriptPath := filepath.Join(os.TempDir(), "composer-setup.php")
	downloadConfig := composerutils.DownloadConfig{
		UseProxy: i.config.UseProxy,
		ProxyURL: i.config.ProxyURL,
	}
	if err := composerutils.DownloadFile(i.config.DownloadURL, scriptPath, downloadConfig); err != nil {
		return err
	}
	defer os.Remove(scriptPath)

	// 执行PHP脚本安装Composer
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

	// 创建可执行Composer脚本
	binPath := filepath.Join(i.config.InstallPath, "composer")
	binContent := fmt.Sprintf("#!/bin/sh\nphp \"%s\" \"$@\"", pharPath)

	var err error
	if i.config.UseSudo {
		// 使用echo和sudo tee创建文件
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' | sudo tee %s > /dev/null", binContent, binPath))
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("使用sudo创建可执行文件失败: %w", err)
		}

		// 设置可执行权限
		chmodCmd := exec.Command("sudo", "chmod", "755", binPath)
		if err = chmodCmd.Run(); err != nil {
			return fmt.Errorf("设置可执行权限失败: %w", err)
		}
	} else {
		// 直接创建文件
		if err = composerutils.CreateFileWithContent(binPath, []byte(binContent), 0755); err != nil {
			return fmt.Errorf("创建可执行文件失败: %w", err)
		}
	}

	return nil
}

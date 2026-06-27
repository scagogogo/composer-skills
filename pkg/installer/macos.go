package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composerutils"
)

// MacOSInstaller 是 MacOS 平台的 Composer 安装器
type MacOSInstaller struct {
	config Config
}

// NewMacOSInstaller 创建MacOS平台的安装器
func NewMacOSInstaller(config Config) *MacOSInstaller {
	return &MacOSInstaller{
		config: config,
	}
}

// Install 在 MacOS 上安装 Composer
func (i *MacOSInstaller) Install() error {
	// 先尝试使用Homebrew安装
	if i.tryBrewInstall() {
		return nil
	}

	// 如果brew安装失败，继续使用原来的安装方法
	// 确保安装目录存在并且可写
	if err := composerutils.CheckWritePermission(i.config.InstallPath); err != nil {
		return fmt.Errorf("安装目录无法写入: %w", err)
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
	if err := composerutils.CreateFileWithContent(binPath, []byte(binContent), 0755); err != nil {
		return fmt.Errorf("创建可执行文件失败: %w", err)
	}

	return nil
}

// tryBrewInstall 尝试使用Homebrew安装Composer
// 返回值: 安装是否成功
func (i *MacOSInstaller) tryBrewInstall() bool {
	// 如果配置不允许使用brew，则跳过
	if !i.config.PreferBrewOnMac {
		return false
	}

	fmt.Println("正在尝试使用Homebrew安装Composer...")

	// 检查brew是否已安装
	_, err := exec.LookPath("brew")
	if err != nil {
		fmt.Println("未检测到Homebrew，将使用传统方式安装Composer")
		return false
	}

	// 尝试使用brew安装composer
	cmd := exec.Command("brew", "install", "composer")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Homebrew安装失败: %s, 将使用传统方式安装\n", string(output))
		return false
	}

	fmt.Println("已成功通过Homebrew安装Composer")
	return true
}

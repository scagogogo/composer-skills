package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composerutils"
)

// WindowsInstaller 是 Windows 平台的 Composer 安装器
type WindowsInstaller struct {
	config Config
}

// NewWindowsInstaller 创建Windows平台的安装器
func NewWindowsInstaller(config Config) *WindowsInstaller {
	return &WindowsInstaller{
		config: config,
	}
}

// Install 在 Windows 上安装 Composer
func (i *WindowsInstaller) Install() error {
	// 确保安装目录存在
	if err := composerutils.EnsureDirectoryExists(i.config.InstallPath); err != nil {
		return fmt.Errorf("创建安装目录失败: %w", err)
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
	cmd := exec.Command("php", scriptPath, "--install-dir="+i.config.InstallPath, "--filename=composer.phar")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w: %s, 错误: %v", ErrInstallationFailed, string(output), err)
	}

	// 创建批处理文件以便直接调用composer
	batPath := filepath.Join(i.config.InstallPath, "composer.bat")
	batContent := fmt.Sprintf("@php \"%s\" %%*", pharPath)
	if err := composerutils.CreateFileWithContent(batPath, []byte(batContent), 0755); err != nil {
		return fmt.Errorf("创建批处理文件失败: %w", err)
	}

	// 将安装目录添加到PATH环境变量（需要管理员权限）
	// 这一步需要用户手动进行或使用管理员权限的方法
	fmt.Println("安装成功。请将以下目录添加到您的PATH环境变量中:")
	fmt.Println(i.config.InstallPath)

	return nil
}

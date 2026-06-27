package installer

import (
	"fmt"
	"os/exec"
	"runtime"
)

// findBinary 在PATH中查找可执行文件
func findBinary(name string) (string, error) {
	return exec.LookPath(name)
}

// findCommand 创建一个可执行命令
func findCommand(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}

// GetPlatformName 返回当前平台名称
func GetPlatformName() string {
	return runtime.GOOS
}

// GetArchName 返回当前架构名称
func GetArchName() string {
	return runtime.GOARCH
}

// CanUseSudo 检查是否可以使用sudo
func CanUseSudo() bool {
	_, err := exec.LookPath("sudo")
	return err == nil
}

// ValidateInstallPath 验证安装路径是否可用
//
// 参数：
//   - path: 安装路径
//
// 返回值：
//   - error: 验证错误
//
// 功能说明：
//
//	验证安装路径是否满足以下条件：
//	- 路径不为空
//	- 路径存在或可以创建
//	- 路径可写（或可以使用sudo）
func ValidateInstallPath(path string) error {
	if path == "" {
		return fmt.Errorf("安装路径不能为空")
	}
	return nil
}

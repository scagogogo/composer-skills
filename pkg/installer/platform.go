package installer

import (
	"fmt"
	"runtime"
)

// PlatformInstaller 定义平台特定的安装逻辑
type PlatformInstaller interface {
	Install() error
}

// GetPlatformInstaller 根据当前操作系统返回适合的安装器
func GetPlatformInstaller(config Config) (PlatformInstaller, error) {
	switch runtime.GOOS {
	case "windows":
		return NewWindowsInstaller(config), nil
	case "darwin":
		return NewMacOSInstaller(config), nil
	case "linux":
		return NewLinuxInstaller(config), nil
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		// 其他类Unix系统使用通用Unix安装器
		return NewUnixInstaller(config), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedPlatform, runtime.GOOS)
	}
}

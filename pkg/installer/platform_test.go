package installer

import (
	"runtime"
	"testing"
)

func TestGetPlatformInstaller(t *testing.T) {
	// 创建一个默认配置
	config := DefaultConfig()

	// 获取当前平台的安装器
	installer, err := GetPlatformInstaller(config)
	if err != nil {
		t.Fatalf("GetPlatformInstaller() error = %v", err)
	}

	// 确保安装器不为空
	if installer == nil {
		t.Fatal("GetPlatformInstaller() 返回了nil安装器")
	}

	// 检查返回的安装器类型是否与当前平台匹配
	switch runtime.GOOS {
	case "windows":
		_, ok := installer.(*WindowsInstaller)
		if !ok {
			t.Errorf("GetPlatformInstaller() 在Windows平台应该返回WindowsInstaller")
		}
	case "darwin":
		_, ok := installer.(*MacOSInstaller)
		if !ok {
			t.Errorf("GetPlatformInstaller() 在macOS平台应该返回MacOSInstaller")
		}
	case "linux":
		_, ok := installer.(*LinuxInstaller)
		if !ok {
			t.Errorf("GetPlatformInstaller() 在Linux平台应该返回LinuxInstaller")
		}
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		_, ok := installer.(*UnixInstaller)
		if !ok {
			t.Errorf("GetPlatformInstaller() 在其他Unix平台应该返回UnixInstaller")
		}
	}
}

func TestGetPlatformInstallerUnsupported(t *testing.T) {
	// 跳过实际测试，因为无法轻易模拟不支持的平台
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "linux" ||
		runtime.GOOS == "freebsd" || runtime.GOOS == "openbsd" || runtime.GOOS == "netbsd" || runtime.GOOS == "dragonfly" {
		t.Skip("跳过不支持的平台测试，因为当前在支持的平台上")
	}

	// 只有在不受支持的平台上才会执行以下代码
	config := DefaultConfig()
	_, err := GetPlatformInstaller(config)
	if err == nil {
		t.Error("GetPlatformInstaller() 应该在不支持的平台上返回错误")
	}
}

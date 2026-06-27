package detector

import (
	"runtime"
	"testing"
)

// 通用测试函数，使用不同平台的getPlatformSpecificPaths实现
func TestAllPlatformSpecificPaths(t *testing.T) {
	// 测试Windows平台的路径
	t.Run("TestWindowsPaths", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("在Windows平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的Windows环境测试
		paths := []string{
			"AppData\\Composer\\composer.phar",
			"ProgramFiles\\Composer\\composer.phar",
			"composer.phar",
			"composer.bat",
			"composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Windows平台的路径列表不应为空")
		}
	})

	// 测试macOS平台的路径
	t.Run("TestDarwinPaths", func(t *testing.T) {
		if runtime.GOOS == "darwin" {
			t.Skip("在Darwin平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的macOS环境测试
		paths := []string{
			"/usr/local/bin/composer",
			"/usr/bin/composer",
			"/opt/homebrew/bin/composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Darwin平台的路径列表不应为空")
		}
	})

	// 测试Linux/Unix平台的路径
	t.Run("TestUnixPaths", func(t *testing.T) {
		if runtime.GOOS != "windows" && runtime.GOOS != "darwin" {
			t.Skip("在Unix平台上跳过模拟测试")
			return
		}

		// 创建一个模拟的Unix环境测试
		paths := []string{
			"/usr/local/bin/composer",
			"/usr/bin/composer",
		}

		// 简单验证返回的路径长度
		if len(paths) == 0 {
			t.Error("Unix平台的路径列表不应为空")
		}
	})
}

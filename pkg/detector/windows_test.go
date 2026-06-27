//go:build windows

package detector

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWindowsSpecificPaths(t *testing.T) {
	paths := getPlatformSpecificPaths()

	// 验证是否包含 Windows 特定的路径
	expectedPaths := []string{
		filepath.Join(os.Getenv("APPDATA"), "Composer", "composer.phar"),
		filepath.Join(os.Getenv("ProgramFiles"), "Composer", "composer.phar"),
		"composer.bat",
	}

	for _, expected := range expectedPaths {
		found := false
		for _, path := range paths {
			if path == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Windows路径中应包含 %s", expected)
		}
	}
}

func TestWindowsExecutableCheck(t *testing.T) {
	// 在 Windows 上创建测试文件
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "composer.bat")

	if err := os.WriteFile(tmpFile, []byte("@echo off\necho Fake Composer"), 0644); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// Windows 上所有文件都应该被认为是可执行的
	if !isExecutable(tmpFile) {
		t.Errorf("Windows上文件 %s 应该被识别为可执行文件", tmpFile)
	}
}

func TestWindowsDetectWithWhereCommand(t *testing.T) {
	// 注意：这个测试依赖于系统环境，如果系统中没有安装 where 命令，测试可能会失败
	t.Skip("跳过依赖于系统环境的测试")

	// 创建一个假的 where 命令输出的测试环境：未实现
	// 完整测试需要模拟系统环境，这超出了简单单元测试的范围
}

//go:build darwin

package detector

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDarwinSpecificPaths(t *testing.T) {
	paths := getPlatformSpecificPaths()

	// 验证是否包含 macOS 特定的路径
	expectedPaths := []string{
		"/usr/local/bin/composer",
		"/usr/bin/composer",
		"/opt/homebrew/bin/composer",
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
			t.Errorf("macOS路径中应包含 %s", expected)
		}
	}

	// 验证是否包含特定于用户的路径
	homeDir := os.Getenv("HOME")
	if homeDir != "" {
		userSpecificPath := filepath.Join(homeDir, ".composer/vendor/bin/composer")
		found := false
		for _, path := range paths {
			if path == userSpecificPath {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("macOS路径中应包含用户路径 %s", userSpecificPath)
		}
	}
}

func TestDarwinExecutablePermissions(t *testing.T) {
	// 创建一个临时文件用于测试
	tmpDir := t.TempDir()
	nonExecutableFile := filepath.Join(tmpDir, "non-executable")
	executableFile := filepath.Join(tmpDir, "executable")

	// 创建非可执行文件
	if err := os.WriteFile(nonExecutableFile, []byte("test"), 0644); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 创建可执行文件
	if err := os.WriteFile(executableFile, []byte("#!/bin/sh\necho test"), 0755); err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 验证非可执行文件
	if isExecutable(nonExecutableFile) {
		t.Errorf("文件 %s 不应该被识别为可执行文件", nonExecutableFile)
	}

	// 验证可执行文件
	if !isExecutable(executableFile) {
		t.Errorf("文件 %s 应该被识别为可执行文件", executableFile)
	}
}

func TestDarwinDetectWithWhichCommand(t *testing.T) {
	// 注意：这个测试依赖于系统环境，如果系统中没有安装 which 命令，测试可能会失败
	t.Skip("跳过依赖于系统环境的测试")

	// 创建一个假的 which 命令输出的测试环境：未实现
	// 完整测试需要模拟系统环境，这超出了简单单元测试的范围
}

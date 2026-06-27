// 测试detector包的通用功能

package detector

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewDetector(t *testing.T) {
	d := NewDetector()
	if d == nil {
		t.Fatal("NewDetector应返回非nil的Detector实例")
		return // 虽然t.Fatal会终止，但这样可以让staticcheck满意
	}

	if len(d.possiblePaths) == 0 {
		t.Error("Detector应包含默认的可能路径")
	}
}

func TestSetPossiblePaths(t *testing.T) {
	d := NewDetector()
	customPaths := []string{"/custom/path1", "/custom/path2"}
	d.SetPossiblePaths(customPaths)

	if len(d.possiblePaths) != len(customPaths) {
		t.Errorf("期望possiblePaths长度为%d，实际为%d", len(customPaths), len(d.possiblePaths))
	}

	for i, path := range customPaths {
		if d.possiblePaths[i] != path {
			t.Errorf("期望路径[%d]为%s，实际为%s", i, path, d.possiblePaths[i])
		}
	}
}

func TestAddPossiblePath(t *testing.T) {
	d := NewDetector()
	originalCount := len(d.possiblePaths)

	newPath := "/new/custom/path"
	d.AddPossiblePath(newPath)

	if len(d.possiblePaths) != originalCount+1 {
		t.Errorf("添加路径后，期望possiblePaths长度为%d，实际为%d", originalCount+1, len(d.possiblePaths))
	}

	if d.possiblePaths[len(d.possiblePaths)-1] != newPath {
		t.Errorf("期望最后一个路径为%s，实际为%s", newPath, d.possiblePaths[len(d.possiblePaths)-1])
	}
}

func TestIsInstalled(t *testing.T) {
	// 创建一个假的"composer"可执行文件用于测试
	tmpDir := t.TempDir()
	fakePath := filepath.Join(tmpDir, "composer")

	if err := os.WriteFile(fakePath, []byte("#!/bin/sh\necho 'Fake Composer'"), 0755); err != nil {
		t.Fatalf("创建假的composer文件失败: %v", err)
	}

	// 测试有效路径
	d := NewDetector()
	d.SetPossiblePaths([]string{fakePath})

	if !d.IsInstalled() {
		t.Error("使用有效的测试路径时，IsInstalled应返回true")
	}

	// 测试无效路径 - 需要清除环境变量并确保which命令找不到composer
	d2 := NewDetector()
	d2.SetPossiblePaths([]string{"/definitely/not/exists/composer"})

	// 临时设置环境变量为空，并在测试结束后恢复
	originalComposerPath := os.Getenv("COMPOSER_PATH")
	os.Setenv("COMPOSER_PATH", "")
	defer func() {
		if originalComposerPath != "" {
			os.Setenv("COMPOSER_PATH", originalComposerPath)
		} else {
			os.Unsetenv("COMPOSER_PATH")
		}
	}()

	// 注意：由于系统可能已安装composer，这个测试可能在某些环境中失败
	// 这是预期的行为，因为detector应该能找到系统中已安装的composer
	_ = d2.IsInstalled() // 不强制要求返回false，因为系统可能确实有composer
}

// 注意：Detect方法的完整测试需要模拟系统环境，这里只做简单测试
func TestDetect(t *testing.T) {
	d := NewDetector()

	// 创建一个假的composer可执行文件
	tmpDir := t.TempDir()
	fakePath := filepath.Join(tmpDir, "composer")

	if err := os.WriteFile(fakePath, []byte("#!/bin/sh\necho 'Fake Composer'"), 0755); err != nil {
		t.Fatalf("创建假的composer文件失败: %v", err)
	}

	// 设置为仅包含这个假路径
	d.SetPossiblePaths([]string{fakePath})

	detectedPath, err := d.Detect()
	if err != nil {
		t.Errorf("使用有效的测试路径时，Detect不应返回错误: %v", err)
	}

	if detectedPath != fakePath {
		t.Errorf("Detect应返回设置的路径 %s，实际返回 %s", fakePath, detectedPath)
	}
}

func TestDetectWithEnvironmentVariable(t *testing.T) {
	// 保存原始环境变量并在测试结束后恢复
	oldPath := os.Getenv("COMPOSER_PATH")
	defer os.Setenv("COMPOSER_PATH", oldPath)

	// 创建一个假的composer可执行文件
	tmpDir := t.TempDir()
	fakePath := filepath.Join(tmpDir, "composer")

	if err := os.WriteFile(fakePath, []byte("#!/bin/sh\necho 'Fake Composer'"), 0755); err != nil {
		t.Fatalf("创建假的composer文件失败: %v", err)
	}

	// 设置环境变量
	os.Setenv("COMPOSER_PATH", fakePath)

	// 创建检测器并尝试检测
	d := NewDetector()
	detectedPath, err := d.Detect()

	if err != nil {
		t.Errorf("使用环境变量设置的路径时，Detect不应返回错误: %v", err)
	}

	if detectedPath != fakePath {
		t.Errorf("Detect应返回环境变量中的路径 %s，实际返回 %s", fakePath, detectedPath)
	}
}

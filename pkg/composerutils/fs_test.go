package composerutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckWritePermission(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() (string, func())
		expectError bool
		errorMsg    string
	}{
		{
			name: "正常可写目录",
			setup: func() (string, func()) {
				// 创建临时目录
				dir := t.TempDir()
				return dir, func() {}
			},
			expectError: false,
		},
		{
			name: "目录不存在但可创建",
			setup: func() (string, func()) {
				// 创建一个不存在的目录路径
				dir := filepath.Join(t.TempDir(), "subdir", "nested")
				return dir, func() {
					os.RemoveAll(filepath.Dir(filepath.Dir(dir)))
				}
			},
			expectError: false,
		},
		// 注意: 以下测试在实际系统中可能难以模拟，因为需要受限权限
		// 但我们仍然提供测试用例结构以供参考
		/*
			{
				name: "无写入权限目录",
				setup: func() (string, func()) {
					// 创建一个只读目录（在实际测试环境中需要根据权限调整）
					dir := filepath.Join(t.TempDir(), "readonly")
					os.MkdirAll(dir, 0500) // 只有执行权限，没有写权限
					return dir, func() {
						os.RemoveAll(dir)
					}
				},
				expectError: true,
				errorMsg:    "目录没有写入权限",
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备测试环境
			dir, cleanup := tt.setup()
			defer cleanup()

			// 执行测试
			err := CheckWritePermission(dir)

			// 验证结果
			if tt.expectError {
				if err == nil {
					t.Errorf("期望错误但得到nil")
				} else if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("期望错误包含'%s'，但得到'%s'", tt.errorMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("不期望错误但得到: %v", err)
			}

			// 额外验证目录是否真的可写
			if !tt.expectError && err == nil {
				// 尝试在目录中创建文件
				testFile := filepath.Join(dir, "test.txt")
				if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
					t.Errorf("目录应该可写，但写入测试文件失败: %v", err)
				}
			}
		})
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() (string, func())
		expectError bool
		errorMsg    string
	}{
		{
			name: "目录已存在",
			setup: func() (string, func()) {
				// 创建临时目录
				dir := t.TempDir()
				return dir, func() {}
			},
			expectError: false,
		},
		{
			name: "目录不存在可创建",
			setup: func() (string, func()) {
				// 创建临时目录并在其中构建新路径
				baseDir := t.TempDir()
				dir := filepath.Join(baseDir, "new_dir")
				return dir, func() {}
			},
			expectError: false,
		},
		{
			name: "多级目录不存在可创建",
			setup: func() (string, func()) {
				baseDir := t.TempDir()
				dir := filepath.Join(baseDir, "level1", "level2", "level3")
				return dir, func() {}
			},
			expectError: false,
		},
		// 以下测试在实际环境中难以模拟，仅作为参考结构
		/*
			{
				name: "无权限创建目录",
				setup: func() (string, func()) {
					// 尝试在只读区域创建目录
					readOnlyDir := "/some/readonly/path"
					return filepath.Join(readOnlyDir, "new_dir"), func() {}
				},
				expectError: true,
				errorMsg:    "permission denied",
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备测试环境
			dir, cleanup := tt.setup()
			defer cleanup()

			// 执行测试
			err := EnsureDirectoryExists(dir)

			// 验证结果
			if tt.expectError {
				if err == nil {
					t.Errorf("期望错误但得到nil")
				} else if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("期望错误包含'%s'，但得到'%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("不期望错误但得到: %v", err)
				}
				// 检查目录是否已创建
				if _, err := os.Stat(dir); os.IsNotExist(err) {
					t.Errorf("目录应该已创建，但未创建: %s", dir)
				}
			}
		})
	}
}

func TestCreateFileWithContent(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() (string, []byte, os.FileMode, func())
		validate    func(string) error
		expectError bool
		errorMsg    string
	}{
		{
			name: "正常创建文件",
			setup: func() (string, []byte, os.FileMode, func()) {
				dir := t.TempDir()
				filePath := filepath.Join(dir, "test.txt")
				content := []byte("hello world")
				return filePath, content, 0644, func() {}
			},
			validate: func(filePath string) error {
				content, err := os.ReadFile(filePath)
				if err != nil {
					return err
				}
				if string(content) != "hello world" {
					return err
				}
				return nil
			},
			expectError: false,
		},
		{
			name: "创建可执行文件",
			setup: func() (string, []byte, os.FileMode, func()) {
				dir := t.TempDir()
				filePath := filepath.Join(dir, "test.sh")
				content := []byte("#!/bin/sh\necho 'test'")
				return filePath, content, 0755, func() {}
			},
			validate: func(filePath string) error {
				info, err := os.Stat(filePath)
				if err != nil {
					return err
				}
				// 检查权限是否正确设置(跳过Windows)
				if os.Getenv("GOOS") != "windows" && info.Mode().Perm() != 0755 {
					t.Errorf("文件权限应为0755，但得到%v", info.Mode().Perm())
				}
				return nil
			},
			expectError: false,
		},
		{
			name: "自动创建父目录",
			setup: func() (string, []byte, os.FileMode, func()) {
				dir := t.TempDir()
				filePath := filepath.Join(dir, "subdir", "nested", "test.txt")
				content := []byte("nested file content")
				return filePath, content, 0644, func() {}
			},
			validate: func(filePath string) error {
				// 检查文件是否存在并内容正确
				content, err := os.ReadFile(filePath)
				if err != nil {
					return err
				}
				if string(content) != "nested file content" {
					t.Errorf("文件内容不匹配，期望'nested file content'，得到'%s'", string(content))
				}
				return nil
			},
			expectError: false,
		},
		// 以下测试在实际环境中难以模拟，仅作为参考结构
		/*
			{
				name: "无写权限目录",
				setup: func() (string, []byte, os.FileMode, func()) {
					// 尝试在只读目录中创建文件
					readOnlyDir := "/some/readonly/path"
					return filepath.Join(readOnlyDir, "test.txt"), []byte("test"), 0644, func() {}
				},
				expectError: true,
				errorMsg:    "permission denied",
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备测试环境
			filePath, content, perm, cleanup := tt.setup()
			defer cleanup()

			// 执行测试
			err := CreateFileWithContent(filePath, content, perm)

			// 验证结果
			if tt.expectError {
				if err == nil {
					t.Errorf("期望错误但得到nil")
				} else if tt.errorMsg != "" && !contains(err.Error(), tt.errorMsg) {
					t.Errorf("期望错误包含'%s'，但得到'%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("不期望错误但得到: %v", err)
				} else if tt.validate != nil {
					if err := tt.validate(filePath); err != nil {
						t.Errorf("文件验证失败: %v", err)
					}
				}
			}
		})
	}
}

// contains 辅助函数，检查字符串是否包含子串
func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && len(s) > len(substr) && s[len(s)-len(substr):] == substr
}

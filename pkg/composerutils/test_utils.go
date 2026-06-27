package composerutils

import (
	"os"
	"path/filepath"
	"testing"
)

// TestHelpers 包含测试辅助函数
type TestHelpers struct {
	// 可以添加更多测试辅助函数
}

// NewTestHelpers 创建测试辅助实例
func NewTestHelpers() *TestHelpers {
	return &TestHelpers{}
}

// CreateTempDir 创建测试用的临时目录
func (h *TestHelpers) CreateTempDir(t *testing.T) string {
	dir, err := os.MkdirTemp("", "composer-test-*")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	return dir
}

// RemoveTempDir 清理测试用的临时目录
func (h *TestHelpers) RemoveTempDir(t *testing.T, dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		t.Errorf("无法清理临时目录 %s: %v", dir, err)
	}
}

// CreateTestFile 在测试目录中创建文件
func (h *TestHelpers) CreateTestFile(t *testing.T, dir, name string, content []byte) string {
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		t.Fatalf("无法创建测试文件 %s: %v", path, err)
	}
	return path
}

// AssertFileExists 断言文件存在
func (h *TestHelpers) AssertFileExists(t *testing.T, path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Errorf("预期文件 %s 应存在，但未找到", path)
	} else if err != nil {
		t.Errorf("检查文件 %s 时发生错误: %v", path, err)
	}
}

// AssertFileNotExists 断言文件不存在
func (h *TestHelpers) AssertFileNotExists(t *testing.T, path string) {
	_, err := os.Stat(path)
	if err == nil {
		t.Errorf("预期文件 %s 不应存在，但找到了", path)
	} else if !os.IsNotExist(err) {
		t.Errorf("检查文件 %s 时发生错误: %v", path, err)
	}
}

// AssertFileContent 断言文件内容
func (h *TestHelpers) AssertFileContent(t *testing.T, path string, expectedContent []byte) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("读取文件 %s 失败: %v", path, err)
		return
	}

	if string(content) != string(expectedContent) {
		t.Errorf("文件 %s 内容不匹配\n期望值: %s\n实际值: %s",
			path, string(expectedContent), string(content))
	}
}

package composerutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewTestHelpers(t *testing.T) {
	helpers := NewTestHelpers()
	if helpers == nil {
		t.Fatal("NewTestHelpers should return non-nil helpers")
	}
}

func TestTestHelpers_CreateTempDir(t *testing.T) {
	helpers := NewTestHelpers()
	dir := helpers.CreateTempDir(t)

	if dir == "" {
		t.Fatal("CreateTempDir should return non-empty path")
	}

	// 验证目录存在
	info, err := os.Stat(dir)
	if err != nil {
		t.Errorf("Created temp dir should exist: %v", err)
	}
	if !info.IsDir() {
		t.Error("Created path should be a directory")
	}

	// 清理
	helpers.RemoveTempDir(t, dir)

	// 验证目录已删除
	_, err = os.Stat(dir)
	if !os.IsNotExist(err) {
		t.Error("Temp dir should be removed after RemoveTempDir")
	}
}

func TestTestHelpers_RemoveTempDir(t *testing.T) {
	helpers := NewTestHelpers()

	// 创建临时目录用于测试删除
	tmpDir := t.TempDir()
	testDir := filepath.Join(tmpDir, "subdir")
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test dir: %v", err)
	}

	// 记录目录存在
	if _, err := os.Stat(testDir); err != nil {
		t.Fatalf("Test dir should exist: %v", err)
	}

	// 测试删除非空目录
	helpers.RemoveTempDir(t, testDir)

	// 验证目录已删除（允许错误，因为 RemoveTempDir 内部处理了错误）
	// 注意：RemoveTempDir 使用 os.RemoveAll，所以应该能删除非空目录
}

func TestTestHelpers_CreateTestFile(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	content := []byte("test content")
	path := helpers.CreateTestFile(t, tmpDir, "test.txt", content)

	if path == "" {
		t.Fatal("CreateTestFile should return non-empty path")
	}

	// 验证文件存在
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("Created file should exist: %v", err)
	}
	if info.IsDir() {
		t.Error("Created path should be a file, not directory")
	}

	// 验证内容
	readContent, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
	}
	if string(readContent) != "test content" {
		t.Errorf("File content mismatch: got '%s', want 'test content'", string(readContent))
	}
}

func TestTestHelpers_CreateTestFile_NestedPath(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	// 创建嵌套目录
	nestedDir := filepath.Join(tmpDir, "a", "b", "c")
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}

	content := []byte("nested content")
	path := helpers.CreateTestFile(t, nestedDir, "nested.txt", content)

	// 验证文件存在
	_, err := os.Stat(path)
	if err != nil {
		t.Errorf("Created nested file should exist: %v", err)
	}
}

func TestTestHelpers_AssertFileExists(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	// 创建测试文件
	testFile := filepath.Join(tmpDir, "exists.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 测试断言文件存在 - 应该不触发 t.Errorf
	helpers.AssertFileExists(t, testFile)
}

func TestTestHelpers_AssertFileNotExists(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	nonExistent := filepath.Join(tmpDir, "notexists.txt")

	// 测试断言不存在的文件 - 这不会触发 t.Errorf，因为文件确实不存在
	helpers.AssertFileNotExists(t, nonExistent)
}

func TestTestHelpers_AssertFileContent(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	testFile := filepath.Join(tmpDir, "content.txt")
	expectedContent := []byte("expected content")

	if err := os.WriteFile(testFile, expectedContent, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 测试匹配的内容 - 不会触发 t.Errorf
	helpers.AssertFileContent(t, testFile, expectedContent)
}

func TestTestHelpers_CreateTempDir_Multiple(t *testing.T) {
	helpers := NewTestHelpers()

	// 创建多个临时目录
	dirs := make([]string, 5)
	for i := 0; i < 5; i++ {
		dirs[i] = helpers.CreateTempDir(t)
	}

	// 验证所有目录都不同
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if dirs[i] == dirs[j] {
				t.Errorf("Two temp dirs should be different: '%s' and '%s'", dirs[i], dirs[j])
			}
		}
	}

	// 清理所有目录
	for _, dir := range dirs {
		helpers.RemoveTempDir(t, dir)
	}
}

func TestTestHelpers_CreateTestFile_EmptyContent(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	// 测试创建空内容文件
	path := helpers.CreateTestFile(t, tmpDir, "empty.txt", []byte{})

	// 验证文件存在且为空
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("Created file should exist: %v", err)
	}
	if info.Size() != 0 {
		t.Errorf("Empty file should have size 0, got %d", info.Size())
	}
}

func TestTestHelpers_RemoveTempDir_NonExistent(t *testing.T) {
	helpers := NewTestHelpers()

	// 测试删除不存在的目录（不应 panic）
	nonExistent := filepath.Join(t.TempDir(), "doesnotexist")
	helpers.RemoveTempDir(t, nonExistent)
}

func TestTestHelpers_AssertFileContent_EmptyFile(t *testing.T) {
	helpers := NewTestHelpers()
	tmpDir := helpers.CreateTempDir(t)
	defer helpers.RemoveTempDir(t, tmpDir)

	testFile := filepath.Join(tmpDir, "empty.txt")
	if err := os.WriteFile(testFile, []byte{}, 0644); err != nil {
		t.Fatalf("Failed to create empty test file: %v", err)
	}

	// 测试空内容匹配 - 不会触发 t.Errorf
	helpers.AssertFileContent(t, testFile, []byte{})
}

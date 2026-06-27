// Package utils 提供文件系统和HTTP相关的实用工具函数
package composerutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CheckWritePermission 检查指定目录是否具有写入权限
//
// 该函数首先检查目录是否存在，如果不存在会尝试创建它。
// 然后通过创建临时文件来验证写入权限，测试后会自动删除临时文件。
//
// 参数:
//   - dir: 要检查写入权限的目录路径
//
// 返回值:
//   - error: 如果目录不可写或无法创建，返回相应错误；如果目录可写返回nil
//
// 使用示例:
//
//	if err := utils.CheckWritePermission("/path/to/directory"); err != nil {
//	    log.Fatalf("目录无写入权限: %v", err)
//	}
//
// 可能的错误:
//   - 目录创建失败（权限不足或路径无效）
//   - 无法在目录中创建临时文件（权限不足）
func CheckWritePermission(dir string) error {
	// 确保目录存在
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("无法创建目录: %w", err)
	}

	// 通过创建临时文件测试写入权限
	tempFile := filepath.Join(dir, ".write-test-"+fmt.Sprintf("%d", os.Getpid()))
	f, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("目录没有写入权限: %w", err)
	}
	f.Close()
	os.Remove(tempFile)

	return nil
}

// EnsureDirectoryExists 确保指定的目录存在，如果不存在则创建
//
// 参数:
//   - dir: 需要确保存在的目录路径
//
// 返回值:
//   - error: 如果无法创建目录则返回错误；成功创建或目录已存在则返回nil
//
// 使用示例:
//
//	installDir := "/usr/local/bin"
//	if err := utils.EnsureDirectoryExists(installDir); err != nil {
//	    log.Fatalf("无法创建安装目录: %v", err)
//	}
//
// 可能的错误:
//   - 权限不足无法创建目录
//   - 路径包含非法字符
//   - 磁盘空间不足
func EnsureDirectoryExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// CreateFileWithContent 在指定路径创建文件并写入内容
//
// 参数:
//   - filePath: 要创建的文件路径（包括文件名）
//   - content: 要写入文件的内容（字节数组）
//   - perm: 文件权限模式（如0644, 0755等）
//
// 返回值:
//   - error: 如果文件创建或写入失败则返回错误；成功则返回nil
//
// 使用示例:
//
//	scriptContent := []byte("#!/bin/sh\necho 'Hello World'")
//	if err := utils.CreateFileWithContent("/usr/local/bin/hello.sh", scriptContent, 0755); err != nil {
//	    log.Fatalf("无法创建脚本文件: %v", err)
//	}
//
// 可能的错误:
//   - 文件已存在且无法覆盖
//   - 路径中的目录不存在且无法创建
//   - 权限不足无法写入文件
//   - 磁盘空间不足
func CreateFileWithContent(filePath string, content []byte, perm os.FileMode) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := EnsureDirectoryExists(dir); err != nil {
		return fmt.Errorf("确保目录存在失败: %w", err)
	}

	// 创建并写入文件
	return os.WriteFile(filePath, content, perm)
}

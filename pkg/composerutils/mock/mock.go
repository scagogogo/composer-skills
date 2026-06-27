package mock

import (
	"fmt"
	"os"
	"os/exec"
)

// CommandExecutor 定义了一个模拟的命令执行器接口
type CommandExecutor interface {
	Execute(name string, args ...string) ([]byte, error)
}

// DefaultCommandExecutor 是默认的命令执行器，实际执行系统命令
type DefaultCommandExecutor struct{}

// Execute 执行实际的系统命令
func (e *DefaultCommandExecutor) Execute(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

// MockCommandExecutor 是模拟的命令执行器，返回预设的结果
type MockCommandExecutor struct {
	// 命令映射：命令+参数 -> 输出和错误
	CommandResults map[string]struct {
		Output []byte
		Err    error
	}
}

// NewMockCommandExecutor 创建一个新的模拟命令执行器
func NewMockCommandExecutor() *MockCommandExecutor {
	return &MockCommandExecutor{
		CommandResults: make(map[string]struct {
			Output []byte
			Err    error
		}),
	}
}

// SetCommandResult 设置指定命令的模拟结果
func (e *MockCommandExecutor) SetCommandResult(name string, args []string, output []byte, err error) {
	cmdKey := name
	for _, arg := range args {
		cmdKey += " " + arg
	}
	e.CommandResults[cmdKey] = struct {
		Output []byte
		Err    error
	}{
		Output: output,
		Err:    err,
	}
}

// Execute 模拟执行命令，返回预设的结果
func (e *MockCommandExecutor) Execute(name string, args ...string) ([]byte, error) {
	cmdKey := name
	for _, arg := range args {
		cmdKey += " " + arg
	}

	result, exists := e.CommandResults[cmdKey]
	if !exists {
		return nil, fmt.Errorf("未找到模拟命令结果: %s", cmdKey)
	}

	return result.Output, result.Err
}

// FileSystemHelper 定义文件系统操作的接口
type FileSystemHelper interface {
	CreateFile(path string, content []byte, perm os.FileMode) error
	CheckWritePermission(path string) error
	EnsureDirectoryExists(path string) error
	RemoveFile(path string) error
}

// DefaultFileSystemHelper 是默认的文件系统操作助手，执行实际的文件操作
type DefaultFileSystemHelper struct{}

// CreateFile 创建文件
func (fs *DefaultFileSystemHelper) CreateFile(path string, content []byte, perm os.FileMode) error {
	return os.WriteFile(path, content, perm)
}

// CheckWritePermission 检查路径是否可写
func (fs *DefaultFileSystemHelper) CheckWritePermission(path string) error {
	// 尝试创建临时文件测试权限
	testFile := path + "/.write_test"
	f, err := os.Create(testFile)
	if err != nil {
		return err
	}
	f.Close()
	return os.Remove(testFile)
}

// EnsureDirectoryExists 确保目录存在
func (fs *DefaultFileSystemHelper) EnsureDirectoryExists(path string) error {
	return os.MkdirAll(path, 0755)
}

// RemoveFile 删除文件
func (fs *DefaultFileSystemHelper) RemoveFile(path string) error {
	return os.Remove(path)
}

// MockFileSystemHelper 是模拟的文件系统操作助手
type MockFileSystemHelper struct {
	CreateFileFunc            func(path string, content []byte, perm os.FileMode) error
	CheckWritePermissionFunc  func(path string) error
	EnsureDirectoryExistsFunc func(path string) error
	RemoveFileFunc            func(path string) error
}

// NewMockFileSystemHelper 创建一个新的模拟文件系统助手
func NewMockFileSystemHelper() *MockFileSystemHelper {
	return &MockFileSystemHelper{
		CreateFileFunc: func(path string, content []byte, perm os.FileMode) error {
			return nil
		},
		CheckWritePermissionFunc: func(path string) error {
			return nil
		},
		EnsureDirectoryExistsFunc: func(path string) error {
			return nil
		},
		RemoveFileFunc: func(path string) error {
			return nil
		},
	}
}

// CreateFile 模拟创建文件
func (fs *MockFileSystemHelper) CreateFile(path string, content []byte, perm os.FileMode) error {
	return fs.CreateFileFunc(path, content, perm)
}

// CheckWritePermission 模拟检查写权限
func (fs *MockFileSystemHelper) CheckWritePermission(path string) error {
	return fs.CheckWritePermissionFunc(path)
}

// EnsureDirectoryExists 模拟确保目录存在
func (fs *MockFileSystemHelper) EnsureDirectoryExists(path string) error {
	return fs.EnsureDirectoryExistsFunc(path)
}

// RemoveFile 模拟删除文件
func (fs *MockFileSystemHelper) RemoveFile(path string) error {
	return fs.RemoveFileFunc(path)
}

// DownloadHelper 定义下载操作的接口
type DownloadHelper interface {
	DownloadFile(url string, target string, config interface{}) error
}

// MockDownloadHelper 是模拟的下载助手
type MockDownloadHelper struct {
	DownloadFileFunc func(url string, target string, config interface{}) error
}

// NewMockDownloadHelper 创建一个新的模拟下载助手
func NewMockDownloadHelper() *MockDownloadHelper {
	return &MockDownloadHelper{
		DownloadFileFunc: func(url string, target string, config interface{}) error {
			return nil
		},
	}
}

// DownloadFile 模拟下载文件
func (dh *MockDownloadHelper) DownloadFile(url string, target string, config interface{}) error {
	return dh.DownloadFileFunc(url, target, config)
}

// RuntimeInfo 定义运行时环境信息
type RuntimeInfo struct {
	GOOS   string // 操作系统
	GOARCH string // 系统架构
}

// MockRuntime 提供模拟的运行时信息
type MockRuntime struct {
	Info RuntimeInfo
}

// NewMockRuntime 创建一个新的模拟运行时
func NewMockRuntime() *MockRuntime {
	return &MockRuntime{
		Info: RuntimeInfo{
			GOOS:   "linux",
			GOARCH: "amd64",
		},
	}
}

// SetOS 设置模拟的操作系统
func (r *MockRuntime) SetOS(os string) {
	r.Info.GOOS = os
}

// GetOS 获取模拟的操作系统
func (r *MockRuntime) GetOS() string {
	return r.Info.GOOS
}

// SetArch 设置模拟的系统架构
func (r *MockRuntime) SetArch(arch string) {
	r.Info.GOARCH = arch
}

// GetArch 获取模拟的系统架构
func (r *MockRuntime) GetArch() string {
	return r.Info.GOARCH
}

package mock

import (
	"errors"
	"os"
	"testing"
)

func TestNewMockCommandExecutor(t *testing.T) {
	executor := NewMockCommandExecutor()
	if executor == nil {
		t.Fatal("NewMockCommandExecutor should return non-nil executor")
	}
	if executor.CommandResults == nil {
		t.Error("CommandResults map should be initialized")
	}
}

func TestMockCommandExecutor_SetCommandResult(t *testing.T) {
	executor := NewMockCommandExecutor()

	testCases := []struct {
		name    string
		cmd     string
		args    []string
		output  []byte
		err     error
	}{
		{"无参数命令", "version", nil, []byte("2.5.0"), nil},
		{"带参数命令", "require", []string{"--dev"}, []byte("success"), nil},
		{"错误命令", "invalid", nil, nil, errors.New("command not found")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor.SetCommandResult(tc.cmd, tc.args, tc.output, tc.err)
		})
	}

	if len(executor.CommandResults) != 3 {
		t.Errorf("Expected 3 command results, got %d", len(executor.CommandResults))
	}
}

func TestMockCommandExecutor_Execute(t *testing.T) {
	executor := NewMockCommandExecutor()

	// 设置一个已知命令的结果
	executor.SetCommandResult("version", nil, []byte("2.5.0"), nil)

	// 测试执行已知命令
	output, err := executor.Execute("version")
	if err != nil {
		t.Errorf("Execute failed: %v", err)
	}
	if string(output) != "2.5.0" {
		t.Errorf("Expected '2.5.0', got '%s'", string(output))
	}

	// 测试执行未知命令
	_, err = executor.Execute("unknown")
	if err == nil {
		t.Error("Execute unknown command should return error")
	}
}

func TestMockCommandExecutor_ExecuteWithArgs(t *testing.T) {
	executor := NewMockCommandExecutor()

	executor.SetCommandResult("require", []string{"--dev", "pkg"}, []byte("installed"), nil)

	output, err := executor.Execute("require", "--dev", "pkg")
	if err != nil {
		t.Errorf("Execute with args failed: %v", err)
	}
	if string(output) != "installed" {
		t.Errorf("Expected 'installed', got '%s'", string(output))
	}
}

func TestNewMockFileSystemHelper(t *testing.T) {
	helper := NewMockFileSystemHelper()
	if helper == nil {
		t.Fatal("NewMockFileSystemHelper should return non-nil helper")
	}

	// 验证默认函数不返回错误
	err := helper.CreateFile("/tmp/test", []byte("content"), 0644)
	if err != nil {
		t.Errorf("Default CreateFile should not return error: %v", err)
	}

	err = helper.CheckWritePermission("/tmp")
	if err != nil {
		t.Errorf("Default CheckWritePermission should not return error: %v", err)
	}

	err = helper.EnsureDirectoryExists("/tmp/testdir")
	if err != nil {
		t.Errorf("Default EnsureDirectoryExists should not return error: %v", err)
	}

	err = helper.RemoveFile("/tmp/testfile")
	if err != nil {
		t.Errorf("Default RemoveFile should not return error: %v", err)
	}
}

func TestMockFileSystemHelper_CustomFunctions(t *testing.T) {
	createErr := errors.New("create error")
	checkErr := errors.New("check error")
	ensureErr := errors.New("ensure error")
	removeErr := errors.New("remove error")

	helper := &MockFileSystemHelper{
		CreateFileFunc: func(path string, content []byte, perm os.FileMode) error {
			return createErr
		},
		CheckWritePermissionFunc: func(path string) error {
			return checkErr
		},
		EnsureDirectoryExistsFunc: func(path string) error {
			return ensureErr
		},
		RemoveFileFunc: func(path string) error {
			return removeErr
		},
	}

	if helper.CreateFile("/tmp/test", nil, 0) != createErr {
		t.Error("CreateFile should return custom error")
	}
	if helper.CheckWritePermission("/tmp") != checkErr {
		t.Error("CheckWritePermission should return custom error")
	}
	if helper.EnsureDirectoryExists("/tmp") != ensureErr {
		t.Error("EnsureDirectoryExists should return custom error")
	}
	if helper.RemoveFile("/tmp") != removeErr {
		t.Error("RemoveFile should return custom error")
	}
}

func TestNewMockDownloadHelper(t *testing.T) {
	helper := NewMockDownloadHelper()
	if helper == nil {
		t.Fatal("NewMockDownloadHelper should return non-nil helper")
	}

	// 验证默认函数不返回错误
	err := helper.DownloadFile("http://example.com/file.zip", "/tmp/file.zip", nil)
	if err != nil {
		t.Errorf("Default DownloadFile should not return error: %v", err)
	}
}

func TestMockDownloadHelper_CustomFunction(t *testing.T) {
	downloadErr := errors.New("download failed")
	helper := &MockDownloadHelper{
		DownloadFileFunc: func(url string, target string, config interface{}) error {
			return downloadErr
		},
	}

	if helper.DownloadFile("http://example.com", "/tmp/file", nil) != downloadErr {
		t.Error("DownloadFile should return custom error")
	}
}

func TestNewMockRuntime(t *testing.T) {
	runtime := NewMockRuntime()
	if runtime == nil {
		t.Fatal("NewMockRuntime should return non-nil runtime")
	}

	if runtime.GetOS() != "linux" {
		t.Errorf("Default OS should be 'linux', got '%s'", runtime.GetOS())
	}
	if runtime.GetArch() != "amd64" {
		t.Errorf("Default Arch should be 'amd64', got '%s'", runtime.GetArch())
	}
}

func TestMockRuntime_SetOS(t *testing.T) {
	runtime := NewMockRuntime()

	testCases := []struct {
		os   string
		want string
	}{
		{"darwin", "darwin"},
		{"windows", "windows"},
		{"linux", "linux"},
		{"freebsd", "freebsd"},
	}

	for _, tc := range testCases {
		runtime.SetOS(tc.os)
		if runtime.GetOS() != tc.want {
			t.Errorf("SetOS(%s): got '%s', want '%s'", tc.os, runtime.GetOS(), tc.want)
		}
	}
}

func TestMockRuntime_SetArch(t *testing.T) {
	runtime := NewMockRuntime()

	testCases := []struct {
		arch string
		want string
	}{
		{"amd64", "amd64"},
		{"arm64", "arm64"},
		{"386", "386"},
		{"arm", "arm"},
	}

	for _, tc := range testCases {
		runtime.SetArch(tc.arch)
		if runtime.GetArch() != tc.want {
			t.Errorf("SetArch(%s): got '%s', want '%s'", tc.arch, runtime.GetArch(), tc.want)
		}
	}
}

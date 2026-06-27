package composer

import (
	"errors"
	"testing"
)

func TestArchive(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for archive command
	SetupMockOutput("archive --format=zip --dir=/tmp", "Archive created successfully", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.Archive("/tmp")
	if err != nil {
		t.Errorf("Archive执行失败: %v", err)
	}

	if !contains(output, "Archive created") {
		t.Errorf("输出应包含归档创建信息，实际为\"%s\"", output)
	}
}

func TestArchiveWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试目标目录不存在
	SetupMockOutput("archive --format=zip --dir=/nonexistent/path", "", errors.New("Directory does not exist"))
	_, err = composer.Archive("/nonexistent/path")
	if err == nil {
		t.Error("不存在的目标目录应该返回错误")
	}

	// 测试权限错误
	SetupMockOutput("archive --format=zip --dir=/no/permission", "", errors.New("Permission denied"))
	_, err = composer.Archive("/no/permission")
	if err == nil {
		t.Error("权限错误应该返回错误")
	}

	// 测试空目标路径
	SetupMockOutput("archive --format=zip --dir=", "", errors.New("Invalid destination"))
	_, err = composer.Archive("")
	if err == nil {
		t.Error("空目标路径应该返回错误")
	}
}

func TestArchiveWithOptions(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试单个选项 - 由于map迭代顺序不确定，我们使用单个选项来确保命令匹配
	SetupMockOutput("archive --dir=/tmp --format=zip", "Archive created in zip format", nil)
	options := map[string]string{
		"format": "zip",
	}
	output, err := composer.ArchiveWithOptions("/tmp", options)
	if err != nil {
		t.Errorf("ArchiveWithOptions(zip)执行失败: %v", err)
	}
	if !contains(output, "zip") {
		t.Errorf("输出应包含格式信息，实际为\"%s\"", output)
	}

	// 测试另一个单个选项
	SetupMockOutput("archive --dir=/tmp --ignore-filters", "Archive created ignoring filters", nil)
	options = map[string]string{
		"ignore-filters": "",
	}
	output, err = composer.ArchiveWithOptions("/tmp", options)
	if err != nil {
		t.Errorf("ArchiveWithOptions(ignore-filters)执行失败: %v", err)
	}
	if !contains(output, "ignoring filters") {
		t.Errorf("输出应包含忽略过滤器信息，实际为\"%s\"", output)
	}
}

func TestArchiveWithFormat(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试ZIP格式
	SetupMockOutput("archive --format=zip --dir=/tmp", "Archive created in zip format", nil)
	output, err := composer.ArchiveWithFormat("/tmp", "zip")
	if err != nil {
		t.Errorf("ArchiveWithFormat(zip)执行失败: %v", err)
	}
	if !contains(output, "zip") {
		t.Errorf("输出应包含zip格式信息，实际为\"%s\"", output)
	}

	// 测试TAR格式
	SetupMockOutput("archive --format=tar --dir=/tmp", "Archive created in tar format", nil)
	output, err = composer.ArchiveWithFormat("/tmp", "tar")
	if err != nil {
		t.Errorf("ArchiveWithFormat(tar)执行失败: %v", err)
	}
	if !contains(output, "tar") {
		t.Errorf("输出应包含tar格式信息，实际为\"%s\"", output)
	}
}

func TestArchivePackage(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试归档指定版本的包
	SetupMockOutput("archive symfony/console=v5.4.0 --dir=/tmp", "Package symfony/console v5.4.0 archived", nil)
	output, err := composer.ArchivePackage("symfony/console", "v5.4.0", "/tmp")
	if err != nil {
		t.Errorf("ArchivePackage执行失败: %v", err)
	}
	if !contains(output, "symfony/console") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}

	// 测试归档最新版本的包
	SetupMockOutput("archive symfony/console --dir=/tmp", "Package symfony/console latest archived", nil)
	output, err = composer.ArchivePackage("symfony/console", "", "/tmp")
	if err != nil {
		t.Errorf("ArchivePackage(latest)执行失败: %v", err)
	}
	if !contains(output, "symfony/console") {
		t.Errorf("输出应包含包名，实际为\"%s\"", output)
	}
}

func TestArchivePackageWithOptions(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试使用自定义选项归档包 - 由于map迭代顺序不确定，使用单个选项
	SetupMockOutput("archive symfony/console=v5.4.0 --dir=/tmp --format=tar", "Package archived with custom options", nil)
	options := map[string]string{
		"format": "tar",
	}
	output, err := composer.ArchivePackageWithOptions("symfony/console", "v5.4.0", "/tmp", options)
	if err != nil {
		t.Errorf("ArchivePackageWithOptions执行失败: %v", err)
	}
	if !contains(output, "custom options") {
		t.Errorf("输出应包含自定义选项信息，实际为\"%s\"", output)
	}
}

func TestArchivePackageWithErrors(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试包不存在
	SetupMockOutput("archive nonexistent/package --dir=/tmp", "", errors.New("Package not found"))
	_, err = composer.ArchivePackage("nonexistent/package", "", "/tmp")
	if err == nil {
		t.Error("不存在的包应该返回错误")
	}

	// 测试无效版本
	SetupMockOutput("archive symfony/console=invalid-version --dir=/tmp", "", errors.New("Invalid version"))
	_, err = composer.ArchivePackage("symfony/console", "invalid-version", "/tmp")
	if err == nil {
		t.Error("无效版本应该返回错误")
	}

	// 测试空包名
	SetupMockOutput("archive  --dir=/tmp", "", errors.New("Package name cannot be empty"))
	_, err = composer.ArchivePackage("", "", "/tmp")
	if err == nil {
		t.Error("空包名应该返回错误")
	}
}

func TestArchiveEdgeCases(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试包含特殊字符的路径
	SetupMockOutput("archive --format=zip --dir=/path with spaces", "Archive created in special path", nil)
	_, err = composer.Archive("/path with spaces")
	if err != nil {
		t.Errorf("特殊字符路径Archive执行失败: %v", err)
	}

	// 测试非常长的路径
	longPath := "/very/long/path/" + "directory/"
	for i := 0; i < 10; i++ {
		longPath += "subdirectory/"
	}
	SetupMockOutput("archive --format=zip --dir="+longPath, "Archive created in long path", nil)
	_, err = composer.Archive(longPath)
	if err != nil {
		t.Errorf("长路径Archive执行失败: %v", err)
	}

	// 测试Unicode路径
	unicodePath := "/路径/with/中文"
	SetupMockOutput("archive --format=zip --dir="+unicodePath, "Archive created in unicode path", nil)
	_, err = composer.Archive(unicodePath)
	if err != nil {
		t.Errorf("Unicode路径Archive执行失败: %v", err)
	}
}

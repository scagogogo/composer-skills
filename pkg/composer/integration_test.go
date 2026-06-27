package composer

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestCompleteProjectWorkflow 测试完整的项目工作流程
func TestCompleteProjectWorkflow(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 1. 初始化项目
	SetupMockOutput("init", "Project initialized successfully", nil)
	err = composer.InitProject()
	if err != nil {
		t.Errorf("初始化项目失败: %v", err)
	}

	// 2. 添加依赖
	SetupMockOutput("require monolog/monolog", "Package monolog/monolog added successfully", nil)
	err = composer.RequirePackage("monolog/monolog", "", false)
	if err != nil {
		t.Errorf("添加依赖失败: %v", err)
	}

	// 3. 安装依赖
	SetupMockOutput("install", "Dependencies installed successfully", nil)
	err = composer.Install(false, false)
	if err != nil {
		t.Errorf("安装依赖失败: %v", err)
	}

	// 4. 运行安全审计
	SetupMockOutput("audit", "No security vulnerabilities found", nil)
	auditOutput, err := composer.Audit()
	if err != nil {
		t.Errorf("安全审计失败: %v", err)
	}
	if !contains(auditOutput, "No security vulnerabilities") {
		t.Errorf("安全审计输出不正确: %s", auditOutput)
	}

	// 5. 更新依赖
	SetupMockOutput("update", "Dependencies updated successfully", nil)
	err = composer.Update([]string{}, false)
	if err != nil {
		t.Errorf("更新依赖失败: %v", err)
	}

	// 6. 生成自动加载文件
	SetupMockOutput("dump-autoload", "Autoload files generated successfully", nil)
	err = composer.DumpAutoload(false)
	if err != nil {
		t.Errorf("生成自动加载文件失败: %v", err)
	}
}

// TestPackageManagementWorkflow 测试包管理工作流程
func TestPackageManagementWorkflow(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 1. 搜索包
	SetupMockOutput("search logger", "Found 5 packages matching logger", nil)
	searchOutput, err := composer.Search("logger")
	if err != nil {
		t.Errorf("搜索包失败: %v", err)
	}
	if !contains(searchOutput, "Found") {
		t.Errorf("搜索输出不正确: %s", searchOutput)
	}

	// 2. 查看包信息
	SetupMockOutput("show monolog/monolog", "Package monolog/monolog\nVersion: 2.8.0", nil)
	showOutput, err := composer.ShowPackage("monolog/monolog")
	if err != nil {
		t.Errorf("查看包信息失败: %v", err)
	}
	if !contains(showOutput, "monolog/monolog") {
		t.Errorf("包信息输出不正确: %s", showOutput)
	}

	// 3. 添加包
	SetupMockOutput("require monolog/monolog:^2.0", "Package added successfully", nil)
	err = composer.RequirePackage("monolog/monolog", "^2.0", false)
	if err != nil {
		t.Errorf("添加包失败: %v", err)
	}

	// 4. 查看依赖树
	SetupMockOutput("show --tree", "Dependency tree:\nmonolog/monolog v2.8.0", nil)
	treeOutput, err := composer.ShowDependencyTree("")
	if err != nil {
		t.Errorf("查看依赖树失败: %v", err)
	}
	if !contains(treeOutput, "Dependency tree") {
		t.Errorf("依赖树输出不正确: %s", treeOutput)
	}

	// 5. 检查过期包
	SetupMockOutput("outdated", "No outdated packages found", nil)
	outdatedOutput, err := composer.OutdatedPackages()
	if err != nil {
		t.Errorf("检查过期包失败: %v", err)
	}
	if !contains(outdatedOutput, "No outdated") {
		t.Errorf("过期包检查输出不正确: %s", outdatedOutput)
	}

	// 6. 移除包
	SetupMockOutput("remove monolog/monolog", "Package removed successfully", nil)
	err = composer.Remove("monolog/monolog", false)
	if err != nil {
		t.Errorf("移除包失败: %v", err)
	}
}

// TestConfigurationWorkflow 测试配置管理工作流程
func TestConfigurationWorkflow(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 设置mock composer.json
	mockJSON := &ComposerJSON{
		Config: map[string]interface{}{
			"vendor-dir": "vendor",
			"bin-dir":    "bin",
		},
	}
	SetMockComposerJSON(mockJSON)
	defer ClearMockComposerJSON()

	// 1. 获取配置
	vendorDir, err := composer.GetConfig("vendor-dir")
	if err != nil {
		t.Errorf("获取vendor-dir配置失败: %v", err)
	}
	if vendorDir != "vendor" {
		t.Errorf("vendor-dir配置不正确，期望: vendor，实际: %v", vendorDir)
	}

	// 2. 设置配置
	err = composer.SetConfig("process-timeout", 600)
	if err != nil {
		t.Errorf("设置process-timeout配置失败: %v", err)
	}

	// 3. 验证配置已设置
	timeout, err := composer.GetConfig("process-timeout")
	if err != nil {
		t.Errorf("获取process-timeout配置失败: %v", err)
	}
	if timeout != 600 {
		t.Errorf("process-timeout配置不正确，期望: 600，实际: %v", timeout)
	}

	// 4. 设置工作目录
	composer.SetWorkingDir("/tmp/test-project")
	if composer.workingDir != "/tmp/test-project" {
		t.Errorf("工作目录设置不正确，期望: /tmp/test-project，实际: %s", composer.workingDir)
	}

	// 5. 验证composer.json
	SetupMockOutput("validate", "composer.json is valid", nil)
	err = composer.Validate()
	if err != nil {
		t.Errorf("验证composer.json失败: %v", err)
	}
}

// TestErrorHandlingWorkflow 测试错误处理工作流程
func TestErrorHandlingWorkflow(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 1. 测试网络错误恢复
	SetupMockOutput("require nonexistent/package", "", errors.New("Network error"))
	err = composer.RequirePackage("nonexistent/package", "", false)
	if err == nil {
		t.Error("网络错误应该被正确处理")
	}

	// 2. 测试权限错误
	SetupMockOutput("install", "", errors.New("Permission denied"))
	err = composer.Install(false, false)
	if err == nil {
		t.Error("权限错误应该被正确处理")
	}

	// 3. 测试超时处理
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	SetupMockOutput("update", "", errors.New("context deadline exceeded"))
	_, err = composer.RunWithContext(ctx, "update")
	if err == nil {
		t.Error("超时错误应该被正确处理")
	}

	// 4. 测试无效JSON处理
	SetupMockOutput("audit --format=json", "invalid json", nil)
	_, err = composer.HasVulnerabilities()
	if err == nil {
		t.Error("无效JSON应该被正确处理")
	}
}

// TestConcurrentOperations 测试并发操作
func TestConcurrentOperations(t *testing.T) {
	ClearMockOutputs()
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 设置多个mock输出
	SetupMockOutput("show package1", "Package package1 info", nil)
	SetupMockOutput("show package2", "Package package2 info", nil)
	SetupMockOutput("show package3", "Package package3 info", nil)

	// 并发执行多个操作
	done := make(chan bool, 3)
	packages := []string{"package1", "package2", "package3"}

	for _, pkg := range packages {
		go func(packageName string) {
			defer func() { done <- true }()
			_, err := composer.ShowPackage(packageName)
			if err != nil {
				t.Errorf("并发查看包 %s 失败: %v", packageName, err)
			}
		}(pkg)
	}

	// 等待所有操作完成
	for i := 0; i < 3; i++ {
		<-done
	}
}

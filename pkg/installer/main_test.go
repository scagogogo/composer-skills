package installer

import (
	"os"
	"testing"
)

// TestMain 是Go测试的入口函数，用于设置和清理全局测试环境
func TestMain(m *testing.M) {
	// 在所有测试运行前进行设置
	setup()

	// 运行所有测试
	exitCode := m.Run()

	// 测试完成后进行清理
	teardown()

	// 使用测试的返回码退出
	os.Exit(exitCode)
}

// setup 为测试准备环境
func setup() {
	// 创建测试所需的临时目录或资源
	// 这里暂时没有全局设置，未来可以根据需要添加
}

// teardown 清理测试环境
func teardown() {
	// 清理创建的临时目录或资源
	// 这里暂时没有全局清理，未来可以根据需要添加
}

// 这个模拟测试包用于集成测试
// 实际项目中可以根据需要使用真实的utils包或创建一个模拟utils包

// 添加一些辅助函数，方便多个测试文件使用

package composer

import (
	"sync"
)

// MockOutput 存储模拟命令的输出和错误
type MockOutput struct {
	Output string
	Error  error
}

// 全局模拟输出映射，用于存储命令到输出的映射
var (
	mockOutputs = make(map[string]MockOutput)
	mockMutex   sync.RWMutex
)

// SetupMockOutputAdvanced 设置模拟命令的输出（高级版本）
func SetupMockOutputAdvanced(command, output string, err error) {
	mockMutex.Lock()
	defer mockMutex.Unlock()
	mockOutputs[command] = MockOutput{
		Output: output,
		Error:  err,
	}
}

// ClearMockOutputsAdvanced 清除所有模拟输出（高级版本）
func ClearMockOutputsAdvanced() {
	mockMutex.Lock()
	defer mockMutex.Unlock()
	mockOutputs = make(map[string]MockOutput)
}

// GetMockOutput 获取模拟命令的输出
func GetMockOutput(command string) (MockOutput, bool) {
	mockMutex.RLock()
	defer mockMutex.RUnlock()
	output, exists := mockOutputs[command]
	return output, exists
}

// createTestComposerLock 创建一个测试用的composer.lock文件

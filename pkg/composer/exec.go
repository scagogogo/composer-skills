package composer

import (
	"strings"
)

// Exec 执行厂商二进制文件
func (c *Composer) Exec(binary string, args ...string) (string, error) {
	execArgs := append([]string{"exec", binary, "--"}, args...)
	return c.Run(execArgs...)
}

// ExecWithList 列出所有可用的厂商二进制文件
func (c *Composer) ExecWithList() ([]string, error) {
	output, err := c.Run("exec", "--list")
	if err != nil {
		return nil, err
	}

	// 分割输出并移除空行
	lines := strings.Split(output, "\n")
	var binaries []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "Available binaries:") {
			binaries = append(binaries, line)
		}
	}

	return binaries, nil
}

// ExecPHP 使用指定的 PHP 二进制文件执行厂商二进制文件
func (c *Composer) ExecPHP(php string, binary string, args ...string) (string, error) {
	execArgs := append([]string{"exec", "--php=" + php, binary, "--"}, args...)
	return c.Run(execArgs...)
}

// ExecWithWorkingDir 在指定的工作目录中执行厂商二进制文件
func (c *Composer) ExecWithWorkingDir(binary string, workingDir string, args ...string) (string, error) {
	// 保存原始工作目录
	originalWorkingDir := c.workingDir
	// 设置新工作目录
	c.workingDir = workingDir

	// 在函数返回时恢复原始工作目录
	defer func() {
		c.workingDir = originalWorkingDir
	}()

	return c.Exec(binary, args...)
}

// ExecAll 执行所有可用的厂商二进制文件
// 返回一个映射，键为二进制文件名，值为执行结果
func (c *Composer) ExecAll(args ...string) (map[string]string, error) {
	// 获取所有可用的二进制文件
	binaries, err := c.ExecWithList()
	if err != nil {
		return nil, err
	}

	// 执行每个二进制文件
	results := make(map[string]string)

	for _, binary := range binaries {
		output, err := c.Exec(binary, args...)
		if err != nil {
			results[binary] = "Error: " + err.Error()
		} else {
			results[binary] = output
		}
	}

	return results, nil
}

// ExecCommand 执行特定命令并返回其输出
// 这是一个更通用的方法，允许执行自定义命令
func (c *Composer) ExecCommand(command string, args ...string) (string, error) {
	allArgs := append([]string{command}, args...)
	return c.Run(allArgs...)
}

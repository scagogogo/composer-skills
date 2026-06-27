package composer

// ShellType 表示 shell 类型
type ShellType string

const (
	// Bash shell
	BashShell ShellType = "bash"
	// Zsh shell
	ZshShell ShellType = "zsh"
	// Fish shell
	FishShell ShellType = "fish"
)

// GenerateCompletion 生成指定 shell 的自动补全脚本
func (c *Composer) GenerateCompletion(shell ShellType) (string, error) {
	return c.Run("completion", string(shell))
}

// GenerateCompletionWithOptions 使用更多选项生成自动补全脚本
func (c *Composer) GenerateCompletionWithOptions(shell ShellType, options map[string]string) (string, error) {
	args := []string{"completion", string(shell)}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ListCommands 列出所有可用的命令
func (c *Composer) ListCommands() (string, error) {
	return c.Run("list")
}

// GetCommandHelp 获取特定命令的帮助信息
func (c *Composer) GetCommandHelp(command string) (string, error) {
	return c.Run("help", command)
}

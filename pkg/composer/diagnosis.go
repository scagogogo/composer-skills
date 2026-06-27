package composer

// Status 显示已安装包的本地修改
func (c *Composer) Status() (string, error) {
	return c.Run("status")
}

// StatusWithOptions 使用更多选项显示已安装包的本地修改
func (c *Composer) StatusWithOptions(options map[string]string) (string, error) {
	args := []string{"status"}

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

// Diagnose 诊断系统以识别常见错误
func (c *Composer) Diagnose() (string, error) {
	return c.Run("diagnose")
}

// DiagnoseWithOptions 使用更多选项诊断系统
func (c *Composer) DiagnoseWithOptions(options map[string]string) (string, error) {
	args := []string{"diagnose"}

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

// LocalExec 执行本地包中的二进制文件
func (c *Composer) LocalExec(command string, args ...string) (string, error) {
	cmdArgs := append([]string{"exec", command}, args...)
	return c.Run(cmdArgs...)
}

// LocalExecWithOptions 使用更多选项执行本地包中的二进制文件
func (c *Composer) LocalExecWithOptions(command string, options map[string]string, args ...string) (string, error) {
	cmdArgs := []string{"exec"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			cmdArgs = append(cmdArgs, "--"+key)
		} else {
			cmdArgs = append(cmdArgs, "--"+key+"="+value)
		}
	}

	cmdArgs = append(cmdArgs, command)
	cmdArgs = append(cmdArgs, args...)

	return c.Run(cmdArgs...)
}

// Check 检查依赖项是否满足要求
func (c *Composer) Check() (string, error) {
	return c.Run("check")
}

// CheckWithOptions 使用更多选项检查依赖项
func (c *Composer) CheckWithOptions(options map[string]string) (string, error) {
	args := []string{"check"}

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

package composer

// Licenses 显示依赖包的许可证信息
func (c *Composer) Licenses() (string, error) {
	return c.Run("licenses")
}

// LicensesWithFormat 显示特定格式的依赖包许可证信息
func (c *Composer) LicensesWithFormat(format string) (string, error) {
	return c.Run("licenses", "--format="+format)
}

// LicensesWithOptions 使用更多选项显示依赖包许可证信息
func (c *Composer) LicensesWithOptions(options map[string]string) (string, error) {
	args := []string{"licenses"}

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

// CheckLicenses 检查依赖包的许可证兼容性
func (c *Composer) CheckLicenses() (string, error) {
	return c.Run("licenses", "--check")
}

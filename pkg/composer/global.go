package composer

// GlobalRequire 全局安装包
func (c *Composer) GlobalRequire(packageName string, version string) error {
	args := []string{"global", "require"}

	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// GlobalUpdate 全局更新包
func (c *Composer) GlobalUpdate(packages []string) error {
	args := []string{"global", "update"}
	args = append(args, packages...)

	_, err := c.Run(args...)
	return err
}

// GlobalRemove 全局移除包
func (c *Composer) GlobalRemove(packageName string) error {
	_, err := c.Run("global", "remove", packageName)
	return err
}

// GlobalInstall 全局安装依赖
func (c *Composer) GlobalInstall() error {
	_, err := c.Run("global", "install")
	return err
}

// GlobalList 列出全局安装的包
func (c *Composer) GlobalList() (string, error) {
	return c.Run("global", "show")
}

// GlobalHome 获取全局目录路径
func (c *Composer) GlobalHome() (string, error) {
	return c.Run("global", "home")
}

// GlobalExecute 执行全局安装的包中的二进制文件
func (c *Composer) GlobalExecute(command string, args ...string) (string, error) {
	cmdArgs := append([]string{"global", "exec", command}, args...)
	return c.Run(cmdArgs...)
}

// GlobalStatus 显示全局安装的包的状态
func (c *Composer) GlobalStatus() (string, error) {
	return c.Run("global", "status")
}

// GlobalDumpAutoload 为全局安装生成自动加载文件
func (c *Composer) GlobalDumpAutoload(optimize bool) error {
	args := []string{"global", "dump-autoload"}

	if optimize {
		args = append(args, "--optimize")
	}

	_, err := c.Run(args...)
	return err
}

// GlobalRequireWithOptions 使用自定义选项全局安装包
//
// 参数：
//   - packageName: 要全局安装的包名，例如"symfony/console"
//   - version: 版本约束，例如"^5.0"，如果为空则使用最新版本
//   - options: 全局安装包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果全局安装包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项全局安装包。支持更多自定义选项，
//	如--prefer-source、--prefer-dist、--no-progress等。
//
// 用法示例：
//
//	// 全局安装包并指定多个选项
//	options := map[string]string{
//	    "prefer-dist": "",
//	    "no-progress": "",
//	    "no-suggest":  "",
//	}
//	err := comp.GlobalRequireWithOptions("symfony/console", "^5.0", options)
//	if err != nil {
//	    log.Fatalf("全局安装包失败: %v", err)
//	}
func (c *Composer) GlobalRequireWithOptions(packageName string, version string, options map[string]string) error {
	args := []string{"global", "require"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}

	_, err := c.Run(args...)
	return err
}

// GlobalUpdateWithOptions 使用自定义选项全局更新包
//
// 参数：
//   - packages: 要更新的包名列表，为空则更新所有包
//   - options: 全局更新包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果全局更新包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项全局更新包。支持更多自定义选项，
//	如--prefer-source、--prefer-dist、--no-dev等。
//
// 用法示例：
//
//	// 全局更新包并指定多个选项
//	options := map[string]string{
//	    "prefer-dist": "",
//	    "no-dev":      "",
//	    "no-progress": "",
//	}
//	err := comp.GlobalUpdateWithOptions([]string{"symfony/console"}, options)
//	if err != nil {
//	    log.Fatalf("全局更新包失败: %v", err)
//	}
func (c *Composer) GlobalUpdateWithOptions(packages []string, options map[string]string) error {
	args := []string{"global", "update"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	args = append(args, packages...)

	_, err := c.Run(args...)
	return err
}

// GlobalRemoveWithOptions 使用自定义选项全局移除包
//
// 参数：
//   - packageName: 要全局移除的包名
//   - options: 全局移除包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果全局移除包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项全局移除包。支持更多自定义选项，
//	如--no-progress、--no-update等。
//
// 用法示例：
//
//	// 全局移除包并指定多个选项
//	options := map[string]string{
//	    "no-progress": "",
//	    "no-update":   "",
//	}
//	err := comp.GlobalRemoveWithOptions("symfony/console", options)
//	if err != nil {
//	    log.Fatalf("全局移除包失败: %v", err)
//	}
func (c *Composer) GlobalRemoveWithOptions(packageName string, options map[string]string) error {
	args := []string{"global", "remove"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	args = append(args, packageName)

	_, err := c.Run(args...)
	return err
}

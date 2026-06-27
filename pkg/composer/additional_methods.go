package composer

// This file contains additional Composer SDK methods that provide
// more convenience wrappers around common Composer CLI operations.

// ==================== Outdated Enhancements ====================

// OutdatedWithOptions displays outdated packages with custom options
//
// 参数：
//   - options: 自定义选项，键为选项名，值为选项值
//
// 返回值：
//   - string: 过时包列表的输出
//   - error: 如果获取过时包列表过程中发生错误
func (c *Composer) OutdatedWithOptions(options map[string]string) (string, error) {
	args := []string{"outdated"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// OutdatedWithFormat displays outdated packages in the specified format
// This is an alias for ShowOutdatedWithFormat with a more intuitive name
//
// 参数：
//   - format: 输出格式，如"json"、"text"、"cli"
//
// 返回值：
//   - string: 过时包列表的输出
//   - error: 如果获取过时包列表过程中发生错误
func (c *Composer) OutdatedWithFormat(format string) (string, error) {
	return c.Run("outdated", "--format", format)
}

// ==================== Dry Run Support ====================

// InstallDryRun simulates an install without actually installing anything
//
// 返回值：
//   - string: 模拟安装的输出结果
//   - error: 如果模拟安装过程中发生错误
func (c *Composer) InstallDryRun() (string, error) {
	return c.Run("install", "--dry-run")
}

// UpdateDryRun simulates an update without actually updating anything
//
// 参数：
//   - packages: 要模拟更新的包名列表，为空则模拟更新所有包
//
// 返回值：
//   - string: 模拟更新的输出结果
//   - error: 如果模拟更新过程中发生错误
func (c *Composer) UpdateDryRun(packages []string) (string, error) {
	args := []string{"update", "--dry-run"}
	args = append(args, packages...)
	return c.Run(args...)
}

// RequireDryRun simulates adding a package without actually adding it
//
// 参数：
//   - packageName: 要模拟添加的包名
//   - version: 版本约束
//
// 返回值：
//   - string: 模拟添加的输出结果
//   - error: 如果模拟添加过程中发生错误
func (c *Composer) RequireDryRun(packageName string, version string) (string, error) {
	args := []string{"require", "--dry-run"}
	if version != "" {
		args = append(args, packageName+":"+version)
	} else {
		args = append(args, packageName)
	}
	return c.Run(args...)
}

// RemoveDryRun simulates removing a package without actually removing it
//
// 参数：
//   - packageName: 要模拟移除的包名
//
// 返回值：
//   - string: 模拟移除的输出结果
//   - error: 如果模拟移除过程中发生错误
func (c *Composer) RemoveDryRun(packageName string) (string, error) {
	return c.Run("remove", "--dry-run", packageName)
}

// ==================== Require Multiple ====================

// RequireMultiple adds multiple packages as dependencies at once
//
// 参数：
//   - packages: 包名到版本约束的映射，例如 map[string]string{"symfony/console": "^5.0", "monolog/monolog": "^2.0"}
//   - dev: 是否作为开发依赖添加
//
// 返回值：
//   - error: 如果添加包过程中发生错误
func (c *Composer) RequireMultiple(packages map[string]string, dev bool) error {
	args := []string{"require"}

	if dev {
		args = append(args, "--dev")
	}

	for pkg, version := range packages {
		if version != "" {
			args = append(args, pkg+":"+version)
		} else {
			args = append(args, pkg)
		}
	}

	_, err := c.Run(args...)
	return err
}

// ==================== Depends / Why Enhancements ====================

// DependsWithOptions shows which packages depend on the given package with options
//
// 参数：
//   - packageName: 要查询反向依赖的包名
//   - options: 自定义选项
//
// 返回值：
//   - string: 反向依赖的输出结果
//   - error: 如果获取反向依赖过程中发生错误
func (c *Composer) DependsWithOptions(packageName string, options map[string]string) (string, error) {
	args := []string{"depends", packageName}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// WhyWithOptions explains why a package is installed with custom options
//
// 参数：
//   - packageName: 要查询原因的包名
//   - options: 自定义选项
//
// 返回值：
//   - string: 安装原因的输出结果
//   - error: 如果查询安装原因过程中发生错误
func (c *Composer) WhyWithOptions(packageName string, options map[string]string) (string, error) {
	args := []string{"why", packageName}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// WhyNotWithOptions explains why a package version cannot be installed with options
//
// 参数：
//   - packageName: 要查询的包名
//   - version: 要查询的版本
//   - options: 自定义选项
//
// 返回值：
//   - string: 解释结果的输出
//   - error: 如果查询过程中发生错误
func (c *Composer) WhyNotWithOptions(packageName string, version string, options map[string]string) (string, error) {
	args := []string{"why-not", packageName, version}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ==================== Suggests Enhancements ====================

// SuggestsWithOptions shows suggested packages with custom options
//
// 参数：
//   - options: 自定义选项
//
// 返回值：
//   - string: 建议包的输出
//   - error: 如果获取建议包过程中发生错误
func (c *Composer) SuggestsWithOptions(options map[string]string) (string, error) {
	args := []string{"suggests"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// SuggestsForPackage shows suggested packages for a specific package
//
// 参数：
//   - packageName: 要查询建议包的包名
//
// 返回值：
//   - string: 建议包的输出
//   - error: 如果获取建议包过程中发生错误
func (c *Composer) SuggestsForPackage(packageName string) (string, error) {
	return c.Run("suggests", packageName)
}

// ==================== Reinstall Enhancements ====================

// ReinstallMultipleWithOptions reinstalls multiple packages with custom options
//
// 参数：
//   - packages: 要重新安装的包名列表
//   - options: 自定义选项
//
// 返回值：
//   - error: 如果重新安装过程中发生错误
func (c *Composer) ReinstallMultipleWithOptions(packages []string, options map[string]string) error {
	args := []string{"reinstall"}

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

// ==================== Global Enhancements ====================

// GlobalInit initializes a new project in the global directory
//
// 参数：
//   - name: 项目名称
//
// 返回值：
//   - error: 如果初始化过程中发生错误
func (c *Composer) GlobalInit(name string) error {
	args := []string{"global", "init", "--name=" + name, "--no-interaction"}
	_, err := c.Run(args...)
	return err
}

// GlobalRequireMultiple globally requires multiple packages at once
//
// 参数：
//   - packages: 包名到版本约束的映射
//
// 返回值：
//   - error: 如果全局安装包过程中发生错误
func (c *Composer) GlobalRequireMultiple(packages map[string]string) error {
	args := []string{"global", "require"}

	for pkg, version := range packages {
		if version != "" {
			args = append(args, pkg+":"+version)
		} else {
			args = append(args, pkg)
		}
	}

	_, err := c.Run(args...)
	return err
}

// GlobalRemoveMultiple globally removes multiple packages at once
//
// 参数：
//   - packages: 要全局移除的包名列表
//
// 返回值：
//   - error: 如果全局移除包过程中发生错误
func (c *Composer) GlobalRemoveMultiple(packages []string) error {
	args := []string{"global", "remove"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// ==================== Show Enhancements ====================

// ShowWithOptions shows package information with custom options
//
// 参数：
//   - options: 自定义选项
//
// 返回值：
//   - string: 包信息的输出
//   - error: 如果获取包信息过程中发生错误
func (c *Composer) ShowWithOptions(options map[string]string) (string, error) {
	args := []string{"show"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ShowLatestVersions shows the latest available versions of packages
//
// 返回值：
//   - string: 最新版本信息的输出
//   - error: 如果获取最新版本过程中发生错误
func (c *Composer) ShowLatestVersions() (string, error) {
	return c.Run("show", "--latest")
}

// ShowOutdatedMinorOnly shows only packages with minor version updates available
//
// 返回值：
//   - string: 过时包列表的输出
//   - error: 如果获取过时包列表过程中发生错误
func (c *Composer) ShowOutdatedMinorOnly() (string, error) {
	return c.Run("outdated", "--minor-only")
}

// ==================== Install/Update with specific options ====================

// InstallNoDev installs dependencies without development packages
//
// 返回值：
//   - error: 如果安装过程中发生错误
func (c *Composer) InstallNoDev() error {
	_, err := c.Run("install", "--no-dev")
	return err
}

// UpdateNoDev updates dependencies without development packages
//
// 参数：
//   - packages: 要更新的包名列表，为空则更新所有包
//
// 返回值：
//   - error: 如果更新过程中发生错误
func (c *Composer) UpdateNoDev(packages []string) error {
	args := []string{"update", "--no-dev"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// UpdateWithDependencies updates packages and their dependencies
//
// 参数：
//   - packages: 要更新的包名列表
//
// 返回值：
//   - error: 如果更新过程中发生错误
func (c *Composer) UpdateWithDependencies(packages []string) error {
	args := []string{"update", "--with-dependencies"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// UpdateWithAllDependencies updates packages and all their dependencies recursively
//
// 参数：
//   - packages: 要更新的包名列表
//
// 返回值：
//   - error: 如果更新过程中发生错误
func (c *Composer) UpdateWithAllDependencies(packages []string) error {
	args := []string{"update", "--with-all-dependencies"}
	args = append(args, packages...)
	_, err := c.Run(args...)
	return err
}

// InstallWithWorkingDir installs dependencies in a specific working directory
//
// 参数：
//   - workingDir: 工作目录路径
//   - noDev: 是否跳过开发依赖
//   - optimize: 是否优化自动加载器
//
// 返回值：
//   - error: 如果安装过程中发生错误
func (c *Composer) InstallWithWorkingDir(workingDir string, noDev bool, optimize bool) error {
	originalDir := c.workingDir
	c.workingDir = workingDir
	defer func() { c.workingDir = originalDir }()

	return c.Install(noDev, optimize)
}

// ==================== Check Platform Enhancements ====================

// CheckPlatformReqsWithFormat checks platform requirements in a specific format
//
// 参数：
//   - format: 输出格式，如"json"或"text"
//
// 返回值：
//   - string: 平台需求检查结果的输出
//   - error: 如果检查过程中发生错误
func (c *Composer) CheckPlatformReqsWithFormat(format string) (string, error) {
	return c.Run("check-platform-reqs", "--format", format)
}

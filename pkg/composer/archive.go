package composer

// Archive 创建项目的归档文件（不包含开发依赖）
//
// 参数：
//   - destination: 存储归档文件的目标目录路径
//
// 返回值：
//   - string: 归档命令的输出结果
//   - error: 如果创建归档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法创建当前项目的ZIP格式归档文件，不包含开发依赖。
//	相当于执行`composer archive --format=zip --dir=destination`命令。
//
// 用法示例：
//
//	output, err := comp.Archive("/path/to/output/dir")
//	if err != nil {
//	    log.Fatalf("创建归档文件失败: %v", err)
//	}
//	fmt.Println("归档创建结果:", output)
func (c *Composer) Archive(destination string) (string, error) {
	return c.Run("archive", "--format=zip", "--dir="+destination)
}

// ArchiveWithFormat 使用指定格式创建项目的归档文件
//
// 参数：
//   - destination: 存储归档文件的目标目录路径
//   - format: 归档文件格式，可以是"zip"或"tar"
//
// 返回值：
//   - string: 归档命令的输出结果
//   - error: 如果创建归档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法创建当前项目的归档文件，可以指定归档格式，不包含开发依赖。
//	相当于执行`composer archive --format=[format] --dir=destination`命令。
//
// 用法示例：
//
//	// 创建TAR格式归档
//	output, err := comp.ArchiveWithFormat("/path/to/output/dir", "tar")
//	if err != nil {
//	    log.Fatalf("创建TAR归档文件失败: %v", err)
//	}
//	fmt.Println("归档创建结果:", output)
func (c *Composer) ArchiveWithFormat(destination string, format string) (string, error) {
	return c.Run("archive", "--format="+format, "--dir="+destination)
}

// ArchiveWithOptions 使用自定义选项创建项目的归档文件
//
// 参数：
//   - destination: 存储归档文件的目标目录路径
//   - options: 归档选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - string: 归档命令的输出结果
//   - error: 如果创建归档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项创建当前项目的归档文件，提供最大的灵活性。
//
// 用法示例：
//
//	// 使用多个自定义选项创建归档
//	options := map[string]string{
//	    "format": "tar",
//	    "file": "my-project.tar",
//	    "ignore-filters": "",
//	}
//	output, err := comp.ArchiveWithOptions("/path/to/output/dir", options)
//	if err != nil {
//	    log.Fatalf("创建归档文件失败: %v", err)
//	}
//	fmt.Println("归档创建结果:", output)
func (c *Composer) ArchiveWithOptions(destination string, options map[string]string) (string, error) {
	args := []string{"archive", "--dir=" + destination}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ArchivePackage 创建指定包的归档文件
//
// 参数：
//   - packageName: 要归档的包名，例如"symfony/console"
//   - version: 要归档的包版本，例如"v5.4.0"，为空则使用最新版本
//   - destination: 存储归档文件的目标目录路径
//
// 返回值：
//   - string: 归档命令的输出结果
//   - error: 如果创建归档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法创建指定包的归档文件，可以指定包版本。
//	相当于执行`composer archive packageName[=version] --dir=destination`命令。
//
// 用法示例：
//
//	// 归档指定版本的包
//	output, err := comp.ArchivePackage("symfony/console", "v5.4.0", "/path/to/output/dir")
//	if err != nil {
//	    log.Fatalf("创建包归档文件失败: %v", err)
//	}
//	fmt.Println("包归档创建结果:", output)
//
//	// 归档最新版本的包
//	output, err = comp.ArchivePackage("symfony/console", "", "/path/to/output/dir")
//	if err != nil {
//	    log.Fatalf("创建包归档文件失败: %v", err)
//	}
func (c *Composer) ArchivePackage(packageName string, version string, destination string) (string, error) {
	args := []string{"archive"}

	if version != "" {
		args = append(args, packageName+"="+version)
	} else {
		args = append(args, packageName)
	}

	args = append(args, "--dir="+destination)

	return c.Run(args...)
}

// ArchivePackageWithOptions 使用自定义选项创建指定包的归档文件
//
// 参数：
//   - packageName: 要归档的包名，例如"symfony/console"
//   - version: 要归档的包版本，例如"v5.4.0"，为空则使用最新版本
//   - destination: 存储归档文件的目标目录路径
//   - options: 归档选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - string: 归档命令的输出结果
//   - error: 如果创建归档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项创建指定包的归档文件，提供最大的灵活性。
//
// 用法示例：
//
//	// 使用自定义选项归档指定版本的包
//	options := map[string]string{
//	    "format": "tar",
//	    "ignore-filters": "",
//	}
//	output, err := comp.ArchivePackageWithOptions(
//	    "symfony/console",
//	    "v5.4.0",
//	    "/path/to/output/dir",
//	    options,
//	)
//	if err != nil {
//	    log.Fatalf("创建包归档文件失败: %v", err)
//	}
//	fmt.Println("包归档创建结果:", output)
func (c *Composer) ArchivePackageWithOptions(packageName string, version string, destination string, options map[string]string) (string, error) {
	args := []string{"archive"}

	if version != "" {
		args = append(args, packageName+"="+version)
	} else {
		args = append(args, packageName)
	}

	args = append(args, "--dir="+destination)

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

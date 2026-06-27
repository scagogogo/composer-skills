package composer

import "strings"

// ValidateStrict 严格验证 composer.json 是否有效
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用严格模式验证composer.json文件是否有效。相当于执行
//	`composer validate --strict`。严格模式会进行更严格的检查，
//	确保composer.json符合最佳实践。
//
// 用法示例：
//
//	output, err := comp.ValidateStrict()
//	if err != nil {
//	    log.Fatalf("严格验证失败: %v", err)
//	}
//	fmt.Println("验证结果:", output)
func (c *Composer) ValidateStrict() (string, error) {
	return c.Run("validate", "--strict")
}

// ValidateWithNoCheck 验证 composer.json 而不检查平台需求
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法验证composer.json文件的格式和结构，但不检查平台需求和其他约束。
//	相当于执行`composer validate --no-check-all`。
//
// 用法示例：
//
//	output, err := comp.ValidateWithNoCheck()
//	if err != nil {
//	    log.Fatalf("验证失败: %v", err)
//	}
//	fmt.Println("仅格式验证结果:", output)
func (c *Composer) ValidateWithNoCheck() (string, error) {
	return c.Run("validate", "--no-check-all")
}

// ValidateWithNoCheckPublish 验证 composer.json 而不检查发布要求
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法验证composer.json文件，但不检查发布到Packagist所需的字段。
//	相当于执行`composer validate --no-check-publish`。
//
// 用法示例：
//
//	output, err := comp.ValidateWithNoCheckPublish()
//	if err != nil {
//	    log.Fatalf("验证失败: %v", err)
//	}
//	fmt.Println("不检查发布要求的验证结果:", output)
func (c *Composer) ValidateWithNoCheckPublish() (string, error) {
	return c.Run("validate", "--no-check-publish")
}

// ValidateWithCheckVersion 验证 composer.json 并检查版本约束
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法验证composer.json文件并检查依赖项的版本约束，包括间接依赖。
//	相当于执行`composer validate --with-dependencies`。
//
// 用法示例：
//
//	output, err := comp.ValidateWithCheckVersion()
//	if err != nil {
//	    log.Fatalf("版本约束验证失败: %v", err)
//	}
//	fmt.Println("版本约束验证结果:", output)
func (c *Composer) ValidateWithCheckVersion() (string, error) {
	return c.Run("validate", "--with-dependencies")
}

// CheckPlatformReqsLock 检查 composer.lock 文件中的平台需求
//
// 返回值：
//   - string: 检查命令的输出结果
//   - error: 如果检查失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查系统是否满足composer.lock文件中指定的平台需求。
//	相当于执行`composer check-platform-reqs --lock`。
//
// 用法示例：
//
//	output, err := comp.CheckPlatformReqsLock()
//	if err != nil {
//	    log.Fatalf("平台需求检查失败: %v", err)
//	}
//	fmt.Println("平台需求检查结果:", output)
func (c *Composer) CheckPlatformReqsLock() (string, error) {
	return c.Run("check-platform-reqs", "--lock")
}

// CheckForOutdatedPackages 检查过时的包，并显示可能的更新
//
// 参数：
//   - direct: 如果为true，则只检查直接依赖
//   - minor: 如果为true，则只显示次要更新
//   - format: 输出格式，可以是"text"、"json"或其他支持的格式
//
// 返回值：
//   - string: 检查命令的输出结果
//   - error: 如果检查失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查项目依赖中过时的包，并显示可用的更新。
//	相当于执行`composer outdated [--direct] [--minor-only] [--format format]`。
//
// 用法示例：
//
//	// 检查所有依赖的更新
//	output, err := comp.CheckForOutdatedPackages(false, false, "")
//	if err != nil {
//	    log.Fatalf("检查过时包失败: %v", err)
//	}
//	fmt.Println("过时包列表:", output)
//
//	// 只检查直接依赖的次要更新，并以JSON格式输出
//	output, err = comp.CheckForOutdatedPackages(true, true, "json")
//	if err != nil {
//	    log.Fatalf("检查过时包失败: %v", err)
//	}
//	fmt.Println("直接依赖的次要更新:", output)
func (c *Composer) CheckForOutdatedPackages(direct bool, minor bool, format string) (string, error) {
	args := []string{"outdated"}

	if direct {
		args = append(args, "--direct")
	}

	if minor {
		args = append(args, "--minor-only")
	}

	if format != "" {
		args = append(args, "--format", format)
	}

	return c.Run(args...)
}

// ValidateSchema 验证 composer.json 和 composer.lock 是否符合架构
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法仅验证composer.json和composer.lock文件的结构是否符合JSON架构，
//	不检查其他约束。相当于执行
//	`composer validate --no-check-all --no-check-publish --no-check-version`。
//
// 用法示例：
//
//	output, err := comp.ValidateSchema()
//	if err != nil {
//	    log.Fatalf("架构验证失败: %v", err)
//	}
//	fmt.Println("架构验证结果:", output)
func (c *Composer) ValidateSchema() (string, error) {
	return c.Run("validate", "--no-check-all", "--no-check-publish", "--no-check-version")
}

// ValidateWithOptions 使用自定义选项验证 composer.json
//
// 参数：
//   - options: 验证选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项验证composer.json文件。可以组合多个验证选项。
//
// 用法示例：
//
//	// 严格验证并检查依赖关系
//	options := map[string]string{
//	    "strict": "",
//	    "with-dependencies": "",
//	}
//	output, err := comp.ValidateWithOptions(options)
//	if err != nil {
//	    log.Fatalf("自定义验证失败: %v", err)
//	}
//	fmt.Println("自定义验证结果:", output)
func (c *Composer) ValidateWithOptions(options map[string]string) (string, error) {
	args := []string{"validate"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

// ValidateQuiet 静默验证 composer.json，只在错误时输出
//
// 返回值：
//   - string: 验证命令的输出结果，成功时为空
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行静默验证，只在发现错误时输出详细信息，成功时不显示输出。
//	相当于执行`composer validate --quiet`命令。
//
// 用法示例：
//
//	output, err := comp.ValidateQuiet()
//	if err != nil {
//	    log.Fatalf("验证失败: %v\n输出: %s", err, output)
//	} else {
//	    fmt.Println("验证成功")
//	}
func (c *Composer) ValidateQuiet() (string, error) {
	return c.Run("validate", "--quiet")
}

// CheckNormalization 检查 composer.json 是否格式化
//
// 返回值：
//   - string: 检查命令的输出结果
//   - error: 如果检查失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查composer.json文件是否已经被规范化格式化。
//	相当于执行`composer validate --no-check-all --check-normalized`命令。
//
// 用法示例：
//
//	output, err := comp.CheckNormalization()
//	if err != nil {
//	    if strings.Contains(output, "not normalized") {
//	        fmt.Println("composer.json需要格式化")
//	    } else {
//	        log.Fatalf("检查格式化失败: %v", err)
//	    }
//	} else {
//	    fmt.Println("composer.json已正确格式化")
//	}
func (c *Composer) CheckNormalization() (string, error) {
	return c.Run("validate", "--no-check-all", "--check-normalized")
}

// NormalizeComposerJson 格式化 composer.json 文件
//
// 返回值：
//   - string: 格式化命令的输出结果
//   - error: 如果格式化失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法格式化composer.json文件，使其符合规范的格式。
//	相当于执行`composer normalize`命令（需要安装normalize插件）。
//
// 用法示例：
//
//	output, err := comp.NormalizeComposerJson()
//	if err != nil {
//	    if strings.Contains(err.Error(), "command not found") {
//	        fmt.Println("请先安装normalize插件: composer global require ergebnis/composer-normalize")
//	    } else {
//	        log.Fatalf("格式化失败: %v", err)
//	    }
//	}
//	fmt.Println("格式化结果:", output)
func (c *Composer) NormalizeComposerJson() (string, error) {
	return c.Run("normalize")
}

// CheckForSecurityVulnerabilities 检查是否存在已知的安全漏洞
//
// 返回值：
//   - string: 检查命令的输出结果
//   - bool: 如果存在安全漏洞，则返回true；否则返回false
//   - error: 如果检查过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查项目的依赖是否包含已知的安全漏洞。
//	相当于执行`composer audit`命令并分析结果。
//
// 用法示例：
//
//	output, hasVulnerabilities, err := comp.CheckForSecurityVulnerabilities()
//	if err != nil {
//	    log.Fatalf("安全检查失败: %v", err)
//	}
//
//	if hasVulnerabilities {
//	    fmt.Println("警告：发现安全漏洞!")
//	    fmt.Println(output)
//	} else {
//	    fmt.Println("没有发现安全漏洞")
//	}
func (c *Composer) CheckForSecurityVulnerabilities() (string, bool, error) {
	output, err := c.Run("audit")

	// 分析输出判断是否存在漏洞
	hasVulnerabilities := false
	if err != nil {
		// composer audit在发现漏洞时会返回非零退出码
		if output != "" && (
		// 检查常见的漏洞存在提示
		strings.Contains(output, "Found") && strings.Contains(output, "vulnerability") ||
			strings.Contains(output, "vulnerabilities") ||
			strings.Contains(output, "Security vulnerability")) {

			hasVulnerabilities = true
			// 保留输出但不返回错误，因为这是预期的行为
			return output, hasVulnerabilities, nil
		}
		// 如果是其他错误则正常返回
		return output, hasVulnerabilities, err
	}

	// 如果没有报错但输出中包含漏洞提示，也认为有漏洞
	if strings.Contains(output, "Found") && (strings.Contains(output, "vulnerability") ||
		strings.Contains(output, "vulnerabilities")) {
		hasVulnerabilities = true
	}

	return output, hasVulnerabilities, nil
}

// ValidateComposerLock 验证 composer.lock 文件
//
// 返回值：
//   - string: 验证命令的输出结果
//   - error: 如果验证失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法验证composer.lock文件是否存在并与composer.json同步。
//	相当于执行`composer validate --check-lock`命令。
//
// 用法示例：
//
//	output, err := comp.ValidateComposerLock()
//	if err != nil {
//	    if strings.Contains(output, "not found") {
//	        fmt.Println("缺少composer.lock文件")
//	    } else if strings.Contains(output, "not up to date") {
//	        fmt.Println("composer.lock需要更新，请运行composer update")
//	    } else {
//	        log.Fatalf("验证composer.lock失败: %v", err)
//	    }
//	} else {
//	    fmt.Println("composer.lock有效:", output)
//	}
func (c *Composer) ValidateComposerLock() (string, error) {
	return c.Run("validate", "--check-lock")
}

// Prohibit 显示被当前平台需求禁止的包
//
// 返回值：
//   - string: 被禁止的包列表
//   - error: 如果查询被禁止包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示哪些包被当前平台需求（platform requirements）所禁止。
//	相当于执行`composer prohibit`命令。
//	此命令可以帮助开发者了解当前平台配置对包安装的限制。
//
// 用法示例：
//
//	output, err := comp.Prohibit()
//	if err != nil {
//	    log.Fatalf("查询被禁止包失败: %v", err)
//	}
//	fmt.Println("被禁止的包:", output)
func (c *Composer) Prohibit() (string, error) {
	return c.Run("prohibit")
}

// ProhibitWithFormat 使用指定格式显示被禁止的包
//
// 参数：
//   - format: 输出格式，可以是"text"或"json"
//
// 返回值：
//   - string: 被禁止的包列表
//   - error: 如果查询被禁止包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法以指定的格式显示被当前平台需求禁止的包。
//	相当于执行`composer prohibit --format=format`命令。
//
// 用法示例：
//
//	// 以JSON格式输出被禁止的包
//	output, err := comp.ProhibitWithFormat("json")
//	if err != nil {
//	    log.Fatalf("查询被禁止包失败: %v", err)
//	}
//	fmt.Println("被禁止的包(JSON):", output)
func (c *Composer) ProhibitWithFormat(format string) (string, error) {
	return c.Run("prohibit", "--format", format)
}

// ProhibitWithOptions 使用自定义选项显示被禁止的包
//
// 参数：
//   - options: 查询被禁止包时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - string: 被禁止的包列表
//   - error: 如果查询被禁止包过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项显示被当前平台需求禁止的包，支持更多自定义选项。
//
// 用法示例：
//
//	// 以JSON格式输出被禁止的包，并显示完整路径
//	options := map[string]string{
//	    "format": "json",
//	    "fixed":  "",
//	}
//	output, err := comp.ProhibitWithOptions(options)
//	if err != nil {
//	    log.Fatalf("查询被禁止包失败: %v", err)
//	}
//	fmt.Println("被禁止的包:", output)
func (c *Composer) ProhibitWithOptions(options map[string]string) (string, error) {
	args := []string{"prohibit"}

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

package cli_security

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01Audit 演示如何进行安全审计
func Example01Audit() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：运行安全审计
	fmt.Println("1. 运行安全审计...")
	result, err := c.AuditWithJSON()
	if err != nil {
		log.Printf("安全审计失败: %v", err)
	} else {
		fmt.Printf("发现 %d 个漏洞\n", result.Found)
	}

	// 示例2：检查是否存在漏洞
	fmt.Println("\n2. 检查是否存在漏洞...")
	hasVuln, err := c.HasVulnerabilities()
	if err != nil {
		log.Printf("检查漏洞失败: %v", err)
	} else if hasVuln {
		fmt.Println("⚠️ 发现安全漏洞！")
	} else {
		fmt.Println("✅ 未发现安全漏洞")
	}

	// 示例3：获取高危漏洞
	fmt.Println("\n3. 获取高危漏洞...")
	highSeverity, err := c.GetHighSeverityVulnerabilities()
	if err != nil {
		log.Printf("获取高危漏洞失败: %v", err)
	} else {
		fmt.Printf("高危漏洞数量: %d\n", len(highSeverity))
		for _, v := range highSeverity {
			fmt.Printf("  - %s: %s\n", v.Title, v.Link)
		}
	}

	// 示例4：获取废弃的包
	fmt.Println("\n4. 获取废弃的包...")
	abandoned, err := c.GetAbandonedPackages()
	if err != nil {
		log.Printf("获取废弃包失败: %v", err)
	} else {
		fmt.Printf("废弃包数量: %d\n", len(abandoned))
		for _, v := range abandoned {
			fmt.Printf("  - %s: %s\n", v.Package, v.Title)
		}
	}

	// 示例5：不包含开发依赖的审计
	fmt.Println("\n5. 不包含开发依赖的审计...")
	output, err := c.AuditWithoutDev()
	if err != nil {
		log.Printf("审计失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例6：审计指定格式的输出
	fmt.Println("\n6. 审计（summary格式）...")
	output, err = c.AuditWithFormat("summary")
	if err != nil {
		log.Printf("审计失败: %v", err)
	} else {
		fmt.Println(output)
	}
}

// Example02PlatformCheck 演示如何检查平台要求
func Example02PlatformCheck() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：检查平台要求
	fmt.Println("1. 检查平台要求...")
	platformInfo, err := c.CheckPlatform()
	if err != nil {
		log.Printf("检查平台要求失败: %v", err)
	} else {
		for _, info := range platformInfo {
			status := "✅"
			if !info.Available {
				status = "❌"
			}
			fmt.Printf("  %s %s %s\n", status, info.Name, info.Version)
		}
	}

	// 示例2：获取PHP版本
	fmt.Println("\n2. 获取PHP版本...")
	phpVersion, err := c.GetPHPVersion()
	if err != nil {
		log.Printf("获取PHP版本失败: %v", err)
	} else {
		fmt.Printf("PHP版本: %s\n", phpVersion)
	}

	// 示例3：检查扩展是否可用
	fmt.Println("\n3. 检查扩展是否可用...")
	hasExt, err := c.HasExtension("mbstring")
	if err != nil {
		log.Printf("检查扩展失败: %v", err)
	} else if hasExt {
		fmt.Println("✅ mbstring 扩展已安装")
	} else {
		fmt.Println("❌ mbstring 扩展未安装")
	}

	// 示例4：获取已安装的扩展列表
	fmt.Println("\n4. 获取已安装的扩展列表...")
	extensions, err := c.GetExtensions()
	if err != nil {
		log.Printf("获取扩展列表失败: %v", err)
	} else {
		fmt.Printf("已安装 %d 个扩展\n", len(extensions))
	}

	// 示例5：检查特定平台是否可用
	fmt.Println("\n5. 检查特定平台是否可用...")
	available, err := c.IsPlatformAvailable("php", "8.1.0")
	if err != nil {
		log.Printf("检查平台可用性失败: %v", err)
	} else if available {
		fmt.Println("✅ PHP 8.1.0 可用")
	} else {
		fmt.Println("❌ PHP 8.1.0 不可用")
	}
}

// Example03Validation 演示如何验证项目配置
func Example03Validation() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：验证composer.json
	fmt.Println("1. 验证composer.json...")
	err = c.Validate()
	if err != nil {
		fmt.Printf("❌ composer.json 验证失败: %v\n", err)
	} else {
		fmt.Println("✅ composer.json 验证通过")
	}

	// 示例2：严格验证
	fmt.Println("\n2. 严格验证composer.json...")
	output, err := c.ValidateStrict()
	if err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println(output)
	}

	// 示例3：验证composer.lock
	fmt.Println("\n3. 验证composer.lock...")
	output, err = c.ValidateComposerLock()
	if err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println(output)
	}

	// 示例4：验证schema
	fmt.Println("\n4. 验证schema...")
	output, err = c.ValidateSchema()
	if err != nil {
		fmt.Printf("验证失败: %v\n", err)
	} else {
		fmt.Println(output)
	}

	// 示例5：规范化composer.json
	fmt.Println("\n5. 规范化composer.json...")
	output, err = c.NormalizeComposerJson()
	if err != nil {
		fmt.Printf("规范化失败: %v\n", err)
	} else {
		fmt.Println(output)
	}

	// 示例6：检查安全漏洞
	fmt.Println("\n6. 检查安全漏洞...")
	output, hasVuln, err := c.CheckForSecurityVulnerabilities()
	if err != nil {
		fmt.Printf("检查失败: %v\n", err)
	} else if hasVuln {
		fmt.Printf("⚠️ 发现安全漏洞:\n%s\n", output)
	} else {
		fmt.Println("✅ 未发现安全漏洞")
	}
}

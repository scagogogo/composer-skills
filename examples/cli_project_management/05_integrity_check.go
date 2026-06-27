package cli_project_management

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example05IntegrityCheck 演示如何使用 Composer 进行完整性检查和诊断
func Example05IntegrityCheck() {
	// 创建 Composer 实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}

	// 创建一个示例项目目录
	tempDir, err := os.MkdirTemp("", "composer-integrity-example")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	// 设置工作目录
	comp.SetWorkingDir(tempDir)

	// 创建一个示例 composer.json 文件
	composerJsonContent := `{
		"name": "example/integrity-check",
		"description": "Integrity check and diagnosis example",
		"type": "project",
		"require": {
			"php": ">=7.4",
			"monolog/monolog": "^2.0"
		}
	}`

	composerJsonPath := filepath.Join(tempDir, "composer.json")
	if err := os.WriteFile(composerJsonPath, []byte(composerJsonContent), 0644); err != nil {
		log.Fatalf("创建 composer.json 失败: %v", err)
	}

	fmt.Printf("创建了示例 composer.json 文件: %s\n\n", composerJsonPath)

	// 注意：以下部分实际执行时可能需要先安装依赖
	// 为演示目的，我们将说明每个函数的用途

	fmt.Println("=== 完整性检查和诊断示例 ===")
	fmt.Println("注意: 某些示例可能需要先安装依赖，此处仅演示 API 的使用方法")

	// 示例 1: 检查项目完整性
	fmt.Println("\n=== 示例 1: 检查项目完整性 ===")
	fmt.Println("功能: 检查依赖项是否满足要求")
	fmt.Println("命令: comp.Check()")
	fmt.Println("说明: 此命令会验证 composer.json 和 composer.lock 中的依赖关系是否一致")

	// 示例 2: 验证 composer.json 文件
	fmt.Println("\n=== 示例 2: 验证 composer.json 文件 ===")
	if err := comp.Validate(); err != nil {
		fmt.Printf("校验失败: %v\n", err)
	} else {
		fmt.Printf("校验成功: 文件格式有效\n")
	}

	// 示例 3: 严格验证 composer.json 文件
	fmt.Println("\n=== 示例 3: 严格验证 composer.json 文件 ===")
	fmt.Println("功能: 严格验证 composer.json 是否有效")
	fmt.Println("命令: comp.ValidateComposerJson(true, false)")
	fmt.Println("说明: 此命令使用严格模式验证 composer.json，但不检查依赖项")

	// 示例 4: 验证 composer.json Schema
	fmt.Println("\n=== 示例 4: 验证 composer.json Schema ===")
	fmt.Println("功能: 验证 composer.json 是否符合架构")
	fmt.Println("命令: comp.ValidateSchema()")
	fmt.Println("说明: 此命令仅验证 composer.json 的架构，不检查版本约束等")

	// 示例 5: 系统诊断
	fmt.Println("\n=== 示例 5: 系统诊断 ===")
	fmt.Println("功能: 诊断系统以识别常见错误")
	fmt.Println("命令: comp.Diagnose()")
	fmt.Println("说明: 此命令会检查系统配置和环境，识别可能的问题")

	// 示例 6: 显示已安装包的修改状态
	fmt.Println("\n=== 示例 6: 显示已安装包的修改状态 ===")
	fmt.Println("功能: 显示已安装包的本地修改")
	fmt.Println("命令: comp.Status()")
	fmt.Println("说明: 此命令会检查已安装包是否有本地修改")

	// 示例 7: 清除缓存
	fmt.Println("\n=== 示例 7: 清除缓存 ===")
	fmt.Println("功能: 清除 Composer 缓存")
	fmt.Println("命令: comp.ClearCache()")
	fmt.Println("说明: 此命令会清除 Composer 的缓存目录")

	// 示例 8: 获取 Composer 主目录
	fmt.Println("\n=== 示例 8: 获取 Composer 主目录 ===")
	homeDir, err := comp.GetComposerHome()
	if err != nil {
		fmt.Printf("获取 Composer 主目录失败: %v\n", err)
	} else {
		fmt.Printf("Composer 主目录: %s\n", homeDir)
	}

	// 示例 9: 读取 composer.json 配置
	fmt.Println("\n=== 示例 9: 读取 composer.json 配置 ===")
	composerJSON, err := comp.ReadComposerJSON()
	if err != nil {
		fmt.Printf("读取 composer.json 失败: %v\n", err)
	} else {
		fmt.Printf("项目名称: %s\n", composerJSON.Name)
		fmt.Printf("项目描述: %s\n", composerJSON.Description)
		fmt.Printf("项目类型: %s\n", composerJSON.Type)
		fmt.Printf("PHP 版本要求: %s\n", composerJSON.Require["php"])
	}

	// 示例 10: 获取 PHP 版本
	fmt.Println("\n=== 示例 10: 获取 PHP 版本 ===")
	fmt.Println("功能: 获取当前 PHP 版本")
	fmt.Println("命令: comp.GetPHPVersion()")
	fmt.Println("说明: 此命令会返回当前环境的 PHP 版本")

	fmt.Println("\n所有示例执行完毕!")
}

package cli_project_management

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example03PlatformCheck 演示如何使用 Composer 检查平台需求
func Example03PlatformCheck() {
	// 创建 Composer 实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}

	// 创建一个示例项目目录
	tempDir, err := os.MkdirTemp("", "composer-platform-example")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	// 设置工作目录
	comp.SetWorkingDir(tempDir)

	// 创建一个包含平台需求的 composer.json 文件
	composerJsonContent := `{
		"name": "example/platform-check",
		"description": "Platform requirements check example",
		"type": "project",
		"require": {
			"php": ">=7.4",
			"ext-json": "*",
			"ext-mbstring": "*",
			"ext-ctype": "*"
		}
	}`

	composerJsonPath := filepath.Join(tempDir, "composer.json")
	if err := os.WriteFile(composerJsonPath, []byte(composerJsonContent), 0644); err != nil {
		log.Fatalf("创建 composer.json 失败: %v", err)
	}

	fmt.Printf("创建了示例 composer.json 文件，包含平台需求: %s\n\n", composerJsonPath)

	// 示例 1: 检查平台需求
	fmt.Println("=== 示例 1: 检查平台需求 ===")
	output, err := comp.CheckPlatformReqs()
	if err != nil {
		log.Printf("检查平台需求失败: %v", err)
	} else {
		fmt.Printf("平台需求检查结果:\n%s\n", output)
	}

	// 示例 2: 验证 composer.json 文件
	fmt.Println("\n=== 示例 2: 验证 composer.json 文件 ===")
	if err := comp.Validate(); err != nil {
		fmt.Printf("校验失败: %v\n", err)
	} else {
		fmt.Printf("校验成功: 文件格式有效\n")
	}

	// 示例 3: 严格验证 composer.json 文件
	fmt.Println("\n=== 示例 3: 严格验证 composer.json 文件 ===")
	if err := comp.ValidateComposerJson(true, false); err != nil {
		fmt.Printf("严格校验失败: %v\n", err)
	} else {
		fmt.Printf("严格校验成功: 文件格式和内容都有效\n")
	}

	// 示例 4: 使用其他验证方法
	fmt.Println("\n=== 示例 4: 使用其他验证方法 ===")
	output, err = comp.ValidateStrict()
	if err != nil {
		fmt.Printf("ValidateStrict 失败: %v\n", err)
	} else {
		fmt.Printf("ValidateStrict 成功: %s\n", output)
	}

	output, err = comp.ValidateSchema()
	if err != nil {
		fmt.Printf("ValidateSchema 失败: %v\n", err)
	} else {
		fmt.Printf("ValidateSchema 成功: %s\n", output)
	}

	// 示例 5: 检查特定平台是否可用
	fmt.Println("\n=== 示例 5: 检查特定平台是否可用 ===")
	available, err := comp.IsPlatformAvailable("php", "7.4")
	if err != nil {
		fmt.Printf("检查 PHP 7.4 失败: %v\n", err)
	} else {
		fmt.Printf("PHP 7.4 是否可用: %v\n", available)
	}

	available, err = comp.IsPlatformAvailable("ext-json", "")
	if err != nil {
		fmt.Printf("检查 ext-json 失败: %v\n", err)
	} else {
		fmt.Printf("ext-json 是否可用: %v\n", available)
	}

	// 检查一个可能不可用的扩展
	available, err = comp.IsPlatformAvailable("ext-imagick", "")
	if err != nil {
		fmt.Printf("检查 ext-imagick 失败: %v\n", err)
	} else {
		fmt.Printf("ext-imagick 是否可用: %v\n", available)
	}

	fmt.Println("\n所有示例执行完毕!")
}

package cli_project_management

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01CreateProject 演示如何使用 Composer 创建和验证项目
func Example01CreateProject() {
	// 创建 Composer 实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}

	// 创建一个示例项目目录
	tempDir, err := os.MkdirTemp("", "composer-project-example")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	// 设置工作目录为临时目录
	comp.SetWorkingDir(tempDir)

	// 示例 1: 使用 init 命令初始化一个新项目
	fmt.Println("=== 示例 1: 初始化一个新项目 ===")
	// 注意: 实际执行时会提示输入信息，这里仅作为示例
	// _, err = comp.Run("init")

	// 示例 2: 创建项目 (使用 create-project 命令)
	fmt.Println("\n=== 示例 2: 使用 create-project 命令创建项目 ===")
	fmt.Println("注意: 此命令会从远程下载项目模板，需要网络连接")
	fmt.Println("示例命令: composer create-project laravel/laravel my-project")
	fmt.Println("为避免实际执行，此处仅展示 API 用法")

	// 示例 3: 手动创建和验证 composer.json 文件
	fmt.Println("\n=== 示例 3: 手动创建和验证 composer.json 文件 ===")
	composerJsonPath := filepath.Join(tempDir, "composer.json")
	sampleContent := `{
		"name": "example/project",
		"description": "A sample project",
		"type": "project",
		"license": "MIT",
		"authors": [
			{
				"name": "Your Name",
				"email": "your.email@example.com"
			}
		],
		"minimum-stability": "stable",
		"require": {
			"php": ">=7.4"
		}
	}`
	if err := os.WriteFile(composerJsonPath, []byte(sampleContent), 0644); err != nil {
		log.Fatalf("创建示例 composer.json 失败: %v", err)
	}

	fmt.Printf("校验文件: %s\n", composerJsonPath)
	if err := comp.Validate(); err != nil {
		fmt.Printf("校验失败: %v\n", err)
	} else {
		fmt.Printf("校验成功: 文件格式有效\n")
	}

	// 示例 4: 严格校验 composer.json 文件
	fmt.Println("\n=== 示例 4: 严格校验 composer.json 文件 ===")
	if err := comp.ValidateComposerJson(true, false); err != nil {
		fmt.Printf("严格校验失败: %v\n", err)
	} else {
		fmt.Printf("严格校验成功: 文件格式和内容都有效\n")
	}

	fmt.Println("\n所有示例执行完毕!")
}

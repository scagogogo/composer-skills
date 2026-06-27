package cli_project_management

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example02RunScript 演示如何使用 Composer 运行脚本
func Example02RunScript() {
	// 创建 Composer 实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}

	// 创建一个示例项目目录
	tempDir, err := os.MkdirTemp("", "composer-script-example")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	// 设置工作目录
	comp.SetWorkingDir(tempDir)

	// 创建一个包含脚本的 composer.json 文件
	composerJsonContent := `{
		"name": "example/run-script",
		"description": "Script execution example",
		"type": "project",
		"require": {
			"php": ">=7.4"
		},
		"scripts": {
			"hello": "echo 'Hello from Composer script!'",
			"list-files": "ls -la",
			"custom-php": "php -r 'echo PHP_VERSION . \"\\n\";'",
			"combined": [
				"@hello",
				"@list-files"
			]
		}
	}`

	composerJsonPath := filepath.Join(tempDir, "composer.json")
	if err := os.WriteFile(composerJsonPath, []byte(composerJsonContent), 0644); err != nil {
		log.Fatalf("创建 composer.json 失败: %v", err)
	}

	fmt.Printf("创建了示例 composer.json 文件，内含多个脚本: %s\n\n", composerJsonPath)

	// 示例 1: 列出所有可用的脚本
	fmt.Println("=== 示例 1: 列出所有可用的脚本 ===")
	scriptList, err := comp.ListScripts()
	if err != nil {
		log.Printf("列出脚本失败: %v", err)
	} else {
		fmt.Printf("可用脚本列表:\n%s\n", scriptList)
	}

	// 示例 2: 运行简单的脚本
	fmt.Println("\n=== 示例 2: 运行简单的 hello 脚本 ===")
	output, err := comp.ExecuteScript("hello")
	if err != nil {
		log.Printf("运行脚本失败: %v", err)
	} else {
		fmt.Printf("脚本输出:\n%s\n", output)
	}

	// 示例 3: 运行带有系统命令的脚本
	fmt.Println("\n=== 示例 3: 运行带有系统命令的脚本 ===")
	output, err = comp.ExecuteScript("list-files")
	if err != nil {
		log.Printf("运行脚本失败: %v", err)
	} else {
		fmt.Printf("脚本输出:\n%s\n", output)
	}

	// 示例 4: 运行带有 PHP 代码的脚本
	fmt.Println("\n=== 示例 4: 运行带有 PHP 代码的脚本 ===")
	output, err = comp.ExecuteScript("custom-php")
	if err != nil {
		log.Printf("运行脚本失败: %v", err)
	} else {
		fmt.Printf("脚本输出 (PHP 版本):\n%s\n", output)
	}

	// 示例 5: 运行组合脚本
	fmt.Println("\n=== 示例 5: 运行组合脚本 ===")
	output, err = comp.ExecuteScript("combined")
	if err != nil {
		log.Printf("运行脚本失败: %v", err)
	} else {
		fmt.Printf("组合脚本输出:\n%s\n", output)
	}

	// 示例 6: 使用 RunScript 方法运行脚本并传递额外参数
	fmt.Println("\n=== 示例 6: 使用 RunScript 方法运行脚本并传递额外参数 ===")
	output, err = comp.RunScript("hello", "--verbose")
	if err != nil {
		log.Printf("运行脚本失败: %v", err)
	} else {
		fmt.Printf("带参数的脚本输出:\n%s\n", output)
	}

	fmt.Println("\n所有示例执行完毕!")
}

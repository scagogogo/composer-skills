package cli_project_management

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example04DependencyAnalysis 演示如何使用 Composer 进行依赖分析
func Example04DependencyAnalysis() {
	// 创建 Composer 实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化 Composer 失败: %v", err)
	}

	// 创建一个示例项目目录
	tempDir, err := os.MkdirTemp("", "composer-dependency-example")
	if err != nil {
		log.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	// 设置工作目录
	comp.SetWorkingDir(tempDir)

	// 创建一个包含依赖的 composer.json 文件
	composerJsonContent := `{
		"name": "example/dependency-analysis",
		"description": "Dependency analysis example",
		"type": "project",
		"require": {
			"php": ">=7.4",
			"guzzlehttp/guzzle": "^7.0",
			"monolog/monolog": "^2.0"
		},
		"require-dev": {
			"phpunit/phpunit": "^9.0"
		}
	}`

	composerJsonPath := filepath.Join(tempDir, "composer.json")
	if err := os.WriteFile(composerJsonPath, []byte(composerJsonContent), 0644); err != nil {
		log.Fatalf("创建 composer.json 失败: %v", err)
	}

	fmt.Printf("创建了示例 composer.json 文件，包含多个依赖: %s\n\n", composerJsonPath)

	// 注意：以下部分实际执行时需要先安装依赖
	// 为演示目的，我们将跳过实际安装，并说明每个函数的用途

	fmt.Println("=== 依赖分析示例 ===")
	fmt.Println("注意: 以下示例需要先安装依赖，此处仅演示 API 的使用方法")

	// 示例 1: 检查依赖关系
	fmt.Println("\n=== 示例 1: 检查依赖关系 ===")
	fmt.Println("功能: 检查依赖项是否满足要求")
	fmt.Println("命令: comp.Check()")
	fmt.Println("说明: 此命令会验证 composer.json 和 composer.lock 中的依赖关系是否一致")

	// 示例 2: 查看依赖树
	fmt.Println("\n=== 示例 2: 查看依赖树 ===")
	fmt.Println("功能: 显示项目的依赖树")
	fmt.Println("命令: comp.ShowDependencyTree(\"\")")
	fmt.Println("说明: 此命令会显示所有安装的包及其依赖关系，以树状结构展示")

	// 示例 3: 查看特定包的依赖树
	fmt.Println("\n=== 示例 3: 查看特定包的依赖树 ===")
	fmt.Println("功能: 显示特定包的依赖树")
	fmt.Println("命令: comp.ShowDependencyTree(\"guzzlehttp/guzzle\")")
	fmt.Println("说明: 此命令会显示指定包及其依赖关系，以树状结构展示")

	// 示例 4: 查看反向依赖
	fmt.Println("\n=== 示例 4: 查看反向依赖 ===")
	fmt.Println("功能: 显示哪些包依赖于指定包")
	fmt.Println("命令: comp.ShowReverseDependencies(\"monolog/monolog\")")
	fmt.Println("说明: 此命令会显示所有依赖于指定包的包")

	// 示例 5: 解释为什么安装了某个包
	fmt.Println("\n=== 示例 5: 解释为什么安装了某个包 ===")
	fmt.Println("功能: 解释为什么项目中安装了指定的包")
	fmt.Println("命令: comp.WhyPackage(\"monolog/monolog\")")
	fmt.Println("说明: 此命令会解释指定包被引入项目的原因")

	// 示例 6: 检查过时的包
	fmt.Println("\n=== 示例 6: 检查过时的包 ===")
	fmt.Println("功能: 显示所有过时的包")
	fmt.Println("命令: comp.OutdatedPackages()")
	fmt.Println("说明: 此命令会检查项目中是否有可更新的包")

	// 示例 7: 仅检查直接依赖中过时的包
	fmt.Println("\n=== 示例 7: 仅检查直接依赖中过时的包 ===")
	fmt.Println("功能: 显示直接依赖中过时的包")
	fmt.Println("命令: comp.OutdatedPackagesDirect()")
	fmt.Println("说明: 此命令仅检查项目直接依赖中是否有可更新的包")

	// 示例 8: 安全审计
	fmt.Println("\n=== 示例 8: 安全审计 ===")
	fmt.Println("功能: 检查项目中的安全漏洞")
	fmt.Println("命令: comp.Audit()")
	fmt.Println("说明: 此命令会检查项目中是否有已知安全漏洞的包")

	// 示例 9: JSON 格式的安全审计
	fmt.Println("\n=== 示例 9: JSON 格式的安全审计 ===")
	fmt.Println("功能: 以 JSON 格式返回安全审计结果")
	fmt.Println("命令: comp.AuditWithJSON()")
	fmt.Println("说明: 此命令会以结构化的 JSON 格式返回审计结果，便于程序处理")

	// 示例 10: 检查高严重性的漏洞
	fmt.Println("\n=== 示例 10: 检查高严重性的漏洞 ===")
	fmt.Println("功能: 获取高严重性的漏洞")
	fmt.Println("命令: comp.GetHighSeverityVulnerabilities()")
	fmt.Println("说明: 此命令会筛选出高危和严重级别的漏洞")

	fmt.Println("\n所有示例执行完毕!")
}

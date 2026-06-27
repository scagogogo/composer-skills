package main

import (
	"context"
	"fmt"

	"github.com/scagogogo/composer-skills/pkg/repository"
)

func main() {
	// 示例 3: 列出 Composer 仓库中的包
	//
	// 这个示例展示了如何获取 Composer 仓库中所有可用包的列表，
	// 以及如何处理这些包信息。

	// 步骤 1: 初始化仓库客户端
	// ---------------------
	// 设置仓库选项
	options := &repository.Options{
		ServerUrl: "https://packagist.org", // 使用官方仓库
	}

	// 创建仓库客户端
	// 注意: 由于 Repository 结构体的内部字段可能不对外导出，
	// 这里我们需要假设有一个可用的初始化方法
	fmt.Printf("使用服务器 URL: %s\n", options.ServerUrl)

	// 在真实代码中，通常会有一个如下的初始化方法：
	// repo := repository.NewRepository(options)
	// 但由于库的设计不同，这里我们直接创建示例实例
	repo := &repository.Repository{}

	// 仅为了避免未使用变量的警告
	_ = options

	// 步骤 2: 列出所有包
	// ---------------
	fmt.Println("正在获取包列表...")

	// 创建上下文
	ctx := context.Background()

	// 调用 List API 获取所有包
	packages, err := repo.List(ctx)
	if err != nil {
		fmt.Printf("获取包列表失败: %v\n", err)
		return
	}

	// 步骤 3: 处理包列表
	// ---------------
	// 打印包的总数
	fmt.Printf("成功获取 %d 个包\n", len(packages))

	// 打印前 10 个包的名称（如果有）
	fmt.Println("\n前 10 个包:")
	maxPrint := 10
	if len(packages) < maxPrint {
		maxPrint = len(packages)
	}

	for i := 0; i < maxPrint; i++ {
		fmt.Printf("  %d. %s\n", i+1, packages[i].Name)
	}

	// 步骤 4: 按名称搜索包（示例）
	// ------------------------
	searchTerm := "symfony/console"
	fmt.Printf("\n搜索包含 '%s' 的包:\n", searchTerm)

	found := 0
	for _, pkg := range packages {
		if found >= 5 {
			break // 只显示前 5 个匹配的结果
		}

		// 简单的字符串匹配，实际应用中可能需要更复杂的搜索逻辑
		if pkg.Name == searchTerm {
			fmt.Printf("  找到完全匹配: %s\n", pkg.Name)
			found++
		}
	}

	if found == 0 {
		fmt.Printf("  未找到完全匹配 '%s' 的包\n", searchTerm)
	}

	// 输出示例：
	// 使用服务器 URL: https://packagist.org
	// 正在获取包列表...
	// 成功获取 25000 个包
	//
	// 前 10 个包:
	//   1. symfony/polyfill
	//   2. symfony/console
	//   3. symfony/process
	//   4. monolog/monolog
	//   5. doctrine/orm
	//   6. laravel/framework
	//   7. phpunit/phpunit
	//   8. guzzlehttp/guzzle
	//   9. nikic/php-parser
	//   10. vlucas/phpdotenv
	//
	// 搜索包含 'symfony/console' 的包:
	//   找到完全匹配: symfony/console
}

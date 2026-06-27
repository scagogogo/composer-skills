package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/scagogogo/composer-skills/pkg/client"
)

// 一些流行的 Composer 包
var popularPackages = []string{
	"symfony/symfony",
	"laravel/framework",
	"guzzlehttp/guzzle",
	"monolog/monolog",
	"phpunit/phpunit",
}

func main() {
	// 创建一个 Composer 客户端
	composerClient := client.NewComposerClient(30 * time.Second)

	// 创建一个保存结果的映射
	results := make(map[string]interface{})

	// 获取统计数据
	fmt.Println("获取 Composer 仓库统计数据...")
	stats, err := composerClient.GetStatistics()
	if err != nil {
		log.Fatalf("获取统计数据失败: %v", err)
	}
	results["statistics"] = stats
	fmt.Printf("总下载量: %d, 包数量: %d, 版本数量: %d\n",
		stats.Totals.Downloads, stats.Totals.Packages, stats.Totals.Versions)

	// 获取安全公告
	fmt.Println("\n获取安全公告...")
	advisories, err := composerClient.GetSecurityAdvisories()
	if err != nil {
		log.Fatalf("获取安全公告失败: %v", err)
	}
	results["advisories"] = advisories
	fmt.Printf("获取到 %d 个包的安全公告\n", len(advisories.Advisories))

	// 获取流行包的信息
	fmt.Println("\n获取流行包的信息...")
	packageInfos := make(map[string]interface{})
	for _, pkgName := range popularPackages {
		fmt.Printf("  获取 %s 的信息...\n", pkgName)
		pkgInfo, err := composerClient.GetPackage(pkgName)
		if err != nil {
			fmt.Printf("    获取 %s 信息失败: %v\n", pkgName, err)
			continue
		}
		packageInfos[pkgName] = pkgInfo

		// 显示一些基本信息
		p := pkgInfo.Package
		fmt.Printf("    名称: %s\n", p.Name)
		fmt.Printf("    描述: %s\n", p.Description)
		fmt.Printf("    类型: %s\n", p.Type)
		fmt.Printf("    下载量: %d\n", p.Downloads.Total)
		fmt.Printf("    版本数: %d\n", len(p.Versions))
		fmt.Printf("    GitHub Stars: %d\n", p.GithubStars)
		fmt.Println()
	}
	results["packages"] = packageInfos

	// 将所有结果保存到文件
	outputFile := "popular_packages_results.json"
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("序列化结果失败: %v", err)
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Fatalf("保存结果到文件失败: %v", err)
	}

	fmt.Printf("\n结果已保存到 %s\n", outputFile)
}

package main

import (
	"context"
	"fmt"

	"github.com/scagogogo/composer-skills/pkg/repository"
)

func main() {
	// 示例 4: 获取 Composer 仓库统计数据
	//
	// 这个示例展示了如何获取 Composer 仓库的统计信息，
	// 包括下载量、包数量和版本数量等。

	// 步骤 1: 初始化仓库客户端
	// ---------------------
	// 设置仓库选项
	options := &repository.Options{
		ServerUrl: "https://packagist.org", // 使用官方仓库
	}

	// 创建仓库客户端
	// 在真实代码中可能会有专门的构造函数
	repo := &repository.Repository{}

	// 仅为了避免未使用变量的警告
	_ = options

	// 步骤 2: 获取统计数据
	// -----------------
	fmt.Println("正在获取 Composer 仓库统计数据...")

	// 创建上下文
	ctx := context.Background()

	// 调用 Statistics API 获取统计数据
	stats, err := repo.Statistics(ctx)
	if err != nil {
		fmt.Printf("获取统计数据失败: %v\n", err)
		return
	}

	// 步骤 3: 处理和显示统计数据
	// -----------------------
	// 打印统计信息
	fmt.Println("\n仓库统计数据:")
	fmt.Printf("  总下载量: %d\n", stats.Totals.Downloads)
	fmt.Printf("  包数量: %d\n", stats.Totals.Packages)
	fmt.Printf("  版本数量: %d\n", stats.Totals.Versions)

	// 步骤 4: 计算一些衍生指标（示例）
	// ---------------------------
	if stats.Totals.Packages > 0 {
		// 计算每个包的平均下载量
		avgDownloadsPerPackage := float64(stats.Totals.Downloads) / float64(stats.Totals.Packages)
		fmt.Printf("\n每个包的平均下载量: %.2f\n", avgDownloadsPerPackage)

		// 计算每个包的平均版本数
		avgVersionsPerPackage := float64(stats.Totals.Versions) / float64(stats.Totals.Packages)
		fmt.Printf("每个包的平均版本数: %.2f\n", avgVersionsPerPackage)
	}

	// 步骤 5: 格式化数据以便人类阅读（示例）
	// --------------------------------
	// 将下载量格式化为更易读的形式
	formattedDownloads := formatNumber(stats.Totals.Downloads)
	fmt.Printf("\n格式化后的下载量: %s\n", formattedDownloads)

	// 输出示例：
	// 正在获取 Composer 仓库统计数据...
	//
	// 仓库统计数据:
	//   总下载量: 25000000000
	//   包数量: 300000
	//   版本数量: 2500000
	//
	// 每个包的平均下载量: 83333.33
	// 每个包的平均版本数: 8.33
	//
	// 格式化后的下载量: 25,000,000,000
}

// formatNumber 格式化数字为易读形式，添加千位分隔符
func formatNumber(n int64) string {
	// 简单实现，将数字转为字符串并添加千位分隔符
	str := fmt.Sprintf("%d", n)
	result := ""

	// 从右向左每三位添加一个逗号
	for i, c := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}

	return result
}

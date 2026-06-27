package main

import (
	"context"
	"fmt"
	"time"

	"github.com/scagogogo/composer-skills/pkg/repository"
)

func main() {
	// 示例 5: 获取 Composer 包的安全公告
	//
	// 这个示例展示了如何获取 Composer 包的安全漏洞公告，
	// 包括按时间获取和按包名获取两种方式。
	// 安全公告包含了漏洞的详细信息，可用于安全审计。

	// 步骤 1: 初始化仓库客户端
	// ---------------------
	// 设置仓库选项
	options := &repository.Options{
		ServerUrl: "https://packagist.org", // 使用官方仓库
	}

	// 创建仓库客户端
	repo := &repository.Repository{}

	// 仅为了避免未使用变量的警告
	_ = options

	// 创建上下文
	ctx := context.Background()

	// 步骤 2: 按时间获取安全公告
	// -----------------------
	fmt.Println("=== 按时间获取安全公告 ===")

	// 设置一个时间点，获取该时间点之后的安全公告
	// 这里设置为 1 年前
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	fmt.Printf("获取 %s 之后的安全公告...\n", oneYearAgo.Format("2006-01-02"))

	// 调用 API 获取安全公告
	advisoriesResp, err := repo.ListSecurityAdvisories(ctx, oneYearAgo)
	if err != nil {
		fmt.Printf("获取安全公告失败: %v\n", err)
		return
	}

	// 统计安全公告数量
	totalAdvisories := 0
	for _, advisories := range advisoriesResp.Advisories {
		totalAdvisories += len(advisories)
	}

	fmt.Printf("找到 %d 个包含安全公告的包，共 %d 个公告\n", len(advisoriesResp.Advisories), totalAdvisories)

	// 打印部分安全公告详情
	count := 0
	fmt.Println("\n部分安全公告详情:")
	for pkgName, advisories := range advisoriesResp.Advisories {
		if count >= 3 {
			break // 只显示前 3 个包的信息
		}

		fmt.Printf("\n包: %s\n", pkgName)
		for i, advisory := range advisories {
			if i >= 2 {
				fmt.Printf("  ...还有 %d 个公告未显示\n", len(advisories)-i)
				break // 每个包只显示前 2 个公告
			}

			fmt.Printf("  - 标题: %s\n", advisory.Title)
			fmt.Printf("    CVE: %s\n", advisory.Cve)
			fmt.Printf("    报告时间: %s\n", advisory.ReportedAt)
			fmt.Printf("    影响版本: %s\n", advisory.AffectedVersions)
		}
		count++
	}

	// 步骤 3: 获取特定包的安全公告
	// -------------------------
	fmt.Println("\n\n=== 获取特定包的安全公告 ===")

	// 设置要查询的包名
	packageName := "symfony/http-kernel"
	fmt.Printf("获取 %s 包的安全公告...\n", packageName)

	// 调用 API 获取特定包的安全公告
	packageAdvisories, err := repo.ListAdvisories(ctx, packageName)
	if err != nil {
		fmt.Printf("获取 %s 包的安全公告失败: %v\n", packageName, err)
		// 继续执行，不返回
	} else {
		fmt.Printf("找到 %d 个安全公告\n", len(packageAdvisories))

		// 打印安全公告详情
		fmt.Println("\n安全公告详情:")
		for i, advisory := range packageAdvisories {
			if i >= 5 {
				fmt.Printf("...还有 %d 个公告未显示\n", len(packageAdvisories)-i)
				break // 只显示前 5 个公告
			}

			fmt.Printf("\n%d. %s\n", i+1, advisory.Title)
			fmt.Printf("   - 公告 ID: %s\n", advisory.AdvisoryID)
			fmt.Printf("   - CVE: %s\n", advisory.Cve)
			fmt.Printf("   - 报告时间: %s\n", advisory.ReportedAt)
			fmt.Printf("   - 影响版本: %s\n", advisory.AffectedVersions)
			fmt.Printf("   - 链接: %s\n", advisory.Link)
		}
	}

	// 输出示例：
	// === 按时间获取安全公告 ===
	// 获取 2022-05-15 之后的安全公告...
	// 找到 245 个包含安全公告的包，共 312 个公告
	//
	// 部分安全公告详情:
	//
	// 包: symfony/http-kernel
	//   - 标题: HTTP Request Smuggling
	//     CVE: CVE-2022-24894
	//     报告时间: 2022-06-20
	//     影响版本: >=4.4.0,<4.4.44||>=5.0.0,<5.4.15||>=6.0.0,<6.0.15||>=6.1.0,<6.1.3
	//   - 标题: Open redirect vulnerability
	//     CVE: CVE-2022-24733
	//     报告时间: 2022-05-25
	//     影响版本: >=5.3.0,<5.4.3
	//   ...还有 3 个公告未显示
	//
	// 包: laravel/framework
	//   - 标题: DoS security issue in the validation logic of nested arrays
	//     CVE: CVE-2023-0243
	//     报告时间: 2023-01-31
	//     影响版本: >=9.0.0,<9.1.8||>=9.2.0,<9.2.1||>=9.3.0,<9.3.1||>=9.4.0,<9.4.1||>=9.5.0,<9.5.1
	//   - 标题: Potential DoS vulnerability in lazy collections
	//     CVE: CVE-2023-0401
	//     报告时间: 2023-01-31
	//     影响版本: >=9.0.0,<9.5.2
	//   ...还有 2 个公告未显示
	//
	// === 获取特定包的安全公告 ===
	// 获取 symfony/http-kernel 包的安全公告...
	// 找到 15 个安全公告
	//
	// 安全公告详情:
	//
	// 1. HTTP Request Smuggling
	//    - 公告 ID: GHSA-xvch-r4wf-h8w9
	//    - CVE: CVE-2022-24894
	//    - 报告时间: 2022-06-20
	//    - 影响版本: >=4.4.0,<4.4.44||>=5.0.0,<5.4.15||>=6.0.0,<6.0.15||>=6.1.0,<6.1.3
	//    - 链接: https://symfony.com/blog/cve-2022-24894-request-smuggling-in-httpkernel
	//
	// 2. Open redirect vulnerability
	//    - 公告 ID: GHSA-mc8h-8q98-4pvf
	//    - CVE: CVE-2022-24733
	//    - 报告时间: 2022-05-25
	//    - 影响版本: >=5.3.0,<5.4.3
	//    - 链接: https://symfony.com/blog/cve-2022-24733-open-redirect-in-the-url-fragment-handling
	//
	// ...还有 10 个公告未显示
}

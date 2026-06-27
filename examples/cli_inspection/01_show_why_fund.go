package cli_inspection

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01ShowPackages 演示如何查看包信息
func Example01ShowPackages() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：显示所有已安装的包
	fmt.Println("1. 显示所有已安装的包...")
	output, err := c.ShowAllPackages()
	if err != nil {
		log.Printf("获取包列表失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：显示特定包的详细信息
	fmt.Println("\n2. 显示特定包的详细信息...")
	output, err = c.ShowPackage("symfony/console")
	if err != nil {
		log.Printf("获取包信息失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例3：显示依赖树
	fmt.Println("\n3. 显示依赖树...")
	output, err = c.ShowDependencyTree("symfony/console")
	if err != nil {
		log.Printf("获取依赖树失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例4：显示反向依赖
	fmt.Println("\n4. 显示反向依赖（哪些包依赖此包）...")
	output, err = c.ShowReverseDependencies("symfony/polyfill-mbstring")
	if err != nil {
		log.Printf("获取反向依赖失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例5：查看过时的包
	fmt.Println("\n5. 查看过时的包...")
	output, err = c.OutdatedPackages()
	if err != nil {
		log.Printf("获取过时包列表失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例6：只查看直接依赖中过时的包
	fmt.Println("\n6. 只查看直接依赖中过时的包...")
	output, err = c.OutdatedPackagesDirect()
	if err != nil {
		log.Printf("获取过时直接依赖失败: %v", err)
	} else {
		fmt.Println(output)
	}
}

// Example02WhyAnalysis 演示why/why-not分析
func Example02WhyAnalysis() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：解释为什么安装了某个包
	fmt.Println("1. 解释为什么安装了某个包...")
	output, err := c.WhyPackage("symfony/polyfill-mbstring")
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：解释为什么不能安装某个版本
	fmt.Println("\n2. 解释为什么不能安装某个版本...")
	output, err = c.WhyNotPackage("symfony/console", "v6.0.0")
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Println(output)
	}
}

// Example03FundAndLicenses 演示资金和许可证信息
func Example03FundAndLicenses() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：查看资金信息
	fmt.Println("1. 查看资金信息...")
	output, err := c.Fund()
	if err != nil {
		log.Printf("获取资金信息失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：获取JSON格式的资金信息
	fmt.Println("\n2. 获取JSON格式的资金信息...")
	fundingInfo, err := c.FundWithJSON()
	if err != nil {
		log.Printf("获取资金信息失败: %v", err)
	} else {
		for _, info := range fundingInfo {
			if info.Funding {
				fmt.Printf("  包: %s, URL: %v\n", info.Name, info.URLs)
			}
		}
	}

	// 示例3：检查是否有资金支持
	fmt.Println("\n3. 检查是否有资金支持...")
	hasFunding, err := c.HasFunding()
	if err != nil {
		log.Printf("检查失败: %v", err)
	} else if hasFunding {
		fmt.Println("项目中有可以捐赠的包")
	} else {
		fmt.Println("项目中没有可以捐赠的包")
	}

	// 示例4：查看许可证信息
	fmt.Println("\n4. 查看许可证信息...")
	output, err = c.Licenses()
	if err != nil {
		log.Printf("获取许可证信息失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例5：获取JSON格式的许可证信息
	fmt.Println("\n5. 获取JSON格式的许可证信息...")
	output, err = c.LicensesWithFormat("json")
	if err != nil {
		log.Printf("获取许可证信息失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例6：检查许可证兼容性
	fmt.Println("\n6. 检查许可证兼容性...")
	output, err = c.CheckLicenses()
	if err != nil {
		log.Printf("检查许可证失败: %v", err)
	} else {
		fmt.Println(output)
	}
}

// Example04Search 演示搜索功能
func Example04Search() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：搜索包
	fmt.Println("1. 搜索包...")
	output, err := c.Search("logger")
	if err != nil {
		log.Printf("搜索失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：查看建议安装的包
	fmt.Println("\n2. 查看建议安装的包...")
	err = c.Suggests()
	if err != nil {
		log.Printf("获取建议失败: %v", err)
	}
}

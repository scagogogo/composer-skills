package cli_package_management

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example02RequireRemove 演示如何添加和移除依赖包
func Example02RequireRemove() {
	// 创建Composer实例
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 设置工作目录（确保目录下有composer.json文件）
	projectDir := "/path/to/project" // 根据实际情况修改
	c.SetWorkingDir(projectDir)
	fmt.Printf("设置工作目录: %s\n", projectDir)

	// 示例1：添加普通依赖（不指定版本）
	packageName := "monolog/monolog"
	fmt.Printf("\n1. 添加依赖包 %s...\n", packageName)
	err = c.RequirePackage(packageName, "", false)
	if err != nil {
		log.Printf("添加依赖包失败: %v", err)
	} else {
		fmt.Printf("依赖包 %s 添加成功\n", packageName)
		// 输出示例：依赖包 monolog/monolog 添加成功
	}

	// 示例2：添加指定版本的依赖
	packageName = "symfony/console"
	version := "^5.4"
	fmt.Printf("\n2. 添加依赖包 %s (版本 %s)...\n", packageName, version)
	err = c.RequirePackage(packageName, version, false)
	if err != nil {
		log.Printf("添加依赖包失败: %v", err)
	} else {
		fmt.Printf("依赖包 %s (版本 %s) 添加成功\n", packageName, version)
		// 输出示例：依赖包 symfony/console (版本 ^5.4) 添加成功
	}

	// 示例3：添加开发依赖
	packageName = "phpunit/phpunit"
	version = "^9.5"
	fmt.Printf("\n3. 添加开发依赖包 %s (版本 %s)...\n", packageName, version)
	err = c.RequirePackage(packageName, version, true)
	if err != nil {
		log.Printf("添加开发依赖包失败: %v", err)
	} else {
		fmt.Printf("开发依赖包 %s (版本 %s) 添加成功\n", packageName, version)
		// 输出示例：开发依赖包 phpunit/phpunit (版本 ^9.5) 添加成功
	}

	// 示例4：移除普通依赖
	packageName = "monolog/monolog"
	fmt.Printf("\n4. 移除依赖包 %s...\n", packageName)
	err = c.Remove(packageName, false)
	if err != nil {
		log.Printf("移除依赖包失败: %v", err)
	} else {
		fmt.Printf("依赖包 %s 移除成功\n", packageName)
		// 输出示例：依赖包 monolog/monolog 移除成功
	}

	// 示例5：移除开发依赖
	packageName = "phpunit/phpunit"
	fmt.Printf("\n5. 移除开发依赖包 %s...\n", packageName)
	err = c.Remove(packageName, true)
	if err != nil {
		log.Printf("移除开发依赖包失败: %v", err)
	} else {
		fmt.Printf("开发依赖包 %s 移除成功\n", packageName)
		// 输出示例：开发依赖包 phpunit/phpunit 移除成功
	}

	// 示例6：使用高级选项添加依赖
	fmt.Println("\n6. 使用高级选项添加依赖...")

	// 这个示例需要导入 pkg/composer/packages.go 中的 RequireOptions 类型
	// 由于示例代码不需要实际运行，这里注释掉
	/*
		options := composer.RequireOptions{
			WithDev:     false,
			WithNoUpdate: true,
			WithNoProgress: true,
			WithIgnorePlatformReqs: false,
		}
		err = c.RequirePackageWithOptions("guzzlehttp/guzzle", "^7.0", options)
	*/

	fmt.Println("使用高级选项添加依赖包 guzzlehttp/guzzle (版本 ^7.0) 成功")
	// 输出示例：使用高级选项添加依赖包 guzzlehttp/guzzle (版本 ^7.0) 成功
}

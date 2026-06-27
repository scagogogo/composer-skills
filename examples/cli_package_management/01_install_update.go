package cli_package_management

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01InstallUpdate 演示如何安装和更新依赖包
func Example01InstallUpdate() {
	// 创建Composer实例
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 设置工作目录（确保目录下有composer.json文件）
	projectDir := "/path/to/project" // 根据实际情况修改
	c.SetWorkingDir(projectDir)
	fmt.Printf("设置工作目录: %s\n", projectDir)

	// 示例1：安装依赖
	fmt.Println("\n1. 安装依赖...")
	err = c.Install(false, false)
	if err != nil {
		log.Printf("安装依赖失败: %v", err)
	} else {
		fmt.Println("依赖安装成功")
		// 输出示例：依赖安装成功
	}

	// 示例2：安装依赖（不包含开发依赖）
	fmt.Println("\n2. 安装依赖（不包含开发依赖）...")
	err = c.Install(true, false)
	if err != nil {
		log.Printf("安装依赖（不包含开发依赖）失败: %v", err)
	} else {
		fmt.Println("依赖安装成功（不包含开发依赖）")
		// 输出示例：依赖安装成功（不包含开发依赖）
	}

	// 示例3：安装依赖并优化自动加载
	fmt.Println("\n3. 安装依赖并优化自动加载...")
	err = c.Install(false, true)
	if err != nil {
		log.Printf("安装依赖并优化自动加载失败: %v", err)
	} else {
		fmt.Println("依赖安装成功（优化自动加载）")
		// 输出示例：依赖安装成功（优化自动加载）
	}

	// 示例4：更新所有依赖
	fmt.Println("\n4. 更新所有依赖...")
	err = c.Update([]string{}, false)
	if err != nil {
		log.Printf("更新所有依赖失败: %v", err)
	} else {
		fmt.Println("所有依赖更新成功")
		// 输出示例：所有依赖更新成功
	}

	// 示例5：更新特定依赖
	fmt.Println("\n5. 更新特定依赖...")
	packagesToUpdate := []string{"monolog/monolog", "symfony/console"}
	err = c.Update(packagesToUpdate, false)
	if err != nil {
		log.Printf("更新特定依赖失败: %v", err)
	} else {
		fmt.Printf("依赖 %v 更新成功\n", packagesToUpdate)
		// 输出示例：依赖 [monolog/monolog symfony/console] 更新成功
	}

	// 示例6：更新依赖（不包含开发依赖）
	fmt.Println("\n6. 更新特定依赖（不包含开发依赖）...")
	err = c.Update(packagesToUpdate, true)
	if err != nil {
		log.Printf("更新特定依赖（不包含开发依赖）失败: %v", err)
	} else {
		fmt.Printf("依赖 %v 更新成功（不包含开发依赖）\n", packagesToUpdate)
		// 输出示例：依赖 [monolog/monolog symfony/console] 更新成功（不包含开发依赖）
	}

	// 示例7：只更新自动加载配置
	fmt.Println("\n7. 更新自动加载配置...")
	err = c.DumpAutoload(false)
	if err != nil {
		log.Printf("更新自动加载配置失败: %v", err)
	} else {
		fmt.Println("自动加载配置更新成功")
		// 输出示例：自动加载配置更新成功
	}

	// 示例8：更新并优化自动加载配置
	fmt.Println("\n8. 更新并优化自动加载配置...")
	err = c.DumpAutoload(true)
	if err != nil {
		log.Printf("更新并优化自动加载配置失败: %v", err)
	} else {
		fmt.Println("自动加载配置更新并优化成功")
		// 输出示例：自动加载配置更新并优化成功
	}
}

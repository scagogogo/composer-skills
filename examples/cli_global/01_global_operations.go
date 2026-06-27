package cli_global

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01GlobalOperations 演示如何使用全局Composer操作
func Example01GlobalOperations() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 示例1：全局安装包
	fmt.Println("1. 全局安装包...")
	err = c.GlobalRequire("phpstan/phpstan", "^1.0")
	if err != nil {
		log.Printf("全局安装包失败: %v", err)
	} else {
		fmt.Println("全局安装成功")
	}

	// 示例2：列出全局安装的包
	fmt.Println("\n2. 列出全局安装的包...")
	output, err := c.GlobalList()
	if err != nil {
		log.Printf("获取全局包列表失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例3：全局更新包
	fmt.Println("\n3. 全局更新包...")
	err = c.GlobalUpdate([]string{"phpstan/phpstan"})
	if err != nil {
		log.Printf("全局更新失败: %v", err)
	} else {
		fmt.Println("全局更新成功")
	}

	// 示例4：全局移除包
	fmt.Println("\n4. 全局移除包...")
	err = c.GlobalRemove("phpstan/phpstan")
	if err != nil {
		log.Printf("全局移除失败: %v", err)
	} else {
		fmt.Println("全局移除成功")
	}

	// 示例5：获取全局目录路径
	fmt.Println("\n5. 获取全局目录路径...")
	output, err = c.GlobalHome()
	if err != nil {
		log.Printf("获取全局目录失败: %v", err)
	} else {
		fmt.Printf("全局目录: %s\n", output)
	}

	// 示例6：执行全局安装的二进制文件
	fmt.Println("\n6. 执行全局安装的二进制文件...")
	output, err = c.GlobalExecute("phpstan", "analyse", "--no-progress")
	if err != nil {
		log.Printf("执行全局命令失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例7：查看全局包状态
	fmt.Println("\n7. 查看全局包状态...")
	output, err = c.GlobalStatus()
	if err != nil {
		log.Printf("获取全局状态失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例8：全局安装依赖
	fmt.Println("\n8. 全局安装依赖...")
	err = c.GlobalInstall()
	if err != nil {
		log.Printf("全局安装依赖失败: %v", err)
	} else {
		fmt.Println("全局依赖安装成功")
	}

	// 示例9：全局生成自动加载文件
	fmt.Println("\n9. 全局生成自动加载文件...")
	err = c.GlobalDumpAutoload(true)
	if err != nil {
		log.Printf("全局生成自动加载失败: %v", err)
	} else {
		fmt.Println("全局自动加载文件已生成")
	}
}

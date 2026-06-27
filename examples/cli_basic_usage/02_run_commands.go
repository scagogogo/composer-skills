package cli_basic_usage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example02RunCommands 演示如何运行基本的Composer命令
func Example02RunCommands() {
	// 创建Composer实例
	options := composer.DefaultOptions()
	c, err := composer.New(options)
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 示例1：使用Run方法执行简单命令
	output, err := c.Run("--version")
	if err != nil {
		log.Fatalf("执行命令失败: %v", err)
	}
	fmt.Printf("Run方法输出: %s\n", output)
	// 输出示例：Run方法输出: Composer version 2.5.7 2023-12-01 11:43:14

	// 示例2：使用带超时的上下文执行命令
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	output, err = c.RunWithContext(ctx, "diagnose")
	if err != nil {
		log.Printf("执行带上下文的命令失败: %v", err)
	} else {
		fmt.Println("RunWithContext方法执行成功，输出省略...")
		// 输出示例：RunWithContext方法执行成功，输出省略...
	}

	// 示例3：设置工作目录后执行命令
	c.SetWorkingDir("/path/to/your/project")
	fmt.Printf("已设置工作目录: %s\n", "/path/to/your/project")

	// 示例4：设置环境变量后执行命令
	c.SetEnv([]string{"COMPOSER_MEMORY_LIMIT=2G", "COMPOSER_NO_INTERACTION=1"})
	fmt.Println("已设置环境变量: COMPOSER_MEMORY_LIMIT=2G, COMPOSER_NO_INTERACTION=1")

	// 示例5: 执行更新自身命令
	fmt.Println("执行 self-update 命令...")
	err = c.SelfUpdate()
	if err != nil {
		log.Printf("更新Composer失败: %v", err)
	} else {
		fmt.Println("Composer自更新成功")
		// 输出示例：Composer自更新成功
	}
}

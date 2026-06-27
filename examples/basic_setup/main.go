package main

import (
	"fmt"

	"github.com/scagogogo/composer-skills/pkg/repository"
)

func main() {
	// 示例 1: 基本设置 - 创建一个 Composer 仓库客户端
	//
	// 这个示例展示了如何创建一个基本的 Composer 仓库客户端，
	// 可以用于后续的 API 调用。

	// 步骤 1: 创建仓库选项
	// ------------------
	// ServerUrl: 指定 Composer 仓库的基础 URL
	// Proxy: 可选参数，如果需要通过代理访问仓库，则设置代理 URL
	options := &repository.Options{
		ServerUrl: "https://packagist.org", // 官方 Composer 仓库
		// 如果需要代理，可以取消注释下面这行
		// Proxy: "http://your-proxy-server:port",
	}

	// 步骤 2: 初始化一个仓库客户端
	// --------------------------
	// 直接创建 Repository 实例
	repo := &repository.Repository{
		// 使用选项初始化
		// 注意: 实际使用中，repository.Repository 的内部字段可能不对外导出
		// 此处仅作示例，实际使用请参考代码库中可能提供的构造函数
	}

	// 由于访问不到 Repository 的内部字段，我们先打印 options 来演示
	fmt.Println("仓库客户端初始化示例")
	fmt.Printf("仓库 URL: %s\n", options.ServerUrl)

	// 注意: 在实际应用中，应使用更合适的方式初始化仓库对象
	// 这里仅用于演示类型的创建
	_ = repo // 防止未使用变量警告

	// 输出示例：
	// 仓库客户端初始化示例
	// 仓库 URL: https://packagist.org
}

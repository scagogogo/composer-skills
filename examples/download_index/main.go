package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scagogogo/composer-skills/pkg/repository"
)

func main() {
	// 示例 2: 下载 Composer 包索引
	//
	// 这个示例展示了如何下载完整的 Composer 包索引。
	// 索引文件包含仓库中所有可用包的列表。

	// 步骤 1: 创建上下文
	// ---------------
	// 创建一个上下文用于传递给 API 调用，可用于控制请求超时等
	ctx := context.Background()

	// 步骤 2: 直接下载索引
	// -----------------
	fmt.Println("正在下载包索引...")

	// 下载整个索引文件
	// 这会返回一个包含所有包名的 JSON 数据
	indexBytes, err := repository.DownloadIndex(ctx)
	if err != nil {
		fmt.Printf("下载索引失败: %v\n", err)
		return
	}

	// 打印索引的大小
	fmt.Printf("索引下载成功: %d 字节\n", len(indexBytes))

	// 索引数据是一个 JSON 格式的字符串，例如:
	// {"packageNames":["vendor1/package1","vendor2/package2",...]}
	// 可以通过 JSON 解析获取所有包名

	// 步骤 3: 将索引保存到文件
	// ---------------------
	// 创建一个临时目录来保存索引文件
	tempDir, err := os.MkdirTemp("", "composer-index")
	if err != nil {
		fmt.Printf("创建临时目录失败: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // 示例结束时清理

	// 构建文件路径
	indexPath := filepath.Join(tempDir, "composer-index.json")

	// 使用便捷方法直接下载并保存到文件
	fmt.Printf("将索引保存到文件: %s\n", indexPath)

	err = repository.DownloadIndexToFile(ctx, indexPath)
	if err != nil {
		fmt.Printf("保存索引文件失败: %v\n", err)
		return
	}

	// 获取文件信息以验证保存成功
	fileInfo, err := os.Stat(indexPath)
	if err != nil {
		fmt.Printf("获取文件信息失败: %v\n", err)
		return
	}

	fmt.Printf("索引文件保存成功，文件大小: %d 字节\n", fileInfo.Size())

	// 输出示例：
	// 正在下载包索引...
	// 索引下载成功: 1234567 字节
	// 将索引保存到文件: /tmp/composer-index-123456/composer-index.json
	// 索引文件保存成功，文件大小: 1234567 字节
}

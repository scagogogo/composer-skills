package cli_package_management

import (
	"fmt"
	"log"
	"strings"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example04SearchPackage 演示如何使用Go Composer SDK搜索包
func Example04SearchPackage() {
	fmt.Println("\n=== 搜索包示例 ===")

	// 创建一个Composer实例
	comp, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("初始化Composer失败: %v", err)
	}

	// 搜索关键字
	keyword := "logger"
	fmt.Printf("搜索包: '%s'\n", keyword)

	// 执行搜索
	output, err := comp.Search(keyword)
	if err != nil {
		log.Fatalf("搜索包失败: %v", err)
	}

	// 打印搜索结果
	fmt.Println("搜索结果:")
	fmt.Println(output)

	// 示例：解析和处理搜索结果
	fmt.Println("\n提取搜索结果中的包名:")

	// 这只是一个简单的示例，实际解析可能需要更复杂的逻辑
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "/") && !strings.HasPrefix(line, " ") {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				fmt.Printf("- %s\n", parts[0])
			}
		}
	}

	// 搜索另一个关键字（更具体的搜索）
	specificKeyword := "monolog"
	fmt.Printf("\n搜索特定包: '%s'\n", specificKeyword)

	specificOutput, err := comp.Search(specificKeyword)
	if err != nil {
		log.Fatalf("搜索特定包失败: %v", err)
	}

	// 打印搜索结果
	fmt.Println("搜索结果:")
	fmt.Println(specificOutput)
}

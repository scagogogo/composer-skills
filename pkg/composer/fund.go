package composer

import (
	"encoding/json"
	"strings"
)

// FundingInfo 表示一个包的资金信息
type FundingInfo struct {
	Name    string   `json:"name"`
	URLs    []string `json:"urls"`
	Funding bool     `json:"funding"`
}

// Fund 显示项目的资金信息
//
// 返回值：
//   - string: 资金信息命令的输出结果
//   - error: 如果执行命令过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示项目中可以捐赠的软件包的资金信息。
//	相当于执行`composer fund`命令。
//
// 用法示例：
//
//	output, err := comp.Fund()
//	if err != nil {
//	    log.Fatalf("获取资金信息失败: %v", err)
//	}
//	fmt.Println("项目资金信息:", output)
func (c *Composer) Fund() (string, error) {
	return c.Run("fund")
}

// FundWithJSON 获取 JSON 格式的资金信息
//
// 返回值：
//   - []FundingInfo: 解析后的资金信息结构体切片
//   - error: 如果执行命令或解析过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取项目中可以捐赠的软件包的资金信息，并以结构化的JSON格式返回。
//	相当于执行`composer fund --format=json`命令，并将输出解析为结构体。
//
// 用法示例：
//
//	fundingInfo, err := comp.FundWithJSON()
//	if err != nil {
//	    log.Fatalf("获取资金信息失败: %v", err)
//	}
//
//	for _, info := range fundingInfo {
//	    if info.Funding {
//	        fmt.Printf("包: %s\n", info.Name)
//	        fmt.Printf("资金链接: %v\n\n", info.URLs)
//	    }
//	}
func (c *Composer) FundWithJSON() ([]FundingInfo, error) {
	output, err := c.Run("fund", "--format=json")
	if err != nil {
		return nil, err
	}

	// 解析 JSON 输出
	var fundingInfo []FundingInfo
	if err := json.Unmarshal([]byte(output), &fundingInfo); err != nil {
		return nil, err
	}

	return fundingInfo, nil
}

// FundWithPackage 显示特定包的资金信息
//
// 参数：
//   - packageName: 要查询资金信息的包名
//
// 返回值：
//   - string: 资金信息命令的输出结果
//   - error: 如果执行命令过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示特定包的资金信息。
//	相当于执行`composer fund packageName`命令。
//
// 用法示例：
//
//	// 显示symfony/console包的资金信息
//	output, err := comp.FundWithPackage("symfony/console")
//	if err != nil {
//	    log.Fatalf("获取包资金信息失败: %v", err)
//	}
//	fmt.Println("包资金信息:", output)
func (c *Composer) FundWithPackage(packageName string) (string, error) {
	return c.Run("fund", packageName)
}

// GetFundingURLs 获取所有支持资金的包的 URL
//
// 返回值：
//   - map[string][]string: 包名到资金URL的映射
//   - error: 如果获取过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取项目中所有支持资金的包的URL，返回一个包名到URL数组的映射。
//	这对于在应用程序中集成资金支持功能很有用。
//
// 用法示例：
//
//	urlsMap, err := comp.GetFundingURLs()
//	if err != nil {
//	    log.Fatalf("获取资金URL失败: %v", err)
//	}
//
//	if len(urlsMap) > 0 {
//	    fmt.Println("可支持资金的包:")
//	    for pkg, urls := range urlsMap {
//	        fmt.Printf("%s:\n", pkg)
//	        for _, url := range urls {
//	            fmt.Printf("  - %s\n", url)
//	        }
//	    }
//	} else {
//	    fmt.Println("没有找到可支持资金的包")
//	}
func (c *Composer) GetFundingURLs() (map[string][]string, error) {
	fundingInfo, err := c.FundWithJSON()
	if err != nil {
		return nil, err
	}

	// 创建包名到 URL 的映射
	urls := make(map[string][]string)

	for _, info := range fundingInfo {
		if info.Funding && len(info.URLs) > 0 {
			urls[info.Name] = info.URLs
		}
	}

	return urls, nil
}

// HasFunding 检查项目是否有任何资金支持
//
// 返回值：
//   - bool: 如果项目中有可以捐赠的包，则返回true；否则返回false
//   - error: 如果检查过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查项目中是否有任何支持资金的包。
//	这对于决定是否在应用程序中显示资金支持信息很有用。
//
// 用法示例：
//
//	hasFunding, err := comp.HasFunding()
//	if err != nil {
//	    log.Fatalf("检查资金支持失败: %v", err)
//	}
//
//	if hasFunding {
//	    fmt.Println("项目中有可以捐赠的包，可以运行`composer fund`查看详情")
//	} else {
//	    fmt.Println("项目中没有可以捐赠的包")
//	}
func (c *Composer) HasFunding() (bool, error) {
	output, err := c.Run("fund", "--format=text")
	if err != nil {
		return false, err
	}

	// 如果输出包含 "No funding"，则没有资金支持
	return !strings.Contains(output, "No funding"), nil
}

// FundWithOptions 使用自定义选项显示资金信息
//
// 参数：
//   - options: 资金命令选项的映射，键为选项名，值为选项值
//
// 返回值：
//   - string: 资金信息命令的输出结果
//   - error: 如果执行命令过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用自定义选项显示项目的资金信息，提供最大的灵活性。
//
// 用法示例：
//
//	// 使用多个自定义选项
//	options := map[string]string{
//	    "format": "json",
//	    "no-dev": "",
//	}
//	output, err := comp.FundWithOptions(options)
//	if err != nil {
//	    log.Fatalf("获取资金信息失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) FundWithOptions(options map[string]string) (string, error) {
	args := []string{"fund"}

	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	return c.Run(args...)
}

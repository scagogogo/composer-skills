package composer

import (
	"encoding/json"
	"strings"
)

// PlatformInfo 表示平台信息
type PlatformInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Available bool   `json:"available"`
	Required  string `json:"required,omitempty"`
}

// PlatformRequirements 表示平台需求
type PlatformRequirements struct {
	Platform map[string]PlatformInfo `json:"platform"`
	Lock     map[string]PlatformInfo `json:"lock,omitempty"`
}

// CheckPlatform 检查平台需求
//
// 返回值：
//   - []PlatformInfo: 平台需求信息的切片
//   - error: 如果检查平台需求过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查当前系统是否满足composer.json中定义的平台需求。
//	相当于执行`composer check-platform --format=json`并解析结果。
//
// 用法示例：
//
//	// 检查平台需求
//	platforms, err := comp.CheckPlatform()
//	if err != nil {
//	    log.Fatalf("检查平台需求失败: %v", err)
//	}
//
//	// 显示平台需求信息
//	for _, platform := range platforms {
//	    status := "不满足"
//	    if platform.Available {
//	        status = "满足"
//	    }
//	    fmt.Printf("%s %s: %s\n", platform.Name, platform.Version, status)
//	}
func (c *Composer) CheckPlatform() ([]PlatformInfo, error) {
	output, err := c.Run("check-platform", "--format=json")
	if err != nil {
		return nil, err
	}

	var platformReqs PlatformRequirements
	if err := json.Unmarshal([]byte(output), &platformReqs); err != nil {
		return nil, err
	}

	var result []PlatformInfo
	for _, info := range platformReqs.Platform {
		result = append(result, info)
	}

	return result, nil
}

// CheckPlatformWithLock 检查 composer.lock 文件中的平台需求
//
// 返回值：
//   - []PlatformInfo: 平台需求信息的切片
//   - error: 如果检查平台需求过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查当前系统是否满足composer.lock中定义的平台需求。
//	相当于执行`composer check-platform --lock --format=json`并解析结果。
//
// 用法示例：
//
//	// 检查lock文件中的平台需求
//	platforms, err := comp.CheckPlatformWithLock()
//	if err != nil {
//	    log.Fatalf("检查lock文件平台需求失败: %v", err)
//	}
//
//	// 显示不满足的平台需求
//	for _, platform := range platforms {
//	    if !platform.Available {
//	        fmt.Printf("警告: %s %s 需求不满足\n", platform.Name, platform.Required)
//	    }
//	}
func (c *Composer) CheckPlatformWithLock() ([]PlatformInfo, error) {
	output, err := c.Run("check-platform", "--lock", "--format=json")
	if err != nil {
		return nil, err
	}

	var platformReqs PlatformRequirements
	if err := json.Unmarshal([]byte(output), &platformReqs); err != nil {
		return nil, err
	}

	var result []PlatformInfo
	for _, info := range platformReqs.Lock {
		result = append(result, info)
	}

	return result, nil
}

// IsPlatformAvailable 检查指定的平台需求是否满足
//
// 参数：
//   - platform: 平台名称，例如"php"或"ext-mbstring"
//   - version: 版本约束，例如">=7.4"，可以为空
//
// 返回值：
//   - bool: 如果平台需求满足则返回true，否则返回false
//   - error: 如果检查平台需求过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查指定的平台需求是否满足。如果在composer.json中找不到对应的需求，
//	则直接检查该平台要求。
//
// 用法示例：
//
//	// 检查PHP版本
//	available, err := comp.IsPlatformAvailable("php", ">=7.4")
//	if err != nil {
//	    log.Fatalf("检查PHP版本失败: %v", err)
//	}
//	if available {
//	    fmt.Println("PHP版本满足需求")
//	} else {
//	    fmt.Println("PHP版本不满足需求")
//	}
//
//	// 检查扩展
//	available, err = comp.IsPlatformAvailable("ext-mbstring", "")
//	if err != nil {
//	    log.Fatalf("检查扩展失败: %v", err)
//	}
//	if available {
//	    fmt.Println("mbstring扩展已安装")
//	} else {
//	    fmt.Println("mbstring扩展未安装")
//	}
func (c *Composer) IsPlatformAvailable(platform string, version string) (bool, error) {
	platformReqs, err := c.CheckPlatform()
	if err != nil {
		return false, err
	}

	for _, info := range platformReqs {
		if info.Name == platform {
			return info.Available, nil
		}
	}

	// 平台未在需求中列出，尝试直接检查
	versionCheck := platform
	if version != "" {
		versionCheck += ":" + version
	}

	output, err := c.Run("check-platform", versionCheck)
	if err != nil {
		return false, err
	}

	// 如果输出包含 "is not available"，则平台不可用
	return !strings.Contains(output, "is not available"), nil
}

// GetPHPVersion 获取当前 PHP 版本
//
// 返回值：
//   - string: 当前系统使用的PHP版本号
//   - error: 如果获取PHP版本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取当前系统使用的PHP版本号。
//	相当于执行`composer run --php-show-version`并解析输出。
//
// 用法示例：
//
//	phpVersion, err := comp.GetPHPVersion()
//	if err != nil {
//	    log.Fatalf("获取PHP版本失败: %v", err)
//	}
//	fmt.Printf("当前PHP版本: %s\n", phpVersion)
func (c *Composer) GetPHPVersion() (string, error) {
	output, err := c.Run("run", "--php-show-version")
	if err != nil {
		return "", err
	}

	// 解析输出以提取 PHP 版本
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PHP ") {
			parts := strings.Split(line, " ")
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}

	return "", nil
}

// GetExtensions 获取已安装的 PHP 扩展
//
// 返回值：
//   - []string: 已安装的PHP扩展名列表
//   - error: 如果获取扩展列表过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取当前PHP环境中已安装的所有扩展。
//	相当于执行`composer run --show-extensions`并解析输出。
//
// 用法示例：
//
//	extensions, err := comp.GetExtensions()
//	if err != nil {
//	    log.Fatalf("获取PHP扩展失败: %v", err)
//	}
//	fmt.Println("已安装的PHP扩展:")
//	for _, ext := range extensions {
//	    fmt.Println("- " + ext)
//	}
func (c *Composer) GetExtensions() ([]string, error) {
	output, err := c.Run("run", "--show-extensions")
	if err != nil {
		return nil, err
	}

	// 解析输出以提取扩展名
	extensions := []string{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "Loaded extensions:") {
			extensions = append(extensions, line)
		}
	}

	return extensions, nil
}

// HasExtension 检查是否安装了指定的 PHP 扩展
//
// 参数：
//   - extension: 要检查的PHP扩展名，例如"mbstring"或"pdo"
//
// 返回值：
//   - bool: 如果扩展已安装则返回true，否则返回false
//   - error: 如果检查扩展过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法检查当前PHP环境中是否安装了指定的扩展。
//
// 用法示例：
//
//	// 检查是否安装了JSON扩展
//	hasJson, err := comp.HasExtension("json")
//	if err != nil {
//	    log.Fatalf("检查扩展失败: %v", err)
//	}
//	if hasJson {
//	    fmt.Println("JSON扩展已安装")
//	} else {
//	    fmt.Println("JSON扩展未安装")
//	}
func (c *Composer) HasExtension(extension string) (bool, error) {
	extensions, err := c.GetExtensions()
	if err != nil {
		return false, err
	}

	for _, ext := range extensions {
		if ext == extension {
			return true, nil
		}
	}

	return false, nil
}

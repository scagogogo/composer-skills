package composer

import (
	"fmt"
	"strings"
)

// Validate 验证composer.json是否有效
func (c *Composer) Validate() error {
	_, err := c.Run("validate")
	return err
}

// GetComposerHome 获取Composer主目录
func (c *Composer) GetComposerHome() (string, error) {
	output, err := c.Run("config", "--global", "home")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// ClearCache 清除Composer缓存
func (c *Composer) ClearCache() error {
	_, err := c.Run("clear-cache")
	return err
}

// GetConfigWithGlobal 获取Composer配置项的值，可选择是否获取全局配置
func (c *Composer) GetConfigWithGlobal(setting string, global bool) (string, error) {
	args := []string{"config"}

	if global {
		args = append(args, "--global")
	}

	args = append(args, setting)

	output, err := c.Run(args...)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}

// SetConfigWithGlobal 设置Composer配置项的值，可选择是否设置全局配置
func (c *Composer) SetConfigWithGlobal(setting string, value string, global bool) error {
	args := []string{"config"}

	if global {
		args = append(args, "--global")
	}

	args = append(args, setting, value)

	_, err := c.Run(args...)
	return err
}

// ValidateComposerJson 验证composer.json，可选参数：strict和with-dependencies
func (c *Composer) ValidateComposerJson(strict bool, withDependencies bool) error {
	args := []string{"validate"}

	if strict {
		args = append(args, "--strict")
	}

	if withDependencies {
		args = append(args, "--with-dependencies")
	}

	_, err := c.Run(args...)
	return err
}

// CheckPlatformReqs 检查平台要求
func (c *Composer) CheckPlatformReqs() (string, error) {
	return c.Run("check-platform-reqs")
}

// ListConfig 列出所有配置值
//
// 返回值：
//   - string: 所有配置值的列表
//   - error: 如果列出配置值过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出Composer的所有配置值。相当于执行`composer config --list`。
//	可以查看当前项目的所有Composer配置项及其值。
//
// 用法示例：
//
//	output, err := comp.ListConfig()
//	if err != nil {
//	    log.Fatalf("列出配置失败: %v", err)
//	}
//	fmt.Println("配置列表:", output)
func (c *Composer) ListConfig() (string, error) {
	output, err := c.Run("config", "--list")
	if err != nil {
		return "", fmt.Errorf("列出配置失败: %w", err)
	}
	return output, nil
}

// ListConfigWithGlobal 列出所有配置值，可选择是否列出全局配置
//
// 参数：
//   - global: 如果为true，则列出全局配置；如果为false，则列出项目配置
//
// 返回值：
//   - string: 所有配置值的列表
//   - error: 如果列出配置值过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出Composer的所有配置值，可以选择列出全局配置或项目配置。
//	相当于执行`composer config --list [--global]`。
//
// 用法示例：
//
//	// 列出全局配置
//	output, err := comp.ListConfigWithGlobal(true)
//	if err != nil {
//	    log.Fatalf("列出全局配置失败: %v", err)
//	}
//	fmt.Println("全局配置:", output)
//
//	// 列出项目配置
//	output, err = comp.ListConfigWithGlobal(false)
//	if err != nil {
//	    log.Fatalf("列出项目配置失败: %v", err)
//	}
//	fmt.Println("项目配置:", output)
func (c *Composer) ListConfigWithGlobal(global bool) (string, error) {
	args := []string{"config", "--list"}

	if global {
		args = append(args, "--global")
	}

	output, err := c.Run(args...)
	if err != nil {
		return "", fmt.Errorf("列出配置失败: %w", err)
	}
	return output, nil
}

// GetConfigSource 获取配置项的来源
//
// 参数：
//   - key: 要查询来源的配置项名称
//
// 返回值：
//   - string: 配置项的来源信息
//   - error: 如果获取配置来源过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取指定配置项的来源信息，显示该配置值是从哪个文件加载的。
//	相当于执行`composer config key --source`。
//	可以帮助调试配置问题，了解某个配置值是来自项目、全局还是默认值。
//
// 用法示例：
//
//	output, err := comp.GetConfigSource("preferred-install")
//	if err != nil {
//	    log.Fatalf("获取配置来源失败: %v", err)
//	}
//	fmt.Println("配置来源:", output)
func (c *Composer) GetConfigSource(key string) (string, error) {
	output, err := c.Run("config", key, "--source")
	if err != nil {
		return "", fmt.Errorf("获取配置来源失败: %w", err)
	}
	return strings.TrimSpace(output), nil
}

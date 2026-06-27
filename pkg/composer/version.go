package composer

import (
	"fmt"
	"strings"
)

// GetVersion 获取composer版本
//
// 返回值：
//   - string: Composer的版本号，例如"2.1.6"
//   - error: 如果获取版本信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行`composer --version`命令，并解析输出以获取Composer的版本号。
//	它会从类似"Composer version 2.1.6 2021-08-19 17:11:08"的输出中提取版本号。
//
// 用法示例：
//
//	version, err := comp.GetVersion()
//	if err != nil {
//	    log.Fatalf("获取Composer版本失败: %v", err)
//	}
//	fmt.Printf("当前Composer版本: %s\n", version)
func (c *Composer) GetVersion() (string, error) {
	output, err := c.Run("--version")
	if err != nil {
		return "", err
	}

	// 解析版本信息
	// 示例输出: "Composer version 2.1.6 2021-08-19 17:11:08"
	parts := strings.Split(output, " ")
	if len(parts) >= 3 {
		return parts[2], nil
	}

	return "", fmt.Errorf("无法解析版本信息: %s", output)
}

// SelfUpdate 执行composer self-update命令
//
// 返回值：
//   - error: 如果自我更新过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法更新Composer本身到最新版本。
//	相当于执行`composer self-update`。
//
// 用法示例：
//
//	err := comp.SelfUpdate()
//	if err != nil {
//	    log.Fatalf("更新Composer失败: %v", err)
//	} else {
//	    fmt.Println("Composer已更新到最新版本")
//	}
func (c *Composer) SelfUpdate() error {
	_, err := c.Run("self-update")
	return err
}

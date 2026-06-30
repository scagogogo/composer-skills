package composer

import (
	"fmt"
	"strings"
)

// Home 打开项目的主页
//
// 返回值：
//   - string: 命令的输出结果
//   - error: 如果打开主页过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法打开当前项目的主页URL。相当于执行`composer home`。
//	Composer会根据composer.json中的homepage字段确定主页地址。
//
// 用法示例：
//
//	output, err := comp.Home()
//	if err != nil {
//	    log.Fatalf("打开项目主页失败: %v", err)
//	}
//	fmt.Println("主页信息:", output)
func (c *Composer) Home() (string, error) {
	output, err := c.Run("home")
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCommandExecution, err)
	}
	return strings.TrimSpace(output), nil
}

// HomePackage 打开指定包的主页
//
// 参数：
//   - packageName: 要打开主页的包名，例如"symfony/console"
//
// 返回值：
//   - string: 命令的输出结果
//   - error: 如果打开包主页过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法打开指定包的主页URL。相当于执行`composer home package/name`。
//	Composer会根据包的composer.json中的homepage字段确定主页地址。
//
// 用法示例：
//
//	output, err := comp.HomePackage("symfony/console")
//	if err != nil {
//	    log.Fatalf("打开包主页失败: %v", err)
//	}
//	fmt.Println("包主页信息:", output)
func (c *Composer) HomePackage(packageName string) (string, error) {
	output, err := c.Run("home", packageName)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCommandExecution, err)
	}
	return strings.TrimSpace(output), nil
}

// HomeWithOptions 使用更多选项打开主页
//
// 参数：
//   - options: 打开主页时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - string: 命令的输出结果
//   - error: 如果打开主页过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法打开项目或包的主页，支持更多自定义选项。
//	例如可以使用--show选项仅显示URL而不在浏览器中打开。
//
// 用法示例：
//
//	// 仅显示主页URL而不打开浏览器
//	options := map[string]string{
//	    "show": "",
//	}
//	output, err := comp.HomeWithOptions(options)
//	if err != nil {
//	    log.Fatalf("打开主页失败: %v", err)
//	}
//	fmt.Println("主页URL:", output)
func (c *Composer) HomeWithOptions(options map[string]string) (string, error) {
	args := []string{"home"}

	// 添加选项
	args = append(args, buildOptionsArgs(options)...)

	output, err := c.Run(args...)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCommandExecution, err)
	}
	return strings.TrimSpace(output), nil
}

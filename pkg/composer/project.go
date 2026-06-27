package composer

import (
	"encoding/json"
)

// CreateProject 创建一个新的项目
//
// 参数：
//   - packageName: 包名，例如"laravel/laravel"
//   - directory: 要创建项目的目录
//   - version: 包的版本，如果为空则使用最新版本
//
// 返回值：
//   - error: 如果创建项目过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用Composer创建一个基于指定包的新项目。相当于执行
//	`composer create-project package/name directory version`
//
// 用法示例：
//
//	// 创建最新版本的Laravel项目
//	err := comp.CreateProject("laravel/laravel", "my-project", "")
//	if err != nil {
//	    log.Fatalf("创建项目失败: %v", err)
//	}
//
//	// 创建指定版本的Symfony项目
//	err = comp.CreateProject("symfony/website-skeleton", "symfony-project", "^5.0")
func (c *Composer) CreateProject(packageName string, directory string, version string) error {
	args := []string{"create-project"}

	if version != "" {
		args = append(args, packageName+":"+version, directory)
	} else {
		args = append(args, packageName, directory)
	}

	_, err := c.Run(args...)
	return err
}

// CreateProjectWithOptions 使用更多选项创建一个新的项目
//
// 参数：
//   - packageName: 包名，例如"laravel/laravel"
//   - directory: 要创建项目的目录
//   - version: 包的版本，如果为空则使用最新版本
//   - options: 创建项目时的额外选项，键为选项名，值为选项值
//
// 返回值：
//   - error: 如果创建项目过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法使用Composer创建一个基于指定包的新项目，支持更多自定义选项。
//
// 用法示例：
//
//	// 创建无开发依赖的Laravel项目
//	options := map[string]string{
//	    "no-dev": "",
//	    "prefer-dist": "",
//	}
//	err := comp.CreateProjectWithOptions("laravel/laravel", "my-project", "", options)
//	if err != nil {
//	    log.Fatalf("创建项目失败: %v", err)
//	}
//
//	// 创建指定稳定性的Symfony项目
//	options = map[string]string{
//	    "stability": "dev",
//	    "repository": "https://example.org/private-repo",
//	}
//	err = comp.CreateProjectWithOptions("symfony/website-skeleton", "symfony-project", "^5.0", options)
func (c *Composer) CreateProjectWithOptions(packageName string, directory string, version string, options map[string]string) error {
	args := []string{"create-project"}

	// 添加选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	if version != "" {
		args = append(args, packageName+":"+version, directory)
	} else {
		args = append(args, packageName, directory)
	}

	_, err := c.Run(args...)
	return err
}

// InitProject 初始化一个新的项目
//
// 返回值：
//   - error: 如果初始化项目过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法初始化一个新的Composer项目，创建一个基本的composer.json文件。
//	相当于执行`composer init`，会通过交互方式询问项目信息。
//
// 用法示例：
//
//	err := comp.InitProject()
//	if err != nil {
//	    log.Fatalf("初始化项目失败: %v", err)
//	}
func (c *Composer) InitProject() error {
	_, err := c.Run("init")
	return err
}

// InitProjectWithOptions 使用更多选项初始化一个新的项目
//
// 参数：
//   - name: 项目名称，格式为"vendor/name"
//   - description: 项目描述
//   - author: 作者信息，格式为"Name <email>"
//   - options: 初始化项目时的额外选项
//
// 返回值：
//   - error: 如果初始化项目过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法初始化一个新的Composer项目，创建一个自定义的composer.json文件，
//	支持指定项目名称、描述、作者和其他选项。
//
// 用法示例：
//
//	// 初始化一个库项目
//	options := map[string]string{
//	    "type": "library",
//	    "license": "MIT",
//	    "no-interaction": "",
//	}
//	err := comp.InitProjectWithOptions(
//	    "myvendor/awesome-lib",
//	    "一个很棒的PHP库",
//	    "张三 <zhangsan@example.com>",
//	    options,
//	)
//	if err != nil {
//	    log.Fatalf("初始化项目失败: %v", err)
//	}
func (c *Composer) InitProjectWithOptions(name string, description string, author string, options map[string]string) error {
	args := []string{"init"}

	// 添加命名选项
	if name != "" {
		args = append(args, "--name="+name)
	}

	if description != "" {
		args = append(args, "--description="+description)
	}

	if author != "" {
		args = append(args, "--author="+author)
	}

	// 添加其他选项
	for key, value := range options {
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}

	_, err := c.Run(args...)
	return err
}

// RunScript 执行composer脚本
//
// 参数：
//   - scriptName: 要执行的脚本名称
//   - args: 传递给脚本的额外参数
//
// 返回值：
//   - string: 命令执行的输出结果
//   - error: 如果执行脚本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行composer.json中定义的脚本，可以传递额外参数。
//	相当于执行`composer run-script script-name -- arg1 arg2`
//
// 用法示例：
//
//	// 执行测试脚本并传递参数
//	output, err := comp.RunScript("test", "--filter=UserTest")
//	if err != nil {
//	    log.Fatalf("执行脚本失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) RunScript(scriptName string, args ...string) (string, error) {
	cmdArgs := append([]string{"run-script", scriptName}, args...)
	return c.Run(cmdArgs...)
}

// ExecuteScript 执行composer.json中定义的自定义脚本
//
// 参数：
//   - scriptName: 要执行的脚本名称
//
// 返回值：
//   - string: 命令执行的输出结果
//   - error: 如果执行脚本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法执行composer.json中定义的脚本，使用`composer run`命令。
//	相当于执行`composer run script-name`
//
// 用法示例：
//
//	// 执行部署脚本
//	output, err := comp.ExecuteScript("deploy")
//	if err != nil {
//	    log.Fatalf("执行脚本失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) ExecuteScript(scriptName string) (string, error) {
	return c.Run("run", scriptName)
}

// Archive 创建项目的存档
//
// 参数：
//   - directory: 存档输出目录，为空则使用当前目录
//   - format: 存档格式，如"zip"或"tar"，为空则使用默认格式
//
// 返回值：
//   - error: 如果创建存档过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法创建当前项目的存档文件。相当于执行
//	`composer archive --dir=directory --format=format`
//
// 用法示例：
//
//	// 创建ZIP格式的存档
//	err := comp.ArchiveProject("./dist", "zip")
//	if err != nil {
//	    log.Fatalf("创建存档失败: %v", err)
//	}
func (c *Composer) ArchiveProject(directory string, format string) error {
	args := []string{"archive"}

	if directory != "" {
		args = append(args, "--dir="+directory)
	}

	if format != "" {
		args = append(args, "--format="+format)
	}

	_, err := c.Run(args...)
	return err
}

// ComposerJsonInfo 表示composer.json文件的部分信息
type ComposerJsonInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Require     map[string]string `json:"require"`
	RequireDev  map[string]string `json:"require-dev"`
}

// GetProjectInfo 获取项目信息
//
// 返回值：
//   - *ComposerJsonInfo: 包含项目基本信息的结构体指针
//   - error: 如果获取项目信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取当前项目的基本信息，包括名称、描述、类型和依赖。
//	相当于执行`composer config --list --json`并解析结果。
//
// 用法示例：
//
//	info, err := comp.GetProjectInfo()
//	if err != nil {
//	    log.Fatalf("获取项目信息失败: %v", err)
//	}
//	fmt.Printf("项目名称: %s\n", info.Name)
//	fmt.Printf("项目描述: %s\n", info.Description)
//	fmt.Printf("依赖数量: %d\n", len(info.Require))
func (c *Composer) GetProjectInfo() (*ComposerJsonInfo, error) {
	output, err := c.Run("config", "--list", "--json")
	if err != nil {
		return nil, err
	}

	var info ComposerJsonInfo
	err = json.Unmarshal([]byte(output), &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// ListScripts 列出composer.json中定义的所有脚本
//
// 返回值：
//   - string: 命令执行的输出结果，包含所有脚本列表
//   - error: 如果列出脚本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出composer.json中定义的所有脚本。
//	相当于执行`composer run-script --list`
//
// 用法示例：
//
//	output, err := comp.ListScripts()
//	if err != nil {
//	    log.Fatalf("列出脚本失败: %v", err)
//	}
//	fmt.Println("可用脚本:")
//	fmt.Println(output)
func (c *Composer) ListScripts() (string, error) {
	return c.Run("run-script", "--list")
}

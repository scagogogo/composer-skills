package composer

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// ComposerJSON 表示 composer.json 文件的结构
type ComposerJSON struct {
	Name                string                 `json:"name,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Type                string                 `json:"type,omitempty"`
	Keywords            []string               `json:"keywords,omitempty"`
	Homepage            string                 `json:"homepage,omitempty"`
	License             interface{}            `json:"license,omitempty"`
	Authors             []map[string]string    `json:"authors,omitempty"`
	Support             map[string]string      `json:"support,omitempty"`
	Require             map[string]string      `json:"require,omitempty"`
	RequireDev          map[string]string      `json:"require-dev,omitempty"`
	Suggest             map[string]string      `json:"suggest,omitempty"`
	Autoload            map[string]interface{} `json:"autoload,omitempty"`
	AutoloadDev         map[string]interface{} `json:"autoload-dev,omitempty"`
	Repositories        map[string]interface{} `json:"repositories,omitempty"`
	Config              map[string]interface{} `json:"config,omitempty"`
	Scripts             map[string]interface{} `json:"scripts,omitempty"`
	ScriptsDescriptions map[string]string      `json:"scripts-descriptions,omitempty"`
	Extra               map[string]interface{} `json:"extra,omitempty"`
	Bin                 []string               `json:"bin,omitempty"`
	Archive             map[string]interface{} `json:"archive,omitempty"`
	NonFeatureBranches  []string               `json:"non-feature-branches,omitempty"`
	MinimumStability    string                 `json:"minimum-stability,omitempty"`
	PreferStable        bool                   `json:"prefer-stable,omitempty"`
	Replace             map[string]string      `json:"replace,omitempty"`
	Conflict            map[string]string      `json:"conflict,omitempty"`
	Provide             map[string]string      `json:"provide,omitempty"`
}

// ErrComposerJSONNotFound 表示未找到 composer.json 文件
var ErrComposerJSONNotFound = errors.New("composer.json not found")

// mockComposerJSON stores a mock ComposerJSON for testing
var mockComposerJSON *ComposerJSON
var useMockComposerJSON bool

// SetMockComposerJSON sets a mock ComposerJSON for testing
func SetMockComposerJSON(composerJSON *ComposerJSON) {
	mockComposerJSON = composerJSON
	useMockComposerJSON = true
}

// ClearMockComposerJSON clears the mock ComposerJSON
func ClearMockComposerJSON() {
	mockComposerJSON = nil
	useMockComposerJSON = false
}

// ReadComposerJSON 读取并解析 composer.json 文件
//
// 返回值：
//   - *ComposerJSON: 解析后的composer.json结构体指针
//   - error: 如果读取或解析过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法从工作目录读取composer.json文件并解析为结构体。如果未指定工作目录，
//	则使用当前目录。
//
// 用法示例：
//
//	composerJSON, err := comp.ReadComposerJSON()
//	if err != nil {
//	    log.Fatalf("读取composer.json失败: %v", err)
//	}
//	fmt.Printf("项目名称: %s\n", composerJSON.Name)
func (c *Composer) ReadComposerJSON() (*ComposerJSON, error) {
	// 如果设置了mock，则使用mock
	if useMockComposerJSON && mockComposerJSON != nil {
		return mockComposerJSON, nil
	}

	// 如果指定了工作目录，则从工作目录中读取
	workDir := c.workingDir
	if workDir == "" {
		// 如果未指定工作目录，则使用当前目录
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	// 读取 composer.json 文件
	composerJSONPath := filepath.Join(workDir, "composer.json")
	if _, err := os.Stat(composerJSONPath); os.IsNotExist(err) {
		return nil, ErrComposerJSONNotFound
	}

	content, err := os.ReadFile(composerJSONPath)
	if err != nil {
		return nil, err
	}

	// 解析 JSON
	var composerJSON ComposerJSON
	if err := json.Unmarshal(content, &composerJSON); err != nil {
		return nil, err
	}

	return &composerJSON, nil
}

// WriteComposerJSON 将 ComposerJSON 写入 composer.json 文件
//
// 参数：
//   - composerJSON: 要写入的composer.json结构体
//
// 返回值：
//   - error: 如果写入过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法将ComposerJSON结构体序列化为JSON格式，并写入到工作目录下的composer.json文件。
//	如果未指定工作目录，则使用当前目录。
//
// 用法示例：
//
//	composerJSON, _ := comp.ReadComposerJSON()
//	composerJSON.Name = "my/new-project"
//	err := comp.WriteComposerJSON(composerJSON)
//	if err != nil {
//	    log.Fatalf("写入composer.json失败: %v", err)
//	}
func (c *Composer) WriteComposerJSON(composerJSON *ComposerJSON) error {
	// 如果指定了工作目录，则写入工作目录
	workDir := c.workingDir
	if workDir == "" {
		// 如果未指定工作目录，则使用当前目录
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	// 序列化为 JSON
	content, err := json.MarshalIndent(composerJSON, "", "    ")
	if err != nil {
		return err
	}

	// 写入 composer.json 文件
	composerJSONPath := filepath.Join(workDir, "composer.json")
	return os.WriteFile(composerJSONPath, content, 0644)
}

// AddRequire 添加依赖到 composer.json
//
// 参数：
//   - packageName: 包名，例如"symfony/console"
//   - version: 版本约束，例如"^5.0"
//   - isDev: 是否为开发依赖
//
// 返回值：
//   - error: 如果添加依赖过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向composer.json文件添加一个新的依赖包。如果isDev为true，则添加到require-dev；
//	否则添加到require。
//
// 用法示例：
//
//	// 添加生产依赖
//	err := comp.AddRequire("symfony/console", "^5.0", false)
//	if err != nil {
//	    log.Fatalf("添加依赖失败: %v", err)
//	}
//
//	// 添加开发依赖
//	err = comp.AddRequire("phpunit/phpunit", "^9.0", true)
func (c *Composer) AddRequire(packageName, version string, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if isDev {
		if composerJSON.RequireDev == nil {
			composerJSON.RequireDev = make(map[string]string)
		}
		composerJSON.RequireDev[packageName] = version
	} else {
		if composerJSON.Require == nil {
			composerJSON.Require = make(map[string]string)
		}
		composerJSON.Require[packageName] = version
	}

	return c.WriteComposerJSON(composerJSON)
}

// RemoveRequire 从 composer.json 中移除依赖
//
// 参数：
//   - packageName: 要移除的包名
//   - isDev: 是否为开发依赖
//
// 返回值：
//   - error: 如果移除依赖过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法从composer.json文件中移除指定的依赖包。如果isDev为true，则从require-dev中移除；
//	否则从require中移除。
//
// 用法示例：
//
//	// 移除生产依赖
//	err := comp.RemoveRequire("symfony/console", false)
//	if err != nil {
//	    log.Fatalf("移除依赖失败: %v", err)
//	}
//
//	// 移除开发依赖
//	err = comp.RemoveRequire("phpunit/phpunit", true)
func (c *Composer) RemoveRequire(packageName string, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if isDev {
		if composerJSON.RequireDev != nil {
			delete(composerJSON.RequireDev, packageName)
		}
	} else {
		if composerJSON.Require != nil {
			delete(composerJSON.Require, packageName)
		}
	}

	return c.WriteComposerJSON(composerJSON)
}

// AddScript 添加脚本到 composer.json
//
// 参数：
//   - name: 脚本名称，例如"post-install-cmd"或自定义脚本名
//   - script: 脚本内容，可以是字符串或字符串数组
//   - description: 脚本描述，可选
//
// 返回值：
//   - error: 如果添加脚本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向composer.json文件添加一个新的脚本。脚本可以是单个命令字符串、
//	PHP类方法调用或命令数组。如果提供了描述，则会添加到scripts-descriptions字段。
//
// 用法示例：
//
//	// 添加单个命令脚本
//	err := comp.AddScript("post-install-cmd", "php -r \"echo 'Installation completed!';\"", "安装后执行")
//	if err != nil {
//	    log.Fatalf("添加脚本失败: %v", err)
//	}
//
//	// 添加多条命令脚本
//	commands := []string{
//	    "php -r \"echo 'Starting tests...';\"",
//	    "phpunit"
//	}
//	err = comp.AddScript("test", commands, "运行测试")
func (c *Composer) AddScript(name string, script interface{}, description string) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 添加脚本
	if composerJSON.Scripts == nil {
		composerJSON.Scripts = make(map[string]interface{})
	}
	composerJSON.Scripts[name] = script

	// 添加脚本描述（如果提供）
	if description != "" {
		if composerJSON.ScriptsDescriptions == nil {
			composerJSON.ScriptsDescriptions = make(map[string]string)
		}
		composerJSON.ScriptsDescriptions[name] = description
	}

	return c.WriteComposerJSON(composerJSON)
}

// RemoveScript 从 composer.json 中移除脚本
//
// 参数：
//   - name: 要移除的脚本名称
//
// 返回值：
//   - error: 如果移除脚本过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法从composer.json文件中移除指定名称的脚本及其描述。
//
// 用法示例：
//
//	err := comp.RemoveScript("post-install-cmd")
//	if err != nil {
//	    log.Fatalf("移除脚本失败: %v", err)
//	}
func (c *Composer) RemoveScript(name string) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 移除脚本
	if composerJSON.Scripts != nil {
		delete(composerJSON.Scripts, name)
	}

	// 移除脚本描述
	if composerJSON.ScriptsDescriptions != nil {
		delete(composerJSON.ScriptsDescriptions, name)
	}

	return c.WriteComposerJSON(composerJSON)
}

// AddAutoload 添加自动加载配置到 composer.json
//
// 参数：
//   - type_: 自动加载类型，可以是"psr-4"、"psr-0"、"classmap"或"files"
//   - namespace: 命名空间，如"App\\"
//   - paths: 路径，可以是字符串或字符串数组
//   - isDev: 是否为开发自动加载
//
// 返回值：
//   - error: 如果添加自动加载配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向composer.json文件添加一个新的自动加载配置。如果isDev为true，则添加到autoload-dev；
//	否则添加到autoload。
//
// 用法示例：
//
//	// 添加PSR-4自动加载
//	err := comp.AddAutoload("psr-4", "App\\", "src/", false)
//	if err != nil {
//	    log.Fatalf("添加自动加载失败: %v", err)
//	}
//
//	// 添加多目录PSR-4自动加载
//	err = comp.AddAutoload("psr-4", "Tests\\", []string{"tests/", "test-framework/"}, true)
func (c *Composer) AddAutoload(type_ string, namespace string, paths interface{}, isDev bool) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 选择正确的自动加载配置
	var autoload map[string]interface{}
	if isDev {
		if composerJSON.AutoloadDev == nil {
			composerJSON.AutoloadDev = make(map[string]interface{})
		}
		autoload = composerJSON.AutoloadDev
	} else {
		if composerJSON.Autoload == nil {
			composerJSON.Autoload = make(map[string]interface{})
		}
		autoload = composerJSON.Autoload
	}

	// 确保类型的配置存在
	if autoload[type_] == nil {
		autoload[type_] = make(map[string]interface{})
	}

	// 添加命名空间和路径
	typeConfig, ok := autoload[type_].(map[string]interface{})
	if !ok {
		return errors.New("invalid autoload configuration")
	}
	typeConfig[namespace] = paths

	return c.WriteComposerJSON(composerJSON)
}

// SetConfig 设置 composer.json 中的配置选项
//
// 参数：
//   - key: 配置键名
//   - value: 配置值
//
// 返回值：
//   - error: 如果设置配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置composer.json中config部分的配置选项。
//
// 用法示例：
//
//	// 设置进程超时时间
//	err := comp.SetConfig("process-timeout", 500)
//	if err != nil {
//	    log.Fatalf("设置配置失败: %v", err)
//	}
//
//	// 设置禁用插件
//	err = comp.SetConfig("disable-plugins", true)
func (c *Composer) SetConfig(key string, value interface{}) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	if composerJSON.Config == nil {
		composerJSON.Config = make(map[string]interface{})
	}
	composerJSON.Config[key] = value

	return c.WriteComposerJSON(composerJSON)
}

// GetConfig 获取 composer.json 中的配置选项
//
// 参数：
//   - key: 配置键名
//
// 返回值：
//   - interface{}: 配置值，可能是字符串、数字、布尔值或复杂结构
//   - error: 如果获取配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取composer.json中config部分的特定配置选项的值。
//	如果配置不存在，则返回nil。
//
// 用法示例：
//
//	// 获取进程超时时间
//	timeout, err := comp.GetConfig("process-timeout")
//	if err != nil {
//	    log.Fatalf("获取配置失败: %v", err)
//	}
//	if timeout != nil {
//	    fmt.Printf("进程超时时间: %v\n", timeout)
//	}
//
//	// 获取vendor目录配置
//	vendorDir, err := comp.GetConfig("vendor-dir")
//	if err != nil {
//	    log.Fatalf("获取配置失败: %v", err)
//	}
//	if vendorDir != nil {
//	    fmt.Printf("Vendor目录: %s\n", vendorDir)
//	}
func (c *Composer) GetConfig(key string) (interface{}, error) {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return nil, err
	}

	if composerJSON.Config == nil {
		return nil, nil
	}
	return composerJSON.Config[key], nil
}

// SetProperty 设置 composer.json 中的顶级属性
//
// 参数：
//   - property: 属性名，例如"name"、"description"、"type"、"license"等
//   - value: 属性值，类型取决于属性
//
// 返回值：
//   - error: 如果设置属性过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置composer.json文件中的顶级属性。目前支持以下属性：
//	- name: 项目名称，格式为"vendor/package"
//	- description: 项目描述
//	- type: 项目类型，如"library"、"project"、"metapackage"等
//	- keywords: 关键词数组
//	- homepage: 项目主页URL
//	- license: 许可证，可以是字符串或字符串数组
//	- minimum-stability: 最小稳定性，如"stable"、"RC"、"beta"、"alpha"、"dev"
//	- prefer-stable: 是否优先使用稳定版本
//
// 用法示例：
//
//	// 设置项目名称
//	err := comp.SetProperty("name", "myvendor/mypackage")
//	if err != nil {
//	    log.Fatalf("设置项目名称失败: %v", err)
//	}
//
//	// 设置项目描述
//	err = comp.SetProperty("description", "一个很棒的PHP库")
//	if err != nil {
//	    log.Fatalf("设置项目描述失败: %v", err)
//	}
//
//	// 设置关键词
//	err = comp.SetProperty("keywords", []string{"php", "library", "awesome"})
//	if err != nil {
//	    log.Fatalf("设置关键词失败: %v", err)
//	}
//
//	// 设置许可证
//	err = comp.SetProperty("license", "MIT")
//	if err != nil {
//	    log.Fatalf("设置许可证失败: %v", err)
//	}
//
//	// 设置是否优先使用稳定版本
//	err = comp.SetProperty("prefer-stable", true)
//	if err != nil {
//	    log.Fatalf("设置prefer-stable失败: %v", err)
//	}
func (c *Composer) SetProperty(property string, value interface{}) error {
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		return err
	}

	// 使用反射来设置属性可能更好，但为了简单起见，我们这里使用硬编码
	switch property {
	case "name":
		if strValue, ok := value.(string); ok {
			composerJSON.Name = strValue
		}
	case "description":
		if strValue, ok := value.(string); ok {
			composerJSON.Description = strValue
		}
	case "type":
		if strValue, ok := value.(string); ok {
			composerJSON.Type = strValue
		}
	case "keywords":
		if strArray, ok := value.([]string); ok {
			composerJSON.Keywords = strArray
		}
	case "homepage":
		if strValue, ok := value.(string); ok {
			composerJSON.Homepage = strValue
		}
	case "license":
		composerJSON.License = value
	case "minimum-stability":
		if strValue, ok := value.(string); ok {
			composerJSON.MinimumStability = strValue
		}
	case "prefer-stable":
		if boolValue, ok := value.(bool); ok {
			composerJSON.PreferStable = boolValue
		}
	default:
		return errors.New("unsupported property")
	}

	return c.WriteComposerJSON(composerJSON)
}

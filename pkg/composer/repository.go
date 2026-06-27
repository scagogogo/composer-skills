package composer

import (
	"encoding/json"
	"fmt"
)

// RepositoryType 表示 Composer 仓库类型
type RepositoryType string

const (
	// VcsRepository 版本控制系统仓库类型
	VcsRepository RepositoryType = "vcs"
	// ComposerRepository Composer 仓库类型
	ComposerRepository RepositoryType = "composer"
	// PackagistRepository Packagist 仓库类型
	PackagistRepository RepositoryType = "packagist"
	// PathRepository 路径仓库类型
	PathRepository RepositoryType = "path"
	// ArtifactRepository artifact 仓库类型
	ArtifactRepository RepositoryType = "artifact"
	// PearRepository PEAR 仓库类型
	PearRepository RepositoryType = "pear"
)

// Repository 表示一个 Composer 仓库
type Repository struct {
	Type    RepositoryType         `json:"type"`
	URL     string                 `json:"url,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// AddRepository 添加一个仓库到 composer.json
//
// 参数：
//   - name: 仓库名称
//   - repo: 仓库结构体，包含类型、URL等信息
//
// 返回值：
//   - error: 如果添加仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法向composer.json中添加一个自定义仓库配置。
//	相当于执行`composer config repositories.name '{"type":"...","url":"..."}'`
//
// 用法示例：
//
//	// 添加一个私有的Composer仓库
//	repo := composer.Repository{
//	    Type: composer.ComposerRepository,
//	    URL:  "https://composer.example.org",
//	}
//	err := comp.AddRepository("private", repo)
//	if err != nil {
//	    log.Fatalf("添加仓库失败: %v", err)
//	}
func (c *Composer) AddRepository(name string, repo Repository) error {
	repoJSON, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	args := []string{"config", "repositories." + name, string(repoJSON)}
	_, err = c.Run(args...)
	return err
}

// RemoveRepository 从 composer.json 中移除仓库
//
// 参数：
//   - name: 要移除的仓库名称
//
// 返回值：
//   - error: 如果移除仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法从composer.json中移除指定名称的仓库配置。
//	相当于执行`composer config --unset repositories.name`
//
// 用法示例：
//
//	// 移除名为"private"的仓库
//	err := comp.RemoveRepository("private")
//	if err != nil {
//	    log.Fatalf("移除仓库失败: %v", err)
//	}
func (c *Composer) RemoveRepository(name string) error {
	args := []string{"config", "--unset", "repositories." + name}
	_, err := c.Run(args...)
	return err
}

// ListRepositories 列出当前项目中配置的所有仓库
//
// 返回值：
//   - string: 列出所有仓库的输出结果
//   - error: 如果列出仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出项目中配置的所有仓库。
//	相当于执行`composer config repositories`
//
// 用法示例：
//
//	output, err := comp.ListRepositories()
//	if err != nil {
//	    log.Fatalf("列出仓库失败: %v", err)
//	}
//	fmt.Println("已配置的仓库:", output)
func (c *Composer) ListRepositories() (string, error) {
	return c.Run("config", "repositories")
}

// AddPackagistRepository 添加 Packagist.org 仓库
//
// 参数：
//   - url: Packagist.org的URL
//
// 返回值：
//   - error: 如果添加Packagist仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加Packagist.org作为Composer包源。
//
// 用法示例：
//
//	// 添加官方Packagist仓库
//	err := comp.AddPackagistRepository("https://repo.packagist.org")
//	if err != nil {
//	    log.Fatalf("添加Packagist仓库失败: %v", err)
//	}
//
//	// 添加Packagist镜像
//	err = comp.AddPackagistRepository("https://mirrors.aliyun.com/composer")
//	if err != nil {
//	    log.Fatalf("添加Packagist镜像失败: %v", err)
//	}
func (c *Composer) AddPackagistRepository(url string) error {
	repo := Repository{
		Type: PackagistRepository,
		URL:  url,
	}
	return c.AddRepository("packagist.org", repo)
}

// DisablePackagistRepository 禁用 Packagist.org 仓库
//
// 返回值：
//   - error: 如果禁用Packagist仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法禁用官方Packagist.org仓库，适用于只使用私有仓库的场景。
//	相当于执行`composer config repositories.packagist.org.url false`
//
// 用法示例：
//
//	err := comp.DisablePackagistRepository()
//	if err != nil {
//	    log.Fatalf("禁用Packagist仓库失败: %v", err)
//	}
func (c *Composer) DisablePackagistRepository() error {
	args := []string{"config", "repositories.packagist.org.url", "false"}
	_, err := c.Run(args...)
	return err
}

// EnablePackagistRepository 启用 Packagist.org 仓库
//
// 返回值：
//   - error: 如果启用Packagist仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法启用官方Packagist.org仓库。
//	相当于执行`composer config repositories.packagist.org.url https://repo.packagist.org`
//
// 用法示例：
//
//	err := comp.EnablePackagistRepository()
//	if err != nil {
//	    log.Fatalf("启用Packagist仓库失败: %v", err)
//	}
func (c *Composer) EnablePackagistRepository() error {
	args := []string{"config", "repositories.packagist.org.url", "https://repo.packagist.org"}
	_, err := c.Run(args...)
	return err
}

// AddVcsRepository 添加版本控制系统仓库
//
// 参数：
//   - name: 仓库名称
//   - url: 版本控制系统仓库的URL
//
// 返回值：
//   - error: 如果添加VCS仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加一个版本控制系统（如Git、SVN）仓库作为包源。
//
// 用法示例：
//
//	// 添加GitHub仓库
//	err := comp.AddVcsRepository("my-lib", "https://github.com/vendor/package")
//	if err != nil {
//	    log.Fatalf("添加VCS仓库失败: %v", err)
//	}
func (c *Composer) AddVcsRepository(name string, url string) error {
	repo := Repository{
		Type: VcsRepository,
		URL:  url,
	}
	return c.AddRepository(name, repo)
}

// AddPathRepository 添加本地路径仓库
//
// 参数：
//   - name: 仓库名称
//   - path: 本地路径
//   - options: 仓库选项，如"symlink"等
//
// 返回值：
//   - error: 如果添加路径仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加一个本地路径作为包源，通常用于开发多个相互依赖的包。
//
// 用法示例：
//
//	// 添加本地路径仓库
//	options := map[string]interface{}{
//	    "symlink": true,
//	}
//	err := comp.AddPathRepository("local", "../my-package", options)
//	if err != nil {
//	    log.Fatalf("添加路径仓库失败: %v", err)
//	}
func (c *Composer) AddPathRepository(name string, path string, options map[string]interface{}) error {
	repo := Repository{
		Type:    PathRepository,
		URL:     path,
		Options: options,
	}
	return c.AddRepository(name, repo)
}

// AddComposerRepository 添加 Composer 仓库
//
// 参数：
//   - name: 仓库名称
//   - url: Composer仓库的URL
//
// 返回值：
//   - error: 如果添加Composer仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加一个Composer格式的仓库作为包源。
//
// 用法示例：
//
//	// 添加私有Composer仓库
//	err := comp.AddComposerRepository("private", "https://composer.example.org")
//	if err != nil {
//	    log.Fatalf("添加Composer仓库失败: %v", err)
//	}
func (c *Composer) AddComposerRepository(name string, url string) error {
	repo := Repository{
		Type: ComposerRepository,
		URL:  url,
	}
	return c.AddRepository(name, repo)
}

// GetPreferredInstall 获取 preferred-install 配置
//
// 返回值：
//   - string: 当前的preferred-install配置
//   - error: 如果获取配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取项目的preferred-install配置，指定是优先使用dist还是source安装。
//	相当于执行`composer config preferred-install`
//
// 用法示例：
//
//	value, err := comp.GetPreferredInstall()
//	if err != nil {
//	    log.Fatalf("获取preferred-install失败: %v", err)
//	}
//	fmt.Printf("当前的preferred-install: %s\n", value)
func (c *Composer) GetPreferredInstall() (string, error) {
	return c.Run("config", "preferred-install")
}

// SetPreferredInstall 设置 preferred-install 配置
//
// 参数：
//   - value: preferred-install的值，可以是"dist"、"source"或"auto"
//
// 返回值：
//   - error: 如果设置配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置项目的preferred-install配置，控制Composer安装包的方式。
//	- "dist": 优先使用打包的分发版本
//	- "source": 优先使用源代码
//	- "auto": 自动选择最合适的方式
//	相当于执行`composer config preferred-install value`
//
// 用法示例：
//
//	// 设置优先使用打包版本
//	err := comp.SetPreferredInstall("dist")
//	if err != nil {
//	    log.Fatalf("设置preferred-install失败: %v", err)
//	}
func (c *Composer) SetPreferredInstall(value string) error {
	if value != "dist" && value != "source" && value != "auto" {
		return fmt.Errorf("invalid preferred-install value: %s, must be 'dist', 'source' or 'auto'", value)
	}

	args := []string{"config", "preferred-install", value}
	_, err := c.Run(args...)
	return err
}

// SetMinimumStability 设置最小稳定性配置
//
// 参数：
//   - stability: 最小稳定性级别，如"stable"、"RC"、"beta"、"alpha"或"dev"
//
// 返回值：
//   - error: 如果设置最小稳定性过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置项目的最小稳定性级别，控制Composer可以安装的包的稳定性级别。
//	相当于执行`composer config minimum-stability stability`
//
// 用法示例：
//
//	// 允许安装beta版本的包
//	err := comp.SetMinimumStability("beta")
//	if err != nil {
//	    log.Fatalf("设置最小稳定性失败: %v", err)
//	}
func (c *Composer) SetMinimumStability(stability string) error {
	args := []string{"config", "minimum-stability", stability}
	_, err := c.Run(args...)
	return err
}

// GetMinimumStability 获取最小稳定性配置
//
// 返回值：
//   - string: 当前的最小稳定性配置
//   - error: 如果获取配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取项目的最小稳定性级别配置。
//	相当于执行`composer config minimum-stability`
//
// 用法示例：
//
//	stability, err := comp.GetMinimumStability()
//	if err != nil {
//	    log.Fatalf("获取最小稳定性失败: %v", err)
//	}
//	fmt.Printf("当前的最小稳定性: %s\n", stability)
func (c *Composer) GetMinimumStability() (string, error) {
	return c.Run("config", "minimum-stability")
}

// GetPreferStable 获取是否优先使用稳定版本的配置
//
// 返回值：
//   - string: 当前的prefer-stable配置（"0"表示false，"1"表示true）
//   - error: 如果获取配置过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取项目是否优先使用稳定版本的配置。
//	相当于执行`composer config prefer-stable`
//
// 用法示例：
//
//	value, err := comp.GetPreferStable()
//	if err != nil {
//	    log.Fatalf("获取prefer-stable失败: %v", err)
//	}
//	preferStable := value == "1"
//	fmt.Printf("当前是否优先使用稳定版本: %v\n", preferStable)
func (c *Composer) GetPreferStable() (string, error) {
	return c.Run("config", "prefer-stable")
}

// AddArtifactRepository 添加一个本地制品仓库
//
// 参数：
//   - name: 仓库名称
//   - path: 制品目录路径，通常包含 .zip 或 .tar 文件
//
// 返回值：
//   - error: 如果添加制品仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加一个本地制品仓库作为包源，该目录下应包含包的压缩文件。
//	相当于在composer.json中添加一个类型为"artifact"的仓库。
//
// 用法示例：
//
//	// 添加本地制品仓库
//	err := comp.AddArtifactRepository("artifacts", "./packages")
//	if err != nil {
//	    log.Fatalf("添加制品仓库失败: %v", err)
//	}
func (c *Composer) AddArtifactRepository(name string, path string) error {
	repo := Repository{
		Type: ArtifactRepository,
		URL:  path,
	}
	return c.AddRepository(name, repo)
}

// SetConfigParameter 设置composer配置项
//
// 参数：
//   - key: 配置项名称
//   - value: 配置项值
//
// 返回值：
//   - error: 如果设置配置项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置composer的配置项，可以是全局配置或项目配置。
//	相当于执行`composer config key value`
//
// 用法示例：
//
//	// 设置项目描述
//	err := comp.SetConfigParameter("description", "我的PHP项目")
//	if err != nil {
//	    log.Fatalf("设置配置项失败: %v", err)
//	}
//
//	// 设置作者信息
//	err = comp.SetConfigParameter("authors.0.name", "张三")
//	if err != nil {
//	    log.Fatalf("设置配置项失败: %v", err)
//	}
//	err = comp.SetConfigParameter("authors.0.email", "zhangsan@example.com")
//	if err != nil {
//	    log.Fatalf("设置配置项失败: %v", err)
//	}
func (c *Composer) SetConfigParameter(key string, value string) error {
	args := []string{"config", key, value}
	_, err := c.Run(args...)
	return err
}

// GetConfigParameter 获取composer配置项
//
// 参数：
//   - key: 配置项名称
//
// 返回值：
//   - string: 配置项的值
//   - error: 如果获取配置项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法获取composer的配置项值。
//	相当于执行`composer config key`
//
// 用法示例：
//
//	// 获取项目名称
//	name, err := comp.GetConfigParameter("name")
//	if err != nil {
//	    log.Fatalf("获取配置项失败: %v", err)
//	}
//	fmt.Printf("项目名称: %s\n", name)
//
//	// 获取项目类型
//	typ, err := comp.GetConfigParameter("type")
//	if err != nil {
//	    log.Fatalf("获取配置项失败: %v", err)
//	}
//	fmt.Printf("项目类型: %s\n", typ)
func (c *Composer) GetConfigParameter(key string) (string, error) {
	return c.Run("config", key)
}

// UnsetConfig 删除composer配置项
//
// 参数：
//   - key: 要删除的配置项名称
//
// 返回值：
//   - error: 如果删除配置项过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法删除composer的配置项。
//	相当于执行`composer config --unset key`
//
// 用法示例：
//
//	// 删除一个不再需要的仓库
//	err := comp.UnsetConfig("repositories.old-repo")
//	if err != nil {
//	    log.Fatalf("删除配置项失败: %v", err)
//	}
func (c *Composer) UnsetConfig(key string) error {
	args := []string{"config", "--unset", key}
	_, err := c.Run(args...)
	return err
}

// AddGlobalRepository 添加全局仓库
//
// 参数：
//   - name: 仓库名称
//   - repo: 仓库结构体，包含类型、URL等信息
//
// 返回值：
//   - error: 如果添加全局仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法添加一个全局仓库，对所有项目生效。
//	相当于执行`composer config --global repositories.name '{"type":"...","url":"..."}'`
//
// 用法示例：
//
//	// 添加全局Composer仓库
//	repo := composer.Repository{
//	    Type: composer.ComposerRepository,
//	    URL:  "https://composer.example.org",
//	}
//	err := comp.AddGlobalRepository("global-private", repo)
//	if err != nil {
//	    log.Fatalf("添加全局仓库失败: %v", err)
//	}
func (c *Composer) AddGlobalRepository(name string, repo Repository) error {
	repoJSON, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	args := []string{"config", "--global", "repositories." + name, string(repoJSON)}
	_, err = c.Run(args...)
	return err
}

// RemoveGlobalRepository 删除全局仓库
//
// 参数：
//   - name: 要删除的全局仓库名称
//
// 返回值：
//   - error: 如果删除全局仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法删除一个全局仓库配置。
//	相当于执行`composer config --global --unset repositories.name`
//
// 用法示例：
//
//	// 删除全局仓库
//	err := comp.RemoveGlobalRepository("global-private")
//	if err != nil {
//	    log.Fatalf("删除全局仓库失败: %v", err)
//	}
func (c *Composer) RemoveGlobalRepository(name string) error {
	args := []string{"config", "--global", "--unset", "repositories." + name}
	_, err := c.Run(args...)
	return err
}

// ListGlobalRepositories 列出所有全局仓库
//
// 返回值：
//   - string: 全局仓库列表的输出
//   - error: 如果列出全局仓库过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法列出所有已配置的全局仓库。
//	相当于执行`composer config --global repositories`
//
// 用法示例：
//
//	// 列出所有全局仓库
//	output, err := comp.ListGlobalRepositories()
//	if err != nil {
//	    log.Fatalf("列出全局仓库失败: %v", err)
//	}
//	fmt.Println("全局仓库列表:", output)
func (c *Composer) ListGlobalRepositories() (string, error) {
	return c.Run("config", "--global", "repositories")
}

// SetPreferStable 设置是否优先使用稳定版本包
//
// 参数：
//   - preferStable: 是否优先使用稳定版本，true表示优先使用稳定版本
//
// 返回值：
//   - error: 如果设置prefer-stable过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法设置composer.json中的prefer-stable选项。
//	如果设置为true，则Composer将优先选择稳定版本的包，即使minimum-stability允许不稳定版本。
//	相当于执行`composer config prefer-stable <0|1>`
//
// 用法示例：
//
//	// 设置优先使用稳定版本
//	err := comp.SetPreferStable(true)
//	if err != nil {
//	    log.Fatalf("设置prefer-stable失败: %v", err)
//	}
func (c *Composer) SetPreferStable(preferStable bool) error {
	value := "0"
	if preferStable {
		value = "1"
	}

	args := []string{"config", "prefer-stable", value}
	_, err := c.Run(args...)
	return err
}

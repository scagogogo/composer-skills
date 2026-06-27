package composer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SatisConfig 表示 Satis 配置
type SatisConfig struct {
	Name                   string                 `json:"name"`
	Homepage               string                 `json:"homepage"`
	Repositories           []map[string]string    `json:"repositories"`
	OutputDir              string                 `json:"output-dir"`
	RequireAll             bool                   `json:"require-all,omitempty"`
	RequireDependencies    bool                   `json:"require-dependencies,omitempty"`
	RequireDevDependencies bool                   `json:"require-dev-dependencies,omitempty"`
	Require                map[string]string      `json:"require,omitempty"`
	Archive                map[string]interface{} `json:"archive,omitempty"`
	MinimumStability       string                 `json:"minimum-stability,omitempty"`
	Providers              bool                   `json:"providers,omitempty"`
	ProvidersURL           string                 `json:"providers-url,omitempty"`
	Config                 map[string]interface{} `json:"config,omitempty"`
	Notify                 map[string]interface{} `json:"notify,omitempty"`
	TwigTemplate           string                 `json:"twig-template,omitempty"`
}

// CreateSatisConfig 创建一个新的 Satis 配置文件
func (c *Composer) CreateSatisConfig(configPath string, name string, homepage string) error {
	// 创建基本配置
	config := SatisConfig{
		Name:         name,
		Homepage:     homepage,
		Repositories: []map[string]string{},
		OutputDir:    "public",
		RequireAll:   true,
	}

	// 序列化为 JSON
	content, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 确保目录存在
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, content, 0644)
}

// AddSatisRepository 向 Satis 配置中添加仓库
func (c *Composer) AddSatisRepository(configPath string, type_ string, url string) error {
	// 读取现有配置
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SatisConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return err
	}

	// 添加仓库
	repo := map[string]string{
		"type": type_,
		"url":  url,
	}

	config.Repositories = append(config.Repositories, repo)

	// 序列化为 JSON
	updatedContent, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, updatedContent, 0644)
}

// BuildSatis 使用指定的配置文件构建 Satis 仓库
func (c *Composer) BuildSatis(configPath string, outputDir string) (string, error) {
	if outputDir == "" {
		return c.Run("satis", "build", configPath)
	}
	return c.Run("satis", "build", configPath, outputDir)
}

// InitSatis 初始化一个新的 Satis 仓库
func (c *Composer) InitSatis(name string, homepage string, dir string) error {
	if dir == "" {
		dir = "satis"
	}

	// 创建目录
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建配置文件
	configPath := filepath.Join(dir, "satis.json")
	return c.CreateSatisConfig(configPath, name, homepage)
}

// UpdateSatisStability 更新 Satis 配置中的最小稳定性
func (c *Composer) UpdateSatisStability(configPath string, stability string) error {
	// 验证稳定性值
	validStabilities := map[string]bool{
		"dev":    true,
		"alpha":  true,
		"beta":   true,
		"RC":     true,
		"stable": true,
	}

	if !validStabilities[stability] {
		return fmt.Errorf("invalid stability: %s", stability)
	}

	// 读取现有配置
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SatisConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return err
	}

	// 更新稳定性
	config.MinimumStability = stability

	// 序列化为 JSON
	updatedContent, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, updatedContent, 0644)
}

// EnableSatisArchive 启用 Satis 的归档功能
func (c *Composer) EnableSatisArchive(configPath string, format string) error {
	// 读取现有配置
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SatisConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return err
	}

	// 设置归档选项
	if config.Archive == nil {
		config.Archive = make(map[string]interface{})
	}

	config.Archive["directory"] = "dist"

	if format != "" {
		config.Archive["format"] = format
	} else {
		config.Archive["format"] = "zip"
	}

	config.Archive["skip-dev"] = false

	// 序列化为 JSON
	updatedContent, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, updatedContent, 0644)
}

// AddSatisRequire 向 Satis 配置中添加依赖
func (c *Composer) AddSatisRequire(configPath string, packageName string, version string) error {
	// 读取现有配置
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var config SatisConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return err
	}

	// 设置依赖，并关闭 require-all
	if config.Require == nil {
		config.Require = make(map[string]string)
	}

	config.Require[packageName] = version
	config.RequireAll = false

	// 序列化为 JSON
	updatedContent, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configPath, updatedContent, 0644)
}

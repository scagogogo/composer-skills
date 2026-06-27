package composer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// AuthConfig 表示 Composer 认证配置
type AuthConfig struct {
	GitHub       map[string]string `json:"github-oauth,omitempty"`
	GitLab       map[string]string `json:"gitlab-oauth,omitempty"`
	GitLabToken  map[string]string `json:"gitlab-token,omitempty"`
	Bitbucket    map[string]string `json:"bitbucket-oauth,omitempty"`
	Bearer       map[string]string `json:"bearer,omitempty"`
	HTTPBasic    map[string]string `json:"http-basic,omitempty"`
	AWSAccessKey map[string]string `json:"aws-access-key,omitempty"`
}

// ErrInvalidAuthType 表示无效的认证类型错误
var ErrInvalidAuthType = errors.New("invalid authentication type")

// GetAuthConfig 获取当前认证配置
func (c *Composer) GetAuthConfig() (*AuthConfig, error) {
	// 获取 auth.json 文件路径
	homeDir, err := c.GetComposerHome()
	if err != nil {
		return nil, err
	}

	authFilePath := filepath.Join(homeDir, "auth.json")

	// 检查文件是否存在
	if _, err := os.Stat(authFilePath); os.IsNotExist(err) {
		// 文件不存在则返回空配置
		return &AuthConfig{}, nil
	}

	// 读取文件内容
	content, err := os.ReadFile(authFilePath)
	if err != nil {
		return nil, err
	}

	// 解析 JSON
	var config AuthConfig
	if err := json.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveAuthConfig 保存认证配置到 auth.json 文件
func (c *Composer) SaveAuthConfig(config *AuthConfig) error {
	// 获取 auth.json 文件路径
	homeDir, err := c.GetComposerHome()
	if err != nil {
		return err
	}

	authFilePath := filepath.Join(homeDir, "auth.json")

	// 序列化为 JSON
	content, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	// 确保目录存在
	if err := os.MkdirAll(homeDir, 0755); err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(authFilePath, content, 0600)
}

// AddGitHubToken 添加 GitHub OAuth 令牌
func (c *Composer) AddGitHubToken(domain string, token string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	if config.GitHub == nil {
		config.GitHub = make(map[string]string)
	}

	config.GitHub[domain] = token

	return c.SaveAuthConfig(config)
}

// AddGitLabToken 添加 GitLab OAuth 令牌
func (c *Composer) AddGitLabToken(domain string, token string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	if config.GitLab == nil {
		config.GitLab = make(map[string]string)
	}

	config.GitLab[domain] = token

	return c.SaveAuthConfig(config)
}

// AddBitbucketToken 添加 Bitbucket OAuth 令牌
func (c *Composer) AddBitbucketToken(domain string, consumer string, token string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	if config.Bitbucket == nil {
		config.Bitbucket = make(map[string]string)
	}

	// Bitbucket 使用 consumer:token 格式
	config.Bitbucket[domain] = fmt.Sprintf("%s:%s", consumer, token)

	return c.SaveAuthConfig(config)
}

// AddBearerToken 添加 Bearer 令牌
func (c *Composer) AddBearerToken(domain string, token string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	if config.Bearer == nil {
		config.Bearer = make(map[string]string)
	}

	config.Bearer[domain] = token

	return c.SaveAuthConfig(config)
}

// AddHTTPBasicAuth 添加 HTTP Basic 认证
func (c *Composer) AddHTTPBasicAuth(domain string, username string, password string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	if config.HTTPBasic == nil {
		config.HTTPBasic = make(map[string]string)
	}

	// HTTP Basic 使用 username:password 格式
	config.HTTPBasic[domain] = fmt.Sprintf("%s:%s", username, password)

	return c.SaveAuthConfig(config)
}

// RemoveToken 移除指定类型和域名的令牌
func (c *Composer) RemoveToken(authType string, domain string) error {
	config, err := c.GetAuthConfig()
	if err != nil {
		return err
	}

	switch authType {
	case "github-oauth":
		if config.GitHub != nil {
			delete(config.GitHub, domain)
		}
	case "gitlab-oauth":
		if config.GitLab != nil {
			delete(config.GitLab, domain)
		}
	case "bitbucket-oauth":
		if config.Bitbucket != nil {
			delete(config.Bitbucket, domain)
		}
	case "bearer":
		if config.Bearer != nil {
			delete(config.Bearer, domain)
		}
	case "http-basic":
		if config.HTTPBasic != nil {
			delete(config.HTTPBasic, domain)
		}
	default:
		return ErrInvalidAuthType
	}

	return c.SaveAuthConfig(config)
}

// GetToken 获取指定类型和域名的令牌
func (c *Composer) GetToken(authType string, domain string) (string, error) {
	config, err := c.GetAuthConfig()
	if err != nil {
		return "", err
	}

	switch authType {
	case "github-oauth":
		if config.GitHub != nil {
			return config.GitHub[domain], nil
		}
	case "gitlab-oauth":
		if config.GitLab != nil {
			return config.GitLab[domain], nil
		}
	case "bitbucket-oauth":
		if config.Bitbucket != nil {
			return config.Bitbucket[domain], nil
		}
	case "bearer":
		if config.Bearer != nil {
			return config.Bearer[domain], nil
		}
	case "http-basic":
		if config.HTTPBasic != nil {
			return config.HTTPBasic[domain], nil
		}
	default:
		return "", ErrInvalidAuthType
	}

	return "", nil
}

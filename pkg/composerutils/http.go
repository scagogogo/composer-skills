// Package utils 提供文件系统和HTTP相关的实用工具函数
package composerutils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// ErrDownloadFailed 表示下载失败的错误
var ErrDownloadFailed = errors.New("下载失败")

// DownloadConfig 定义下载配置选项
//
// 该结构体包含配置下载行为的选项，如是否使用代理及代理URL。
// 可以根据需要扩展此结构体，添加超时、重试次数等选项。
type DownloadConfig struct {
	// UseProxy 指示是否使用代理
	UseProxy bool
	// ProxyURL 指定代理服务器URL，仅当UseProxy为true时有效
	ProxyURL string
	// TimeoutSeconds 指定下载超时时间（秒），默认为60秒
	TimeoutSeconds int
}

// DownloadFile 从指定URL下载文件到目标路径
//
// 支持HTTP和HTTPS协议，可选择是否通过代理下载。
// 如果目标文件已存在，将被覆盖。
//
// 参数:
//   - sourceURL: 要下载文件的URL地址
//   - destPath: 下载文件存储的本地路径
//   - config: 下载配置，包含代理设置等
//
// 返回值:
//   - error: 如果下载过程中出现错误则返回，成功则返回nil
//
// 使用示例:
//
//	// 直接下载
//	config := utils.DownloadConfig{UseProxy: false}
//	err := utils.DownloadFile("https://example.com/file.zip", "/tmp/file.zip", config)
//
//	// 使用代理下载
//	proxyConfig := utils.DownloadConfig{
//	    UseProxy: true,
//	    ProxyURL: "http://proxy.example.com:8080",
//	    TimeoutSeconds: 120, // 设置更长的超时时间
//	}
//	err := utils.DownloadFile("https://example.com/large-file.tar.gz", "/downloads/file.tar.gz", proxyConfig)
//
// 可能的错误:
//   - URL格式无效
//   - 网络连接失败
//   - HTTP响应状态码非200
//   - 本地文件创建或写入失败
//   - 超时
func DownloadFile(sourceURL, destPath string, config DownloadConfig) error {
	// 创建HTTP客户端
	client := &http.Client{}

	// 设置超时
	timeout := 60 // 默认60秒
	if config.TimeoutSeconds > 0 {
		timeout = config.TimeoutSeconds
	}
	client.Timeout = time.Duration(timeout) * time.Second

	// 如果需要使用代理
	if config.UseProxy && config.ProxyURL != "" {
		proxyURL, err := url.Parse(config.ProxyURL)
		if err != nil {
			return fmt.Errorf("%w: 代理URL格式无效: %v", ErrDownloadFailed, err)
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	// 发起HTTP请求
	resp, err := client.Get(sourceURL)
	if err != nil {
		return fmt.Errorf("%w: 请求失败: %v", ErrDownloadFailed, err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: 服务器返回状态码 %d", ErrDownloadFailed, resp.StatusCode)
	}

	// 创建目标文件
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("%w: 无法创建目标文件: %v", ErrDownloadFailed, err)
	}
	defer out.Close()

	// 将响应内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("%w: 写入文件失败: %v", ErrDownloadFailed, err)
	}

	return nil
}

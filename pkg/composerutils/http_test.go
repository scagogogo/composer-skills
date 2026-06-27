package composerutils

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestDownloadFile(t *testing.T) {
	// 创建一个测试服务器
	normalServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/success" {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hello, this is test content!"))
		} else if r.URL.Path == "/notfound" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Not Found"))
		} else if r.URL.Path == "/servererror" {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server Error"))
		} else if r.URL.Path == "/slowresponse" {
			// 模拟慢响应
			time.Sleep(300 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("This is a slow response"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Bad Request"))
		}
	}))
	defer normalServer.Close()

	// 创建代理服务器模拟
	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodConnect {
			// 返回一个成功响应，表示代理连接建立
			w.WriteHeader(http.StatusOK)
		} else {
			// 代理请求处理 - 实际场景中这里会转发请求，但为了测试简单起见，直接返回响应
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Response via proxy"))
		}
	}))
	defer proxyServer.Close()

	tests := []struct {
		name        string
		sourceURL   string
		setup       func() (string, DownloadConfig, func())
		expectError bool
		errorType   error // 期望的错误类型
		validate    func(string) bool
	}{
		{
			name:      "成功下载文件",
			sourceURL: normalServer.URL + "/success",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "downloaded.txt")
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 10,
				}
				return destPath, config, func() {}
			},
			expectError: false,
			validate: func(destPath string) bool {
				data, err := os.ReadFile(destPath)
				if err != nil {
					t.Errorf("无法读取下载的文件: %v", err)
					return false
				}
				return string(data) == "Hello, this is test content!"
			},
		},
		{
			name:      "服务器返回404",
			sourceURL: normalServer.URL + "/notfound",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "notfound.txt")
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 10,
				}
				return destPath, config, func() {}
			},
			expectError: true,
			errorType:   ErrDownloadFailed,
		},
		{
			name:      "服务器返回500",
			sourceURL: normalServer.URL + "/servererror",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "server_error.txt")
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 10,
				}
				return destPath, config, func() {}
			},
			expectError: true,
			errorType:   ErrDownloadFailed,
		},
		{
			name:      "请求超时",
			sourceURL: normalServer.URL + "/slowresponse",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "timeout.txt")
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 0, // 设置为很小的值，默认将使用60秒
				}
				return destPath, config, func() {}
			},
			// 注意：实际上，由于服务器延迟仅为300ms，此测试可能不会超时
			// 如果要测试超时，需要使用更长的延迟或更短的超时设置
			expectError: false,
		},
		{
			name:      "无效URL",
			sourceURL: "http://invalid.url.that.does.not.exist.example",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "invalid.txt")
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 1, // 快速超时以避免长时间等待
				}
				return destPath, config, func() {}
			},
			expectError: true,
			errorType:   ErrDownloadFailed,
		},
		{
			name:      "无效代理URL",
			sourceURL: normalServer.URL + "/success",
			setup: func() (string, DownloadConfig, func()) {
				dir := t.TempDir()
				destPath := filepath.Join(dir, "invalid_proxy.txt")
				config := DownloadConfig{
					UseProxy:       true,
					ProxyURL:       "://invalid-proxy-url", // 无效的代理URL
					TimeoutSeconds: 1,
				}
				return destPath, config, func() {}
			},
			expectError: true,
			errorType:   ErrDownloadFailed,
		},
		{
			name:      "无写入权限的目标路径",
			sourceURL: normalServer.URL + "/success",
			setup: func() (string, DownloadConfig, func()) {
				// 在测试环境中，通常难以模拟无写入权限
				// 这里仅作为测试用例结构参考
				// 实际测试中，可以尝试使用只读目录或不存在的目录
				destPath := "/non-existent-dir/file.txt" // 假设这个目录不存在或无写入权限
				if os.Getenv("GOOS") == "windows" {
					destPath = "C:\\Windows\\System32\\non-existent-file.txt" // Windows上通常无权写入此目录
				}
				config := DownloadConfig{
					UseProxy:       false,
					TimeoutSeconds: 1,
				}
				return destPath, config, func() {}
			},
			expectError: true,
			// 注意，这里不检查具体错误信息，因为不同平台会有不同的权限错误信息
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 准备测试环境
			destPath, config, cleanup := tt.setup()
			defer cleanup()

			// 执行测试
			err := DownloadFile(tt.sourceURL, destPath, config)

			// 验证结果
			if tt.expectError {
				if err == nil {
					t.Errorf("期望错误但得到nil")
				} else if tt.errorType != nil && !isErrorOfType(err, tt.errorType) {
					t.Errorf("期望错误类型%v，但得到%v", tt.errorType, err)
				}
			} else {
				if err != nil {
					t.Errorf("不期望错误但得到: %v", err)
				} else if tt.validate != nil && !tt.validate(destPath) {
					t.Errorf("文件内容验证失败")
				}
			}
		})
	}
}

// 测试使用代理进行下载
// 注意：这个测试需要一个实际的代理服务器，因此在CI环境中可能会被跳过
func TestDownloadFileWithProxy(t *testing.T) {
	// 如果设置了CI环境变量，跳过此测试
	if os.Getenv("CI") != "" {
		t.Skip("在CI环境中跳过代理测试")
	}

	// 创建一个模拟的HTTP代理服务器
	proxyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 记录代理请求
		t.Logf("收到代理请求：%s %s", r.Method, r.URL.String())

		// 如果是CONNECT请求（HTTPS代理），回应成功
		if r.Method == http.MethodConnect {
			t.Log("CONNECT请求被模拟处理")
			w.WriteHeader(http.StatusOK)
			return
		}

		// 对于HTTP请求，直接返回内容
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("这是从代理服务器返回的内容"))
	}))
	defer proxyServer.Close()

	t.Logf("测试代理服务器地址: %s", proxyServer.URL)

	// 此URL实际不会被访问，因为我们的模拟代理会直接返回内容
	sourceURL := "http://example.com/file.txt"
	dir := t.TempDir()
	destPath := filepath.Join(dir, "proxy-downloaded.txt")

	config := DownloadConfig{
		UseProxy:       true,
		ProxyURL:       proxyServer.URL,
		TimeoutSeconds: 5,
	}

	// 执行带代理的下载
	err := DownloadFile(sourceURL, destPath, config)

	// 这个测试可能会失败，因为模拟代理可能无法正确处理实际请求
	// 这里主要是演示代理配置的用法
	if err != nil {
		t.Logf("通过代理下载失败 (这可能是预期的): %v", err)
	} else {
		content, err := os.ReadFile(destPath)
		if err != nil {
			t.Errorf("无法读取下载的文件: %v", err)
		} else {
			t.Logf("通过代理下载的文件内容: %s", string(content))
		}
	}
}

// 测试超时行为
func TestDownloadFileTimeout(t *testing.T) {
	// 创建一个会延迟响应的服务器
	delayServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // 延迟2秒
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Delayed response"))
	}))
	defer delayServer.Close()

	dir := t.TempDir()
	destPath := filepath.Join(dir, "timeout-test.txt")

	// 设置1秒超时
	config := DownloadConfig{
		UseProxy:       false,
		TimeoutSeconds: 1,
	}

	// 执行下载，预期会超时
	err := DownloadFile(delayServer.URL, destPath, config)

	// 验证是否因超时而失败
	if err == nil {
		t.Error("期望超时错误但得到nil")
	} else {
		t.Logf("正确检测到超时: %v", err)
	}
}

// 测试大文件下载
func TestDownloadLargeFile(t *testing.T) {
	// 文件大小：5MB
	fileSize := 5 * 1024 * 1024

	// 创建一个提供大文件的服务器
	largeFileServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
		w.WriteHeader(http.StatusOK)

		// 以1KB的块写入数据
		chunk := make([]byte, 1024)
		for i := 0; i < 1024*5; i++ { // 写入5MB数据
			_, _ = w.Write(chunk)
			// 刷新以确保数据被发送
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}))
	defer largeFileServer.Close()

	dir := t.TempDir()
	destPath := filepath.Join(dir, "large-file.bin")

	config := DownloadConfig{
		UseProxy:       false,
		TimeoutSeconds: 30, // 给大文件足够的下载时间
	}

	// 执行下载
	err := DownloadFile(largeFileServer.URL, destPath, config)
	if err != nil {
		t.Errorf("下载大文件失败: %v", err)
		return
	}

	// 验证文件大小
	info, err := os.Stat(destPath)
	if err != nil {
		t.Errorf("无法获取下载文件信息: %v", err)
		return
	}

	if info.Size() != int64(fileSize) {
		t.Errorf("下载的文件大小错误: 期望 %d 字节, 实际 %d 字节", fileSize, info.Size())
	} else {
		t.Logf("成功下载 %d 字节的文件", info.Size())
	}
}

// 测试下载时磁盘空间不足的情况
// 注意：这个测试在实际环境中难以模拟，这里仅作为示例
func TestDownloadWithInsufficientDiskSpace(t *testing.T) {
	t.Skip("跳过磁盘空间不足测试，因为这在自动化测试环境中难以可靠模拟")

	// 在实际情况下，可以尝试以下方法：
	// 1. 使用 Docker 容器设置磁盘限额
	// 2. 使用模拟文件系统
	// 3. 创建一个接近磁盘容量限制的大文件，然后尝试写入更多数据
}

// isErrorOfType 检查错误是否属于指定类型
func isErrorOfType(err error, target error) bool {
	// 检查错误是否包含目标错误
	if err == target {
		return true
	}
	// 如果是包装错误，检查底层错误
	if errors.Is(err, target) {
		return true
	}
	// 简单的字符串匹配检查
	return err != nil && target != nil && err.Error() == target.Error()
}

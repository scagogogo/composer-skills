package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClientIntegration 测试客户端集成功能
func TestClientIntegration(t *testing.T) {
	t.Run("package fetch integration", func(t *testing.T) {
		// 创建模拟服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/packages/symfony/console.json") {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"name": "symfony/console",
					"description": "Eases the creation of beautiful and testable command line interfaces",
					"time": "2023-01-01T00:00:00Z",
					"maintainers": [],
					"versions": {},
					"type": "library",
					"downloads": {
						"total": 1000000,
						"monthly": 50000,
						"daily": 2000
					}
				}`))
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer server.Close()

		// 创建使用测试服务器的客户端
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		// 测试获取包信息
		data, err := testClient.GetPackage("symfony/console")
		require.NoError(t, err)

		assert.Equal(t, "symfony/console", data.Package.Name)
		assert.Contains(t, data.Package.Description, "command line interfaces")
		assert.Equal(t, 1000000, data.Package.Downloads.Total)

		// 测试 JSON 序列化
		jsonData, err := json.MarshalIndent(data, "", "  ")
		require.NoError(t, err)
		assert.Contains(t, string(jsonData), "symfony/console")
	})

	t.Run("package not found", func(t *testing.T) {
		// 创建模拟服务器，返回 404
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status": "error", "message": "Package not found"}`))
		}))
		defer server.Close()

		// 创建使用测试服务器的客户端
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		_, err := testClient.GetPackage("nonexistent/package")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected status code: 404")
	})
}

// TestStatisticsIntegration 测试统计信息集成功能
func TestStatisticsIntegration(t *testing.T) {
	t.Run("successful stats fetch", func(t *testing.T) {
		// 创建模拟服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/statistics.json") {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"totals": {
						"downloads": 25000000000,
						"packages": 400000,
						"versions": 3000000
					}
				}`))
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer server.Close()

		// 创建使用测试服务器的客户端
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		data, err := testClient.GetStatistics()
		require.NoError(t, err)

		assert.Equal(t, int64(25000000000), data.Totals.Downloads)
		assert.Equal(t, 400000, data.Totals.Packages)
		assert.Equal(t, 3000000, data.Totals.Versions)

		// 测试 JSON 序列化
		jsonData, err := json.MarshalIndent(data, "", "  ")
		require.NoError(t, err)
		assert.Contains(t, string(jsonData), "downloads")
		assert.Contains(t, string(jsonData), "packages")
		assert.Contains(t, string(jsonData), "versions")
	})
}

// TestAdvisoriesIntegration 测试安全公告集成功能
func TestAdvisoriesIntegration(t *testing.T) {
	t.Run("successful advisories fetch", func(t *testing.T) {
		// 创建模拟服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/advisories.json") {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"advisories": {
						"symfony/http-foundation": [
							{
								"advisoryId": "PKSA-38s9-s9dj",
								"packageName": "symfony/http-foundation",
								"title": "HTTP Request Smuggling",
								"cve": "CVE-2022-24894"
							}
						]
					}
				}`))
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer server.Close()

		// 创建使用测试服务器的客户端
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		data, err := testClient.GetSecurityAdvisories()
		require.NoError(t, err)

		assert.Contains(t, data.Advisories, "symfony/http-foundation")
		assert.Len(t, data.Advisories["symfony/http-foundation"], 1)
		assert.Equal(t, "PKSA-38s9-s9dj", data.Advisories["symfony/http-foundation"][0].AdvisoryID)

		// 测试 JSON 序列化
		jsonData, err := json.MarshalIndent(data, "", "  ")
		require.NoError(t, err)
		assert.Contains(t, string(jsonData), "advisories")
		assert.Contains(t, string(jsonData), "symfony/http-foundation")
	})
}

// TestFileOutputIntegration 测试文件输出集成功能
func TestFileOutputIntegration(t *testing.T) {
	t.Run("output to file", func(t *testing.T) {
		// 创建临时文件
		tmpFile := "/tmp/test_output.json"
		defer os.Remove(tmpFile)

		// 创建模拟服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"totals": {
					"downloads": 1000000,
					"packages": 500,
					"versions": 2000
				}
			}`))
		}))
		defer server.Close()

		// 创建使用测试服务器的客户端
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		data, err := testClient.GetStatistics()
		require.NoError(t, err)

		// 写入文件
		jsonData, err := json.MarshalIndent(data, "", "  ")
		require.NoError(t, err)

		err = os.WriteFile(tmpFile, jsonData, 0644)
		require.NoError(t, err)

		// 验证文件是否创建并包含正确内容
		assert.FileExists(t, tmpFile)

		content, err := os.ReadFile(tmpFile)
		require.NoError(t, err)

		assert.Contains(t, string(content), "downloads")
		assert.Contains(t, string(content), "packages")
		assert.Contains(t, string(content), "versions")
	})
}

// TestClientConfiguration 测试客户端配置
func TestClientConfiguration(t *testing.T) {
	t.Run("timeout configuration", func(t *testing.T) {
		timeout := 10 * time.Second

		// 创建客户端并验证超时设置
		testClient := client.NewComposerClient(timeout)
		assert.NotNil(t, testClient)
	})

	t.Run("custom base URL configuration", func(t *testing.T) {
		customURL := "https://custom.packagist.org"
		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(customURL))
		assert.NotNil(t, testClient)
	})

	t.Run("API credentials configuration", func(t *testing.T) {
		username := "testuser"
		apiToken := "testtoken"
		testClient := client.NewComposerClient(30*time.Second,
			client.WithAPICredentials(username, apiToken))
		assert.NotNil(t, testClient)
	})
}

// TestMultipleOperationsIntegration 测试多个操作的集成
func TestMultipleOperationsIntegration(t *testing.T) {
	t.Run("multiple API calls", func(t *testing.T) {
		// 创建模拟服务器，支持多个端点
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch {
			case strings.Contains(r.URL.Path, "/statistics.json"):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"totals": {"downloads": 1000000, "packages": 500, "versions": 2000}}`))
			case strings.Contains(r.URL.Path, "/packages/list.json"):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"packageNames": ["symfony/console", "laravel/framework"]}`))
			case strings.Contains(r.URL.Path, "/search.json"):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"results": [{"name": "test/package"}], "total": 1}`))
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer server.Close()

		testClient := client.NewComposerClient(30*time.Second, client.WithBaseURL(server.URL))

		// 测试统计信息
		stats, err := testClient.GetStatistics()
		require.NoError(t, err)
		assert.Equal(t, int64(1000000), stats.Totals.Downloads)

		// 测试包列表
		packages, err := testClient.ListPackages()
		require.NoError(t, err)
		assert.Len(t, packages.PackageNames, 2)

		// 测试搜索
		searchResults, err := testClient.SearchPackages("test", 10, 1)
		require.NoError(t, err)
		assert.Equal(t, 1, searchResults.Total)
	})
}

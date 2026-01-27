package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/domain"
)

// BenchmarkComposerClient_GetPackage 基准测试获取包信息
func BenchmarkComposerClient_GetPackage(b *testing.B) {
	// Create mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetPackage("symfony/console")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_GetStatistics 基准测试获取统计信息
func BenchmarkComposerClient_GetStatistics(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"totals": {
				"downloads": 25000000000,
				"packages": 400000,
				"versions": 3000000
			}
		}`))
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetStatistics()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_SearchPackages 基准测试搜索包
func BenchmarkComposerClient_SearchPackages(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"results": [
				{
					"name": "symfony/console",
					"description": "Eases the creation of beautiful and testable command line interfaces"
				},
				{
					"name": "symfony/http-foundation",
					"description": "Defines an object-oriented layer for the HTTP specification"
				}
			],
			"total": 2
		}`))
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.SearchPackages("symfony", 10, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_ListPackages 基准测试获取包列表
func BenchmarkComposerClient_ListPackages(b *testing.B) {
	// Generate a large package list for realistic testing
	var packageNames []string
	for i := 0; i < 1000; i++ {
		packageNames = append(packageNames, fmt.Sprintf("vendor%d/package%d", i, i))
	}
	
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string][]string{"packageNames": packageNames}
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.ListPackages()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_GetSecurityAdvisories 基准测试获取安全公告
func BenchmarkComposerClient_GetSecurityAdvisories(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
				],
				"laravel/framework": [
					{
						"advisoryId": "PKSA-12ab-cd34",
						"packageName": "laravel/framework",
						"title": "SQL Injection",
						"cve": "CVE-2023-0001"
					}
				]
			}
		}`))
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetSecurityAdvisories()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_CreatePackage 基准测试创建包
func BenchmarkComposerClient_CreatePackage(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "success",
			"message": "Package created successfully"
		}`))
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, 
		WithBaseURL(server.URL),
		WithAPICredentials("testuser", "testtoken"))

	request := &domain.PackageCreateRequest{
		Repository: "https://github.com/test/package",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.CreatePackage(context.Background(), request)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkComposerClient_ConcurrentRequests 基准测试并发请求
func BenchmarkComposerClient_ConcurrentRequests(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate some processing time
		time.Sleep(1 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"name": "test/package",
			"description": "Test package",
			"downloads": {"total": 1000}
		}`))
	}))
	defer server.Close()

	client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.GetPackage("test/package")
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkJSON_Marshal 基准测试JSON序列化
func BenchmarkJSON_Marshal(b *testing.B) {
	packageInfo := &domain.ComposerPackageInfo{
		PackageName: "symfony/console",
		Package: domain.PackageInfo{
			Name:        "symfony/console",
			Description: "Eases the creation of beautiful and testable command line interfaces",
			Type:        "library",
			Downloads: domain.PackageDownloads{
				Total:   1000000,
				Monthly: 50000,
				Daily:   2000,
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(packageInfo)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkJSON_Unmarshal 基准测试JSON反序列化
func BenchmarkJSON_Unmarshal(b *testing.B) {
	jsonData := []byte(`{
		"packageName": "symfony/console",
		"package": {
			"name": "symfony/console",
			"description": "Eases the creation of beautiful and testable command line interfaces",
			"type": "library",
			"downloads": {
				"total": 1000000,
				"monthly": 50000,
				"daily": 2000
			}
		}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var packageInfo domain.ComposerPackageInfo
		err := json.Unmarshal(jsonData, &packageInfo)
		if err != nil {
			b.Fatal(err)
		}
	}
}

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/scagogogo/composer-skills/pkg/domain"
	"github.com/stretchr/testify/assert"
)

// TestNewComposerClient 测试客户端创建
func TestNewComposerClient(t *testing.T) {
	t.Run("default configuration", func(t *testing.T) {
		timeout := 30 * time.Second
		client := NewComposerClient(timeout)

		assert.NotNil(t, client)
		assert.Equal(t, "https://packagist.org", client.baseURL)
		assert.Equal(t, "https://repo.packagist.org", client.repoURL)
		assert.Equal(t, timeout, client.httpClient.Timeout)
		assert.Empty(t, client.username)
		assert.Empty(t, client.apiToken)
	})

	t.Run("with custom options", func(t *testing.T) {
		timeout := 10 * time.Second
		customBaseURL := "https://custom.packagist.org"
		customRepoURL := "https://custom.repo.packagist.org"
		username := "testuser"
		apiToken := "testtoken"

		client := NewComposerClient(timeout,
			WithBaseURL(customBaseURL),
			WithRepoURL(customRepoURL),
			WithAPICredentials(username, apiToken),
		)

		assert.NotNil(t, client)
		assert.Equal(t, customBaseURL, client.baseURL)
		assert.Equal(t, customRepoURL, client.repoURL)
		assert.Equal(t, timeout, client.httpClient.Timeout)
		assert.Equal(t, username, client.username)
		assert.Equal(t, apiToken, client.apiToken)
	})
}

// TestComposerClient_GetPackage 测试获取包信息
func TestComposerClient_GetPackage(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		// 创建模拟服务器
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/symfony/console.json", r.URL.Path)
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
		result, err := client.GetPackage("symfony/console")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "symfony/console", result.Package.Name)
		assert.Equal(t, "Eases the creation of beautiful and testable command line interfaces", result.Package.Description)
		assert.Equal(t, 1000000, result.Package.Downloads.Total)
	})

	t.Run("package not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status": "error", "message": "Package not found"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackage("nonexistent/package")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 404")
	})

	t.Run("invalid JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackage("symfony/console")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal package")
	})

	t.Run("network error", func(t *testing.T) {
		client := NewComposerClient(1*time.Millisecond, WithBaseURL("http://invalid-url-that-does-not-exist.test"))
		result, err := client.GetPackage("symfony/console")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to get package")
	})
}

// TestComposerClient_GetStatistics 测试获取统计信息
func TestComposerClient_GetStatistics(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/statistics.json", r.URL.Path)
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
		result, err := client.GetStatistics()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int64(25000000000), result.Totals.Downloads)
		assert.Equal(t, 400000, result.Totals.Packages)
		assert.Equal(t, 3000000, result.Totals.Versions)
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetStatistics()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 500")
	})
}

// TestComposerClient_GetSecurityAdvisories 测试获取安全公告
func TestComposerClient_GetSecurityAdvisories(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/advisories.json", r.URL.Path)
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
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisories()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Advisories, "symfony/http-foundation")
		assert.Len(t, result.Advisories["symfony/http-foundation"], 1)
		assert.Equal(t, "PKSA-38s9-s9dj", result.Advisories["symfony/http-foundation"][0].AdvisoryID)
	})

	t.Run("empty advisories response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisories()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Advisories)
	})

	t.Run("server timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond) // Simulate slow response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL)) // Short timeout
		_, err := client.GetSecurityAdvisories()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})

	t.Run("invalid JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisories()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})
}

// TestComposerClient_ListPackages 测试获取包列表
func TestComposerClient_ListPackages(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/list.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packageNames": [
					"symfony/console",
					"laravel/framework",
					"guzzlehttp/guzzle"
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackages()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 3)
		assert.Contains(t, result.PackageNames, "symfony/console")
		assert.Contains(t, result.PackageNames, "laravel/framework")
		assert.Contains(t, result.PackageNames, "guzzlehttp/guzzle")
	})

	t.Run("empty package list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackages()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("large package list", func(t *testing.T) {
		// Generate a large list of package names
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
		result, err := client.ListPackages()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 1000)
		assert.Equal(t, "vendor0/package0", result.PackageNames[0])
		assert.Equal(t, "vendor999/package999", result.PackageNames[999])
	})

	t.Run("malformed JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": [invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackages()

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

// TestComposerClient_SearchPackages 测试搜索包
func TestComposerClient_SearchPackages(t *testing.T) {
	t.Run("successful search", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/search.json", r.URL.Path)
			assert.Equal(t, "symfony", r.URL.Query().Get("q"))
			assert.Equal(t, "10", r.URL.Query().Get("per_page"))
			assert.Equal(t, "1", r.URL.Query().Get("page"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"results": [
					{
						"name": "symfony/console",
						"description": "Eases the creation of beautiful and testable command line interfaces"
					}
				],
				"total": 1
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.SearchPackages("symfony", 10, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Total)
		assert.Len(t, result.Results, 1)
		assert.Equal(t, "symfony/console", result.Results[0].Name)
	})

	t.Run("search with no parameters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/search.json", r.URL.Path)
			assert.Equal(t, "test", r.URL.Query().Get("q"))
			assert.Empty(t, r.URL.Query().Get("per_page"))
			assert.Empty(t, r.URL.Query().Get("page"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"results": [], "total": 0}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.SearchPackages("test", 0, 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.Total)
		assert.Len(t, result.Results, 0)
	})
}

// TestComposerClient_GetPackageWithV2Metadata 测试获取V2元数据
func TestComposerClient_GetPackageWithV2Metadata(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		expectedData := `{"packages": {"symfony/console": {"1.0.0": {}}}}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/p2/symfony/console.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedData))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithRepoURL(server.URL))
		result, err := client.GetPackageWithV2Metadata("symfony/console")

		assert.NoError(t, err)
		assert.Equal(t, expectedData, string(result))
	})

	t.Run("package not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithRepoURL(server.URL))
		result, err := client.GetPackageWithV2Metadata("nonexistent/package")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 404")
	})
}

// TestComposerClient_GetSecurityAdvisoriesForPackages 测试获取指定包的安全公告
func TestComposerClient_GetSecurityAdvisoriesForPackages(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/security-advisories/", r.URL.Path)
			packages := r.URL.Query()["packages[]"]
			assert.Contains(t, packages, "symfony/console")
			assert.Contains(t, packages, "laravel/framework")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"advisories": {
					"symfony/console": [
						{
							"advisoryId": "TEST-001",
							"packageName": "symfony/console",
							"title": "Test Advisory"
						}
					]
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{"symfony/console", "laravel/framework"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Advisories, "symfony/console")
	})

	t.Run("empty package list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			packages := r.URL.Query()["packages[]"]
			assert.Empty(t, packages)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Advisories)
	})

	t.Run("single package", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			packages := r.URL.Query()["packages[]"]
			assert.Len(t, packages, 1)
			assert.Equal(t, "symfony/console", packages[0])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"advisories": {
					"symfony/console": [
						{
							"advisoryId": "TEST-001",
							"packageName": "symfony/console",
							"title": "Test Advisory"
						}
					]
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{"symfony/console"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Advisories, "symfony/console")
		assert.Len(t, result.Advisories["symfony/console"], 1)
	})

	t.Run("large package list", func(t *testing.T) {
		// Generate a large list of package names
		var packages []string
		for i := 0; i < 100; i++ {
			packages = append(packages, fmt.Sprintf("vendor%d/package%d", i, i))
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedPackages := r.URL.Query()["packages[]"]
			assert.Len(t, receivedPackages, 100)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages(packages)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Advisories)
	})

	t.Run("packages with special characters", func(t *testing.T) {
		packages := []string{
			"vendor/package-with-dashes",
			"vendor/package_with_underscores",
			"vendor/package.with.dots",
			"vendor/package with spaces",
			"供应商/包名",
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedPackages := r.URL.Query()["packages[]"]
			assert.Len(t, receivedPackages, 5)
			for _, pkg := range packages {
				assert.Contains(t, receivedPackages, pkg)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages(packages)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Advisories)
	})

	t.Run("packages with empty strings", func(t *testing.T) {
		packages := []string{"", "valid/package", "", "another/package"}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedPackages := r.URL.Query()["packages[]"]
			// Empty strings should still be included in the request
			assert.Len(t, receivedPackages, 4)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages(packages)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Advisories)
	})

	t.Run("server returns HTTP error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid package names"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{"invalid/package"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 400")
	})

	t.Run("server returns malformed JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{"test/package"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("network timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesForPackages([]string{"test/package"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})
}

// TestComposerClient_GetSecurityAdvisoriesSince 测试获取指定时间后的安全公告
func TestComposerClient_GetSecurityAdvisoriesSince(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/security-advisories/", r.URL.Path)
			assert.Equal(t, "1672531200", r.URL.Query().Get("updatedSince"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesSince(testTime)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("different timestamp formats", func(t *testing.T) {
		testCases := []struct {
			name string
			time time.Time
		}{
			{
				name: "unix epoch",
				time: time.Unix(0, 0).UTC(),
			},
			{
				name: "year 2000",
				time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				name: "recent date",
				time: time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC),
			},
			{
				name: "future date",
				time: time.Date(2030, 12, 31, 23, 59, 59, 0, time.UTC),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				expectedUnix := fmt.Sprintf("%d", tc.time.Unix())

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedUnix, r.URL.Query().Get("updatedSince"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"advisories": {}}`))
				}))
				defer server.Close()

				client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
				result, err := client.GetSecurityAdvisoriesSince(tc.time)

				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Empty(t, result.Advisories)
			})
		}
	})

	t.Run("timezone handling", func(t *testing.T) {
		// Test different timezones - should all convert to UTC unix timestamp
		locations := []struct {
			name     string
			location *time.Location
		}{
			{"UTC", time.UTC},
			{"EST", time.FixedZone("EST", -5*3600)},
			{"JST", time.FixedZone("JST", 9*3600)},
		}

		baseTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
		expectedUnix := fmt.Sprintf("%d", baseTime.Unix())

		for _, loc := range locations {
			t.Run(loc.name, func(t *testing.T) {
				// Convert to different timezone but same moment
				testTime := baseTime.In(loc.location)

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedUnix, r.URL.Query().Get("updatedSince"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"advisories": {}}`))
				}))
				defer server.Close()

				client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
				result, err := client.GetSecurityAdvisoriesSince(testTime)

				assert.NoError(t, err)
				assert.NotNil(t, result)
			})
		}
	})

	t.Run("server returns advisories", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"advisories": {
					"symfony/console": [
						{
							"advisoryId": "RECENT-001",
							"packageName": "symfony/console",
							"title": "Recent Advisory",
							"cve": "CVE-2023-0001"
						}
					],
					"laravel/framework": [
						{
							"advisoryId": "RECENT-002",
							"packageName": "laravel/framework",
							"title": "Another Recent Advisory"
						}
					]
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesSince(testTime)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Advisories, 2)
		assert.Contains(t, result.Advisories, "symfony/console")
		assert.Contains(t, result.Advisories, "laravel/framework")
		assert.Len(t, result.Advisories["symfony/console"], 1)
		assert.Equal(t, "RECENT-001", result.Advisories["symfony/console"][0].AdvisoryID)
	})

	t.Run("server returns HTTP error", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid timestamp"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesSince(testTime)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 400")
	})

	t.Run("server returns malformed JSON", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {malformed`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesSince(testTime)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("network timeout", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"advisories": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL))
		result, err := client.GetSecurityAdvisoriesSince(testTime)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})
}

// TestComposerClient_ListPackagesByVendor 测试按供应商获取包列表
func TestComposerClient_ListPackagesByVendor(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/list.json", r.URL.Path)
			assert.Equal(t, "symfony", r.URL.Query().Get("vendor"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packageNames": [
					"symfony/console",
					"symfony/http-foundation"
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("symfony")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 2)
		assert.Contains(t, result.PackageNames, "symfony/console")
		assert.Contains(t, result.PackageNames, "symfony/http-foundation")
	})

	t.Run("empty vendor name", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "", r.URL.Query().Get("vendor"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("nonexistent vendor", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "nonexistent-vendor", r.URL.Query().Get("vendor"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("nonexistent-vendor")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("vendor with special characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "vendor-with-dashes", r.URL.Query().Get("vendor"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packageNames": [
					"vendor-with-dashes/package1",
					"vendor-with-dashes/package2"
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("vendor-with-dashes")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 2)
	})

	t.Run("vendor with URL encoding characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// The vendor name should be URL encoded in the query
			vendor := r.URL.Query().Get("vendor")
			assert.Equal(t, "vendor with spaces", vendor)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": ["vendor-space/package"]}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("vendor with spaces")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 1)
	})

	t.Run("vendor with unicode characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vendor := r.URL.Query().Get("vendor")
			assert.Equal(t, "供应商", vendor)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": ["供应商/包"]}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("供应商")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 1)
	})

	t.Run("very long vendor name", func(t *testing.T) {
		longVendor := strings.Repeat("a", 500)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vendor := r.URL.Query().Get("vendor")
			assert.Equal(t, longVendor, vendor)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor(longVendor)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("vendor with slashes", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vendor := r.URL.Query().Get("vendor")
			assert.Equal(t, "vendor/with/slashes", vendor)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("vendor/with/slashes")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("server returns malformed JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": [invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("test-vendor")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("server returns HTTP error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("test-vendor")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 500")
	})

	t.Run("network timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL))
		result, err := client.ListPackagesByVendor("test-vendor")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})
}

// TestComposerClient_ListPackagesByType 测试按类型获取包列表
func TestComposerClient_ListPackagesByType(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/list.json", r.URL.Path)
			assert.Equal(t, "library", r.URL.Query().Get("type"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packageNames": [
					"symfony/console",
					"guzzlehttp/guzzle"
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("library")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 2)
	})

	t.Run("different package types", func(t *testing.T) {
		testCases := []struct {
			packageType string
			expected    []string
		}{
			{"library", []string{"vendor/library1", "vendor/library2"}},
			{"symfony-bundle", []string{"vendor/bundle1"}},
			{"wordpress-plugin", []string{"vendor/plugin1", "vendor/plugin2"}},
			{"composer-plugin", []string{"vendor/composer-plugin"}},
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("type_%s", tc.packageType), func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, tc.packageType, r.URL.Query().Get("type"))
					response := map[string][]string{"packageNames": tc.expected}
					jsonData, _ := json.Marshal(response)
					w.WriteHeader(http.StatusOK)
					w.Write(jsonData)
				}))
				defer server.Close()

				client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
				result, err := client.ListPackagesByType(tc.packageType)

				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expected, result.PackageNames)
			})
		}
	})

	t.Run("empty type parameter", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "", r.URL.Query().Get("type"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("invalid package type", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "invalid-type", r.URL.Query().Get("type"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("invalid-type")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("case sensitivity test", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			typeParam := r.URL.Query().Get("type")
			assert.Equal(t, "Library", typeParam) // Capital L
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": ["vendor/package"]}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("Library")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.PackageNames, 1)
	})

	t.Run("type with special characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			typeParam := r.URL.Query().Get("type")
			assert.Equal(t, "type-with-dashes", typeParam)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("type-with-dashes")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("numeric type", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			typeParam := r.URL.Query().Get("type")
			assert.Equal(t, "123", typeParam)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("123")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("very long type name", func(t *testing.T) {
		longType := strings.Repeat("type", 100)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			typeParam := r.URL.Query().Get("type")
			assert.Equal(t, longType, typeParam)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType(longType)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.PackageNames)
	})

	t.Run("server error response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid type parameter"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("invalid")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 400")
	})

	t.Run("malformed JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packageNames": [malformed`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesByType("library")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})
}

// TestComposerClient_CreatePackage 测试创建包
func TestComposerClient_CreatePackage(t *testing.T) {
	t.Run("successful creation with credentials", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/create-package", r.URL.Path)
			assert.Equal(t, "testuser", r.URL.Query().Get("username"))
			assert.Equal(t, "testtoken", r.URL.Query().Get("apiToken"))
			assert.Equal(t, "POST", r.Method)

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

		result, err := client.CreatePackage(context.Background(), request)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("missing credentials", func(t *testing.T) {
		client := NewComposerClient(30 * time.Second)
		request := &domain.PackageCreateRequest{}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "API credentials are required")
	})

	t.Run("invalid repository URL", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{
				"status": "error",
				"message": "Invalid repository URL"
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageCreateRequest{
			Repository: "invalid-url",
		}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Invalid repository URL")
	})

	t.Run("unauthorized access", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{
				"status": "error",
				"message": "Invalid credentials"
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("invalid", "credentials"))

		request := &domain.PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Invalid credentials")
	})

	t.Run("server internal error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"status": "error",
				"message": "Internal server error"
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "Internal server error")
	})

	t.Run("malformed JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("network timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status": "success"}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.CreatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})

	t.Run("context cancellation", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status": "success"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		request := &domain.PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.CreatePackage(ctx, request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "context canceled")
	})
}

// TestComposerClient_SearchPackagesByTags 测试按标签搜索包
func TestComposerClient_SearchPackagesByTags(t *testing.T) {
	t.Run("successful search", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/search.json", r.URL.Path)
			tags := r.URL.Query()["tags"]
			assert.Contains(t, tags, "framework")
			assert.Contains(t, tags, "php")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"results": [
					{
						"name": "laravel/framework",
						"description": "The Laravel Framework"
					}
				],
				"total": 1
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.SearchPackagesByTags([]string{"framework", "php"}, 10, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Total)
		assert.Len(t, result.Results, 1)
		assert.Equal(t, "laravel/framework", result.Results[0].Name)
	})
}

// TestComposerClient_SearchPackagesByType 测试按类型搜索包
func TestComposerClient_SearchPackagesByType(t *testing.T) {
	t.Run("successful search", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/search.json", r.URL.Path)
			assert.Equal(t, "test", r.URL.Query().Get("q"))
			assert.Equal(t, "library", r.URL.Query().Get("type"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"results": [
					{
						"name": "phpunit/phpunit",
						"description": "The PHP Unit Testing framework"
					}
				],
				"total": 1
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.SearchPackagesByType("test", "library", 10, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Total)
		assert.Len(t, result.Results, 1)
		assert.Equal(t, "phpunit/phpunit", result.Results[0].Name)
	})
}

// TestComposerClient_ListPopularPackages 测试获取流行包列表
func TestComposerClient_ListPopularPackages(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/explore/popular.json", r.URL.Path)
			assert.Equal(t, "50", r.URL.Query().Get("per_page"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packages": [
					{
						"name": "symfony/console",
						"description": "Eases the creation of beautiful and testable command line interfaces",
						"downloads": 1000000
					}
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(50)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Packages, 1)
		assert.Equal(t, "symfony/console", result.Packages[0].Name)
	})

	t.Run("different per_page values", func(t *testing.T) {
		testCases := []struct {
			perPage  int
			expected string
		}{
			{10, "10"},
			{25, "25"},
			{100, "100"},
			{1, "1"},
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("per_page_%d", tc.perPage), func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, tc.expected, r.URL.Query().Get("per_page"))
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{"packages": []}`))
				}))
				defer server.Close()

				client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
				result, err := client.ListPopularPackages(tc.perPage)

				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Empty(t, result.Packages)
			})
		}
	})

	t.Run("zero per_page parameter", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "0", r.URL.Query().Get("per_page"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("negative per_page parameter", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "-1", r.URL.Query().Get("per_page"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(-1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("very large per_page parameter", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "10000", r.URL.Query().Get("per_page"))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(10000)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("server returns empty packages", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(50)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("server returns multiple packages", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"packages": [
					{
						"name": "symfony/console",
						"description": "Console component",
						"downloads": 1000000
					},
					{
						"name": "guzzlehttp/guzzle",
						"description": "HTTP client",
						"downloads": 900000
					},
					{
						"name": "monolog/monolog",
						"description": "Logging library",
						"downloads": 800000
					}
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(3)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Packages, 3)
		assert.Equal(t, "symfony/console", result.Packages[0].Name)
		assert.Equal(t, "guzzlehttp/guzzle", result.Packages[1].Name)
		assert.Equal(t, "monolog/monolog", result.Packages[2].Name)
	})

	t.Run("server returns HTTP error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(50)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 500")
	})

	t.Run("server returns malformed JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": [invalid json`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(50)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("network timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"packages": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL))
		result, err := client.ListPopularPackages(50)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})
}

// TestComposerClient_GetPackageChanges 测试获取包变更信息
func TestComposerClient_GetPackageChanges(t *testing.T) {
	t.Run("successful request with timestamp", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/metadata/changes.json", r.URL.Path)
			assert.Equal(t, "1672531200", r.URL.Query().Get("since"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"actions": [
					{
						"type": "update",
						"package": "symfony/console",
						"time": 1672531200
					}
				]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageChanges(context.Background(), 1672531200)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Actions, 1)
		assert.Equal(t, "update", result.Actions[0].Type)
		assert.Equal(t, "symfony/console", result.Actions[0].Package)
	})

	t.Run("successful request without timestamp", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/metadata/changes.json", r.URL.Path)
			assert.Empty(t, r.URL.Query().Get("since"))

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"actions": []}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageChanges(context.Background(), 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Actions, 0)
	})
}

// TestComposerClient_GetPackageDevVersions 测试获取包开发版本
func TestComposerClient_GetPackageDevVersions(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		expectedData := `{"packages": {"symfony/console": {"dev-master": {}}}}`
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/p2/symfony/console~dev.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedData))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithRepoURL(server.URL))
		result, err := client.GetPackageDevVersions("symfony/console")

		assert.NoError(t, err)
		assert.Equal(t, expectedData, string(result))
	})

	t.Run("package not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithRepoURL(server.URL))
		result, err := client.GetPackageDevVersions("nonexistent/package")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 404")
	})
}

// TestComposerClient_ListPackagesWithData 测试获取带数据的包列表
func TestComposerClient_ListPackagesWithData(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/list.json", r.URL.Path)
			fields := r.URL.Query()["fields[]"]
			assert.Contains(t, fields, "repository")
			assert.Contains(t, fields, "type")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"package": {
					"symfony/console": {
						"repository": "https://github.com/symfony/console",
						"type": "library"
					}
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"repository", "type"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Packages, "symfony/console")
	})

	t.Run("empty fields list", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fields := r.URL.Query()["fields[]"]
			assert.Empty(t, fields)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("single field", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fields := r.URL.Query()["fields[]"]
			assert.Len(t, fields, 1)
			assert.Equal(t, "repository", fields[0])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"package": {
					"test/package": {
						"repository": "https://github.com/test/package"
					}
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"repository"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Packages, "test/package")
	})

	t.Run("multiple fields", func(t *testing.T) {
		fields := []string{"repository", "type", "description", "keywords", "homepage"}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedFields := r.URL.Query()["fields[]"]
			assert.Len(t, receivedFields, len(fields))
			for _, field := range fields {
				assert.Contains(t, receivedFields, field)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"package": {
					"test/package": {
						"repository": "https://github.com/test/package",
						"type": "library",
						"description": "Test package",
						"keywords": ["test", "example"],
						"homepage": "https://example.com"
					}
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData(fields)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Contains(t, result.Packages, "test/package")
	})

	t.Run("fields with special characters", func(t *testing.T) {
		fields := []string{"field-with-dashes", "field_with_underscores", "field.with.dots"}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedFields := r.URL.Query()["fields[]"]
			for _, field := range fields {
				assert.Contains(t, receivedFields, field)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData(fields)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("fields with empty strings", func(t *testing.T) {
		fields := []string{"", "valid-field", "", "another-field"}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedFields := r.URL.Query()["fields[]"]
			// Empty strings should still be included in the request
			assert.Len(t, receivedFields, 4)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData(fields)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("large number of fields", func(t *testing.T) {
		// Generate a large list of field names
		var fields []string
		for i := 0; i < 50; i++ {
			fields = append(fields, fmt.Sprintf("field%d", i))
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			receivedFields := r.URL.Query()["fields[]"]
			assert.Len(t, receivedFields, 50)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData(fields)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Packages)
	})

	t.Run("server returns multiple packages", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"package": {
					"symfony/console": {
						"repository": "https://github.com/symfony/console",
						"type": "library"
					},
					"laravel/framework": {
						"repository": "https://github.com/laravel/framework",
						"type": "library"
					},
					"guzzlehttp/guzzle": {
						"repository": "https://github.com/guzzle/guzzle",
						"type": "library"
					}
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"repository", "type"})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Packages, 3)
		assert.Contains(t, result.Packages, "symfony/console")
		assert.Contains(t, result.Packages, "laravel/framework")
		assert.Contains(t, result.Packages, "guzzlehttp/guzzle")
	})

	t.Run("server returns HTTP error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid fields"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"invalid-field"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 400")
	})

	t.Run("server returns malformed JSON", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {malformed`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"repository"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to unmarshal")
	})

	t.Run("network timeout", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"package": {}}`))
		}))
		defer server.Close()

		client := NewComposerClient(50*time.Millisecond, WithBaseURL(server.URL))
		result, err := client.ListPackagesWithData([]string{"repository"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "deadline exceeded")
	})
}

// TestComposerClient_GetPackageStats 测试获取包统计信息
func TestComposerClient_GetPackageStats(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/symfony/console/stats.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"downloads": {
					"total": 1000000,
					"monthly": 50000,
					"daily": 2000
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageStats("symfony/console")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1000000, result.Downloads.Total)
	})

	t.Run("package with zero downloads", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/test/newpackage/stats.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"downloads": {
					"total": 0,
					"monthly": 0,
					"daily": 0
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageStats("test/newpackage")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 0, result.Downloads.Total)
		assert.Equal(t, 0, result.Downloads.Monthly)
		assert.Equal(t, 0, result.Downloads.Daily)
	})

	t.Run("invalid package name", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status": "error", "message": "Package not found"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageStats("invalid/package")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "unexpected status code: 404")
	})

	t.Run("malformed stats response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"downloads": "invalid"}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageStats("symfony/console")

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("package name with special characters", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/packages/vendor-name/package-name/stats.json", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"downloads": {
					"total": 5000,
					"monthly": 500,
					"daily": 20
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))
		result, err := client.GetPackageStats("vendor-name/package-name")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 5000, result.Downloads.Total)
	})
}

// TestComposerClient_EditPackage 测试编辑包
func TestComposerClient_EditPackage(t *testing.T) {
	t.Run("successful edit with credentials", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/packages/test/package", r.URL.Path)
			assert.Equal(t, "testuser", r.URL.Query().Get("username"))
			assert.Equal(t, "testtoken", r.URL.Query().Get("apiToken"))
			assert.Equal(t, "PUT", r.Method)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success"
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageEditRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.EditPackage(context.Background(), "test/package", request)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "success", result.Status)
	})

	t.Run("missing credentials", func(t *testing.T) {
		client := NewComposerClient(30 * time.Second)
		request := &domain.PackageEditRequest{}

		result, err := client.EditPackage(context.Background(), "test/package", request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "API credentials are required")
	})
}

// TestComposerClient_UpdatePackage 测试更新包
func TestComposerClient_UpdatePackage(t *testing.T) {
	t.Run("successful update with credentials", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "/api/update-package", r.URL.Path)
			assert.Equal(t, "testuser", r.URL.Query().Get("username"))
			assert.Equal(t, "testtoken", r.URL.Query().Get("apiToken"))
			assert.Equal(t, "POST", r.Method)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success",
				"jobs": ["job1", "job2"]
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second,
			WithBaseURL(server.URL),
			WithAPICredentials("testuser", "testtoken"))

		request := &domain.PackageUpdateRequest{
			Repository: "https://github.com/test/package",
		}

		result, err := client.UpdatePackage(context.Background(), request)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "success", result.Status)
		assert.Len(t, result.Jobs, 2)
	})

	t.Run("missing credentials", func(t *testing.T) {
		client := NewComposerClient(30 * time.Second)
		request := &domain.PackageUpdateRequest{}

		result, err := client.UpdatePackage(context.Background(), request)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "API credentials are required")
	})
}

// TestComposerClient_ConcurrentAccess 测试并发访问安全性
func TestComposerClient_ConcurrentAccess(t *testing.T) {
	t.Run("concurrent package requests", func(t *testing.T) {
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			// Simulate some processing time
			time.Sleep(10 * time.Millisecond)

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"name": "test/package",
				"description": "Test package",
				"downloads": {
					"total": 1000,
					"monthly": 100,
					"daily": 10
				}
			}`))
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

		// Launch multiple concurrent requests
		const numGoroutines = 10
		results := make(chan error, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				packageName := fmt.Sprintf("test/package%d", id)
				_, err := client.GetPackage(packageName)
				results <- err
			}(i)
		}

		// Collect results
		for i := 0; i < numGoroutines; i++ {
			err := <-results
			assert.NoError(t, err, "Goroutine %d failed", i)
		}

		// Verify all requests were processed
		assert.Equal(t, numGoroutines, requestCount)
	})

	t.Run("concurrent different operations", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/packages/test/package.json":
				w.Write([]byte(`{"name": "test/package", "downloads": {"total": 1000}}`))
			case "/statistics.json":
				w.Write([]byte(`{"totals": {"downloads": 1000000, "packages": 500}}`))
			case "/packages/list.json":
				w.Write([]byte(`{"packageNames": ["test/package1", "test/package2"]}`))
			case "/search.json":
				w.Write([]byte(`{"results": [{"name": "test/package"}], "total": 1}`))
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		}))
		defer server.Close()

		client := NewComposerClient(30*time.Second, WithBaseURL(server.URL))

		// Launch different types of concurrent requests
		const numOperations = 4
		results := make(chan error, numOperations)

		// GetPackage
		go func() {
			_, err := client.GetPackage("test/package")
			results <- err
		}()

		// GetStatistics
		go func() {
			_, err := client.GetStatistics()
			results <- err
		}()

		// ListPackages
		go func() {
			_, err := client.ListPackages()
			results <- err
		}()

		// SearchPackages
		go func() {
			_, err := client.SearchPackages("test", 10, 1)
			results <- err
		}()

		// Collect results
		for i := 0; i < numOperations; i++ {
			err := <-results
			assert.NoError(t, err, "Operation %d failed", i)
		}
	})
}

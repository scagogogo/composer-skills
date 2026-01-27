package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStatisticsResponse_Serialization 测试统计响应的序列化/反序列化
func TestStatisticsResponse_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &StatisticsResponse{
			Totals: Totals{
				Downloads: 25000000000,
				Packages:  400000,
				Versions:  3000000,
			},
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result StatisticsResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Totals.Downloads, result.Totals.Downloads)
		assert.Equal(t, original.Totals.Packages, result.Totals.Packages)
		assert.Equal(t, original.Totals.Versions, result.Totals.Versions)
	})

	t.Run("unmarshal from JSON string", func(t *testing.T) {
		jsonStr := `{
			"totals": {
				"downloads": 25000000000,
				"packages": 400000,
				"versions": 3000000
			}
		}`

		var result StatisticsResponse
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Equal(t, int64(25000000000), result.Totals.Downloads)
		assert.Equal(t, 400000, result.Totals.Packages)
		assert.Equal(t, 3000000, result.Totals.Versions)
	})

	t.Run("handle missing fields", func(t *testing.T) {
		jsonStr := `{"totals": {}}`

		var result StatisticsResponse
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Equal(t, int64(0), result.Totals.Downloads)
		assert.Equal(t, 0, result.Totals.Packages)
		assert.Equal(t, 0, result.Totals.Versions)
	})
}

// TestAdvisoriesResponse_Serialization 测试安全公告响应的序列化/反序列化
func TestAdvisoriesResponse_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &AdvisoriesResponse{
			Advisories: map[string][]*Advisory{
				"symfony/http-foundation": {
					{
						AdvisoryID:       "PKSA-38s9-s9dj",
						PackageName:      "symfony/http-foundation",
						RemoteID:         "CVE-2022-24894",
						Title:            "HTTP Request Smuggling",
						Link:             "https://github.com/advisories/GHSA-rc93-5vf2-xh7q",
						Cve:              "CVE-2022-24894",
						AffectedVersions: ">=5.4.0,<5.4.19|>=6.0.0,<6.0.4",
						Source:           "GitHub",
						ReportedAt:       "2022-03-10T12:00:00Z",
						Sources: []*Source{
							{
								Name:     "GitHub",
								RemoteID: "GHSA-rc93-5vf2-xh7q",
							},
						},
					},
				},
			},
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result AdvisoriesResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Contains(t, result.Advisories, "symfony/http-foundation")
		advisory := result.Advisories["symfony/http-foundation"][0]
		assert.Equal(t, "PKSA-38s9-s9dj", advisory.AdvisoryID)
		assert.Equal(t, "symfony/http-foundation", advisory.PackageName)
		assert.Equal(t, "CVE-2022-24894", advisory.Cve)
		assert.Len(t, advisory.Sources, 1)
		assert.Equal(t, "GitHub", advisory.Sources[0].Name)
	})

	t.Run("unmarshal from JSON string", func(t *testing.T) {
		jsonStr := `{
			"advisories": {
				"test/package": [
					{
						"advisoryId": "TEST-001",
						"packageName": "test/package",
						"title": "Test Advisory",
						"cve": "CVE-2023-0001",
						"sources": []
					}
				]
			}
		}`

		var result AdvisoriesResponse
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Contains(t, result.Advisories, "test/package")
		assert.Len(t, result.Advisories["test/package"], 1)
		assert.Equal(t, "TEST-001", result.Advisories["test/package"][0].AdvisoryID)
	})
}

// TestPackageInfo_Serialization 测试包信息的序列化/反序列化
func TestPackageInfo_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		original := &PackageInfo{
			Name:        "symfony/console",
			Description: "Eases the creation of beautiful and testable command line interfaces",
			Time:        testTime,
			Maintainers: []*Maintainer{
				{
					Name:      "Fabien Potencier",
					AvatarURL: "https://avatars.githubusercontent.com/u/47313",
				},
			},
			Versions: map[string]*Version{
				"v6.0.0": {
					Name:        "symfony/console",
					Version:     "v6.0.0",
					Description: "Eases the creation of beautiful and testable command line interfaces",
				},
			},
			Type:             "library",
			Repository:       "https://github.com/symfony/console",
			GithubStars:      1000,
			GithubWatchers:   100,
			GithubForks:      200,
			GithubOpenIssues: 50,
			Language:         "PHP",
			Dependents:       5000,
			Suggesters:       100,
			Downloads: PackageDownloads{
				Total:   1000000,
				Monthly: 50000,
				Daily:   2000,
			},
			Favers: 500,
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageInfo
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证基本字段
		assert.Equal(t, original.Name, result.Name)
		assert.Equal(t, original.Description, result.Description)
		assert.Equal(t, original.Type, result.Type)
		assert.Equal(t, original.GithubStars, result.GithubStars)
		assert.Equal(t, original.Downloads.Total, result.Downloads.Total)

		// 验证维护者
		assert.Len(t, result.Maintainers, 1)
		assert.Equal(t, "Fabien Potencier", result.Maintainers[0].Name)

		// 验证版本
		assert.Contains(t, result.Versions, "v6.0.0")
		assert.Equal(t, "v6.0.0", result.Versions["v6.0.0"].Version)
	})
}

// TestComposerPackageInfo_Serialization 测试完整包信息的序列化/反序列化
func TestComposerPackageInfo_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		now := time.Now()
		original := &ComposerPackageInfo{
			PackageName:          "symfony/console",
			PackageNameLowercase: "symfony/console",
			Package: PackageInfo{
				Name:        "symfony/console",
				Description: "Console component",
				Type:        "library",
				Downloads: PackageDownloads{
					Total:   1000000,
					Monthly: 50000,
					Daily:   2000,
				},
			},
			PackageInfoMd5: "abc123",
			CreateTime:     &now,
			UpdateTime:     &now,
			ChangeTime:     &now,
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result ComposerPackageInfo
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.PackageName, result.PackageName)
		assert.Equal(t, original.Package.Name, result.Package.Name)
		assert.Equal(t, original.Package.Downloads.Total, result.Package.Downloads.Total)
		assert.Equal(t, original.PackageInfoMd5, result.PackageInfoMd5)

		// 时间字段可能有精度差异，所以使用 Unix 时间戳比较
		assert.Equal(t, original.CreateTime.Unix(), result.CreateTime.Unix())
	})
}

// TestPackageListResponse_Serialization 测试包列表响应的序列化/反序列化
func TestPackageListResponse_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &PackageListResponse{
			PackageNames: []string{
				"symfony/console",
				"laravel/framework",
				"guzzlehttp/guzzle",
			},
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageListResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Len(t, result.PackageNames, 3)
		assert.Contains(t, result.PackageNames, "symfony/console")
		assert.Contains(t, result.PackageNames, "laravel/framework")
		assert.Contains(t, result.PackageNames, "guzzlehttp/guzzle")
	})
}

// TestPackageCreateRequest_Serialization 测试包创建请求的序列化/反序列化
func TestPackageCreateRequest_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &PackageCreateRequest{
			Repository: "https://github.com/test/package",
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageCreateRequest
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Repository, result.Repository)
	})

	t.Run("unmarshal from JSON string", func(t *testing.T) {
		jsonStr := `{"repository": "https://github.com/example/repo"}`

		var result PackageCreateRequest
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Equal(t, "https://github.com/example/repo", result.Repository)
	})
}

// TestPackageCreateResponse_Serialization 测试包创建响应的序列化/反序列化
func TestPackageCreateResponse_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &PackageCreateResponse{
			Status: "success",
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageCreateResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Status, result.Status)
	})
}

// TestPackageUpdateResponse_Serialization 测试包更新响应的序列化/反序列化
func TestPackageUpdateResponse_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal with jobs", func(t *testing.T) {
		original := &PackageUpdateResponse{
			Status: "success",
			Jobs:   []string{"job1", "job2", "job3"},
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageUpdateResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Status, result.Status)
		assert.Len(t, result.Jobs, 3)
		assert.Contains(t, result.Jobs, "job1")
		assert.Contains(t, result.Jobs, "job2")
		assert.Contains(t, result.Jobs, "job3")
	})

	t.Run("marshal and unmarshal without jobs", func(t *testing.T) {
		original := &PackageUpdateResponse{
			Status: "success",
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result PackageUpdateResponse
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Status, result.Status)
		assert.Nil(t, result.Jobs)
	})
}

// TestMaintainer_Serialization 测试维护者信息的序列化/反序列化
func TestMaintainer_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &Maintainer{
			Name:      "John Doe",
			AvatarURL: "https://avatars.githubusercontent.com/u/123456",
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result Maintainer
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Name, result.Name)
		assert.Equal(t, original.AvatarURL, result.AvatarURL)
	})

	t.Run("handle missing optional fields", func(t *testing.T) {
		jsonStr := `{"name": "Jane Doe"}`

		var result Maintainer
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Equal(t, "Jane Doe", result.Name)
		assert.Empty(t, result.AvatarURL)
	})
}

// TestVersion_Serialization 测试版本信息的序列化/反序列化
func TestVersion_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		testTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		original := &Version{
			Name:              "symfony/console",
			Version:           "v6.0.0",
			VersionNormalized: "6.0.0.0",
			Description:       "Console component",
			Keywords:          []string{"console", "cli"},
			Homepage:          "https://symfony.com",
			License:           []string{"MIT"},
			Type:              "library",
			Time:              testTime,
			Require: map[string]string{
				"php": ">=8.0",
			},
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result Version
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Name, result.Name)
		assert.Equal(t, original.Version, result.Version)
		assert.Equal(t, original.Description, result.Description)
		assert.Equal(t, original.Keywords, result.Keywords)
		assert.Equal(t, original.License, result.License)
		assert.Equal(t, original.Type, result.Type)
		assert.Equal(t, original.Time.Unix(), result.Time.Unix())
		assert.Equal(t, original.Require, result.Require)
	})

	t.Run("unmarshal from JSON string", func(t *testing.T) {
		jsonStr := `{
			"name": "test/package",
			"version": "1.0.0",
			"description": "Test package",
			"type": "library",
			"time": "2023-01-01T00:00:00Z"
		}`

		var result Version
		err := json.Unmarshal([]byte(jsonStr), &result)
		require.NoError(t, err)

		assert.Equal(t, "test/package", result.Name)
		assert.Equal(t, "1.0.0", result.Version)
		assert.Equal(t, "Test package", result.Description)
		assert.Equal(t, "library", result.Type)
	})
}

// TestSource_Serialization 测试安全公告来源的序列化/反序列化
func TestSource_Serialization(t *testing.T) {
	t.Run("marshal and unmarshal", func(t *testing.T) {
		original := &Source{
			Name:     "GitHub",
			RemoteID: "GHSA-rc93-5vf2-xh7q",
		}

		// 序列化
		data, err := json.Marshal(original)
		require.NoError(t, err)

		// 反序列化
		var result Source
		err = json.Unmarshal(data, &result)
		require.NoError(t, err)

		// 验证
		assert.Equal(t, original.Name, result.Name)
		assert.Equal(t, original.RemoteID, result.RemoteID)
	})
}

package domain

import (
	"encoding/json"
	"testing"
)

// FuzzPackageInfo 模糊测试包信息反序列化
func FuzzPackageInfo(f *testing.F) {
	// Add seed inputs
	f.Add(`{"name": "test/package", "description": "Test package"}`)
	f.Add(`{"name": "symfony/console", "type": "library", "downloads": {"total": 1000}}`)
	f.Add(`{}`)
	f.Add(`{"name": "", "description": ""}`)
	f.Add(`{"downloads": {"total": -1, "monthly": 0, "daily": 0}}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var packageInfo PackageInfo
		err := json.Unmarshal([]byte(jsonData), &packageInfo)
		// We don't expect errors to crash the program
		// Just ensure the unmarshaling doesn't panic
		_ = err
	})
}

// FuzzComposerPackageInfo 模糊测试完整包信息反序列化
func FuzzComposerPackageInfo(f *testing.F) {
	// Add seed inputs
	f.Add(`{"packageName": "test/package", "package": {"name": "test/package"}}`)
	f.Add(`{"package": {"downloads": {"total": 1000000}}}`)
	f.Add(`{"packageName": ""}`)
	f.Add(`{"package": null}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var packageInfo ComposerPackageInfo
		err := json.Unmarshal([]byte(jsonData), &packageInfo)
		_ = err
	})
}

// FuzzStatisticsResponse 模糊测试统计响应反序列化
func FuzzStatisticsResponse(f *testing.F) {
	// Add seed inputs
	f.Add(`{"totals": {"downloads": 1000, "packages": 500, "versions": 2000}}`)
	f.Add(`{"totals": {"downloads": -1}}`)
	f.Add(`{"totals": {}}`)
	f.Add(`{}`)
	f.Add(`{"totals": null}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var stats StatisticsResponse
		err := json.Unmarshal([]byte(jsonData), &stats)
		_ = err
	})
}

// FuzzAdvisoriesResponse 模糊测试安全公告响应反序列化
func FuzzAdvisoriesResponse(f *testing.F) {
	// Add seed inputs
	f.Add(`{"advisories": {"test/package": [{"advisoryId": "TEST-001"}]}}`)
	f.Add(`{"advisories": {}}`)
	f.Add(`{}`)
	f.Add(`{"advisories": null}`)
	f.Add(`{"advisories": {"": []}}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var advisories AdvisoriesResponse
		err := json.Unmarshal([]byte(jsonData), &advisories)
		_ = err
	})
}

// FuzzAdvisory 模糊测试单个安全公告反序列化
func FuzzAdvisory(f *testing.F) {
	// Add seed inputs
	f.Add(`{"advisoryId": "TEST-001", "packageName": "test/package", "title": "Test Advisory"}`)
	f.Add(`{"advisoryId": "", "packageName": "", "title": ""}`)
	f.Add(`{}`)
	f.Add(`{"cve": "CVE-2023-0001"}`)
	f.Add(`{"sources": [{"name": "GitHub", "remoteId": "GHSA-xxxx"}]}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var advisory Advisory
		err := json.Unmarshal([]byte(jsonData), &advisory)
		_ = err
	})
}

// FuzzPackageListResponse 模糊测试包列表响应反序列化
func FuzzPackageListResponse(f *testing.F) {
	// Add seed inputs
	f.Add(`{"packageNames": ["test/package1", "test/package2"]}`)
	f.Add(`{"packageNames": []}`)
	f.Add(`{}`)
	f.Add(`{"packageNames": null}`)
	f.Add(`{"packageNames": [""]}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var packageList PackageListResponse
		err := json.Unmarshal([]byte(jsonData), &packageList)
		_ = err
	})
}

// FuzzSearchResponse 模糊测试搜索响应反序列化
func FuzzSearchResponse(f *testing.F) {
	// Add seed inputs
	f.Add(`{"results": [{"name": "test/package"}], "total": 1}`)
	f.Add(`{"results": [], "total": 0}`)
	f.Add(`{}`)
	f.Add(`{"total": -1}`)
	f.Add(`{"results": null, "total": 0}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var searchResponse SearchResponse
		err := json.Unmarshal([]byte(jsonData), &searchResponse)
		_ = err
	})
}

// FuzzVersion 模糊测试版本信息反序列化
func FuzzVersion(f *testing.F) {
	// Add seed inputs
	f.Add(`{"name": "test/package", "version": "1.0.0", "type": "library"}`)
	f.Add(`{"version": "", "name": ""}`)
	f.Add(`{}`)
	f.Add(`{"time": "invalid-date"}`)
	f.Add(`{"require": {"php": ">=8.0"}}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var version Version
		err := json.Unmarshal([]byte(jsonData), &version)
		_ = err
	})
}

// FuzzMaintainer 模糊测试维护者信息反序列化
func FuzzMaintainer(f *testing.F) {
	// Add seed inputs
	f.Add(`{"name": "John Doe", "avatarUrl": "https://example.com/avatar.jpg"}`)
	f.Add(`{"name": "", "avatarUrl": ""}`)
	f.Add(`{}`)
	f.Add(`{"name": null}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var maintainer Maintainer
		err := json.Unmarshal([]byte(jsonData), &maintainer)
		_ = err
	})
}

// FuzzPackageCreateRequest 模糊测试包创建请求序列化/反序列化
func FuzzPackageCreateRequest(f *testing.F) {
	// Add seed inputs
	f.Add(`{"repository": "https://github.com/test/package"}`)
	f.Add(`{"repository": ""}`)
	f.Add(`{}`)
	f.Add(`{"repository": "invalid-url"}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var request PackageCreateRequest
		err := json.Unmarshal([]byte(jsonData), &request)
		if err == nil {
			// If unmarshaling succeeded, try marshaling back
			_, marshalErr := json.Marshal(request)
			_ = marshalErr
		}
	})
}

// FuzzPackageDownloads 模糊测试下载统计反序列化
func FuzzPackageDownloads(f *testing.F) {
	// Add seed inputs
	f.Add(`{"total": 1000000, "monthly": 50000, "daily": 2000}`)
	f.Add(`{"total": -1, "monthly": -1, "daily": -1}`)
	f.Add(`{}`)
	f.Add(`{"total": 0}`)
	f.Add(`{"total": 9223372036854775807}`) // Max int64
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var downloads PackageDownloads
		err := json.Unmarshal([]byte(jsonData), &downloads)
		_ = err
	})
}

// FuzzTotals 模糊测试总计统计反序列化
func FuzzTotals(f *testing.F) {
	// Add seed inputs
	f.Add(`{"downloads": 25000000000, "packages": 400000, "versions": 3000000}`)
	f.Add(`{"downloads": -1, "packages": -1, "versions": -1}`)
	f.Add(`{}`)
	f.Add(`{"downloads": 0, "packages": 0, "versions": 0}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var totals Totals
		err := json.Unmarshal([]byte(jsonData), &totals)
		_ = err
	})
}

// FuzzSource 模糊测试安全公告来源反序列化
func FuzzSource(f *testing.F) {
	// Add seed inputs
	f.Add(`{"name": "GitHub", "remoteId": "GHSA-xxxx-yyyy-zzzz"}`)
	f.Add(`{"name": "", "remoteId": ""}`)
	f.Add(`{}`)
	f.Add(`{"name": null, "remoteId": null}`)
	
	f.Fuzz(func(t *testing.T, jsonData string) {
		var source Source
		err := json.Unmarshal([]byte(jsonData), &source)
		_ = err
	})
}

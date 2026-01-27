package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/domain"
)

// ComposerClient 是 Composer API 的客户端
// 提供获取包信息、安全公告和统计数据的方法
type ComposerClient struct {
	httpClient *http.Client
	baseURL    string
	repoURL    string
	username   string
	apiToken   string
}

// ComposerClientOption 表示客户端的配置选项
type ComposerClientOption func(*ComposerClient)

// WithBaseURL 设置API基础URL
func WithBaseURL(baseURL string) ComposerClientOption {
	return func(c *ComposerClient) {
		c.baseURL = baseURL
	}
}

// WithRepoURL 设置仓库URL
func WithRepoURL(repoURL string) ComposerClientOption {
	return func(c *ComposerClient) {
		c.repoURL = repoURL
	}
}

// WithAPICredentials 设置API凭据
func WithAPICredentials(username, apiToken string) ComposerClientOption {
	return func(c *ComposerClient) {
		c.username = username
		c.apiToken = apiToken
	}
}

// NewComposerClient 创建一个新的 Composer API 客户端
func NewComposerClient(timeout time.Duration, options ...ComposerClientOption) *ComposerClient {
	client := &ComposerClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: "https://packagist.org",
		repoURL: "https://repo.packagist.org",
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// GetPackage 获取指定包的信息
func (c *ComposerClient) GetPackage(packageName string) (*domain.ComposerPackageInfo, error) {
	url := fmt.Sprintf("%s/packages/%s.json", c.baseURL, packageName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get package info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var packageInfo domain.PackageInfo
	if err := json.Unmarshal(body, &packageInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package info: %w", err)
	}

	// 创建完整的包信息对象
	now := time.Now()
	result := &domain.ComposerPackageInfo{
		PackageName:          packageName,
		PackageNameLowercase: strings.ToLower(packageName),
		Package:              packageInfo,
		CreateTime:           &now,
		UpdateTime:           &now,
	}

	return result, nil
}

// GetPackageWithV2Metadata 使用Composer V2元数据获取包信息
// 对应API: https://repo.packagist.org/p2/[vendor]/[package].json
func (c *ComposerClient) GetPackageWithV2Metadata(packageName string) ([]byte, error) {
	url := fmt.Sprintf("%s/p2/%s.json", c.repoURL, packageName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get v2 metadata: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// GetPackageDevVersions 获取包的开发版本信息
// 对应API: https://repo.packagist.org/p2/[vendor]/[package]~dev.json
func (c *ComposerClient) GetPackageDevVersions(packageName string) ([]byte, error) {
	url := fmt.Sprintf("%s/p2/%s~dev.json", c.repoURL, packageName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get dev versions: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// GetStatistics 获取 Composer 仓库的统计信息
func (c *ComposerClient) GetStatistics() (*domain.StatisticsResponse, error) {
	url := fmt.Sprintf("%s/statistics.json", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var stats domain.StatisticsResponse
	if err := json.Unmarshal(body, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal statistics: %w", err)
	}

	return &stats, nil
}

// GetSecurityAdvisories 获取安全公告信息
func (c *ComposerClient) GetSecurityAdvisories() (*domain.AdvisoriesResponse, error) {
	url := fmt.Sprintf("%s/advisories.json", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get advisories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var advisories domain.AdvisoriesResponse
	if err := json.Unmarshal(body, &advisories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal advisories: %w", err)
	}

	return &advisories, nil
}

// GetSecurityAdvisoriesForPackages 获取指定包的安全公告
// 对应API: https://packagist.org/api/security-advisories/?packages[]=vendor/package
func (c *ComposerClient) GetSecurityAdvisoriesForPackages(packageNames []string) (*domain.AdvisoriesResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/api/security-advisories/", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	for _, pkg := range packageNames {
		params.Add("packages[]", pkg)
	}

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get advisories for packages: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var advisories domain.AdvisoriesResponse
	if err := json.Unmarshal(body, &advisories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal advisories: %w", err)
	}

	return &advisories, nil
}

// GetSecurityAdvisoriesSince 获取自指定时间以来更新的安全公告
// 对应API: https://packagist.org/api/security-advisories/?updatedSince=timestamp
func (c *ComposerClient) GetSecurityAdvisoriesSince(updatedSince time.Time) (*domain.AdvisoriesResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/api/security-advisories/", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	params.Add("updatedSince", fmt.Sprintf("%d", updatedSince.Unix()))

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get advisories since %v: %w", updatedSince, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var advisories domain.AdvisoriesResponse
	if err := json.Unmarshal(body, &advisories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal advisories: %w", err)
	}

	return &advisories, nil
}

// ListPackages 获取所有包名列表
// 对应API: https://packagist.org/packages/list.json
func (c *ComposerClient) ListPackages() (*domain.PackageListResponse, error) {
	url := fmt.Sprintf("%s/packages/list.json", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to list packages: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PackageListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package list: %w", err)
	}

	return &response, nil
}

// ListPackagesByVendor 获取指定供应商的包名列表
// 对应API: https://packagist.org/packages/list.json?vendor=vendor
func (c *ComposerClient) ListPackagesByVendor(vendor string) (*domain.PackageListResponse, error) {
	url := fmt.Sprintf("%s/packages/list.json?vendor=%s", c.baseURL, url.QueryEscape(vendor))

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to list packages for vendor %s: %w", vendor, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PackageListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package list: %w", err)
	}

	return &response, nil
}

// ListPackagesByType 获取指定类型的包名列表
// 对应API: https://packagist.org/packages/list.json?type=type
func (c *ComposerClient) ListPackagesByType(packageType string) (*domain.PackageListResponse, error) {
	url := fmt.Sprintf("%s/packages/list.json?type=%s", c.baseURL, url.QueryEscape(packageType))

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to list packages of type %s: %w", packageType, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PackageListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package list: %w", err)
	}

	return &response, nil
}

// ListPackagesWithData 获取包列表，并附带额外数据
// 对应API: https://packagist.org/packages/list.json?fields[]=repository&fields[]=type
func (c *ComposerClient) ListPackagesWithData(fields []string) (*domain.PackageListWithDataResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/packages/list.json", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	for _, field := range fields {
		params.Add("fields[]", field)
	}

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to list packages with data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PackageListWithDataResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package list with data: %w", err)
	}

	return &response, nil
}

// ListPopularPackages 获取流行包列表
// 对应API: https://packagist.org/explore/popular.json?per_page=100
func (c *ComposerClient) ListPopularPackages(perPage int) (*domain.PopularPackagesResponse, error) {
	url := fmt.Sprintf("%s/explore/popular.json?per_page=%d", c.baseURL, perPage)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to list popular packages: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PopularPackagesResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal popular packages: %w", err)
	}

	return &response, nil
}

// SearchPackages 通过关键词搜索包
// 对应API: https://packagist.org/search.json?q=query
func (c *ComposerClient) SearchPackages(query string, perPage, page int) (*domain.SearchResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/search.json", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	params.Add("q", query)

	if perPage > 0 {
		params.Add("per_page", fmt.Sprintf("%d", perPage))
	}

	if page > 0 {
		params.Add("page", fmt.Sprintf("%d", page))
	}

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to search packages: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search results: %w", err)
	}

	return &response, nil
}

// SearchPackagesByTags 通过标签搜索包
// 对应API: https://packagist.org/search.json?tags=tag
func (c *ComposerClient) SearchPackagesByTags(tags []string, perPage, page int) (*domain.SearchResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/search.json", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	for _, tag := range tags {
		params.Add("tags", tag)
	}

	if perPage > 0 {
		params.Add("per_page", fmt.Sprintf("%d", perPage))
	}

	if page > 0 {
		params.Add("page", fmt.Sprintf("%d", page))
	}

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to search packages by tags: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search results: %w", err)
	}

	return &response, nil
}

// SearchPackagesByType 通过类型搜索包
// 对应API: https://packagist.org/search.json?q=query&type=type
func (c *ComposerClient) SearchPackagesByType(query, packageType string, perPage, page int) (*domain.SearchResponse, error) {
	// 构造URL
	baseURL := fmt.Sprintf("%s/search.json", c.baseURL)

	// 构造查询参数
	params := url.Values{}
	params.Add("q", query)
	params.Add("type", packageType)

	if perPage > 0 {
		params.Add("per_page", fmt.Sprintf("%d", perPage))
	}

	if page > 0 {
		params.Add("page", fmt.Sprintf("%d", page))
	}

	requestURL := baseURL + "?" + params.Encode()

	resp, err := c.httpClient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to search packages by type: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.SearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search results: %w", err)
	}

	return &response, nil
}

// GetPackageStats 获取包的下载统计
// 对应API: https://packagist.org/packages/[vendor]/[package]/stats.json
func (c *ComposerClient) GetPackageStats(packageName string) (*domain.PackageStatsResponse, error) {
	url := fmt.Sprintf("%s/packages/%s/stats.json", c.baseURL, packageName)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get package stats: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.PackageStatsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package stats: %w", err)
	}

	return &response, nil
}

// GetPackageChanges 获取包的变更信息
// 对应API: https://packagist.org/metadata/changes.json?since=timestamp
func (c *ComposerClient) GetPackageChanges(ctx context.Context, since int64) (*domain.ChangeTrackingResponse, error) {
	var url string
	if since > 0 {
		url = fmt.Sprintf("%s/metadata/changes.json?since=%d", c.baseURL, since)
	} else {
		url = fmt.Sprintf("%s/metadata/changes.json", c.baseURL)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get package changes: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response domain.ChangeTrackingResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package changes: %w", err)
	}

	return &response, nil
}

// CreatePackage 创建一个新的包
// 对应API: https://packagist.org/api/create-package?username=username&apiToken=token
func (c *ComposerClient) CreatePackage(ctx context.Context, request *domain.PackageCreateRequest) (*domain.PackageCreateResponse, error) {
	if c.username == "" || c.apiToken == "" {
		return nil, fmt.Errorf("API credentials are required for creating packages")
	}

	url := fmt.Sprintf("%s/api/create-package?username=%s&apiToken=%s", c.baseURL, url.QueryEscape(c.username), url.QueryEscape(c.apiToken))

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create package: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create package: %s", string(body))
	}

	var response domain.PackageCreateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// EditPackage 编辑一个已存在的包
// 对应API: https://packagist.org/api/packages/[package name]?username=username&apiToken=token
func (c *ComposerClient) EditPackage(ctx context.Context, packageName string, request *domain.PackageEditRequest) (*domain.PackageEditResponse, error) {
	if c.username == "" || c.apiToken == "" {
		return nil, fmt.Errorf("API credentials are required for editing packages")
	}

	url := fmt.Sprintf("%s/api/packages/%s?username=%s&apiToken=%s", c.baseURL, url.QueryEscape(packageName), url.QueryEscape(c.username), url.QueryEscape(c.apiToken))

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to edit package: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit package: %s", string(body))
	}

	var response domain.PackageEditResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

// UpdatePackage 更新一个包
// 对应API: https://packagist.org/api/update-package?username=username&apiToken=token
func (c *ComposerClient) UpdatePackage(ctx context.Context, request *domain.PackageUpdateRequest) (*domain.PackageUpdateResponse, error) {
	if c.username == "" || c.apiToken == "" {
		return nil, fmt.Errorf("API credentials are required for updating packages")
	}

	url := fmt.Sprintf("%s/api/update-package?username=%s&apiToken=%s", c.baseURL, url.QueryEscape(c.username), url.QueryEscape(c.apiToken))

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to update package: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update package: %s", string(body))
	}

	var response domain.PackageUpdateResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}

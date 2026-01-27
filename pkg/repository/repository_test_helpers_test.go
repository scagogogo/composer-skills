package repository

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCreateMockServer tests the createMockServer helper function
func TestCreateMockServer(t *testing.T) {
	t.Run("creates server with fixed response", func(t *testing.T) {
		expectedResponse := `{"test": "data"}`
		server := createMockServer(expectedResponse)
		defer server.Close()

		// Test that server returns expected response
		resp, err := http.Get(server.URL)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, expectedResponse, string(body))
	})

	t.Run("handles empty response", func(t *testing.T) {
		server := createMockServer("")
		defer server.Close()

		resp, err := http.Get(server.URL)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Empty(t, string(body))
	})

	t.Run("handles JSON response", func(t *testing.T) {
		jsonResponse := `{"packages": ["test/package1", "test/package2"]}`
		server := createMockServer(jsonResponse)
		defer server.Close()

		resp, err := http.Get(server.URL)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.JSONEq(t, jsonResponse, string(body))
	})
}

// TestCreateMockServerWithRoutes tests the createMockServerWithRoutes helper function
func TestCreateMockServerWithRoutes(t *testing.T) {
	t.Run("serves different responses for different routes", func(t *testing.T) {
		routes := map[string]string{
			"/api/packages":    `{"packages": ["test/package"]}`,
			"/api/statistics":  `{"totals": {"downloads": 1000}}`,
			"/api/advisories": `{"advisories": {}}`,
		}

		server := createMockServerWithRoutes(routes)
		defer server.Close()

		// Test each route
		for path, expectedResponse := range routes {
			resp, err := http.Get(server.URL + path)
			require.NoError(t, err, "Failed to get %s", path)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusOK, resp.StatusCode, "Wrong status for %s", path)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err, "Failed to read body for %s", path)
			assert.Equal(t, expectedResponse, string(body), "Wrong response for %s", path)
		}
	})

	t.Run("returns 404 for unknown routes", func(t *testing.T) {
		routes := map[string]string{
			"/known": "response",
		}

		server := createMockServerWithRoutes(routes)
		defer server.Close()

		// Test unknown route
		resp, err := http.Get(server.URL + "/unknown")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("handles empty routes map", func(t *testing.T) {
		routes := map[string]string{}
		server := createMockServerWithRoutes(routes)
		defer server.Close()

		// Any request should return 404
		resp, err := http.Get(server.URL + "/any")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

// TestNewTestRepository tests the newTestRepository helper function
func TestNewTestRepository(t *testing.T) {
	t.Run("creates repository with correct server URL", func(t *testing.T) {
		serverURL := "https://test.example.com"
		repo := newTestRepository(serverURL)

		assert.NotNil(t, repo)
		assert.NotNil(t, repo.options)
		assert.Equal(t, serverURL, repo.options.ServerUrl)
		assert.Empty(t, repo.options.Proxy) // Should not have proxy by default
	})

	t.Run("creates repository with localhost URL", func(t *testing.T) {
		serverURL := "http://localhost:8080"
		repo := newTestRepository(serverURL)

		assert.NotNil(t, repo)
		assert.Equal(t, serverURL, repo.options.ServerUrl)
	})

	t.Run("creates repository with empty URL", func(t *testing.T) {
		repo := newTestRepository("")

		assert.NotNil(t, repo)
		assert.Equal(t, "", repo.options.ServerUrl)
	})
}

// TestNewTestRepositoryWithProxy tests the newTestRepositoryWithProxy helper function
func TestNewTestRepositoryWithProxy(t *testing.T) {
	t.Run("creates repository with server URL and proxy", func(t *testing.T) {
		serverURL := "https://test.example.com"
		proxy := "http://proxy.example.com:8080"
		repo := newTestRepositoryWithProxy(serverURL, proxy)

		assert.NotNil(t, repo)
		assert.NotNil(t, repo.options)
		assert.Equal(t, serverURL, repo.options.ServerUrl)
		assert.Equal(t, proxy, repo.options.Proxy)
	})

	t.Run("creates repository with empty proxy", func(t *testing.T) {
		serverURL := "https://test.example.com"
		repo := newTestRepositoryWithProxy(serverURL, "")

		assert.NotNil(t, repo)
		assert.Equal(t, serverURL, repo.options.ServerUrl)
		assert.Empty(t, repo.options.Proxy)
	})

	t.Run("creates repository with SOCKS proxy", func(t *testing.T) {
		serverURL := "https://test.example.com"
		proxy := "socks5://127.0.0.1:1080"
		repo := newTestRepositoryWithProxy(serverURL, proxy)

		assert.NotNil(t, repo)
		assert.Equal(t, serverURL, repo.options.ServerUrl)
		assert.Equal(t, proxy, repo.options.Proxy)
	})
}

// TestTestHelperIntegration tests that the helper functions work together
func TestTestHelperIntegration(t *testing.T) {
	t.Run("repository can use mock server", func(t *testing.T) {
		// Create mock server with test data
		testResponse := `{"packageNames": ["test/package1", "test/package2"]}`
		server := createMockServer(testResponse)
		defer server.Close()

		// Create repository pointing to mock server
		repo := newTestRepository(server.URL)

		// Test that repository can make requests to mock server
		// (This would normally use the repository's methods, but we'll test the setup)
		assert.Equal(t, server.URL, repo.options.ServerUrl)
		assert.NotNil(t, repo.options)
	})

	t.Run("repository with proxy can use mock server", func(t *testing.T) {
		testResponse := `{"totals": {"downloads": 1000}}`
		server := createMockServer(testResponse)
		defer server.Close()

		proxy := "http://proxy.test:8080"
		repo := newTestRepositoryWithProxy(server.URL, proxy)

		assert.Equal(t, server.URL, repo.options.ServerUrl)
		assert.Equal(t, proxy, repo.options.Proxy)
	})
}

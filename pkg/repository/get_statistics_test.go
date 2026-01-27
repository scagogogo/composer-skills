package repository

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/scagogogo/composer-crawler/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Statistics(t *testing.T) {
	// Create a test server that returns a mock statistics response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/statistics.json" {
			w.WriteHeader(http.StatusOK)
			// Sample statistics response
			w.Write([]byte(`{
				"totals": {
					"downloads": 1000000,
					"packages": 500,
					"versions": 2000
				}
			}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create repository with test server URL
	repo := &Repository{
		options: &Options{
			ServerUrl: server.URL,
		},
	}

	// Test normal case
	t.Run("successful request", func(t *testing.T) {
		stats, err := repo.Statistics(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, stats)
		assert.Equal(t, int64(1000000), stats.Totals.Downloads)
		assert.Equal(t, 500, stats.Totals.Packages)
		assert.Equal(t, 2000, stats.Totals.Versions)
	})

	// Test failure case
	t.Run("invalid server URL", func(t *testing.T) {
		invalidRepo := &Repository{
			options: &Options{
				ServerUrl: "http://invalid-url-that-doesnt-exist.example",
			},
		}
		stats, err := invalidRepo.Statistics(context.Background())
		assert.Error(t, err)
		assert.Nil(t, stats)
	})

	// Test malformed response
	t.Run("malformed response", func(t *testing.T) {
		malformedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{malformed json}`))
		}))
		defer malformedServer.Close()

		malformedRepo := &Repository{
			options: &Options{
				ServerUrl: malformedServer.URL,
			},
		}
		stats, err := malformedRepo.Statistics(context.Background())
		assert.Error(t, err)
		assert.Nil(t, stats)
	})

	// Test empty response
	t.Run("empty response", func(t *testing.T) {
		emptyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}))
		defer emptyServer.Close()

		emptyRepo := &Repository{
			options: &Options{
				ServerUrl: emptyServer.URL,
			},
		}
		stats, err := emptyRepo.Statistics(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, stats)
		// Verify default values when fields are missing
		assert.Equal(t, domain.Totals{}, stats.Totals)
	})
}

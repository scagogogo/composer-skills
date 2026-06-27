package composer

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestFund(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund", "funding information", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Fund()

	if err != nil {
		t.Errorf("Fund failed: %v", err)
	}
	if output != "funding information" {
		t.Errorf("Expected 'funding information', got '%s'", output)
	}
}

func TestFundWithJSON(t *testing.T) {
	ClearMockOutputs()
	fundingData := []FundingInfo{
		{Name: "pkg1", URLs: []string{"url1"}, Funding: true},
		{Name: "pkg2", URLs: []string{"url2"}, Funding: false},
	}
	jsonData, _ := json.Marshal(fundingData)
	SetupMockOutput("fund --format=json", string(jsonData), nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	info, err := composer.FundWithJSON()

	if err != nil {
		t.Errorf("FundWithJSON failed: %v", err)
	}
	if len(info) != 2 {
		t.Errorf("Expected 2 funding info, got %d", len(info))
	}
}

func TestFundWithJSON_InvalidJSON(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund --format=json", "invalid json", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	_, err := composer.FundWithJSON()

	if err == nil {
		t.Error("FundWithJSON should fail with invalid JSON")
	}
}

func TestFundWithPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund vendor/package", "package funding info", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.FundWithPackage("vendor/package")

	if err != nil {
		t.Errorf("FundWithPackage failed: %v", err)
	}
	if output != "package funding info" {
		t.Errorf("Expected 'package funding info', got '%s'", output)
	}
}

func TestGetFundingURLs(t *testing.T) {
	ClearMockOutputs()
	fundingData := []FundingInfo{
		{Name: "pkg1", URLs: []string{"url1", "url2"}, Funding: true},
		{Name: "pkg2", URLs: []string{"url3"}, Funding: false},
	}
	jsonData, _ := json.Marshal(fundingData)
	SetupMockOutput("fund --format=json", string(jsonData), nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	urls, err := composer.GetFundingURLs()

	if err != nil {
		t.Errorf("GetFundingURLs failed: %v", err)
	}
	if len(urls) != 1 {
		t.Errorf("Expected 1 package with funding URLs, got %d", len(urls))
	}
	if _, ok := urls["pkg1"]; !ok {
		t.Error("pkg1 should be in URLs map")
	}
	if _, ok := urls["pkg2"]; ok {
		t.Error("pkg2 should not be in URLs map (Funding: false)")
	}
}

func TestHasFunding(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund --format=text", "No funding", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	has, err := composer.HasFunding()

	if err != nil {
		t.Errorf("HasFunding failed: %v", err)
	}
	if has {
		t.Error("HasFunding should return false for 'No funding'")
	}
}

func TestHasFunding_True(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund --format=text", "Funding available for packages", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	has, err := composer.HasFunding()

	if err != nil {
		t.Errorf("HasFunding true case failed: %v", err)
	}
	if !has {
		t.Error("HasFunding should return true when funding is available")
	}
}

func TestFundWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund --format=json", "fund output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"format": "json"}
	output, err := composer.FundWithOptions(options)

	if err != nil {
		t.Errorf("FundWithOptions failed: %v", err)
	}
	if output != "fund output" {
		t.Errorf("Expected 'fund output', got '%s'", output)
	}
}

func TestFund_Error(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund", "", errors.New("fund command failed"))

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	_, err := composer.Fund()

	if err == nil {
		t.Error("Fund should return error when command fails")
	}
}
package composer

import (
	"testing"
)

func TestCheckPlatform(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform --format=json", `{"platform":{"php":{"name":"php","version":"8.1.0","available":true}}}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.CheckPlatform()

	if err != nil {
		t.Errorf("CheckPlatform failed: %v", err)
	}
	if len(output) == 0 {
		t.Error("CheckPlatform should return non-empty output")
	}
}

func TestCheckPlatformWithLock(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform --lock --format=json", `{"lock":{"php":{"name":"php","version":"8.1.0","available":true}}}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.CheckPlatformWithLock()

	if err != nil {
		t.Errorf("CheckPlatformWithLock failed: %v", err)
	}
	if len(output) == 0 {
		t.Error("CheckPlatformWithLock should return non-empty output")
	}
}

func TestIsPlatformAvailable(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform --format=json", `{"platform":{"php":{"name":"php","version":"8.1.0","available":true}}}`, nil)
	SetupMockOutput("check-platform php:8.1.0", "php 8.1.0 is available", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	available, err := composer.IsPlatformAvailable("php", "8.1.0")

	if err != nil {
		t.Errorf("IsPlatformAvailable failed: %v", err)
	}
	if !available {
		t.Errorf("Expected platform to be available")
	}
}

func TestGetPHPVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run --php-show-version", "PHP 8.1.0\nphp.ini location: /etc/php.ini", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	version, err := composer.GetPHPVersion()

	if err != nil {
		t.Errorf("GetPHPVersion failed: %v", err)
	}
	if version != "8.1.0" {
		t.Errorf("Expected '8.1.0', got '%s'", version)
	}
}

func TestGetExtensions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run --show-extensions", "Loaded extensions:\npdo\nmbstring\njson", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	extensions, err := composer.GetExtensions()

	if err != nil {
		t.Errorf("GetExtensions failed: %v", err)
	}
	if len(extensions) != 3 {
		t.Errorf("Expected 3 extensions, got %d", len(extensions))
	}
}

func TestHasExtension(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run --show-extensions", "Loaded extensions:\npdo\nmbstring\njson", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	hasExt, err := composer.HasExtension("mbstring")

	if err != nil {
		t.Errorf("HasExtension failed: %v", err)
	}
	if !hasExt {
		t.Errorf("Expected mbstring extension to be present")
	}
}

func TestHasExtension_NotFound(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run --show-extensions", "Loaded extensions:\npdo\nmbstring\njson", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	hasExt, err := composer.HasExtension("curl")

	if err != nil {
		t.Errorf("HasExtension failed: %v", err)
	}
	if hasExt {
		t.Errorf("Expected curl extension to be absent")
	}
}

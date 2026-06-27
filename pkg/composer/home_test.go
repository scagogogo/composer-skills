package composer

import (
	"testing"
)

func TestHome(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home", "https://example.com", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Home()

	if err != nil {
		t.Errorf("Home failed: %v", err)
	}
	if output != "https://example.com" {
		t.Errorf("Expected 'https://example.com', got '%s'", output)
	}
}

func TestHomePackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home vendor/package", "https://github.com/vendor/package", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.HomePackage("vendor/package")

	if err != nil {
		t.Errorf("HomePackage failed: %v", err)
	}
	if output != "https://github.com/vendor/package" {
		t.Errorf("Expected 'https://github.com/vendor/package', got '%s'", output)
	}
}

func TestHomeWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home --show", "https://example.com", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"show": ""}
	output, err := composer.HomeWithOptions(options)

	if err != nil {
		t.Errorf("HomeWithOptions failed: %v", err)
	}
	if output != "https://example.com" {
		t.Errorf("Expected 'https://example.com', got '%s'", output)
	}
}

func TestHomeWithError(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home", "", ErrCommandExecution)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	_, err := composer.Home()

	if err == nil {
		t.Error("Home命令失败时应该返回错误")
	}
}

func TestHomePackageWithError(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home vendor/package", "", ErrCommandExecution)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	_, err := composer.HomePackage("vendor/package")

	if err == nil {
		t.Error("HomePackage命令失败时应该返回错误")
	}
}

func TestHomeWithOptionsWithError(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("home --show", "", ErrCommandExecution)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"show": ""}
	_, err := composer.HomeWithOptions(options)

	if err == nil {
		t.Error("HomeWithOptions命令失败时应该返回错误")
	}
}

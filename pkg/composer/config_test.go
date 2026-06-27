package composer

import (
	"testing"
)

func TestGetComposerHome(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global home", "/home/user/.composer", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	home, err := composer.GetComposerHome()

	if err != nil {
		t.Errorf("GetComposerHome failed: %v", err)
	}
	if home != "/home/user/.composer" {
		t.Errorf("Expected '/home/user/.composer', got '%s'", home)
	}
}

func TestClearCache(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("clear-cache", "Cache cleared", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.ClearCache()

	if err != nil {
		t.Errorf("ClearCache failed: %v", err)
	}
}

func TestGetConfigWithGlobal(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global preferred-install", "dist", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	value, err := composer.GetConfigWithGlobal("preferred-install", true)

	if err != nil {
		t.Errorf("GetConfigWithGlobal failed: %v", err)
	}
	if value != "dist" {
		t.Errorf("Expected 'dist', got '%s'", value)
	}
}

func TestGetConfigWithGlobal_NonGlobal(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config preferred-install", "source", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	value, err := composer.GetConfigWithGlobal("preferred-install", false)

	if err != nil {
		t.Errorf("GetConfigWithGlobal non-global failed: %v", err)
	}
	if value != "source" {
		t.Errorf("Expected 'source', got '%s'", value)
	}
}

func TestSetConfigWithGlobal(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global preferred-install dist", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetConfigWithGlobal("preferred-install", "dist", true)

	if err != nil {
		t.Errorf("SetConfigWithGlobal failed: %v", err)
	}
}

func TestSetConfigWithGlobal_NonGlobal(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config preferred-install source", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetConfigWithGlobal("preferred-install", "source", false)

	if err != nil {
		t.Errorf("SetConfigWithGlobal non-global failed: %v", err)
	}
}

func TestValidateComposerJson(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --strict", "valid", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.ValidateComposerJson(true, false)

	if err != nil {
		t.Errorf("ValidateComposerJson strict failed: %v", err)
	}
}

func TestValidateComposerJson_WithDependencies(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --strict --with-dependencies", "valid with deps", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.ValidateComposerJson(true, true)

	if err != nil {
		t.Errorf("ValidateComposerJson with deps failed: %v", err)
	}
}

func TestCheckPlatformReqs(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform-reqs", "all requirements satisfied", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.CheckPlatformReqs()

	if err != nil {
		t.Errorf("CheckPlatformReqs failed: %v", err)
	}
	if output != "all requirements satisfied" {
		t.Errorf("Expected 'all requirements satisfied', got '%s'", output)
	}
}

func TestListConfig(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --list", "preferred-install dist\nsort-packages true", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListConfig()

	if err != nil {
		t.Errorf("ListConfig failed: %v", err)
	}
	if !contains(output, "preferred-install") {
		t.Errorf("Expected 'preferred-install', got '%s'", output)
	}
}

func TestListConfigWithGlobal_Global(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --list --global", "home /home/user/.composer", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListConfigWithGlobal(true)

	if err != nil {
		t.Errorf("ListConfigWithGlobal global failed: %v", err)
	}
	if !contains(output, "home") {
		t.Errorf("Expected 'home', got '%s'", output)
	}
}

func TestListConfigWithGlobal_NonGlobal(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --list", "preferred-install dist", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListConfigWithGlobal(false)

	if err != nil {
		t.Errorf("ListConfigWithGlobal non-global failed: %v", err)
	}
	if !contains(output, "preferred-install") {
		t.Errorf("Expected 'preferred-install', got '%s'", output)
	}
}

func TestGetConfigSource(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config preferred-install --source", "From ./composer.json", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetConfigSource("preferred-install")

	if err != nil {
		t.Errorf("GetConfigSource failed: %v", err)
	}
	if !contains(output, "composer.json") {
		t.Errorf("Expected 'composer.json', got '%s'", output)
	}
}

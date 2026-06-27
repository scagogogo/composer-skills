package composer

import (
	"testing"
)

func TestReadComposerJSON(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{Name: "test/package"})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	composerJSON, err := composer.ReadComposerJSON()

	if err != nil {
		t.Errorf("ReadComposerJSON failed: %v", err)
	}
	if composerJSON.Name != "test/package" {
		t.Errorf("Expected 'test/package', got '%s'", composerJSON.Name)
	}

	ClearMockComposerJSON()
}

func TestWriteComposerJSON(t *testing.T) {
	ClearMockOutputs()

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	composerJSON := &ComposerJSON{Name: "written/package"}

	// This test just verifies the method doesn't panic
	err := composer.WriteComposerJSON(composerJSON)
	if err != nil {
		t.Errorf("WriteComposerJSON failed: %v", err)
	}
}

func TestAddRequire(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddRequire("vendor/package", "^1.0", false)

	if err != nil {
		t.Errorf("AddRequire failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestAddRequire_Dev(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddRequire("vendor/package", "^1.0", true)

	if err != nil {
		t.Errorf("AddRequire dev failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestRemoveRequire(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{Require: map[string]string{"vendor/package": "^1.0"}})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveRequire("vendor/package", false)

	if err != nil {
		t.Errorf("RemoveRequire failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestAddScript(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddScript("test-script", "php test.php", "Test script")

	if err != nil {
		t.Errorf("AddScript failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestAddScript_MultipleCommands(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	commands := []string{"php -r \"echo 1\"", "phpunit"}
	err := composer.AddScript("test", commands, "Run tests")

	if err != nil {
		t.Errorf("AddScript multiple commands failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestRemoveScript(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{Scripts: map[string]interface{}{"test-script": "php test.php"}})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveScript("test-script")

	if err != nil {
		t.Errorf("RemoveScript failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestAddAutoload(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddAutoload("psr-4", "App\\", "src/", false)

	if err != nil {
		t.Errorf("AddAutoload failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestAddAutoload_Dev(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddAutoload("psr-4", "Tests\\", "tests/", true)

	if err != nil {
		t.Errorf("AddAutoload dev failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestSetConfig(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetConfig("process-timeout", 500)

	if err != nil {
		t.Errorf("SetConfig failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestGetConfig(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{Config: map[string]interface{}{"process-timeout": 500}})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	value, err := composer.GetConfig("process-timeout")

	if err != nil {
		t.Errorf("GetConfig failed: %v", err)
	}
	if value != 500 {
		t.Errorf("Expected 500, got '%v'", value)
	}

	ClearMockComposerJSON()
}

func TestSetProperty(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetProperty("name", "vendor/package")

	if err != nil {
		t.Errorf("SetProperty failed: %v", err)
	}

	ClearMockComposerJSON()
}

func TestSetProperty_Unsupported(t *testing.T) {
	ClearMockOutputs()
	SetMockComposerJSON(&ComposerJSON{})

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetProperty("unsupported", "value")

	if err == nil {
		t.Error("SetProperty should fail for unsupported property")
	}

	ClearMockComposerJSON()
}

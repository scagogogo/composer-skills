package composer

import (
	"os"
	"testing"
)

func TestSetEnvVariable(t *testing.T) {
	testVar := EnvironmentVariable("COMPOSER_TEST_VAR")
	testValue := "test-value"

	err := SetEnvVariable(testVar, testValue)
	if err != nil {
		t.Errorf("SetEnvVariable failed: %v", err)
	}

	got := os.Getenv(string(testVar))
	if got != testValue {
		t.Errorf("Expected '%s', got '%s'", testValue, got)
	}

	os.Unsetenv(string(testVar))
}

func TestGetEnvVariable(t *testing.T) {
	testVar := EnvironmentVariable("COMPOSER_TEST_VAR_2")
	testValue := "test-value-2"

	os.Setenv(string(testVar), testValue)
	defer os.Unsetenv(string(testVar))

	got := GetEnvVariable(testVar)
	if got != testValue {
		t.Errorf("Expected '%s', got '%s'", testValue, got)
	}
}

func TestGetEnvVariable_NotSet(t *testing.T) {
	testVar := EnvironmentVariable("COMPOSER_NOT_SET_VAR")
	got := GetEnvVariable(testVar)
	if got != "" {
		t.Errorf("Expected empty string for unset variable, got '%s'", got)
	}
}

func TestSetProcessTimeout(t *testing.T) {
	err := SetProcessTimeout(300)
	if err != nil {
		t.Errorf("SetProcessTimeout failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerProcessTimeout))
	if got != "300" {
		t.Errorf("Expected '300', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerProcessTimeout))
}

func TestEnableSuperuser(t *testing.T) {
	err := EnableSuperuser()
	if err != nil {
		t.Errorf("EnableSuperuser failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerAllowSuperuser))
	if got != "1" {
		t.Errorf("Expected '1', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerAllowSuperuser))
}

func TestDisableSuperuser(t *testing.T) {
	os.Setenv(string(EnvComposerAllowSuperuser), "1")
	defer os.Unsetenv(string(EnvComposerAllowSuperuser))

	err := DisableSuperuser()
	if err != nil {
		t.Errorf("DisableSuperuser failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerAllowSuperuser))
	if got != "" {
		t.Errorf("Expected empty string after DisableSuperuser, got '%s'", got)
	}
}

func TestSetMemoryLimit(t *testing.T) {
	err := SetMemoryLimit("-1")
	if err != nil {
		t.Errorf("SetMemoryLimit failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerMemoryLimit))
	if got != "-1" {
		t.Errorf("Expected '-1', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerMemoryLimit))
}

func TestDisableInteraction(t *testing.T) {
	err := DisableInteraction()
	if err != nil {
		t.Errorf("DisableInteraction failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerNoInteraction))
	if got != "1" {
		t.Errorf("Expected '1', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerNoInteraction))
}

func TestEnableInteraction(t *testing.T) {
	os.Setenv(string(EnvComposerNoInteraction), "1")
	defer os.Unsetenv(string(EnvComposerNoInteraction))

	err := EnableInteraction()
	if err != nil {
		t.Errorf("EnableInteraction failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerNoInteraction))
	if got != "" {
		t.Errorf("Expected empty string after EnableInteraction, got '%s'", got)
	}
}

func TestSetVendorDir(t *testing.T) {
	err := SetVendorDir("/custom/vendor")
	if err != nil {
		t.Errorf("SetVendorDir failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerVendorDir))
	if got != "/custom/vendor" {
		t.Errorf("Expected '/custom/vendor', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerVendorDir))
}

func TestSetBinDir(t *testing.T) {
	err := SetBinDir("/custom/bin")
	if err != nil {
		t.Errorf("SetBinDir failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerBinDir))
	if got != "/custom/bin" {
		t.Errorf("Expected '/custom/bin', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerBinDir))
}

func TestSetCaFile(t *testing.T) {
	err := SetCaFile("/path/to/ca-bundle.crt")
	if err != nil {
		t.Errorf("SetCaFile failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerCafile))
	if got != "/path/to/ca-bundle.crt" {
		t.Errorf("Expected '/path/to/ca-bundle.crt', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerCafile))
}

func TestDisableDev(t *testing.T) {
	err := DisableDev()
	if err != nil {
		t.Errorf("DisableDev failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerNoDev))
	if got != "1" {
		t.Errorf("Expected '1', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerNoDev))
}

func TestEnableDev(t *testing.T) {
	os.Setenv(string(EnvComposerNoDev), "1")
	defer os.Unsetenv(string(EnvComposerNoDev))

	err := EnableDev()
	if err != nil {
		t.Errorf("EnableDev failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerNoDev))
	if got != "" {
		t.Errorf("Expected empty string after EnableDev, got '%s'", got)
	}
}

func TestSetDiscardChanges(t *testing.T) {
	err := SetDiscardChanges("stash")
	if err != nil {
		t.Errorf("SetDiscardChanges failed: %v", err)
	}

	got := os.Getenv(string(EnvComposerDiscardChanges))
	if got != "stash" {
		t.Errorf("Expected 'stash', got '%s'", got)
	}

	os.Unsetenv(string(EnvComposerDiscardChanges))
}

func TestGetEnvironmentInfo(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --list", "composer.composer.version: 2.0.0\ncomposer.platform.php: 8.1", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	info, err := composer.GetEnvironmentInfo()

	if err != nil {
		t.Errorf("GetEnvironmentInfo failed: %v", err)
	}
	if info["composer.composer.version"] != "2.0.0" {
		t.Errorf("Expected '2.0.0', got '%s'", info["composer.composer.version"])
	}
	if info["composer.platform.php"] != "8.1" {
		t.Errorf("Expected '8.1', got '%s'", info["composer.platform.php"])
	}
}

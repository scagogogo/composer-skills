package composer

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for validate command
	SetupMockOutput("validate", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Validate()
	if err != nil {
		t.Errorf("Validate执行失败: %v", err)
	}
}

func TestValidateWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("validate", "", errors.New("validation failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Validate()
	if err == nil {
		t.Error("验证失败时应该返回错误")
	}
}

func TestValidateStrict(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for strict validate command
	SetupMockOutput("validate --strict", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.ValidateStrict()
	if err != nil {
		t.Errorf("ValidateStrict执行失败: %v", err)
	}

	if !contains(output, "is valid") {
		t.Errorf("输出应包含验证结果，实际为\"%s\"", output)
	}
}

func TestValidateWithNoCheck(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for validate command with no-check-all
	SetupMockOutput("validate --no-check-all", "./composer.json is valid", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	output, err := composer.ValidateWithNoCheck()
	if err != nil {
		t.Errorf("ValidateWithNoCheck执行失败: %v", err)
	}

	if !contains(output, "is valid") {
		t.Errorf("输出应包含验证结果，实际为\"%s\"", output)
	}
}

func TestValidateWithNoCheckPublish(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --no-check-publish", "./composer.json is valid", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ValidateWithNoCheckPublish()

	if err != nil {
		t.Errorf("ValidateWithNoCheckPublish failed: %v", err)
	}
	if !contains(output, "is valid") {
		t.Errorf("Expected 'is valid', got '%s'", output)
	}
}

func TestValidateWithCheckVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --with-dependencies", "./composer.json is valid", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ValidateWithCheckVersion()

	if err != nil {
		t.Errorf("ValidateWithCheckVersion failed: %v", err)
	}
	if !contains(output, "is valid") {
		t.Errorf("Expected 'is valid', got '%s'", output)
	}
}

func TestCheckPlatformReqsLock(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform-reqs --lock", "platform requirements check passed", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.CheckPlatformReqsLock()

	if err != nil {
		t.Errorf("CheckPlatformReqsLock failed: %v", err)
	}
	if !contains(output, "platform requirements") {
		t.Errorf("Expected 'platform requirements', got '%s'", output)
	}
}

func TestCheckForOutdatedPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated --direct --minor-only --format json", "outdated packages in json", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.CheckForOutdatedPackages(true, true, "json")

	if err != nil {
		t.Errorf("CheckForOutdatedPackages failed: %v", err)
	}
	if !contains(output, "outdated") {
		t.Errorf("Expected 'outdated', got '%s'", output)
	}
}

func TestValidateSchema(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --no-check-all --no-check-publish --no-check-version", "schema valid", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ValidateSchema()

	if err != nil {
		t.Errorf("ValidateSchema failed: %v", err)
	}
	if !contains(output, "schema") {
		t.Errorf("Expected 'schema', got '%s'", output)
	}
}

func TestValidateWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --strict", "strict validation passed", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	options := map[string]string{"strict": ""}
	output, err := composer.ValidateWithOptions(options)

	if err != nil {
		t.Errorf("ValidateWithOptions failed: %v", err)
	}
	if !contains(output, "strict validation") {
		t.Errorf("Expected 'strict validation', got '%s'", output)
	}
}

func TestValidateQuiet(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --quiet", "", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ValidateQuiet()

	if err != nil {
		t.Errorf("ValidateQuiet failed: %v", err)
	}
	if output != "" {
		t.Errorf("Expected empty output, got '%s'", output)
	}
}

func TestCheckNormalization(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --no-check-all --check-normalized", "composer.json is normalized", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.CheckNormalization()

	if err != nil {
		t.Errorf("CheckNormalization failed: %v", err)
	}
	if !contains(output, "normalized") {
		t.Errorf("Expected 'normalized', got '%s'", output)
	}
}

func TestNormalizeComposerJson(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("normalize", "composer.json normalized", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.NormalizeComposerJson()

	if err != nil {
		t.Errorf("NormalizeComposerJson failed: %v", err)
	}
	if !contains(output, "normalized") {
		t.Errorf("Expected 'normalized', got '%s'", output)
	}
}

func TestCheckForSecurityVulnerabilities(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit", "No security vulnerability advisories found", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, hasVulns, err := composer.CheckForSecurityVulnerabilities()

	if err != nil {
		t.Errorf("CheckForSecurityVulnerabilities failed: %v", err)
	}
	if hasVulns {
		t.Errorf("Expected no vulnerabilities")
	}
	if !contains(output, "No security") {
		t.Errorf("Expected 'No security', got '%s'", output)
	}
}

func TestValidateComposerLock(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("validate --check-lock", "composer.lock is valid", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ValidateComposerLock()

	if err != nil {
		t.Errorf("ValidateComposerLock failed: %v", err)
	}
	if !contains(output, "valid") {
		t.Errorf("Expected 'valid', got '%s'", output)
	}
}

func TestProhibit(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("prohibit", "no prohibited packages", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.Prohibit()

	if err != nil {
		t.Errorf("Prohibit failed: %v", err)
	}
	if !contains(output, "prohibited") {
		t.Errorf("Expected 'prohibited', got '%s'", output)
	}
}

func TestProhibitWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("prohibit", `{"prohibited": []}`, nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	output, err := composer.ProhibitWithFormat("json")

	if err != nil {
		t.Errorf("ProhibitWithFormat failed: %v", err)
	}
	if !contains(output, "prohibited") {
		t.Errorf("Expected 'prohibited', got '%s'", output)
	}
}

func TestProhibitWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("prohibit", "prohibited packages with options", nil)

	composer, _ := New(Options{ExecutablePath: "/path/to/composer"})
	options := map[string]string{
		"format": "json",
		"fixed":  "",
	}
	output, err := composer.ProhibitWithOptions(options)

	if err != nil {
		t.Errorf("ProhibitWithOptions failed: %v", err)
	}
	if !contains(output, "prohibited") {
		t.Errorf("Expected 'prohibited', got '%s'", output)
	}
}

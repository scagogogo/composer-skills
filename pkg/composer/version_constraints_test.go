package composer

import (
	"testing"
)

func TestFormatVersionConstraint_ExactVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2.3", ExactVersion)
	if result != "1.2.3" {
		t.Errorf("Expected '1.2.3', got '%s'", result)
	}
}

func TestFormatVersionConstraint_CaretVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2.3", CaretVersion)
	if result != "^1.2.3" {
		t.Errorf("Expected '^1.2.3', got '%s'", result)
	}
}

func TestFormatVersionConstraint_TildeVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2.3", TildeVersion)
	if result != "~1.2.3" {
		t.Errorf("Expected '~1.2.3', got '%s'", result)
	}
}

func TestFormatVersionConstraint_WildcardVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2", WildcardVersion)
	if result != "1.2.*" {
		t.Errorf("Expected '1.2.*', got '%s'", result)
	}
}

func TestFormatVersionConstraint_RangeVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2", RangeVersion)
	// Should be >=1.2.0 <2.0.0
	expected := ">=1.2.0 <2.0.0"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFormatVersionConstraint_UnknownVersion(t *testing.T) {
	result := FormatVersionConstraint("1.2.3", "unknown")
	if result != "1.2.3" {
		t.Errorf("Expected '1.2.3' for unknown constraint, got '%s'", result)
	}
}

func TestIncrementMajorVersion(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"1", "2"},
		{"1.2", "2"},
		{"1.2.3", "2"},
		{"10.20.30", "11"},
	}

	for _, tc := range testCases {
		result := incrementMajorVersion(tc.input)
		if result != tc.expected {
			t.Errorf("incrementMajorVersion(%s): expected '%s', got '%s'", tc.input, tc.expected, result)
		}
	}
}

func TestIncrementMajorVersion_InvalidInput(t *testing.T) {
	// 测试无效输入
	result := incrementMajorVersion("")
	if result != "1" {
		t.Errorf("Expected '1' for empty string, got '%s'", result)
	}

	result = incrementMajorVersion("abc")
	if result != "1" {
		t.Errorf("Expected '1' for 'abc', got '%s'", result)
	}
}

func TestUpdatePackageVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require vendor/package:^1.0.0", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UpdatePackageVersion("vendor/package", "1.0.0", CaretVersion)

	if err != nil {
		t.Errorf("UpdatePackageVersion failed: %v", err)
	}
}

func TestLockPackageVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require vendor/package:1.0.0", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.LockPackageVersion("vendor/package", "1.0.0")

	if err != nil {
		t.Errorf("LockPackageVersion failed: %v", err)
	}
}

func TestGetPackageVersions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show --all vendor/package", "package versions output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetPackageVersions("vendor/package")

	if err != nil {
		t.Errorf("GetPackageVersions failed: %v", err)
	}
	if output != "package versions output" {
		t.Errorf("Expected 'package versions output', got '%s'", output)
	}
}

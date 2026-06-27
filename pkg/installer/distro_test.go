package installer

import (
	"os/exec"
	"runtime"
	"testing"
)

func TestDetectLinuxDistro(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping Linux distro detection test on non-Linux platform")
	}

	info, err := DetectLinuxDistro()
	if err != nil {
		t.Logf("DetectLinuxDistro returned error (may be expected): %v", err)
		return
	}

	if info == nil {
		t.Error("Expected non-nil DistroInfo")
		return
	}

	t.Logf("Detected distro: ID=%s, Name=%s, Version=%s, PackageManager=%s",
		info.ID, info.Name, info.Version, info.PackageManager)

	if info.ID == "" {
		t.Error("Expected non-empty distro ID")
	}
	if info.PackageManager == "" {
		t.Error("Expected non-empty package manager")
	}
}

func TestDeterminePkgManager(t *testing.T) {
	tests := []struct {
		distroID string
		expected string
	}{
		{"ubuntu", "apt"},
		{"debian", "apt"},
		{"linuxmint", "apt"},
		{"pop", "apt"},
		{"fedora", "dnf"},
		{"centos", "yum"}, // or dnf depending on what's available
		{"rhel", "yum"},
		{"arch", "pacman"},
		{"manjaro", "pacman"},
		{"alpine", "apk"},
		{"opensuse", "zypper"},
		{"gentoo", "emerge"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.distroID, func(t *testing.T) {
			result := determinePkgManager(tt.distroID)
			// For centos, it could be yum or dnf depending on system
			if tt.distroID == "centos" || tt.distroID == "rhel" {
				if result != "yum" && result != "dnf" {
					t.Errorf("determinePkgManager(%s) = %s, expected yum or dnf", tt.distroID, result)
				}
				return
			}
			if result != tt.expected {
				t.Errorf("determinePkgManager(%s) = %s, expected %s", tt.distroID, result, tt.expected)
			}
		})
	}
}

func TestHasPHP(t *testing.T) {
	result := HasPHP()
	// We can't assert true/false since it depends on the test environment
	t.Logf("HasPHP() = %v", result)
}

func TestGetPHPVersion(t *testing.T) {
	if !HasPHP() {
		t.Skip("PHP not available, skipping version test")
	}

	version, err := GetPHPVersion()
	if err != nil {
		t.Errorf("GetPHPVersion() returned error: %v", err)
	}
	if version == "" {
		t.Error("Expected non-empty PHP version")
	}
	t.Logf("PHP version: %s", version)
}

func TestInstallComposerViaPackageManager(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping package manager test on non-Linux platform")
	}

	// Test with unknown package manager
	distro := &DistroInfo{ID: "unknown", PackageManager: "unknown"}
	attempted, err := InstallComposerViaPackageManager(distro, false)
	if attempted {
		t.Error("Expected no attempt for unknown package manager")
	}
	if err != nil {
		t.Errorf("Expected nil error for unknown package manager, got: %v", err)
	}

	// Test with nil distro
	attempted, err = InstallComposerViaPackageManager(nil, false)
	if attempted {
		t.Error("Expected no attempt for nil distro")
	}
}

func TestCheckComposerVersion(t *testing.T) {
	// Try to find composer in PATH
	path, err := exec.LookPath("composer")
	if err != nil {
		t.Skip("Composer not found in PATH, skipping version check test")
	}

	version, err := CheckComposerVersion(path)
	if err != nil {
		t.Errorf("CheckComposerVersion() returned error: %v", err)
	}
	if version == "" {
		t.Error("Expected non-empty composer version")
	}
	t.Logf("Composer version: %s", version)
}

func TestDistroInfo(t *testing.T) {
	info := &DistroInfo{
		ID:             "ubuntu",
		Name:           "Ubuntu",
		Version:        "22.04",
		PackageManager: "apt",
	}

	if info.ID != "ubuntu" {
		t.Errorf("Expected ID=ubuntu, got %s", info.ID)
	}
	if info.Name != "Ubuntu" {
		t.Errorf("Expected Name=Ubuntu, got %s", info.Name)
	}
	if info.Version != "22.04" {
		t.Errorf("Expected Version=22.04, got %s", info.Version)
	}
	if info.PackageManager != "apt" {
		t.Errorf("Expected PackageManager=apt, got %s", info.PackageManager)
	}
}

func TestDetectViaKnownFiles(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping known files detection on non-Linux platform")
	}

	info := detectViaKnownFiles()
	if info != nil {
		t.Logf("Detected via known files: ID=%s, Name=%s, PackageManager=%s",
			info.ID, info.Name, info.PackageManager)
	} else {
		t.Log("No distro detected via known files (may be expected)")
	}
}

func TestDetectViaAvailablePkgManager(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping package manager detection on non-Linux platform")
	}

	info := detectViaAvailablePkgManager()
	if info == nil {
		t.Error("Expected non-nil result from detectViaAvailablePkgManager")
		return
	}
	t.Logf("Detected package manager: ID=%s, Name=%s, PackageManager=%s",
		info.ID, info.Name, info.PackageManager)
}

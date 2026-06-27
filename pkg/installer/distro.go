package installer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// DistroInfo represents information about a Linux distribution
type DistroInfo struct {
	// ID is the distro identifier (e.g., "ubuntu", "centos", "arch", "alpine")
	ID string
	// Name is the human-readable name (e.g., "Ubuntu", "CentOS")
	Name string
	// Version is the version string (e.g., "22.04", "9")
	Version string
	// PackageManager is the primary package manager (e.g., "apt", "yum", "dnf", "pacman", "apk")
	PackageManager string
}

// DetectLinuxDistro detects the current Linux distribution
// Returns a DistroInfo struct with distribution details
func DetectLinuxDistro() (*DistroInfo, error) {
	if runtime.GOOS != "linux" {
		return nil, fmt.Errorf("not running on Linux")
	}

	// Try /etc/os-release first (modern standard)
	if info, err := parseOSRelease(); err == nil && info != nil {
		return info, nil
	}

	// Fallback: try lsb_release command
	if info, err := detectViaLSBRelease(); err == nil && info != nil {
		return info, nil
	}

	// Fallback: check known files
	if info := detectViaKnownFiles(); info != nil {
		return info, nil
	}

	// Last resort: check which package managers are available
	return detectViaAvailablePkgManager(), nil
}

// parseOSRelease reads /etc/os-release to determine distro
func parseOSRelease() (*DistroInfo, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return nil, err
	}

	info := &DistroInfo{}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := strings.Trim(parts[1], `"'`)

		switch key {
		case "ID":
			info.ID = value
		case "NAME":
			info.Name = value
		case "VERSION_ID":
			info.Version = value
		}
	}

	if info.ID == "" {
		return nil, fmt.Errorf("could not determine distro ID from /etc/os-release")
	}

	info.PackageManager = determinePkgManager(info.ID)
	return info, nil
}

// detectViaLSBRelease uses lsb_release command
func detectViaLSBRelease() (*DistroInfo, error) {
	cmd := exec.Command("lsb_release", "-a")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	info := &DistroInfo{}
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Distributor ID":
			info.ID = strings.ToLower(value)
			info.Name = value
		case "Release":
			info.Version = value
		}
	}

	if info.ID == "" {
		return nil, fmt.Errorf("could not determine distro from lsb_release")
	}

	info.PackageManager = determinePkgManager(info.ID)
	return info, nil
}

// detectViaKnownFiles checks for distro-specific files
func detectViaKnownFiles() *DistroInfo {
	// Red Hat / CentOS / Fedora
	if _, err := os.Stat("/etc/redhat-release"); err == nil {
		data, _ := os.ReadFile("/etc/redhat-release")
		releaseStr := string(data)

		if strings.Contains(releaseStr, "Fedora") {
			return &DistroInfo{ID: "fedora", Name: "Fedora", PackageManager: "dnf"}
		}
		if strings.Contains(releaseStr, "CentOS") {
			return &DistroInfo{ID: "centos", Name: "CentOS", PackageManager: "yum"}
		}
		return &DistroInfo{ID: "rhel", Name: "Red Hat Enterprise Linux", PackageManager: "yum"}
	}

	// Debian
	if _, err := os.Stat("/etc/debian_version"); err == nil {
		return &DistroInfo{ID: "debian", Name: "Debian", PackageManager: "apt"}
	}

	// Arch Linux
	if _, err := os.Stat("/etc/arch-release"); err == nil {
		return &DistroInfo{ID: "arch", Name: "Arch Linux", PackageManager: "pacman"}
	}

	// Alpine Linux
	if _, err := os.Stat("/etc/alpine-release"); err == nil {
		return &DistroInfo{ID: "alpine", Name: "Alpine Linux", PackageManager: "apk"}
	}

	// SuSE
	if _, err := os.Stat("/etc/SuSE-release"); err == nil {
		return &DistroInfo{ID: "suse", Name: "SUSE", PackageManager: "zypper"}
	}

	// Gentoo
	if _, err := os.Stat("/etc/gentoo-release"); err == nil {
		return &DistroInfo{ID: "gentoo", Name: "Gentoo", PackageManager: "emerge"}
	}

	return nil
}

// detectViaAvailablePkgManager detects which package manager is available
func detectViaAvailablePkgManager() *DistroInfo {
	pkgManagers := []struct {
		cmd  string
		info *DistroInfo
	}{
		{"apt-get", &DistroInfo{ID: "debian-like", Name: "Debian-like", PackageManager: "apt"}},
		{"dnf", &DistroInfo{ID: "fedora-like", Name: "Fedora-like", PackageManager: "dnf"}},
		{"yum", &DistroInfo{ID: "rhel-like", Name: "RHEL-like", PackageManager: "yum"}},
		{"pacman", &DistroInfo{ID: "arch-like", Name: "Arch-like", PackageManager: "pacman"}},
		{"apk", &DistroInfo{ID: "alpine-like", Name: "Alpine-like", PackageManager: "apk"}},
		{"zypper", &DistroInfo{ID: "suse-like", Name: "SUSE-like", PackageManager: "zypper"}},
		{"emerge", &DistroInfo{ID: "gentoo-like", Name: "Gentoo-like", PackageManager: "emerge"}},
	}

	for _, pm := range pkgManagers {
		if _, err := exec.LookPath(pm.cmd); err == nil {
			return pm.info
		}
	}

	// Ultimate fallback
	return &DistroInfo{ID: "unknown", Name: "Unknown Linux", PackageManager: "unknown"}
}

// determinePkgManager maps distro IDs to their package managers
func determinePkgManager(distroID string) string {
	switch distroID {
	case "ubuntu", "debian", "linuxmint", "pop", "elementary", "kubuntu", "xubuntu", "linuxlite":
		return "apt"
	case "fedora":
		return "dnf"
	case "centos", "rhel", "redhat", "oracle", "amzn", "almalinux", "rocky":
		// CentOS 8+ uses dnf, but yum is still common
		if _, err := exec.LookPath("dnf"); err == nil {
			return "dnf"
		}
		return "yum"
	case "arch", "manjaro", "endeavouros", "garuda":
		return "pacman"
	case "alpine":
		return "apk"
	case "opensuse", "opensuse-leap", "opensuse-tumbleweed", "sles":
		return "zypper"
	case "gentoo":
		return "emerge"
	case "void":
		return "xbps"
	case "solus":
		return "eopkg"
	default:
		return "unknown"
	}
}

// InstallComposerViaPackageManager attempts to install Composer using the system's
// native package manager. Returns true if installation was attempted (whether
// successful or not), false if no suitable package manager was found.
func InstallComposerViaPackageManager(distro *DistroInfo, useSudo bool) (bool, error) {
	if distro == nil || distro.PackageManager == "unknown" {
		return false, nil
	}

	switch distro.PackageManager {
	case "apt":
		return true, runPkgCommand(useSudo, "apt-get", "install", "-y", "composer")
	case "dnf":
		return true, runPkgCommand(useSudo, "dnf", "install", "-y", "composer")
	case "yum":
		return true, runPkgCommand(useSudo, "yum", "install", "-y", "composer")
	case "pacman":
		return true, runPkgCommand(useSudo, "pacman", "-S", "--noconfirm", "composer")
	case "apk":
		return true, runPkgCommand(useSudo, "apk", "add", "composer")
	case "zypper":
		return true, runPkgCommand(useSudo, "zypper", "-n", "in", "composer")
	default:
		return false, nil
	}
}

// runPkgCommand runs a package manager command, optionally with sudo
func runPkgCommand(useSudo bool, args ...string) error {
	var cmd *exec.Cmd
	if useSudo {
		allArgs := append([]string{"sudo"}, args...)
		cmd = exec.Command(allArgs[0], allArgs[1:]...)
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("package manager command failed: %v, output: %s", err, string(output))
	}
	return nil
}

// HasPHP checks if PHP is available on the system
func HasPHP() bool {
	_, err := exec.LookPath("php")
	if err != nil {
		// Also try common PHP paths
		phpPaths := []string{
			"/usr/bin/php",
			"/usr/local/bin/php",
			"/opt/homebrew/bin/php", // macOS Homebrew
		}
		for _, p := range phpPaths {
			if _, err := os.Stat(p); err == nil {
				return true
			}
		}
		return false
	}
	return true
}

// GetPHPVersion returns the PHP version string if PHP is available
func GetPHPVersion() (string, error) {
	cmd := exec.Command("php", "-r", "echo PHP_VERSION;")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get PHP version: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// InstallPHP attempts to install PHP using the system's package manager
func InstallPHP(distro *DistroInfo, useSudo bool) error {
	if distro == nil || distro.PackageManager == "unknown" {
		return fmt.Errorf("cannot install PHP: unknown package manager")
	}

	switch distro.PackageManager {
	case "apt":
		return runPkgCommand(useSudo, "apt-get", "install", "-y", "php", "php-cli", "php-mbstring", "php-xml", "php-curl")
	case "dnf":
		return runPkgCommand(useSudo, "dnf", "install", "-y", "php", "php-cli", "php-mbstring", "php-xml", "php-curl")
	case "yum":
		return runPkgCommand(useSudo, "yum", "install", "-y", "php", "php-cli", "php-mbstring", "php-xml", "php-curl")
	case "pacman":
		return runPkgCommand(useSudo, "pacman", "-S", "--noconfirm", "php")
	case "apk":
		return runPkgCommand(useSudo, "apk", "add", "php81", "php81-cli", "php81-mbstring", "php81-xml", "php81-curl")
	case "zypper":
		return runPkgCommand(useSudo, "zypper", "-n", "in", "php7", "php7-cli")
	default:
		return fmt.Errorf("unsupported package manager: %s", distro.PackageManager)
	}
}

// CheckComposerVersion runs `composer --version` and returns the version string
func CheckComposerVersion(composerPath string) (string, error) {
	cmd := exec.Command(composerPath, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to check composer version: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// InstallComposerVersion installs a specific version of Composer
// version can be "1", "2", "2.2", "latest", "preview", or a specific version like "2.5.1"
func InstallComposerVersion(version string, installPath string, useSudo bool) error {
	// Map common version aliases to their download URLs
	var downloadURL string
	switch version {
	case "1", "1.x":
		downloadURL = "https://getcomposer.org/composer-1.phar"
	case "2", "2.x", "latest", "":
		downloadURL = "https://getcomposer.org/composer.phar"
	case "preview":
		downloadURL = "https://getcomposer.org/composer-preview.phar"
	case "stable":
		downloadURL = "https://getcomposer.org/composer.phar"
	default:
		// Assume specific version like "2.5.1"
		downloadURL = fmt.Sprintf("https://getcomposer.org/download/%s/composer.phar", version)
	}

	// Download the phar file
	pharPath := fmt.Sprintf("%s/composer.phar", installPath)
	downloadConfig := downloadConfigDefault()
	if err := downloadComposerPhar(downloadURL, pharPath, downloadConfig); err != nil {
		return fmt.Errorf("failed to download composer phar: %w", err)
	}

	// Make it executable and create wrapper script
	return createComposerWrapper(pharPath, installPath, useSudo)
}

// downloadComposerPhar downloads a composer phar file
func downloadComposerPhar(url, destPath string, config downloadConfigInternal) error {
	cmd := exec.Command("curl", "-fsSL", "-o", destPath, url)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("download failed: %v, output: %s", err, string(output))
	}
	return nil
}

// downloadConfigInternal is an internal type for download configuration
type downloadConfigInternal struct {
	UseProxy bool
	ProxyURL string
}

func downloadConfigDefault() downloadConfigInternal {
	return downloadConfigInternal{}
}

// createComposerWrapper creates the composer executable wrapper script
func createComposerWrapper(pharPath, installPath string, useSudo bool) error {
	if runtime.GOOS == "windows" {
		batPath := fmt.Sprintf("%s/composer.bat", installPath)
		batContent := fmt.Sprintf("@php \"%s\" %%*", pharPath)
		if err := writeFile(batPath, []byte(batContent), useSudo); err != nil {
			return err
		}
	} else {
		binPath := fmt.Sprintf("%s/composer", installPath)
		binContent := fmt.Sprintf("#!/bin/sh\nphp \"%s\" \"$@\"", pharPath)
		if err := writeFile(binPath, []byte(binContent), useSudo); err != nil {
			return err
		}
		// Make executable
		if useSudo {
			exec.Command("sudo", "chmod", "+x", binPath).Run()
		} else {
			os.Chmod(binPath, 0755)
		}
	}
	return nil
}

// writeFile writes content to a file, optionally using sudo
func writeFile(path string, content []byte, useSudo bool) error {
	if useSudo {
		cmd := exec.Command("sudo", "tee", path)
		cmd.Stdin = strings.NewReader(string(content))
		cmd.Stdout = nil // discard output
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to write file with sudo: %w", err)
		}
		return nil
	}
	return os.WriteFile(path, content, 0755)
}

// UninstallComposer removes Composer from the system
func UninstallComposer(composerPath string, useSudo bool) error {
	filesToRemove := []string{
		composerPath,
		strings.TrimSuffix(composerPath, "/composer") + "/composer.phar",
	}

	for _, file := range filesToRemove {
		if _, err := os.Stat(file); err == nil {
			if useSudo {
				if output, err := exec.Command("sudo", "rm", "-f", file).CombinedOutput(); err != nil {
					return fmt.Errorf("failed to remove %s: %v, output: %s", file, err, string(output))
				}
			} else {
				if err := os.Remove(file); err != nil {
					return fmt.Errorf("failed to remove %s: %w", file, err)
				}
			}
		}
	}

	return nil
}

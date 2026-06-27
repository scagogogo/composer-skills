package installer

import (
	"os"
	"path/filepath"
	"runtime"
)

// Config holds the installer configuration
type Config struct {
	// DownloadURL is the URL to download the Composer installer script
	DownloadURL string
	// InstallPath is the directory where Composer will be installed
	InstallPath string
	// UseProxy indicates whether to use a proxy for downloads
	UseProxy bool
	// ProxyURL is the proxy server URL
	ProxyURL string
	// TimeoutSeconds is the download/install timeout in seconds
	TimeoutSeconds int
	// UseSudo indicates whether to use sudo/admin privileges (Unix-like systems)
	UseSudo bool
	// PreferBrewOnMac indicates whether to prefer Homebrew installation on macOS
	PreferBrewOnMac bool
	// PreferPackageManager indicates whether to try the OS package manager first (Linux)
	// If true, the installer will try apt/dnf/pacman/etc. before downloading directly
	PreferPackageManager bool
	// AutoInstallPHP indicates whether to automatically install PHP if not found
	// If true, the installer will attempt to install PHP via the system package manager
	// before installing Composer
	AutoInstallPHP bool
	// TargetVersion specifies which Composer version to install
	// Can be "latest", "1", "2", "preview", or a specific version like "2.5.1"
	// Empty means "latest"
	TargetVersion string
}

// DefaultConfig returns the default configuration suitable for the current OS
func DefaultConfig() Config {
	config := Config{
		DownloadURL:         "https://getcomposer.org/installer",
		TimeoutSeconds:      300,
		UseProxy:            false,
		PreferBrewOnMac:     true,
		PreferPackageManager: true,
		AutoInstallPHP:      true,
		TargetVersion:       "latest",
	}

	// Set default install path based on OS
	switch runtime.GOOS {
	case "windows":
		config.InstallPath = filepath.Join(os.Getenv("ProgramFiles"), "Composer")
	case "darwin", "linux":
		config.InstallPath = "/usr/local/bin"
	}

	return config
}

// SmartConfig returns an optimized configuration for the current system
// It detects the OS and adjusts settings for the best installation method
func SmartConfig() Config {
	config := DefaultConfig()

	switch runtime.GOOS {
	case "darwin":
		// macOS: prefer Homebrew
		config.PreferBrewOnMac = true
		config.AutoInstallPHP = true
	case "linux":
		// Linux: prefer native package manager
		config.PreferPackageManager = true
		config.AutoInstallPHP = true
		config.UseSudo = true
	case "windows":
		// Windows: direct download
		config.PreferPackageManager = false
		config.AutoInstallPHP = false
	}

	return config
}

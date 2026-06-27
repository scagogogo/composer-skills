//go:build windows
// +build windows

package detector

import (
	"os"
	"path/filepath"
)

// getPlatformSpecificPaths 返回 Windows 平台上可能的 Composer 路径
func getPlatformSpecificPaths() []string {
	return []string{
		filepath.Join(os.Getenv("APPDATA"), "Composer", "composer.phar"),
		filepath.Join(os.Getenv("ProgramFiles"), "Composer", "composer.phar"),
		filepath.Join(os.Getenv("ProgramFiles(x86)"), "Composer", "composer.phar"),
		"composer.phar",
		"composer.bat",
		"composer",
	}
}

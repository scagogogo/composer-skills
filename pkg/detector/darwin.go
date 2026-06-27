//go:build darwin
// +build darwin

package detector

import (
	"os"
	"path/filepath"
)

// getPlatformSpecificPaths 返回 macOS (Darwin) 平台上可能的 Composer 路径
func getPlatformSpecificPaths() []string {
	return []string{
		"/usr/local/bin/composer",
		"/usr/bin/composer",
		"/opt/homebrew/bin/composer",
		filepath.Join(os.Getenv("HOME"), ".composer/vendor/bin/composer"),
		filepath.Join(os.Getenv("HOME"), "composer.phar"),
	}
}

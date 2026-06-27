package installer

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	// 检查默认URL
	if config.DownloadURL != "https://getcomposer.org/installer" {
		t.Errorf("DefaultConfig() DownloadURL = %v, 期望值 %v",
			config.DownloadURL, "https://getcomposer.org/installer")
	}

	// 检查超时设置
	if config.TimeoutSeconds != 300 {
		t.Errorf("DefaultConfig() TimeoutSeconds = %v, 期望值 %v",
			config.TimeoutSeconds, 300)
	}

	// 检查代理设置
	if config.UseProxy != false {
		t.Errorf("DefaultConfig() UseProxy = %v, 期望值 %v",
			config.UseProxy, false)
	}

	// 检查Brew设置
	if config.PreferBrewOnMac != true {
		t.Errorf("DefaultConfig() PreferBrewOnMac = %v, 期望值 %v",
			config.PreferBrewOnMac, true)
	}

	// 检查平台相关设置
	switch runtime.GOOS {
	case "windows":
		expectedPath := filepath.Join(os.Getenv("ProgramFiles"), "Composer")
		if config.InstallPath != expectedPath {
			t.Errorf("DefaultConfig() Windows InstallPath = %v, 期望值 %v",
				config.InstallPath, expectedPath)
		}
	case "darwin", "linux":
		if config.InstallPath != "/usr/local/bin" {
			t.Errorf("DefaultConfig() Unix InstallPath = %v, 期望值 %v",
				config.InstallPath, "/usr/local/bin")
		}
	}
}

package installer

import (
	"runtime"
	"testing"
)

func TestErrors(t *testing.T) {
	if ErrInstallationFailed.Error() != "安装失败" {
		t.Errorf("ErrInstallationFailed message mismatch, got: %s", ErrInstallationFailed.Error())
	}
	if ErrInsufficientRights.Error() != "权限不足，请使用管理员/sudo权限" {
		t.Errorf("ErrInsufficientRights message mismatch, got: %s", ErrInsufficientRights.Error())
	}
	if ErrUnsupportedPlatform.Error() != "不支持的操作系统平台" {
		t.Errorf("ErrUnsupportedPlatform message mismatch, got: %s", ErrUnsupportedPlatform.Error())
	}
	if ErrDownloadFailed.Error() != "下载失败" {
		t.Errorf("ErrDownloadFailed message mismatch, got: %s", ErrDownloadFailed.Error())
	}
	if ErrPHPNotFound.Error() != "未找到PHP，Composer需要PHP才能运行" {
		t.Errorf("ErrPHPNotFound message mismatch, got: %s", ErrPHPNotFound.Error())
	}
	if ErrComposerAlreadyInstalled.Error() != "Composer已经安装" {
		t.Errorf("ErrComposerAlreadyInstalled message mismatch, got: %s", ErrComposerAlreadyInstalled.Error())
	}
}

func TestNewInstaller(t *testing.T) {
	config := Config{
		DownloadURL:         "http://example.com/composer.php",
		InstallPath:         "/test/path",
		UseProxy:            true,
		ProxyURL:            "http://proxy.example.com",
		TimeoutSeconds:      120,
		UseSudo:             true,
		PreferBrewOnMac:     false,
		PreferPackageManager: true,
		AutoInstallPHP:      true,
		TargetVersion:       "2",
	}

	inst := NewInstaller(config)

	if inst.config.DownloadURL != config.DownloadURL {
		t.Errorf("NewInstaller.config.DownloadURL = %v, 期望值 %v",
			inst.config.DownloadURL, config.DownloadURL)
	}
	if inst.config.InstallPath != config.InstallPath {
		t.Errorf("NewInstaller.config.InstallPath = %v, 期望值 %v",
			inst.config.InstallPath, config.InstallPath)
	}
	if inst.config.PreferPackageManager != config.PreferPackageManager {
		t.Errorf("NewInstaller.config.PreferPackageManager = %v, 期望值 %v",
			inst.config.PreferPackageManager, config.PreferPackageManager)
	}
	if inst.config.AutoInstallPHP != config.AutoInstallPHP {
		t.Errorf("NewInstaller.config.AutoInstallPHP = %v, 期望值 %v",
			inst.config.AutoInstallPHP, config.AutoInstallPHP)
	}
	if inst.config.TargetVersion != config.TargetVersion {
		t.Errorf("NewInstaller.config.TargetVersion = %v, 期望值 %v",
			inst.config.TargetVersion, config.TargetVersion)
	}
}

func TestDefaultInstaller(t *testing.T) {
	inst := DefaultInstaller()

	defaultConfig := DefaultConfig()
	if inst.config.DownloadURL != defaultConfig.DownloadURL {
		t.Errorf("DefaultInstaller().config.DownloadURL = %v, 期望值 %v",
			inst.config.DownloadURL, defaultConfig.DownloadURL)
	}
	if inst.config.PreferPackageManager != defaultConfig.PreferPackageManager {
		t.Errorf("DefaultInstaller().config.PreferPackageManager = %v, 期望值 %v",
			inst.config.PreferPackageManager, defaultConfig.PreferPackageManager)
	}
	if inst.config.AutoInstallPHP != defaultConfig.AutoInstallPHP {
		t.Errorf("DefaultInstaller().config.AutoInstallPHP = %v, 期望值 %v",
			inst.config.AutoInstallPHP, defaultConfig.AutoInstallPHP)
	}
}

func TestSmartConfig(t *testing.T) {
	config := SmartConfig()

	if config.DownloadURL == "" {
		t.Error("SmartConfig() should have non-empty DownloadURL")
	}
	if config.InstallPath == "" {
		t.Error("SmartConfig() should have non-empty InstallPath")
	}

	// Check platform-specific defaults
	switch runtime.GOOS {
	case "darwin":
		if !config.PreferBrewOnMac {
			t.Error("SmartConfig on macOS should have PreferBrewOnMac=true")
		}
		if !config.AutoInstallPHP {
			t.Error("SmartConfig on macOS should have AutoInstallPHP=true")
		}
	case "linux":
		if !config.PreferPackageManager {
			t.Error("SmartConfig on Linux should have PreferPackageManager=true")
		}
		if !config.AutoInstallPHP {
			t.Error("SmartConfig on Linux should have AutoInstallPHP=true")
		}
		if !config.UseSudo {
			t.Error("SmartConfig on Linux should have UseSudo=true")
		}
	case "windows":
		if config.PreferPackageManager {
			t.Error("SmartConfig on Windows should have PreferPackageManager=false")
		}
	}
}

func TestGetConfig(t *testing.T) {
	config := Config{
		DownloadURL:    "http://example.com/composer.php",
		InstallPath:    "/test/path",
		UseProxy:       true,
		ProxyURL:       "http://proxy.example.com",
		TimeoutSeconds: 120,
		UseSudo:        true,
	}

	inst := NewInstaller(config)
	gotConfig := inst.GetConfig()

	if gotConfig.DownloadURL != config.DownloadURL {
		t.Errorf("GetConfig().DownloadURL = %v, 期望值 %v",
			gotConfig.DownloadURL, config.DownloadURL)
	}
	if gotConfig.InstallPath != config.InstallPath {
		t.Errorf("GetConfig().InstallPath = %v, 期望值 %v",
			gotConfig.InstallPath, config.InstallPath)
	}
}

func TestSetConfig(t *testing.T) {
	originalConfig := DefaultConfig()
	inst := NewInstaller(originalConfig)

	newConfig := Config{
		DownloadURL:         "http://example.com/composer.php",
		InstallPath:         "/new/test/path",
		UseProxy:            true,
		ProxyURL:            "http://proxy.example.com",
		TimeoutSeconds:      120,
		UseSudo:             true,
		PreferBrewOnMac:     false,
		PreferPackageManager: true,
		AutoInstallPHP:      false,
		TargetVersion:       "2.5",
	}

	inst.SetConfig(newConfig)
	gotConfig := inst.GetConfig()

	if gotConfig.DownloadURL != newConfig.DownloadURL {
		t.Errorf("SetConfig后, GetConfig().DownloadURL = %v, 期望值 %v",
			gotConfig.DownloadURL, newConfig.DownloadURL)
	}
	if gotConfig.InstallPath != newConfig.InstallPath {
		t.Errorf("SetConfig后, GetConfig().InstallPath = %v, 期望值 %v",
			gotConfig.InstallPath, newConfig.InstallPath)
	}
	if gotConfig.PreferPackageManager != newConfig.PreferPackageManager {
		t.Errorf("SetConfig后, GetConfig().PreferPackageManager = %v, 期望值 %v",
			gotConfig.PreferPackageManager, newConfig.PreferPackageManager)
	}
	if gotConfig.AutoInstallPHP != newConfig.AutoInstallPHP {
		t.Errorf("SetConfig后, GetConfig().AutoInstallPHP = %v, 期望值 %v",
			gotConfig.AutoInstallPHP, newConfig.AutoInstallPHP)
	}
	if gotConfig.TargetVersion != newConfig.TargetVersion {
		t.Errorf("SetConfig后, GetConfig().TargetVersion = %v, 期望值 %v",
			gotConfig.TargetVersion, newConfig.TargetVersion)
	}
}

func TestIsInstalled(t *testing.T) {
	inst := DefaultInstaller()
	result := inst.IsInstalled()
	// Just verify it doesn't panic - result depends on environment
	t.Logf("IsInstalled() = %v", result)
}

func TestGetInstalledVersion(t *testing.T) {
	inst := DefaultInstaller()
	version, err := inst.GetInstalledVersion()
	if err != nil {
		t.Logf("GetInstalledVersion() returned error (expected if composer not installed): %v", err)
		return
	}
	t.Logf("GetInstalledVersion() = %s", version)
}

func TestDefaultConfigNewFields(t *testing.T) {
	config := DefaultConfig()

	if !config.PreferPackageManager {
		t.Error("DefaultConfig().PreferPackageManager should be true")
	}
	if !config.AutoInstallPHP {
		t.Error("DefaultConfig().AutoInstallPHP should be true")
	}
	if config.TargetVersion != "latest" {
		t.Errorf("DefaultConfig().TargetVersion = %v, expected latest", config.TargetVersion)
	}
}

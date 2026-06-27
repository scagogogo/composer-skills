package composer

import (
	"errors"
	"testing"
)

func TestInstallWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install", "Dependencies installed", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"no-dev":              "",
		"optimize-autoloader": "",
	}
	err := composer.InstallWithOptions(options)

	if err != nil {
		t.Errorf("InstallWithOptions failed: %v", err)
	}
}

func TestUpdateWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update --no-dev vendor/package1 vendor/package2", "Packages updated", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-dev": ""}
	err := composer.UpdateWithOptions([]string{"vendor/package1", "vendor/package2"}, options)

	if err != nil {
		t.Errorf("UpdateWithOptions failed: %v", err)
	}
}

func TestDumpAutoloadWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("dump-autoload --optimize", "Autoload dumped and optimized", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"optimize": ""}
	err := composer.DumpAutoloadWithOptions(options)

	if err != nil {
		t.Errorf("DumpAutoloadWithOptions failed: %v", err)
	}
}

func TestCheckDependencies(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check", "Dependencies check passed", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.CheckDependencies()

	if err != nil {
		t.Errorf("CheckDependencies failed: %v", err)
	}
	if output != "Dependencies check passed" {
		t.Errorf("Expected 'Dependencies check passed', got '%s'", output)
	}
}

func TestFundPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("fund", "Funding information", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.FundPackages()

	if err != nil {
		t.Errorf("FundPackages failed: %v", err)
	}
	if output != "Funding information" {
		t.Errorf("Expected 'Funding information', got '%s'", output)
	}
}

func TestRunAudit(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit", "Audit complete", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.RunAudit()

	if err != nil {
		t.Errorf("RunAudit failed: %v", err)
	}
	if output != "Audit complete" {
		t.Errorf("Expected 'Audit complete', got '%s'", output)
	}
}

func TestSuggests(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output for suggests command
	SetupMockOutput("suggests", "vendor/suggested-package: For enhanced functionality\nvendor/another-package: For additional features", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err != nil {
		t.Errorf("Suggests执行失败: %v", err)
	}
}

func TestSuggestsWithError(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock to return an error
	SetupMockOutput("suggests", "", errors.New("suggests command failed"))

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err == nil {
		t.Error("Suggests命令失败时应该返回错误")
	}
}

func TestSuggestsWithEmptyOutput(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Set up mock output with empty result
	SetupMockOutput("suggests", "", nil)

	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	err = composer.Suggests()
	if err != nil {
		t.Errorf("Suggests执行失败: %v", err)
	}
}

func TestInstallWithPreferSource(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install --prefer-source", "Dependencies installed from source", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InstallWithPreferSource()

	if err != nil {
		t.Errorf("InstallWithPreferSource failed: %v", err)
	}
}

func TestInstallWithPreferDist(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install --prefer-dist", "Dependencies installed from dist", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InstallWithPreferDist()

	if err != nil {
		t.Errorf("InstallWithPreferDist failed: %v", err)
	}
}

func TestInstallNoScripts(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install --no-scripts", "Dependencies installed without scripts", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InstallNoScripts()

	if err != nil {
		t.Errorf("InstallNoScripts failed: %v", err)
	}
}

func TestInstallWithClassmapAuthoritative(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install --classmap-authoritative", "Dependencies installed with authoritative classmap", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InstallWithClassmapAuthoritative()

	if err != nil {
		t.Errorf("InstallWithClassmapAuthoritative failed: %v", err)
	}
}

func TestInstallWithAPCu(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install --apcu-autoloader", "Dependencies installed with APCu autoloader", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InstallWithAPCu()

	if err != nil {
		t.Errorf("InstallWithAPCu failed: %v", err)
	}
}

func TestUpdateWithPreferSource(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update --prefer-source vendor/package", "Packages updated from source", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UpdateWithPreferSource([]string{"vendor/package"})

	if err != nil {
		t.Errorf("UpdateWithPreferSource failed: %v", err)
	}
}

func TestUpdateWithPreferDist(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update --prefer-dist vendor/package", "Packages updated from dist", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UpdateWithPreferDist([]string{"vendor/package"})

	if err != nil {
		t.Errorf("UpdateWithPreferDist failed: %v", err)
	}
}

func TestUpdateWithLock(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update --lock", "Lock file hash updated", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UpdateWithLock()

	if err != nil {
		t.Errorf("UpdateWithLock failed: %v", err)
	}
}

func TestUpdateNoScripts(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update --no-scripts vendor/package", "Packages updated without scripts", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UpdateNoScripts([]string{"vendor/package"})

	if err != nil {
		t.Errorf("UpdateNoScripts failed: %v", err)
	}
}

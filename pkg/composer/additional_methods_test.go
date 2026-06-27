package composer

import (
	"testing"
)

// ==================== OutdatedWithOptions Tests ====================

func TestOutdatedWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated", `{"installed": []}`, nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"direct":  "",
		"format":  "json",
		"minor-only": "",
	}
	output, err := comp.OutdatedWithOptions(options)

	if err != nil {
		t.Errorf("OutdatedWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

// ==================== Dry Run Tests ====================

func TestInstallDryRun(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install", "Dry run: nothing would be installed", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.InstallDryRun()

	if err != nil {
		t.Errorf("InstallDryRun failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty dry run output")
	}
}

func TestUpdateDryRun(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update", "Dry run: nothing would be updated", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.UpdateDryRun([]string{"vendor/pkg"})

	if err != nil {
		t.Errorf("UpdateDryRun failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty dry run output")
	}
}

func TestRequireDryRun(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require", "Dry run: would require vendor/pkg", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.RequireDryRun("vendor/pkg", "^1.0")

	if err != nil {
		t.Errorf("RequireDryRun failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty dry run output")
	}
}

func TestRemoveDryRun(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("remove", "Dry run: would remove vendor/pkg", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.RemoveDryRun("vendor/pkg")

	if err != nil {
		t.Errorf("RemoveDryRun failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty dry run output")
	}
}

// ==================== RequireMultiple Tests ====================

func TestRequireMultiple(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require", "", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	packages := map[string]string{
		"vendor/pkg1": "^1.0",
		"vendor/pkg2": "^2.0",
	}
	err := comp.RequireMultiple(packages, false)

	if err != nil {
		t.Errorf("RequireMultiple failed: %v", err)
	}
}

func TestRequireMultiple_Dev(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require", "", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	packages := map[string]string{
		"vendor/pkg1": "^1.0",
	}
	err := comp.RequireMultiple(packages, true)

	if err != nil {
		t.Errorf("RequireMultiple dev failed: %v", err)
	}
}

// ==================== Depends/Why Options Tests ====================

func TestDependsWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("depends", "reverse dependencies output", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"tree": "",
	}
	output, err := comp.DependsWithOptions("vendor/pkg", options)

	if err != nil {
		t.Errorf("DependsWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

func TestWhyWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("why", "why output", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"tree": "",
	}
	output, err := comp.WhyWithOptions("vendor/pkg", options)

	if err != nil {
		t.Errorf("WhyWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

func TestWhyNotWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("why-not", "why-not output", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"tree": "",
	}
	output, err := comp.WhyNotWithOptions("vendor/pkg", "v1.0.0", options)

	if err != nil {
		t.Errorf("WhyNotWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

// ==================== Suggests Options Tests ====================

func TestSuggestsWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("suggests", "suggested packages output", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"no-dev": "",
	}
	output, err := comp.SuggestsWithOptions(options)

	if err != nil {
		t.Errorf("SuggestsWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

func TestSuggestsForPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("suggests", "package suggestions", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.SuggestsForPackage("vendor/pkg")

	if err != nil {
		t.Errorf("SuggestsForPackage failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

// ==================== Reinstall Multiple With Options Tests ====================

func TestReinstallMultipleWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("reinstall", "reinstalled packages", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	packages := []string{"vendor/pkg1", "vendor/pkg2"}
	options := map[string]string{
		"prefer-source": "",
	}
	err := comp.ReinstallMultipleWithOptions(packages, options)

	if err != nil {
		t.Errorf("ReinstallMultipleWithOptions failed: %v", err)
	}
}

// ==================== Global Enhanced Tests ====================

func TestGlobalInit(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "initialized global project", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.GlobalInit("myvendor/myproject")

	if err != nil {
		t.Errorf("GlobalInit failed: %v", err)
	}
}

func TestGlobalRequireMultiple(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "packages globally required", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	packages := map[string]string{
		"vendor/pkg1": "^1.0",
		"vendor/pkg2": "",
	}
	err := comp.GlobalRequireMultiple(packages)

	if err != nil {
		t.Errorf("GlobalRequireMultiple failed: %v", err)
	}
}

func TestGlobalRemoveMultiple(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "packages globally removed", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.GlobalRemoveMultiple([]string{"vendor/pkg1", "vendor/pkg2"})

	if err != nil {
		t.Errorf("GlobalRemoveMultiple failed: %v", err)
	}
}

// ==================== Show Enhanced Tests ====================

func TestShowWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show", "show output with options", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"format": "json",
		"direct": "",
	}
	output, err := comp.ShowWithOptions(options)

	if err != nil {
		t.Errorf("ShowWithOptions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

func TestShowLatestVersions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show", "latest versions output", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.ShowLatestVersions()

	if err != nil {
		t.Errorf("ShowLatestVersions failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

func TestShowOutdatedMinorOnly(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated", "minor only outdated", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.ShowOutdatedMinorOnly()

	if err != nil {
		t.Errorf("ShowOutdatedMinorOnly failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

// ==================== Install/Update NoDev Tests ====================

func TestInstallNoDev(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install", "installed no dev", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.InstallNoDev()

	if err != nil {
		t.Errorf("InstallNoDev failed: %v", err)
	}
}

func TestUpdateNoDev(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update", "updated no dev", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.UpdateNoDev([]string{"vendor/pkg"})

	if err != nil {
		t.Errorf("UpdateNoDev failed: %v", err)
	}
}

func TestUpdateWithDependencies(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update", "updated with deps", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.UpdateWithDependencies([]string{"vendor/pkg"})

	if err != nil {
		t.Errorf("UpdateWithDependencies failed: %v", err)
	}
}

func TestUpdateWithAllDependencies(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("update", "updated with all deps", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.UpdateWithAllDependencies([]string{"vendor/pkg"})

	if err != nil {
		t.Errorf("UpdateWithAllDependencies failed: %v", err)
	}
}

// ==================== InstallWithWorkingDir Tests ====================

func TestInstallWithWorkingDir(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("install", "installed in working dir", nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := comp.InstallWithWorkingDir("/tmp/test-project", true, true)

	if err != nil {
		t.Errorf("InstallWithWorkingDir failed: %v", err)
	}
}

// ==================== CheckPlatformReqsWithFormat Tests ====================

func TestCheckPlatformReqsWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check-platform-reqs", `{"platform": {}}`, nil)

	comp, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := comp.CheckPlatformReqsWithFormat("json")

	if err != nil {
		t.Errorf("CheckPlatformReqsWithFormat failed: %v", err)
	}
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

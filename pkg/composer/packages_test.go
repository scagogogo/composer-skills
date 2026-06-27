package composer

import (
	"testing"
)

func TestShowPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show vendor/package", "Package info", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowPackage("vendor/package")

	if err != nil {
		t.Errorf("ShowPackage failed: %v", err)
	}
	if output != "Package info" {
		t.Errorf("Expected 'Package info', got '%s'", output)
	}
}

func TestSearch(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("search query", "Search results", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Search("query")

	if err != nil {
		t.Errorf("Search failed: %v", err)
	}
	if output != "Search results" {
		t.Errorf("Expected 'Search results', got '%s'", output)
	}
}

func TestBrowsePackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("browse vendor/package", "Browser output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.BrowsePackage("vendor/package")

	if err != nil {
		t.Errorf("BrowsePackage failed: %v", err)
	}
}

func TestShowAllPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show", "all packages list", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowAllPackages()

	if err != nil {
		t.Errorf("ShowAllPackages failed: %v", err)
	}
	if output != "all packages list" {
		t.Errorf("Expected 'all packages list', got '%s'", output)
	}
}

func TestShowDependencyTree(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show --tree vendor/package", "dependency tree output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowDependencyTree("vendor/package")

	if err != nil {
		t.Errorf("ShowDependencyTree failed: %v", err)
	}
	if output != "dependency tree output" {
		t.Errorf("Expected 'dependency tree output', got '%s'", output)
	}
}

func TestShowReverseDependencies(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("depends vendor/package", "reverse deps output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowReverseDependencies("vendor/package")

	if err != nil {
		t.Errorf("ShowReverseDependencies failed: %v", err)
	}
	if output != "reverse deps output" {
		t.Errorf("Expected 'reverse deps output', got '%s'", output)
	}
}

func TestWhyPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("why vendor/package", "why output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.WhyPackage("vendor/package")

	if err != nil {
		t.Errorf("WhyPackage failed: %v", err)
	}
	if output != "why output" {
		t.Errorf("Expected 'why output', got '%s'", output)
	}
}

func TestOutdatedPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated", "outdated output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.OutdatedPackages()

	if err != nil {
		t.Errorf("OutdatedPackages failed: %v", err)
	}
	if output != "outdated output" {
		t.Errorf("Expected 'outdated output', got '%s'", output)
	}
}

func TestBumpPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("bump vendor/package1 vendor/package2", "bump output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.BumpPackages([]string{"vendor/package1", "vendor/package2"})

	if err != nil {
		t.Errorf("BumpPackages failed: %v", err)
	}
}

func TestWhyNotPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("why-not vendor/package v1.0.0", "why-not output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.WhyNotPackage("vendor/package", "v1.0.0")

	if err != nil {
		t.Errorf("WhyNotPackage failed: %v", err)
	}
	if output != "why-not output" {
		t.Errorf("Expected 'why-not output', got '%s'", output)
	}
}

func TestOutdatedPackagesDirect(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated --direct", "direct outdated output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.OutdatedPackagesDirect()

	if err != nil {
		t.Errorf("OutdatedPackagesDirect failed: %v", err)
	}
	if output != "direct outdated output" {
		t.Errorf("Expected 'direct outdated output', got '%s'", output)
	}
}

func TestRequirePackageWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("require --no-update vendor/package:^1.0", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-update": ""}
	err := composer.RequirePackageWithOptions("vendor/package", "^1.0", options)

	if err != nil {
		t.Errorf("RequirePackageWithOptions failed: %v", err)
	}
}

func TestBumpPackagesWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("bump --dry-run vendor/package", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"dry-run": ""}
	err := composer.BumpPackagesWithOptions([]string{"vendor/package"}, options)

	if err != nil {
		t.Errorf("BumpPackagesWithOptions failed: %v", err)
	}
}

func TestBrowsePackageWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("browse vendor/package --docs", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"docs": ""}
	err := composer.BrowsePackageWithOptions("vendor/package", options)

	if err != nil {
		t.Errorf("BrowsePackageWithOptions failed: %v", err)
	}
}

func TestReinstall(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("reinstall", "reinstalled vendor/package", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.Reinstall("vendor/package")

	if err != nil {
		t.Errorf("Reinstall failed: %v", err)
	}
}

func TestShowPackageWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show vendor/package --format json", `{"name": "vendor/package"}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowPackageWithFormat("vendor/package", "json")

	if err != nil {
		t.Errorf("ShowPackageWithFormat failed: %v", err)
	}
	if !contains(output, "vendor/package") {
		t.Errorf("Expected 'vendor/package', got '%s'", output)
	}
}

func TestShowOutdatedWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("outdated --format json", `{"installed": []}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowOutdatedWithFormat("json")

	if err != nil {
		t.Errorf("ShowOutdatedWithFormat failed: %v", err)
	}
	if !contains(output, "installed") {
		t.Errorf("Expected 'installed', got '%s'", output)
	}
}

func TestShowDirectPackages(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show --direct", "direct packages list", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowDirectPackages()

	if err != nil {
		t.Errorf("ShowDirectPackages failed: %v", err)
	}
	if output != "direct packages list" {
		t.Errorf("Expected 'direct packages list', got '%s'", output)
	}
}

func TestShowSelfPackage(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("show --self", "composer info", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ShowSelfPackage()

	if err != nil {
		t.Errorf("ShowSelfPackage failed: %v", err)
	}
	if !contains(output, "composer") {
		t.Errorf("Expected 'composer', got '%s'", output)
	}
}

func TestSearchWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("search logger --format json", `{"results": []}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.SearchWithFormat("logger", "json")

	if err != nil {
		t.Errorf("SearchWithFormat failed: %v", err)
	}
	if !contains(output, "results") {
		t.Errorf("Expected 'results', got '%s'", output)
	}
}

func TestSearchOnlyName(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("search logger --only-name", "monolog/monolog", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.SearchOnlyName("logger")

	if err != nil {
		t.Errorf("SearchOnlyName failed: %v", err)
	}
	if !contains(output, "monolog") {
		t.Errorf("Expected 'monolog', got '%s'", output)
	}
}

func TestSearchWithType(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("search logger --type=composer-plugin", "plugin results", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.SearchWithType("logger", "composer-plugin")

	if err != nil {
		t.Errorf("SearchWithType failed: %v", err)
	}
	if !contains(output, "plugin") {
		t.Errorf("Expected 'plugin', got '%s'", output)
	}
}

func TestRemoveWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("remove --dev --no-update vendor/package", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"dev":       "",
		"no-update": "",
	}
	err := composer.RemoveWithOptions("vendor/package", options)

	if err != nil {
		t.Errorf("RemoveWithOptions failed: %v", err)
	}
}

func TestRemoveMultiple(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("remove vendor/package1 vendor/package2", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveMultiple([]string{"vendor/package1", "vendor/package2"}, false)

	if err != nil {
		t.Errorf("RemoveMultiple failed: %v", err)
	}
}

func TestRemoveMultiple_Dev(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("remove --dev vendor/package1 vendor/package2", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveMultiple([]string{"vendor/package1", "vendor/package2"}, true)

	if err != nil {
		t.Errorf("RemoveMultiple dev failed: %v", err)
	}
}

func TestReinstallWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("reinstall", "reinstalled with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"prefer-source": ""}
	err := composer.ReinstallWithOptions("vendor/package", options)

	if err != nil {
		t.Errorf("ReinstallWithOptions failed: %v", err)
	}
}

func TestReinstallMultiple(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("reinstall", "reinstalled multiple", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.ReinstallMultiple([]string{"vendor/package1", "vendor/package2"})

	if err != nil {
		t.Errorf("ReinstallMultiple failed: %v", err)
	}
}

package composer

import (
	"testing"
)

func TestGlobalRequire(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global require vendor/package", "installed globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalRequire("vendor/package", "")

	if err != nil {
		t.Errorf("GlobalRequire failed: %v", err)
	}
}

func TestGlobalRequire_WithVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global require vendor/package:^1.0", "installed globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalRequire("vendor/package", "^1.0")

	if err != nil {
		t.Errorf("GlobalRequire with version failed: %v", err)
	}
}

func TestGlobalUpdate(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global update vendor/package1 vendor/package2", "updated globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalUpdate([]string{"vendor/package1", "vendor/package2"})

	if err != nil {
		t.Errorf("GlobalUpdate failed: %v", err)
	}
}

func TestGlobalUpdate_Empty(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global update", "updated all globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalUpdate([]string{})

	if err != nil {
		t.Errorf("GlobalUpdate empty failed: %v", err)
	}
}

func TestGlobalRemove(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global remove vendor/package", "removed globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalRemove("vendor/package")

	if err != nil {
		t.Errorf("GlobalRemove failed: %v", err)
	}
}

func TestGlobalInstall(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global install", "installed globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalInstall()

	if err != nil {
		t.Errorf("GlobalInstall failed: %v", err)
	}
}

func TestGlobalList(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global show", "globally installed packages", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GlobalList()

	if err != nil {
		t.Errorf("GlobalList failed: %v", err)
	}
	if output != "globally installed packages" {
		t.Errorf("Expected 'globally installed packages', got '%s'", output)
	}
}

func TestGlobalHome(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global home", "/home/user/.composer", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GlobalHome()

	if err != nil {
		t.Errorf("GlobalHome failed: %v", err)
	}
	if output != "/home/user/.composer" {
		t.Errorf("Expected '/home/user/.composer', got '%s'", output)
	}
}

func TestGlobalExecute(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global exec command arg1", "executed globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GlobalExecute("command", "arg1")

	if err != nil {
		t.Errorf("GlobalExecute failed: %v", err)
	}
	if output != "executed globally" {
		t.Errorf("Expected 'executed globally', got '%s'", output)
	}
}

func TestGlobalStatus(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global status", "global status OK", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GlobalStatus()

	if err != nil {
		t.Errorf("GlobalStatus failed: %v", err)
	}
	if output != "global status OK" {
		t.Errorf("Expected 'global status OK', got '%s'", output)
	}
}

func TestGlobalDumpAutoload(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global dump-autoload", "autoload dumped globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalDumpAutoload(false)

	if err != nil {
		t.Errorf("GlobalDumpAutoload failed: %v", err)
	}
}

func TestGlobalDumpAutoload_Optimize(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global dump-autoload --optimize", "optimized autoload dumped globally", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.GlobalDumpAutoload(true)

	if err != nil {
		t.Errorf("GlobalDumpAutoload optimize failed: %v", err)
	}
}

func TestGlobalRequireWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "installed globally with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"prefer-dist": ""}
	err := composer.GlobalRequireWithOptions("vendor/package", "^1.0", options)

	if err != nil {
		t.Errorf("GlobalRequireWithOptions failed: %v", err)
	}
}

func TestGlobalRequireWithOptions_NoVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "installed globally with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"prefer-dist": ""}
	err := composer.GlobalRequireWithOptions("vendor/package", "", options)

	if err != nil {
		t.Errorf("GlobalRequireWithOptions no version failed: %v", err)
	}
}

func TestGlobalUpdateWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "updated globally with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"prefer-dist": "",
		"no-dev":      "",
	}
	err := composer.GlobalUpdateWithOptions([]string{"vendor/package1", "vendor/package2"}, options)

	if err != nil {
		t.Errorf("GlobalUpdateWithOptions failed: %v", err)
	}
}

func TestGlobalUpdateWithOptions_Empty(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "updated all globally with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-dev": ""}
	err := composer.GlobalUpdateWithOptions([]string{}, options)

	if err != nil {
		t.Errorf("GlobalUpdateWithOptions empty failed: %v", err)
	}
}

func TestGlobalRemoveWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("global", "removed globally with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-progress": ""}
	err := composer.GlobalRemoveWithOptions("vendor/package", options)

	if err != nil {
		t.Errorf("GlobalRemoveWithOptions failed: %v", err)
	}
}
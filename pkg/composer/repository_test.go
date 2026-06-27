package composer

import (
	"testing"
)

func TestAddRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.my-repo {\"type\":\"composer\",\"url\":\"https://example.com\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	repo := Repository{
		Type: ComposerRepository,
		URL:  "https://example.com",
	}
	err := composer.AddRepository("my-repo", repo)

	if err != nil {
		t.Errorf("AddRepository failed: %v", err)
	}
}

func TestRemoveRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --unset repositories.my-repo", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveRepository("my-repo")

	if err != nil {
		t.Errorf("RemoveRepository failed: %v", err)
	}
}

func TestListRepositories(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories", "repo1\nrepo2", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListRepositories()

	if err != nil {
		t.Errorf("ListRepositories failed: %v", err)
	}
	if output != "repo1\nrepo2" {
		t.Errorf("Expected 'repo1\\nrepo2', got '%s'", output)
	}
}

func TestAddPackagistRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.packagist.org {\"type\":\"packagist\",\"url\":\"https://repo.packagist.org\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddPackagistRepository("https://repo.packagist.org")

	if err != nil {
		t.Errorf("AddPackagistRepository failed: %v", err)
	}
}

func TestDisablePackagistRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.packagist.org.url false", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.DisablePackagistRepository()

	if err != nil {
		t.Errorf("DisablePackagistRepository failed: %v", err)
	}
}

func TestEnablePackagistRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.packagist.org.url https://repo.packagist.org", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.EnablePackagistRepository()

	if err != nil {
		t.Errorf("EnablePackagistRepository failed: %v", err)
	}
}

func TestAddVcsRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.my-lib {\"type\":\"vcs\",\"url\":\"https://github.com/vendor/package\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddVcsRepository("my-lib", "https://github.com/vendor/package")

	if err != nil {
		t.Errorf("AddVcsRepository failed: %v", err)
	}
}

func TestAddPathRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.local {\"type\":\"path\",\"url\":\"../my-package\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddPathRepository("local", "../my-package", nil)

	if err != nil {
		t.Errorf("AddPathRepository failed: %v", err)
	}
}

func TestAddComposerRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.private {\"type\":\"composer\",\"url\":\"https://composer.example.org\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddComposerRepository("private", "https://composer.example.org")

	if err != nil {
		t.Errorf("AddComposerRepository failed: %v", err)
	}
}

func TestGetPreferredInstall(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config preferred-install", "dist", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetPreferredInstall()

	if err != nil {
		t.Errorf("GetPreferredInstall failed: %v", err)
	}
	if output != "dist" {
		t.Errorf("Expected 'dist', got '%s'", output)
	}
}

func TestSetPreferredInstall(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config preferred-install dist", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetPreferredInstall("dist")

	if err != nil {
		t.Errorf("SetPreferredInstall failed: %v", err)
	}
}

func TestSetPreferredInstall_Invalid(t *testing.T) {
	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetPreferredInstall("invalid")

	if err == nil {
		t.Errorf("SetPreferredInstall should return error for invalid value")
	}
}

func TestSetMinimumStability(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config minimum-stability beta", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetMinimumStability("beta")

	if err != nil {
		t.Errorf("SetMinimumStability failed: %v", err)
	}
}

func TestGetMinimumStability(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config minimum-stability", "stable", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetMinimumStability()

	if err != nil {
		t.Errorf("GetMinimumStability failed: %v", err)
	}
	if output != "stable" {
		t.Errorf("Expected 'stable', got '%s'", output)
	}
}

func TestGetPreferStable(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config prefer-stable", "1", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetPreferStable()

	if err != nil {
		t.Errorf("GetPreferStable failed: %v", err)
	}
	if output != "1" {
		t.Errorf("Expected '1', got '%s'", output)
	}
}

func TestAddArtifactRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config repositories.artifacts {\"type\":\"artifact\",\"url\":\"./packages\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.AddArtifactRepository("artifacts", "./packages")

	if err != nil {
		t.Errorf("AddArtifactRepository failed: %v", err)
	}
}

func TestSetConfigParameter(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config description My project", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetConfigParameter("description", "My project")

	if err != nil {
		t.Errorf("SetConfigParameter failed: %v", err)
	}
}

func TestGetConfigParameter(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config name", "vendor/package", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.GetConfigParameter("name")

	if err != nil {
		t.Errorf("GetConfigParameter failed: %v", err)
	}
	if output != "vendor/package" {
		t.Errorf("Expected 'vendor/package', got '%s'", output)
	}
}

func TestUnsetConfig(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --unset repositories.old-repo", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.UnsetConfig("repositories.old-repo")

	if err != nil {
		t.Errorf("UnsetConfig failed: %v", err)
	}
}

func TestAddGlobalRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global repositories.global-private {\"type\":\"composer\",\"url\":\"https://composer.example.org\"}", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	repo := Repository{
		Type: ComposerRepository,
		URL:  "https://composer.example.org",
	}
	err := composer.AddGlobalRepository("global-private", repo)

	if err != nil {
		t.Errorf("AddGlobalRepository failed: %v", err)
	}
}

func TestRemoveGlobalRepository(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global --unset repositories.global-private", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.RemoveGlobalRepository("global-private")

	if err != nil {
		t.Errorf("RemoveGlobalRepository failed: %v", err)
	}
}

func TestListGlobalRepositories(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --global repositories", "global-repo1\nglobal-repo2", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListGlobalRepositories()

	if err != nil {
		t.Errorf("ListGlobalRepositories failed: %v", err)
	}
	if output != "global-repo1\nglobal-repo2" {
		t.Errorf("Expected 'global-repo1\\nglobal-repo2', got '%s'", output)
	}
}

func TestSetPreferStable(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config prefer-stable 1", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetPreferStable(true)

	if err != nil {
		t.Errorf("SetPreferStable failed: %v", err)
	}
}

func TestSetPreferStable_False(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config prefer-stable 0", "", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.SetPreferStable(false)

	if err != nil {
		t.Errorf("SetPreferStable(false) failed: %v", err)
	}
}

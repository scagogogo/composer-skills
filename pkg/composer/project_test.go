package composer

import (
	"testing"
)

func TestCreateProject(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("create-project vendor/package:^1.0 /path/to/project", "Project created", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.CreateProject("vendor/package", "/path/to/project", "^1.0")

	if err != nil {
		t.Errorf("CreateProject failed: %v", err)
	}
}

func TestCreateProject_NoVersion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("create-project vendor/package /path/to/project", "Project created", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.CreateProject("vendor/package", "/path/to/project", "")

	if err != nil {
		t.Errorf("CreateProject no version failed: %v", err)
	}
}

func TestCreateProjectWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("create-project --no-dev vendor/package:^1.0 /path/to/project", "Project created with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-dev": ""}
	err := composer.CreateProjectWithOptions("vendor/package", "/path/to/project", "^1.0", options)

	if err != nil {
		t.Errorf("CreateProjectWithOptions failed: %v", err)
	}
}

func TestRunScript(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run-script test-script", "Script output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.RunScript("test-script")

	if err != nil {
		t.Errorf("RunScript failed: %v", err)
	}
	if output != "Script output" {
		t.Errorf("Expected 'Script output', got '%s'", output)
	}
}

func TestRunScript_WithArgs(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run-script test-script arg1 arg2", "Script with args output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.RunScript("test-script", "arg1", "arg2")

	if err != nil {
		t.Errorf("RunScript with args failed: %v", err)
	}
	if output != "Script with args output" {
		t.Errorf("Expected 'Script with args output', got '%s'", output)
	}
}

func TestExecuteScript(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run deploy", "deploy output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ExecuteScript("deploy")

	if err != nil {
		t.Errorf("ExecuteScript failed: %v", err)
	}
	if output != "deploy output" {
		t.Errorf("Expected 'deploy output', got '%s'", output)
	}
}

func TestListScripts(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("run-script --list", "Available scripts:\ntest\nbuild\ndeploy", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListScripts()

	if err != nil {
		t.Errorf("ListScripts failed: %v", err)
	}
	if output == "" {
		t.Error("ListScripts should return non-empty output")
	}
}

func TestArchiveProject(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("archive --dir=/tmp --format=zip", "Archive created", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.ArchiveProject("/tmp", "zip")

	if err != nil {
		t.Errorf("ArchiveProject failed: %v", err)
	}
}

func TestInitProject(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("init", "Project initialized", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	err := composer.InitProject()

	if err != nil {
		t.Errorf("InitProject failed: %v", err)
	}
}

func TestInitProjectWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("init --name=vendor/package --description=test --author=Test --no-interaction", "Project initialized", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"no-interaction": ""}
	err := composer.InitProjectWithOptions("vendor/package", "test", "Test", options)

	if err != nil {
		t.Errorf("InitProjectWithOptions failed: %v", err)
	}
}

func TestGetProjectInfo(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("config --list --json", `{"name":"test/package"}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	info, err := composer.GetProjectInfo()

	if err != nil {
		t.Errorf("GetProjectInfo failed: %v", err)
	}
	if info.Name != "test/package" {
		t.Errorf("Expected 'test/package', got '%s'", info.Name)
	}
}

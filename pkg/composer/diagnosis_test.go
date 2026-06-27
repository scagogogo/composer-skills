package composer

import (
	"testing"
)

func TestStatus(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("status", "no local changes", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Status()

	if err != nil {
		t.Errorf("Status failed: %v", err)
	}
	if output != "no local changes" {
		t.Errorf("Expected 'no local changes', got '%s'", output)
	}
}

func TestStatusWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("status --verbose", "verbose status info", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.StatusWithOptions(map[string]string{"verbose": ""})

	if err != nil {
		t.Errorf("StatusWithOptions failed: %v", err)
	}
	if output != "verbose status info" {
		t.Errorf("Expected 'verbose status info', got '%s'", output)
	}
}

func TestDiagnose(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("diagnose", "diagnosis complete", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Diagnose()

	if err != nil {
		t.Errorf("Diagnose failed: %v", err)
	}
	if output != "diagnosis complete" {
		t.Errorf("Expected 'diagnosis complete', got '%s'", output)
	}
}

func TestDiagnoseWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("diagnose --strict", "strict diagnosis", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.DiagnoseWithOptions(map[string]string{"strict": ""})

	if err != nil {
		t.Errorf("DiagnoseWithOptions failed: %v", err)
	}
	if output != "strict diagnosis" {
		t.Errorf("Expected 'strict diagnosis', got '%s'", output)
	}
}

func TestLocalExec(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("exec command arg1", "local exec output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.LocalExec("command", "arg1")

	if err != nil {
		t.Errorf("LocalExec failed: %v", err)
	}
	if output != "local exec output" {
		t.Errorf("Expected 'local exec output', got '%s'", output)
	}
}

func TestLocalExecWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("exec --dir=./packages command arg1", "local exec with options", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"dir": "./packages"}
	output, err := composer.LocalExecWithOptions("command", options, "arg1")

	if err != nil {
		t.Errorf("LocalExecWithOptions failed: %v", err)
	}
	if output != "local exec with options" {
		t.Errorf("Expected 'local exec with options', got '%s'", output)
	}
}

func TestCheck(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check", "dependencies satisfied", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Check()

	if err != nil {
		t.Errorf("Check failed: %v", err)
	}
	if output != "dependencies satisfied" {
		t.Errorf("Expected 'dependencies satisfied', got '%s'", output)
	}
}

func TestCheckWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("check --strict", "strict check results", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.CheckWithOptions(map[string]string{"strict": ""})

	if err != nil {
		t.Errorf("CheckWithOptions failed: %v", err)
	}
	if output != "strict check results" {
		t.Errorf("Expected 'strict check results', got '%s'", output)
	}
}
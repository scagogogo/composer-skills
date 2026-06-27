package composer

import (
	"testing"
)

func TestGenerateCompletion(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("completion bash", "bash completion script", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	script, err := composer.GenerateCompletion(BashShell)

	if err != nil {
		t.Errorf("GenerateCompletion failed: %v", err)
	}
	if script != "bash completion script" {
		t.Errorf("Expected 'bash completion script', got '%s'", script)
	}
}

func TestGenerateCompletion_Zsh(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("completion zsh", "zsh completion script", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	script, err := composer.GenerateCompletion(ZshShell)

	if err != nil {
		t.Errorf("GenerateCompletion zsh failed: %v", err)
	}
	if script != "zsh completion script" {
		t.Errorf("Expected 'zsh completion script', got '%s'", script)
	}
}

func TestGenerateCompletion_Fish(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("completion fish", "fish completion script", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	script, err := composer.GenerateCompletion(FishShell)

	if err != nil {
		t.Errorf("GenerateCompletion fish failed: %v", err)
	}
	if script != "fish completion script" {
		t.Errorf("Expected 'fish completion script', got '%s'", script)
	}
}

func TestGenerateCompletionWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("completion bash --global", "global bash completion", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"global": ""}
	script, err := composer.GenerateCompletionWithOptions(BashShell, options)

	if err != nil {
		t.Errorf("GenerateCompletionWithOptions failed: %v", err)
	}
	if script != "global bash completion" {
		t.Errorf("Expected 'global bash completion', got '%s'", script)
	}
}

func TestGenerateCompletionWithOptions_WithValue(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("completion bash --type=complete", "typed completion", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{"type": "complete"}
	script, err := composer.GenerateCompletionWithOptions(BashShell, options)

	if err != nil {
		t.Errorf("GenerateCompletionWithOptions with value failed: %v", err)
	}
	if script != "typed completion" {
		t.Errorf("Expected 'typed completion', got '%s'", script)
	}
}

func TestListCommands(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("list", "available commands list", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.ListCommands()

	if err != nil {
		t.Errorf("ListCommands failed: %v", err)
	}
	if output != "available commands list" {
		t.Errorf("Expected 'available commands list', got '%s'", output)
	}
}

func TestGetCommandHelp(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("help require", "help text for require command", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	help, err := composer.GetCommandHelp("require")

	if err != nil {
		t.Errorf("GetCommandHelp failed: %v", err)
	}
	if help != "help text for require command" {
		t.Errorf("Expected 'help text for require command', got '%s'", help)
	}
}

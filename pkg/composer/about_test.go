package composer

import (
	"testing"
)

func TestAbout(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("about", "Composer - Package Manager for PHP", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.About()

	if err != nil {
		t.Errorf("About failed: %v", err)
	}
	if output != "Composer - Package Manager for PHP" {
		t.Errorf("Expected 'Composer - Package Manager for PHP', got '%s'", output)
	}
}

func TestAboutWithError(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("about", "", ErrCommandExecution)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	_, err := composer.About()

	if err == nil {
		t.Error("About should return error when command fails")
	}
}

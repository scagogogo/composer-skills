package composer

import (
	"testing"
)

func TestBuildSatis(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("satis build /path/to/satis.json", "Satis build complete", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.BuildSatis("/path/to/satis.json", "")

	if err != nil {
		t.Errorf("BuildSatis failed: %v", err)
	}
	if output != "Satis build complete" {
		t.Errorf("Expected 'Satis build complete', got '%s'", output)
	}
}

func TestBuildSatisWithOutputDir(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("satis build /path/to/satis.json /output/dir", "Satis build complete", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.BuildSatis("/path/to/satis.json", "/output/dir")

	if err != nil {
		t.Errorf("BuildSatis with output dir failed: %v", err)
	}
	if output != "Satis build complete" {
		t.Errorf("Expected 'Satis build complete', got '%s'", output)
	}
}

package composer

import (
	"testing"
)

func TestAudit(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit", "Audit output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.Audit()

	if err != nil {
		t.Errorf("Audit failed: %v", err)
	}
	if output != "Audit output" {
		t.Errorf("Expected 'Audit output', got '%s'", output)
	}
}

func TestAuditWithoutDev(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit --no-dev", "Audit output without dev", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.AuditWithoutDev()

	if err != nil {
		t.Errorf("AuditWithoutDev failed: %v", err)
	}
	if output != "Audit output without dev" {
		t.Errorf("Expected 'Audit output without dev', got '%s'", output)
	}
}

func TestAuditWithFormat(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit --format=json", `{"vulnerabilities":[]}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.AuditWithFormat("json")

	if err != nil {
		t.Errorf("AuditWithFormat failed: %v", err)
	}
	if output != `{"vulnerabilities":[]}` {
		t.Errorf("Expected JSON output, got '%s'", output)
	}
}

func TestAuditWithFormat_Plain(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit --format=plain", "Audit output in plain text", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.AuditWithFormat("plain")

	if err != nil {
		t.Errorf("AuditWithFormat(plain) failed: %v", err)
	}
	if output != "Audit output in plain text" {
		t.Errorf("Expected plain text output, got '%s'", output)
	}
}

func TestAuditWithOptions(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit --format=json", `{"vulnerabilities":[]}`, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	options := map[string]string{
		"format": "json",
	}
	output, err := composer.AuditWithOptions(options)

	if err != nil {
		t.Errorf("AuditWithOptions failed: %v", err)
	}
	if output != `{"vulnerabilities":[]}` {
		t.Errorf("Expected JSON output, got '%s'", output)
	}
}

func TestAuditLock(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit composer.lock", "Audit of lock file", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.AuditLock("composer.lock")

	if err != nil {
		t.Errorf("AuditLock failed: %v", err)
	}
	if output != "Audit of lock file" {
		t.Errorf("Expected 'Audit of lock file', got '%s'", output)
	}
}

func TestAuditLock_EmptyPath(t *testing.T) {
	ClearMockOutputs()
	SetupMockOutput("audit", "Audit output", nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	output, err := composer.AuditLock("")

	if err != nil {
		t.Errorf("AuditLock with empty path failed: %v", err)
	}
	if output != "Audit output" {
		t.Errorf("Expected 'Audit output', got '%s'", output)
	}
}

func TestGetHighSeverityVulnerabilities(t *testing.T) {
	ClearMockOutputs()
	jsonOutput := `{
		"vulnerabilities": [
			{"package": "pkg1", "severity": "high", "advisory": "adv1"},
			{"package": "pkg2", "severity": "medium", "advisory": "adv2"},
			{"package": "pkg3", "severity": "critical", "advisory": "adv3"}
		],
		"found": 3
	}`
	SetupMockOutput("audit --format=json", jsonOutput, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	vulns, err := composer.GetHighSeverityVulnerabilities()

	if err != nil {
		t.Errorf("GetHighSeverityVulnerabilities failed: %v", err)
	}
	if len(vulns) != 2 {
		t.Errorf("Expected 2 high/critical vulnerabilities, got %d", len(vulns))
	}
}

func TestGetHighSeverityVulnerabilities_Empty(t *testing.T) {
	ClearMockOutputs()
	jsonOutput := `{"vulnerabilities": [], "found": 0}`
	SetupMockOutput("audit --format=json", jsonOutput, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	vulns, err := composer.GetHighSeverityVulnerabilities()

	if err != nil {
		t.Errorf("GetHighSeverityVulnerabilities empty failed: %v", err)
	}
	if len(vulns) != 0 {
		t.Errorf("Expected 0 vulnerabilities, got %d", len(vulns))
	}
}

func TestGetAbandonedPackages(t *testing.T) {
	ClearMockOutputs()
	jsonOutput := `{
		"vulnerabilities": [
			{"package": "pkg1", "abandoned": true, "advisory": "adv1"},
			{"package": "pkg2", "abandoned": false, "advisory": "adv2"},
			{"package": "pkg3", "abandoned": true, "advisory": "adv3"}
		],
		"found": 3
	}`
	SetupMockOutput("audit --format=json", jsonOutput, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	vulns, err := composer.GetAbandonedPackages()

	if err != nil {
		t.Errorf("GetAbandonedPackages failed: %v", err)
	}
	if len(vulns) != 2 {
		t.Errorf("Expected 2 abandoned packages, got %d", len(vulns))
	}
}

func TestGetAbandonedPackages_Empty(t *testing.T) {
	ClearMockOutputs()
	jsonOutput := `{"vulnerabilities": [], "found": 0}`
	SetupMockOutput("audit --format=json", jsonOutput, nil)

	composer, _ := New(Options{ExecutablePath: "/fake/composer"})
	vulns, err := composer.GetAbandonedPackages()

	if err != nil {
		t.Errorf("GetAbandonedPackages empty failed: %v", err)
	}
	if len(vulns) != 0 {
		t.Errorf("Expected 0 abandoned packages, got %d", len(vulns))
	}
}

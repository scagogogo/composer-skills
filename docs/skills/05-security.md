# Security Skills

Audit your PHP project for vulnerabilities, check platform requirements, and validate configuration files.

## Auditing Dependencies

### Basic Audit

Run a security audit on your project's dependencies:

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

output, err := comp.Audit()
fmt.Println(output)
```

CLI:

```bash
composer-skills local audit --working-dir /path/to/project
```

### Audit with JSON Output

Get structured results that you can process programmatically:

```go
result, err := comp.AuditWithJSON()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d vulnerabilities\n", result.Found)
for _, vuln := range result.Vulnerabilities {
    fmt.Printf("Package: %s %s\n", vuln.Package, vuln.Version)
    fmt.Printf("  Title: %s\n", vuln.Title)
    fmt.Printf("  Severity: %s\n", vuln.Severity)
    fmt.Printf("  Link: %s\n", vuln.Link)
    if len(vuln.CVE) > 0 {
        fmt.Printf("  CVE: %v\n", vuln.CVE)
    }
}
```

### Audit without Dev Dependencies

Only audit production dependencies:

```go
output, err := comp.AuditWithoutDev()
fmt.Println(output)
```

### Audit with Custom Format

Specify the output format (`json`, `table`, or `plain`):

```go
output, err := comp.AuditWithFormat("table")
```

### Audit with Custom Options

Full control over audit flags:

```go
output, err := comp.AuditWithOptions(map[string]string{
    "no-dev":  "",
    "format":  "json",
    "locked":  "",
})
```

### Audit a Lock File

Audit a specific `composer.lock` file without installing dependencies:

```go
// Audit the current project's lock file
output, err := comp.AuditLock("")

// Audit a specific lock file
output, err := comp.AuditLock("/path/to/other/project/composer.lock")
```

## Checking for Vulnerabilities

### Has Vulnerabilities

Quick boolean check for whether any vulnerabilities exist:

```go
hasVulns, err := comp.HasVulnerabilities()
if err != nil {
    log.Fatal(err)
}
if hasVulns {
    fmt.Println("WARNING: Security vulnerabilities found!")
} else {
    fmt.Println("No vulnerabilities found.")
}
```

### High Severity Vulnerabilities

Filter for only critical or high-severity issues:

```go
highVulns, err := comp.GetHighSeverityVulnerabilities()
if err != nil {
    log.Fatal(err)
}
if len(highVulns) > 0 {
    fmt.Printf("Found %d high/critical vulnerabilities:\n", len(highVulns))
    for _, vuln := range highVulns {
        fmt.Printf("  %s: %s (%s)\n", vuln.Package, vuln.Title, vuln.Severity)
    }
}
```

### Abandoned Packages

Find packages that have been abandoned by their maintainers:

```go
abandoned, err := comp.GetAbandonedPackages()
if err != nil {
    log.Fatal(err)
}
if len(abandoned) > 0 {
    fmt.Printf("Found %d abandoned packages:\n", len(abandoned))
    for _, pkg := range abandoned {
        fmt.Printf("  %s %s - %s\n", pkg.Package, pkg.Version, pkg.Link)
    }
    fmt.Println("Consider replacing these packages to avoid security risks.")
}
```

### Check for Security Vulnerabilities (verbose)

Get both the raw output and a boolean indicating vulnerability presence:

```go
output, hasVulns, err := comp.CheckForSecurityVulnerabilities()
if hasVulns {
    fmt.Println("Security vulnerabilities detected:")
    fmt.Println(output)
}
```

## Checking Platform Requirements

### Check Platform Requirements

Verify that the current system meets the platform requirements defined in `composer.json`:

```go
output, err := comp.CheckPlatformReqs()
fmt.Println(output)
```

### Check Platform (Structured)

Get structured platform requirement data:

```go
platforms, err := comp.CheckPlatform()
if err != nil {
    log.Fatal(err)
}
for _, p := range platforms {
    status := "NOT met"
    if p.Available {
        status = "met"
    }
    fmt.Printf("%s %s: %s\n", p.Name, p.Version, status)
}
```

### Check Platform with Lock

Check requirements from `composer.lock`:

```go
platforms, err := comp.CheckPlatformWithLock()
for _, p := range platforms {
    if !p.Available {
        fmt.Printf("WARNING: %s %s requirement not met\n", p.Name, p.Required)
    }
}
```

### Check Lock Platform Requirements

```go
output, err := comp.CheckPlatformReqsLock()
fmt.Println(output)
```

## Security Advisories (Packagist API)

Use the Packagist API client to check for known security advisories:

```go
c := client.NewComposerClient(30 * time.Second)

// All advisories
advisories, _ := c.GetSecurityAdvisories()

// For specific packages
advisories, _ = c.GetSecurityAdvisoriesForPackages([]string{
    "symfony/console",
    "laravel/framework",
})

// Recent advisories (last 7 days)
recent, _ := c.GetSecurityAdvisoriesSince(time.Now().AddDate(0, 0, -7))
```

See [Packagist API Skills](02-packagist-api.md#security-advisories) for full details.

## Validating composer.json and composer.lock

### Basic Validation

```go
err := comp.Validate()
if err != nil {
    fmt.Println("composer.json is invalid:", err)
} else {
    fmt.Println("composer.json is valid")
}
```

CLI:

```bash
composer-skills local validate --working-dir /path/to/project
```

### Strict Validation

Perform stricter validation that enforces best practices:

```go
output, err := comp.ValidateStrict()
```

### Validate with Dependencies

Also validate version constraints of all dependencies:

```go
output, err := comp.ValidateWithCheckVersion()
```

### Validate Schema Only

Check only the JSON structure, skip other checks:

```go
output, err := comp.ValidateSchema()
```

### Validate composer.lock

Check that `composer.lock` exists and is in sync with `composer.json`:

```go
output, err := comp.ValidateComposerLock()
```

### Validate Quietly

Only output on errors, silent on success:

```go
output, err := comp.ValidateQuiet()
```

### Validate with Custom Options

Combine multiple validation options:

```go
output, err := comp.ValidateWithOptions(map[string]string{
    "strict":            "",
    "with-dependencies": "",
    "no-check-publish":  "",
})
```

### Check Normalization

Verify that `composer.json` is properly formatted:

```go
output, err := comp.CheckNormalization()
```

### Full Validation with ComposerJson Struct

```go
output, err := comp.ValidateComposerJson(true, true) // strict=true, withDependencies=true
```

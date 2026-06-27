# Platform and Diagnosis Skills

Check PHP versions, extensions, platform requirements, diagnose issues, and keep Composer up to date.

## PHP Version Checking

### Get the PHP Version

Find which PHP version Composer is using:

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

phpVersion, err := comp.GetPHPVersion()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("PHP version: %s\n", phpVersion)
```

## Extension Checking

### List All Installed Extensions

```go
extensions, err := comp.GetExtensions()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Installed PHP extensions:")
for _, ext := range extensions {
    fmt.Println("  - " + ext)
}
```

### Check for a Specific Extension

```go
hasMbstring, err := comp.HasExtension("mbstring")
if err != nil {
    log.Fatal(err)
}
if hasMbstring {
    fmt.Println("mbstring is installed")
} else {
    fmt.Println("mbstring is NOT installed")
}
```

Common extensions to check: `json`, `mbstring`, `xml`, `curl`, `pdo`, `pdo_mysql`, `openssl`, `zip`.

## Platform Requirements

### Check Platform Requirements

Verify the current system meets the platform requirements defined in `composer.json`:

```go
platforms, err := comp.CheckPlatform()
if err != nil {
    log.Fatal(err)
}
for _, p := range platforms {
    status := "NOT available"
    if p.Available {
        status = "available"
    }
    fmt.Printf("%s %s: %s\n", p.Name, p.Version, status)
}
```

### Check Platform with Lock

Verify requirements from `composer.lock`:

```go
platforms, err := comp.CheckPlatformWithLock()
for _, p := range platforms {
    if !p.Available {
        fmt.Printf("WARNING: %s %s not met\n", p.Name, p.Required)
    }
}
```

### Check a Specific Platform Requirement

```go
// Check if PHP meets a minimum version
available, err := comp.IsPlatformAvailable("php", ">=7.4")
if err != nil {
    log.Fatal(err)
}
if available {
    fmt.Println("PHP >= 7.4 is available")
} else {
    fmt.Println("PHP >= 7.4 is NOT available")
}

// Check if an extension is available
available, err = comp.IsPlatformAvailable("ext-mbstring", "")
```

### Check Platform Requirements (raw output)

```go
output, err := comp.CheckPlatformReqs()
fmt.Println(output)
```

### Check Lock Platform Requirements (raw output)

```go
output, err := comp.CheckPlatformReqsLock()
fmt.Println(output)
```

## Diagnose / Check / Status

### Diagnose

Run Composer's built-in diagnostic tool to identify common problems:

```go
output, err := comp.Diagnose()
fmt.Println(output)
```

Composer checks connectivity, Git availability, HTTP settings, and more.

### Diagnose with Options

```go
output, err := comp.DiagnoseWithOptions(map[string]string{
    "no-interaction": "",
})
```

### Check Dependencies

Verify that `composer.json` and `composer.lock` are in sync:

```go
output, err := comp.Check()
fmt.Println(output)
```

### Check with Options

```go
output, err := comp.CheckWithOptions(map[string]string{
    "lock": "",
})
```

### Status

Show local modifications to installed packages:

```go
output, err := comp.Status()
fmt.Println(output)
```

### Status with Options

```go
output, err := comp.StatusWithOptions(map[string]string{
    "verbose": "",
})
```

## Environment Info

### Get Full Environment Information

Retrieve all Composer configuration as a map:

```go
info, err := comp.GetEnvironmentInfo()
if err != nil {
    log.Fatal(err)
}
for key, value := range info {
    fmt.Printf("%s: %s\n", key, value)
}
```

This returns all config values from `composer config --list`.

### Get Composer Executable Path

Find the Composer binary without creating an instance:

```go
path, err := composer.GetComposerPath()
if err != nil {
    fmt.Println("Composer not found:", err)
} else {
    fmt.Println("Composer path:", path)
}
```

### Get Working Directory and Environment

```go
comp, _ := composer.New(composer.DefaultOptions())

// Get current working directory
dir := comp.GetWorkingDir()

// Get configured environment variables
env := comp.GetEnv()

// Get executable path
execPath := comp.GetExecutablePath()

// Check if Composer is installed
if comp.IsInstalled() {
    fmt.Println("Composer is available at:", execPath)
}
```

## Self-Update

Update Composer to the latest version:

```go
err := comp.SelfUpdate()
if err != nil {
    log.Fatal("Self-update failed:", err)
}
fmt.Println("Composer updated successfully")
```

### Get Current Version

```go
version, err := comp.GetVersion()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Composer version:", version)
```

Equivalent Composer commands:

```bash
composer self-update
composer --version
```

## Common Diagnostic Workflows

### Full System Check

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

// 1. Check Composer version
version, _ := comp.GetVersion()
fmt.Printf("Composer: %s\n", version)

// 2. Check PHP version
php, _ := comp.GetPHPVersion()
fmt.Printf("PHP: %s\n", php)

// 3. Check critical extensions
for _, ext := range []string{"json", "mbstring", "xml", "curl"} {
    has, _ := comp.HasExtension(ext)
    fmt.Printf("ext-%s: %v\n", ext, has)
}

// 4. Check platform requirements
platforms, _ := comp.CheckPlatform()
for _, p := range platforms {
    if !p.Available {
        fmt.Printf("WARNING: %s not available\n", p.Name)
    }
}

// 5. Run diagnostics
output, _ := comp.Diagnose()
fmt.Println(output)
```

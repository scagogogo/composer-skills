# Composer Skills

<p align="center">
  <img src="docs/images/header-badge.png" alt="Composer Skills" width="100%">
</p>

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/composer-skills.svg)](https://pkg.go.dev/github.com/scagogogo/composer-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/composer-skills)](https://goreportcard.com/report/github.com/scagogogo/composer-skills)
[![Tests](https://github.com/scagogogo/composer-skills/actions/workflows/go-tests.yml/badge.svg)](https://github.com/scagogogo/composer-skills/actions/workflows/go-tests.yml)
[![License: MIT](https://img.shields.io/github/license/scagogogo/composer-skills)](LICENSE)

**The missing Go SDK for the PHP Composer ecosystem** — Stop parsing `exec.Command` output by hand. One import gives you a typed, tested API for both the Packagist REST API and every Composer CLI command, plus zero-config auto-installation.

[简体中文](README-zh-CN.md) | English

---

## The Problem It Solves

If you write Go code that touches the PHP/Composer world, you've probably done this:

```go
// 😩 The old way — fragile, untyped, no error handling
out, _ := exec.Command("composer", "audit").Output()
lines := strings.Split(string(out), "\n")
// Now try to parse that...
```

Composer Skills gives you this instead:

```go
// 😊 The new way — typed, tested, auto-installing
result, _ := comp.AuditWithJSON()
fmt.Printf("Vulnerabilities: %d\n", result.Found)
```

**What you get:**

| Pain Point | Solution |
|---|---|
| Raw `exec.Command` output parsing | **234 typed Go methods** with structured results |
| Hand-written HTTP requests to Packagist | **20 typed API methods** returning Go structs |
| "Is Composer installed on this machine?" | **Cross-OS detector** finds it anywhere |
| "Composer isn't installed, now what?" | **Auto-installer** downloads Composer + PHP automatically |
| Different code for different OS | **Smart defaults** per platform (brew, apt, direct download) |

> **One import, full power:** `go get github.com/scagogogo/composer-skills`

---

## Architecture

<p align="center">
  <img src="docs/images/architecture.png" alt="Three-Layer Architecture" width="90%">
</p>

| Layer | What it does | Package |
|-------|-------------|---------|
| **Skills Documentation** | Progressive disclosure guides (12 guides) | `docs/skills/` |
| **CLI Tool** | 50+ subcommands via Cobra | `cmd/composer-skills/` |
| **Packagist API SDK** | HTTP calls to Packagist (pure Go) | `pkg/client`, `pkg/repository` |
| **Composer CLI SDK** | Executes local `composer` binary (234 methods) | `pkg/composer` |
| **Foundation** | Domain models, detection, installation, utilities | `pkg/domain`, `pkg/detector`, `pkg/installer`, `pkg/composerutils` |

---

## Two SDKs in One

<p align="center">
  <img src="docs/images/sdk-comparison.png" alt="SDK Comparison" width="90%">
</p>

| | Packagist API SDK | Composer CLI SDK |
|---|---|---|
| **Package** | `pkg/client`, `pkg/repository` | `pkg/composer` |
| **How it works** | HTTP calls to Packagist API | Executes local `composer` binary |
| **Requires PHP?** | **No** (pure Go) | Yes (PHP 7.4+ and Composer 2.0+) |
| **Use cases** | Search packages, get stats, security advisories | Install/update deps, manage projects, audit, run scripts |

---

## Feature Map

<p align="center">
  <img src="docs/images/feature-mindmap.png" alt="Feature Tree" width="100%">
</p>

---

## Auto-Install: Zero Config

<p align="center">
  <img src="docs/images/auto-install-flow.png" alt="Auto-Install Flow" width="90%">
</p>

Composer Skills handles the entire setup chain — detect Composer → check PHP → install if missing → verify → ready. Default options have `AutoInstall: true`, so it just works:

```go
// That's it. If Composer is missing, it gets installed automatically.
comp, err := composer.New(composer.DefaultOptions())
```

**Platform support:**

<p align="center">
  <img src="docs/images/platform-matrix.png" alt="Platform Support" width="90%">
</p>

---

## Security-First

<p align="center">
  <img src="docs/images/security-features.png" alt="Security Features" width="90%">
</p>

```go
// Local audit with structured results
result, _ := comp.AuditWithJSON()
if result.Found > 0 {
    for _, v := range result.Advisories {
        fmt.Printf("⚠ %s: %s (%s)\n", v.Package, v.Title, v.Severity)
    }
}

// Remote advisories from Packagist
advisories, _ := client.GetSecurityAdvisories()

// Validate composer.json before committing
result, _ := comp.ValidateStructured()
```

---

## ✨ Key Features

- **Full Composer CLI Coverage** — 234 SDK methods wrapping every standard Composer command across 20 categories
- **Packagist API Client** — 20 methods to search, browse, and query the PHP package registry from Go (pure Go, no PHP)
- **Security-First** — Audit dependencies, check vulnerabilities, validate schemas, check platform requirements
- **Auto-Detection & Installation** — Automatically finds or installs Composer (with cross-OS detection and PHP auto-install)
- **Cross-Platform** — Windows, macOS, and Linux support with smart defaults
- **CLI Tool** — 50+ subcommands exposing all SDK capabilities from the terminal
- **Structured Results** — Type-safe return values (AuditInfo, OutdatedInfo, VersionInfo, etc.) instead of raw strings
- **Convenience Methods** — `IsPackageInstalled`, `GetDirectDependencyNames`, `GetProjectSummary`, and 18 more helpers
- **Progressive Docs** — From 3-line quickstart to full API reference (12 guides)
- **Well-Tested** — 450+ tests with mock-based isolation

---

## 🚀 Quick Start

### Install

```bash
go get github.com/scagogogo/composer-skills
```

### Packagist API (No PHP Required)

```go
package main

import (
    "fmt"
    "time"
    "github.com/scagogogo/composer-skills/pkg/client"
)

func main() {
    c := client.NewComposerClient(30 * time.Second)

    // Search for packages
    results, _ := c.SearchPackages("logging", 10, 1)
    fmt.Printf("Found %d packages\n", results.Total)

    // Get package details
    pkg, _ := c.GetPackage("monolog/monolog")
    fmt.Printf("%s: %s\n", pkg.Package.Name, pkg.Package.Description)

    // Security advisories
    advisories, _ := c.GetSecurityAdvisories()
    fmt.Printf("%d advisories\n", len(advisories.Advisories))

    // Statistics
    stats, _ := c.GetStatistics()
    fmt.Printf("Total packages: %d\n", stats.Packages)
}
```

### Composer CLI Wrapper (Requires PHP + Composer)

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/composer-skills/pkg/composer"
)

func main() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatal(err)
    }
    comp.SetWorkingDir("/path/to/php/project")

    // Dependency management
    comp.Install(false, true)
    comp.RequirePackage("monolog/monolog", "^3.0", false)
    comp.Update([]string{}, false)

    // Security audit (structured results)
    result, _ := comp.AuditWithJSON()
    fmt.Printf("Vulnerabilities found: %d\n", result.Found)

    // Package inspection
    output, _ := comp.ShowDependencyTree("symfony/console")
    output, _ = comp.WhyPackage("symfony/polyfill-mbstring")
    output, _ = comp.OutdatedPackages()

    // Platform checks
    phpVer, _ := comp.GetPHPVersion()
    hasExt, _ := comp.HasExtension("mbstring")

    // Auth & configuration
    comp.AddGitHubToken("github.com", "your-token")
    config, _ := comp.GetAuthConfig()
}
```

### Auto-Install Composer

```go
package main

import (
    "fmt"
    "github.com/scagogogo/composer-skills/pkg/installer"
    "github.com/scagogogo/composer-skills/pkg/detector"
)

func main() {
    // Detect if Composer is installed
    d := detector.NewDetector()
    if d.IsInstalled() {
        path, _ := d.Detect()
        fmt.Printf("Composer found at: %s\n", path)
        return
    }

    // Auto-install with smart OS detection (can also auto-install PHP)
    inst := installer.NewInstaller(installer.SmartConfig())
    if err := inst.Install(); err != nil {
        fmt.Printf("Install failed: %v\n", err)
    }
}
```

### Convenience Methods

```go
// Quick helpers that combine multiple operations
isInstalled := comp.IsPackageInstalled("monolog/monolog")
isDev := comp.IsPackageDev("monolog/monolog")
deps := comp.GetDirectDependencyNames()
summary := comp.GetProjectSummary()
hasLock := comp.HasComposerLock()
hasVendor := comp.HasVendorDir()
abandoned := comp.GetAbandonedPackagesFromLock()
namespaces := comp.GetNamespaceMap()
scripts := comp.GetScripts()
```

### CLI Tool

```bash
# Build
make build

# Packagist operations (no PHP needed)
./bin/composer-skills search query "logging"
./bin/composer-skills package info symfony/console
./bin/composer-skills security advisories
./bin/composer-skills repo stats

# Local Composer operations
./bin/composer-skills local install --working-dir /path/to/project
./bin/composer-skills local audit --working-dir /path/to/project
./bin/composer-skills local outdated --working-dir /path/to/project
./bin/composer-skills local why monolog/monolog --working-dir /path/to/project
./bin/composer-skills local fund --working-dir /path/to/project
./bin/composer-skills local global require phpstan/phpstan --version "^1.0"
```

---

## 📋 SDK Coverage

### Packagist API (20 methods)

| Category | Methods |
|----------|---------|
| Package Info | `GetPackage` · `GetPackageStats` · `GetPackageWithV2Metadata` · `GetPackageDevVersions` · `GetPackageChanges` |
| Search | `SearchPackages` · `SearchPackagesByTags` · `SearchPackagesByType` |
| Statistics | `GetStatistics` |
| Security | `GetSecurityAdvisories` · `GetSecurityAdvisoriesForPackages` · `GetSecurityAdvisoriesSince` |
| Listing | `ListPackages` · `ListPackagesByVendor` · `ListPackagesByType` · `ListPackagesWithData` · `ListPopularPackages` |
| Management | `CreatePackage` · `EditPackage` · `UpdatePackage` |

### Composer CLI (234 methods across 20 categories)

| Category | Count | Highlights |
|----------|-------|------------|
| Core | 10 | `Run`, `RunWithContext`, `RunWithTimeout`, `GetVersion`, `SelfUpdate` |
| Dependencies | 16 | `Install`, `Update`, `DumpAutoload`, `Suggests` + variants |
| Packages | 20 | `Require`, `Remove`, `Reinstall`, `Bump`, `Search`, `Show`, `Why`, `WhyNot` + variants |
| Audit | 10 | `Audit`, `AuditWithJSON`, `HasVulnerabilities`, `GetHighSeverityVulnerabilities` |
| Project | 10 | `CreateProject`, `InitProject`, `RunScript`, `ListScripts`, `GetProjectInfo` |
| Config | 12 | `GetConfig`, `SetConfig`, `ListConfig`, `ClearCache`, `GetComposerHome` |
| Validation | 14 | `Validate`, `ValidateStrict`, `ValidateSchema`, `NormalizeComposerJson` |
| Platform | 8 | `CheckPlatform`, `GetPHPVersion`, `GetExtensions`, `HasExtension` |
| Repository | 18 | `AddVcsRepository`, `AddComposerRepository`, `SetMinimumStability` |
| Global | 14 | `GlobalRequire`, `GlobalUpdate`, `GlobalRemove`, `GlobalInstall` |
| Auth | 10 | `AddGitHubToken`, `AddGitLabToken`, `AddBearerToken`, `GetAuthConfig` |
| Fund | 7 | `Fund`, `FundWithJSON`, `HasFunding`, `GetFundingURLs` |
| Licenses | 4 | `Licenses`, `LicensesWithFormat`, `CheckLicenses` |
| Diagnosis | 8 | `Diagnose`, `Check`, `Status`, `LocalExec` |
| Exec | 8 | `Exec`, `ExecCommand`, `ExecPHP`, `ExecWithList` |
| Satis | 8 | `InitSatis`, `CreateSatisConfig`, `BuildSatis` |
| Version | 5 | `GetPackageVersions`, `LockPackageVersion`, `UpdatePackageVersion` |
| Environment | 12 | `GetEnvironmentInfo`, `SetMemoryLimit`, `EnableDev`, `DisableInteraction` |
| composer.json | 10 | `ReadComposerJSON`, `WriteComposerJSON`, `AddRequire`, `AddScript`, `AddAutoload` |
| Archive | 6 | `Archive`, `ArchiveWithFormat`, `ArchivePackage` |

### Structured Result Types

Instead of parsing raw CLI output, Composer Skills provides typed results:

```go
// Audit results
auditInfo, _ := comp.GetAuditInfo()

// Outdated packages
outdatedInfo, _ := comp.GetOutdatedInfo()

// Version information
versionInfo, _ := comp.GetVersionInfo()

// Validation results
validateResult, _ := comp.ValidateStructured()

// Platform requirements
platformReqs, _ := comp.CheckPlatformReqsStructured()

// License information
licensesInfo, _ := comp.GetLicensesInfo()

// Configuration
configInfo, _ := comp.GetConfigStructured()

// Search results
searchInfo, _ := comp.SearchInfo("monolog")

// Diagnose results
diagnoseInfo, _ := comp.DiagnoseStructured()
```

---

## 🏗️ Project Structure

```
composer-skills/
├── cmd/composer-skills/        # CLI tool (Cobra-based, 50+ commands)
├── pkg/
│   ├── client/                 # Packagist HTTP API client (20 methods)
│   ├── domain/                 # Data models (Package, Advisory, Statistics, Version...)
│   ├── repository/             # Repository operations layer
│   ├── composer/               # Composer CLI wrapper SDK (234 methods, 20 categories)
│   ├── detector/               # Composer installation detection (cross-OS)
│   ├── installer/              # Composer auto-installation (OS-aware, PHP auto-install)
│   └── composerutils/          # Shared utilities (filesystem, HTTP, mock helpers)
├── examples/                   # 15+ example programs
├── docs/
│   ├── skills/                 # Progressive disclosure documentation (12 guides)
│   └── images/                 # Generated diagrams and visual assets
└── Makefile
```

---

## 📖 Documentation

### Progressive Disclosure Guides

| Guide | Level | Description |
|-------|-------|-------------|
| [Getting Started](docs/skills/01-getting-started.md) | 🟢 Beginner | Installation and first steps |
| [Packagist API](docs/skills/02-packagist-api.md) | 🟢 Beginner | Remote API operations (search, stats, advisories) |
| [Dependency Management](docs/skills/03-dependency-management.md) | 🟡 Intermediate | Install, update, require, remove |
| [Project Management](docs/skills/04-project-management.md) | 🟡 Intermediate | Create, init, scripts, archive |
| [Security](docs/skills/05-security.md) | 🔴 Advanced | Audit, vulnerabilities, validation |
| [Package Inspection](docs/skills/06-package-inspection.md) | 🟡 Intermediate | Show, tree, why, fund, licenses |
| [Configuration](docs/skills/07-configuration.md) | 🔴 Advanced | composer.json, config, auth, repos |
| [Global Operations](docs/skills/08-global-operations.md) | 🟡 Intermediate | Global require, update, remove |
| [Platform & Diagnosis](docs/skills/09-platform-and-diagnosis.md) | 🔴 Advanced | PHP, extensions, diagnose |
| [Advanced](docs/skills/10-advanced.md) | 🔴 Advanced | Satis, exec, version constraints |
| [CLI Reference](docs/skills/11-cli-reference.md) | 📖 Reference | Complete command reference |

---

## Use Cases

Composer Skills is designed for anyone who needs to interact with the PHP/Composer ecosystem from Go:

- **CI/CD pipelines** — Automate `composer install`, run security audits, check for outdated packages
- **Security scanners** — Query Packagist advisories, audit dependencies, check platform requirements
- **Package mirrors** — Download package indexes, list packages, get statistics from Packagist
- **Dependency dashboards** — Show dependency trees, check licenses, track funding, monitor outdated packages
- **DevOps automation** — Auto-detect and install Composer, manage global packages, configure auth tokens
- **Satis builders** — Initialize, configure, and build private Composer repositories

---

## 🧪 Testing

```bash
make test           # Run all tests
make test-race      # Run with race detection
make test-coverage  # Generate coverage report
make check          # Format + vet + test
```

---

## 📋 Requirements

- **Go 1.23+**
- For CLI wrapper: PHP 7.4+ and Composer 2.0+
- For Packagist API client: No external dependencies (pure Go)

---

## 🙏 Acknowledgments

- [Packagist](https://packagist.org/) — The PHP package repository
- [Composer](https://getcomposer.org/) — Dependency Manager for PHP

---

## 📄 License

[MIT](LICENSE)

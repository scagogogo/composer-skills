# Getting Started

Install the library and start using it in under a minute.

## Installation

```bash
go get github.com/scagogogo/composer-skills
```

Requires Go 1.23 or later.

## Basic Usage -- Packagist API Client

The Packagist API client lets you query the Packagist repository without PHP. Three lines to get started:

```go
package main

import (
    "fmt"
    "time"
    "github.com/scagogogo/composer-skills/pkg/client"
)

func main() {
    c := client.NewComposerClient(30 * time.Second)
    pkg, _ := c.GetPackage("symfony/console")
    fmt.Println(pkg.Package.Name, pkg.Package.Downloads.Total)
}
```

## Basic Usage -- Composer CLI Wrapper

The CLI wrapper executes your local `composer` binary from Go. It auto-detects and auto-installs Composer if not found:

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
    comp.Install(false, true)
    version, _ := comp.GetVersion()
    fmt.Println("Composer version:", version)
}
```

## Basic Usage -- CLI Tool

The project also ships a standalone CLI tool:

```bash
# Build the CLI tool
go build -o composer-skills ./cmd/composer-skills

# Get package info
composer-skills package info symfony/console

# Search packages
composer-skills search query "logging"

# Run local Composer operations
composer-skills local install --no-dev --optimize
composer-skills local require monolog/monolog --version "^3.0"
```

## What You Can Do

### Packagist API Client (Remote, No PHP Required)

- Get package information, stats, V2 metadata, dev versions
- Search packages by query, tags, or type
- Get security advisories for all packages, specific packages, or since a timestamp
- List all packages, by vendor, by type, or popular packages
- Get repository statistics
- Create, edit, and update packages (requires API credentials)

### Composer CLI Wrapper (Local, Requires PHP + Composer)

- Install, update, require, and remove dependencies
- Create and initialize projects
- Run scripts and list available scripts
- Audit dependencies for security vulnerabilities
- Show package info, dependency trees, and outdated packages
- Configure repositories, authentication, and settings
- Manage global packages
- Check platform requirements and diagnose issues
- Validate composer.json and composer.lock
- Generate shell completion scripts
- Build Satis private repositories

## Next Steps

- [Packagist API](02-packagist-api.md) -- Learn the full Packagist API client
- [Dependency Management](03-dependency-management.md) -- Manage project dependencies
- [Security](05-security.md) -- Audit and secure your project

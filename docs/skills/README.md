# Composer Skills Documentation

This directory contains progressive disclosure documentation for the Composer Skills library -- a comprehensive Go toolkit for the PHP Composer ecosystem.

## What is Composer Skills?

Composer Skills provides two complementary capabilities:

- **Packagist API Client** (`pkg/client`, `pkg/repository`) -- Pure Go HTTP client for the Packagist API. No PHP required.
- **Composer CLI Wrapper** (`pkg/composer`) -- Go wrapper around the local `composer` binary for managing PHP projects. Requires PHP and Composer.

## Documentation Index

| Document | Description |
|----------|-------------|
| [01 - Getting Started](01-getting-started.md) | Installation, basic usage, and what you can do |
| [02 - Packagist API](02-packagist-api.md) | Package info, search, security advisories, repository operations |
| [03 - Dependency Management](03-dependency-management.md) | Install, require, remove, update, bump, reinstall, dump-autoload |
| [04 - Project Management](04-project-management.md) | Create project, init, run scripts, archive |
| [05 - Security](05-security.md) | Audit, vulnerability checks, platform requirements, validation |
| [06 - Package Inspection](06-package-inspection.md) | Show, dependency tree, why/why-not, outdated, fund, licenses |
| [07 - Configuration](07-configuration.md) | composer.json, config get/set, repositories, auth, environment |
| [08 - Global Operations](08-global-operations.md) | Global require/update/remove, install, list, status |
| [09 - Platform and Diagnosis](09-platform-and-diagnosis.md) | PHP version, extensions, platform reqs, diagnose, self-update |
| [10 - Advanced](10-advanced.md) | Satis, exec, version constraints, browse, completion, normalization |
| [11 - CLI Reference](11-cli-reference.md) | Complete CLI command reference with all flags |

## How to Read This Documentation

Each document follows a progressive disclosure pattern:

1. **Overview** -- What the feature does and why you would use it
2. **Simplest example** -- The minimum code to get something working
3. **More complex examples** -- Additional options and configurations
4. **Both SDK and CLI** -- Every feature shows Go SDK and CLI usage

## Quick Links

- Go package: `github.com/scagogogo/composer-skills`
- Packagist API client: `pkg/client`
- Composer CLI wrapper: `pkg/composer`
- CLI tool: `cmd/composer-skills`
- Examples: `examples/`

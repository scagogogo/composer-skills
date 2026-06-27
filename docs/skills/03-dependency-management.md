# Dependency Management Skills

Manage your PHP project's dependencies using the Composer CLI wrapper (`pkg/composer`).

## Installing Dependencies

### Basic Install

Install all dependencies defined in `composer.json`:

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

// Install all dependencies (including dev)
err := comp.Install(false, false)
```

CLI:

```bash
composer-skills local install --working-dir /path/to/project
```

### Install with Options

```go
// Production only, with optimized autoloader
err := comp.Install(true, true) // noDev=true, optimize=true
```

CLI:

```bash
composer-skills local install --no-dev --optimize --working-dir /path/to/project
```

### Install with Custom Options

Use `InstallWithOptions` for full control over flags:

```go
err := comp.InstallWithOptions(map[string]string{
    "no-dev":             "",
    "optimize-autoloader": "",
    "prefer-dist":        "",
    "no-progress":        "",
    "no-interaction":     "",
})
```

Equivalent Composer command:

```bash
composer install --no-dev --optimize-autoloader --prefer-dist --no-progress --no-interaction
```

## Adding Packages

### Require a Package

```go
// Add a production dependency
err := comp.RequirePackage("monolog/monolog", "^3.0", false)

// Add without specifying version (uses latest)
err := comp.RequirePackage("monolog/monolog", "", false)
```

CLI:

```bash
composer-skills local require monolog/monolog --version "^3.0"
```

### Require Dev Dependency

```go
// Add as a development dependency
err := comp.RequirePackage("phpunit/phpunit", "^10.0", true)
```

CLI:

```bash
composer-skills local require phpunit/phpunit --version "^10.0" --dev
```

### Require with Options

Use `RequirePackageWithOptions` for additional flags:

```go
err := comp.RequirePackageWithOptions("symfony/console", "^6.0", map[string]string{
    "dev":           "",
    "prefer-source": "",
    "no-update":     "",
})
```

Equivalent Composer command:

```bash
composer require --dev --prefer-source --no-update symfony/console:^6.0
```

## Removing Packages

### Remove a Single Package

```go
// Remove from production dependencies
err := comp.Remove("symfony/console", false)

// Remove from dev dependencies
err := comp.Remove("phpunit/phpunit", true)
```

CLI:

```bash
composer-skills local remove symfony/console
composer-skills local remove phpunit/phpunit --dev
```

### Remove Multiple Packages

Call `Remove` for each package, or use the generic `Run` method:

```go
// Remove multiple packages at once
_, err := comp.Run("remove", "pkg/one", "pkg/two", "pkg/three")
```

## Updating Packages

### Update All Dependencies

```go
// Update all dependencies
err := comp.Update([]string{}, false)

// Update all, skip dev dependencies
err := comp.Update([]string{}, true)
```

CLI:

```bash
composer-skills local update --working-dir /path/to/project
composer-skills local update --no-dev --working-dir /path/to/project
```

### Update Specific Packages

```go
err := comp.Update([]string{"symfony/console", "symfony/process"}, false)
```

CLI:

```bash
composer-skills local update symfony/console symfony/process
```

### Update with Options

```go
err := comp.UpdateWithOptions([]string{"symfony/console"}, map[string]string{
    "no-dev":           "",
    "prefer-dist":      "",
    "with-dependencies": "",
    "no-progress":      "",
})
```

Equivalent Composer command:

```bash
composer update --no-dev --prefer-dist --with-dependencies --no-progress symfony/console
```

## Bumping Packages

Bump packages to the latest version that satisfies the version constraint in `composer.json`:

```go
// Bump specific packages
err := comp.BumpPackages([]string{"symfony/console", "symfony/process"})

// Bump all packages
err := comp.BumpPackages([]string{})
```

### Bump with Options

```go
err := comp.BumpPackagesWithOptions([]string{"symfony/console"}, map[string]string{
    "dev-only":      "",
    "prefer-stable": "",
    "dry-run":       "",
})
```

Equivalent Composer command:

```bash
composer bump --dev-only --prefer-stable --dry-run symfony/console
```

## Reinstalling Packages

Remove and re-add a package to force a clean reinstall:

```go
err := comp.Reinstall("symfony/console")
```

This is equivalent to:

```go
comp.Remove("symfony/console", false)
comp.RequirePackage("symfony/console", "", false)
```

## Dumping Autoload

Regenerate the Composer autoloader files after changes to class structure.

### Basic Dump

```go
// Standard autoload generation
err := comp.DumpAutoload(false)

// Optimized autoload (generates class map for faster loading)
err := comp.DumpAutoload(true)
```

### Dump with Options

```go
err := comp.DumpAutoloadWithOptions(map[string]string{
    "optimize":              "",
    "classmap-authoritative": "",
    "apcu":                  "",
    "no-dev":                "",
})
```

Equivalent Composer command:

```bash
composer dump-autoload --optimize --classmap-authoritative --apcu --no-dev
```

## Checking Dependencies

Check for dependency conflicts and sync issues:

```go
output, err := comp.CheckDependencies()
fmt.Println(output)
```

## Suggested Packages

View packages suggested by your installed dependencies:

```go
err := comp.Suggests()
```

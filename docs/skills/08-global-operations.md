# Global Operations Skills

Manage globally installed Composer packages and the Composer home directory.

## Global Require

Install a package globally, making its binaries available system-wide:

```go
comp, _ := composer.New(composer.DefaultOptions())

// Install a global package
err := comp.GlobalRequire("phpstan/phpstan", "^1.0")

// Install without version constraint (uses latest)
err = comp.GlobalRequire("phpcs/phpcs-sniffs", "")
```

Equivalent Composer command:

```bash
composer global require phpstan/phpstan:^1.0
```

## Global Update

Update globally installed packages:

```go
// Update all global packages
err := comp.GlobalUpdate([]string{})

// Update specific global packages
err = comp.GlobalUpdate([]string{"phpstan/phpstan"})
```

Equivalent Composer command:

```bash
composer global update
composer global update phpstan/phpstan
```

## Global Remove

Remove a globally installed package:

```go
err := comp.GlobalRemove("phpstan/phpstan")
```

Equivalent Composer command:

```bash
composer global remove phpstan/phpstan
```

## Global Install

Install dependencies in the global Composer directory (useful after manually editing `~/.composer/composer.json`):

```go
err := comp.GlobalInstall()
```

Equivalent Composer command:

```bash
composer global install
```

## Global List

List all globally installed packages:

```go
output, err := comp.GlobalList()
fmt.Println(output)
```

Equivalent Composer command:

```bash
composer global show
```

## Global Status

Show the status of globally installed packages (e.g., local modifications):

```go
output, err := comp.GlobalStatus()
fmt.Println(output)
```

Equivalent Composer command:

```bash
composer global status
```

## Global Dump-Autoload

Regenerate the autoloader for global packages:

```go
// Standard autoload
err := comp.GlobalDumpAutoload(false)

// Optimized autoload
err = comp.GlobalDumpAutoload(true)
```

Equivalent Composer command:

```bash
composer global dump-autoload
composer global dump-autoload --optimize
```

## Global Execute

Run a binary from a globally installed package:

```go
output, err := comp.GlobalExecute("phpstan", "analyse", "--level=5", "src/")
fmt.Println(output)
```

Equivalent Composer command:

```bash
composer global exec phpstan -- analyse --level=5 src/
```

## Composer Home Directory

### Get Composer Home

Find the Composer home directory (typically `~/.composer` or `~/.config/composer`):

```go
home, err := comp.GetComposerHome()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Composer home:", home)
```

### Global Home

Open the Composer home directory in a browser (if configured):

```go
output, err := comp.GlobalHome()
fmt.Println(output)
```

## Common Global Workflows

### Install and Use a Global Tool

```go
comp, _ := composer.New(composer.DefaultOptions())

// Install PHPStan globally
comp.GlobalRequire("phpstan/phpstan", "^1.0")

// Run it
output, _ := comp.GlobalExecute("phpstan", "analyse", "src/")
fmt.Println(output)
```

### Update All Global Packages

```go
comp, _ := composer.New(composer.DefaultOptions())
err := comp.GlobalUpdate([]string{})
if err != nil {
    log.Fatal(err)
}
```

### Clean Up Global Packages

```go
comp, _ := composer.New(composer.DefaultOptions())

// Remove unused tools
comp.GlobalRemove("old/tool-one")
comp.GlobalRemove("old/tool-two")

// Regenerate autoload
comp.GlobalDumpAutoload(true)
```

# Advanced Skills

Satis private repositories, exec commands, version constraints, browse, completion, composer.json manipulation, archive operations, and normalization.

## Satis (Private Composer Repository)

Satis is a static Composer repository generator. Use it to create a private package repository.

### Initialize a Satis Repository

```go
comp, _ := composer.New(composer.DefaultOptions())

// Create a satis directory with default config
err := comp.InitSatis("My Private Repo", "https://packages.example.org", "satis")
```

This creates `satis/satis.json` with a basic configuration.

### Create a Satis Config

```go
err := comp.CreateSatisConfig("/path/to/satis.json", "My Private Repo", "https://packages.example.org")
```

### Add a Repository to Satis

```go
err := comp.AddSatisRepository("/path/to/satis.json", "vcs", "https://github.com/vendor/private-package")
err = comp.AddSatisRepository("/path/to/satis.json", "composer", "https://internal.packages.org")
```

### Add a Package Requirement

```go
err := comp.AddSatisRequire("/path/to/satis.json", "vendor/package", "^1.0")
```

This disables `require-all` and adds the specific package requirement.

### Update Satis Stability

```go
err := comp.UpdateSatisStability("/path/to/satis.json", "beta")
```

Valid values: `dev`, `alpha`, `beta`, `RC`, `stable`.

### Enable Satis Archive

Enable distribution archive generation:

```go
// Enable with default ZIP format
err := comp.EnableSatisArchive("/path/to/satis.json", "")

// Enable with TAR format
err = comp.EnableSatisArchive("/path/to/satis.json", "tar")
```

### Build the Satis Repository

```go
// Build to default output directory
output, err := comp.BuildSatis("/path/to/satis.json", "")

// Build to a specific directory
output, err = comp.BuildSatis("/path/to/satis.json", "/var/www/packages")
fmt.Println(output)
```

Equivalent Composer commands:

```bash
composer satis build /path/to/satis.json
composer satis build /path/to/satis.json /var/www/packages
```

## Exec Commands

Execute binaries installed via Composer.

### Execute a Binary

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

output, err := comp.Exec("phpunit", "--filter=UserTest")
fmt.Println(output)
```

### List Available Binaries

```go
binaries, err := comp.ExecWithList()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Available binaries:")
for _, bin := range binaries {
    fmt.Println("  - " + bin)
}
```

### Execute with Specific PHP Binary

```go
output, err := comp.ExecPHP("/usr/bin/php8.1", "phpunit", "--version")
```

### Execute in a Different Working Directory

```go
output, err := comp.ExecWithWorkingDir("phpunit", "/other/project", "--testsuite=unit")
```

### Execute All Binaries

Run all available binaries and collect results:

```go
results, err := comp.ExecAll("--version")
for binary, output := range results {
    fmt.Printf("%s: %s\n", binary, output)
}
```

### Execute a Custom Command

```go
output, err := comp.ExecCommand("custom-command", "arg1", "arg2")
```

### Local Exec (from diagnosis package)

```go
output, err := comp.LocalExec("phpstan", "analyse", "src/")

// With options
output, err = comp.LocalExecWithOptions("phpstan", map[string]string{
    "verbose": "",
}, "analyse", "src/")
```

Equivalent Composer command:

```bash
composer exec phpunit -- --filter=UserTest
composer exec --list
```

## Version Constraints

### Formatting Version Constraints

The SDK provides a `FormatVersionConstraint` helper to generate version constraint strings:

```go
// Exact: "1.2.3"
v := composer.FormatVersionConstraint("1.2.3", composer.ExactVersion)

// Caret: "^1.2.3" (>=1.2.3, <2.0.0)
v = composer.FormatVersionConstraint("1.2.3", composer.CaretVersion)

// Tilde: "~1.2.3" (>=1.2.3, <1.3.0)
v = composer.FormatVersionConstraint("1.2.3", composer.TildeVersion)

// Wildcard: "1.2.*"
v = composer.FormatVersionConstraint("1.2", composer.WildcardVersion)

// Range: ">=1.2.0 <2.0.0"
v = composer.FormatVersionConstraint("1.2", composer.RangeVersion)
```

### Update a Package Version Constraint

```go
// Update to a caret constraint
err := comp.UpdatePackageVersion("symfony/console", "6.0", composer.CaretVersion)
// Results in: composer require symfony/console:^6.0

// Update to a tilde constraint
err = comp.UpdatePackageVersion("symfony/console", "6.0.0", composer.TildeVersion)
// Results in: composer require symfony/console:~6.0.0
```

### Lock a Package to an Exact Version

```go
err := comp.LockPackageVersion("symfony/console", "6.0.3")
// Results in: composer require symfony/console:6.0.3
```

### Get Available Versions

```go
output, err := comp.GetPackageVersions("symfony/console")
fmt.Println(output)
```

### Version Constraint Types

| Type | Format | Example | Meaning |
|------|--------|---------|---------|
| Exact | `1.2.3` | `1.2.3` | Exactly this version |
| Caret | `^1.2.3` | `>=1.2.3 <2.0.0` | Compatible with major version |
| Tilde | `~1.2.3` | `>=1.2.3 <1.3.0` | Compatible with minor version |
| Wildcard | `1.2.*` | `>=1.2.0 <1.3.0` | Any patch version |
| Range | `>=1.2.0 <2.0.0` | Explicit range | Custom range |

## Browse Packages

Open a package's homepage in the default browser:

```go
comp, _ := composer.New(composer.DefaultOptions())

// Open the package homepage
err := comp.BrowsePackage("symfony/console")

// Open with options (e.g., docs or issues page)
err = comp.BrowsePackageWithOptions("symfony/console", map[string]string{
    "docs": "",
})
err = comp.BrowsePackageWithOptions("symfony/console", map[string]string{
    "issues": "",
})
```

Equivalent Composer command:

```bash
composer browse symfony/console
composer browse --docs symfony/console
```

## Completion Generation

Generate shell completion scripts for Composer:

```go
comp, _ := composer.New(composer.DefaultOptions())

// Bash completion
output, err := comp.GenerateCompletion(composer.BashShell)

// Zsh completion
output, err = comp.GenerateCompletion(composer.ZshShell)

// Fish completion
output, err = comp.GenerateCompletion(composer.FishShell)

// With options
output, err = comp.GenerateCompletionWithOptions(composer.BashShell, map[string]string{
    "no-interaction": "",
})
```

### List All Commands

```go
output, err := comp.ListCommands()
fmt.Println(output)
```

### Get Command Help

```go
output, err := comp.GetCommandHelp("require")
fmt.Println(output)
```

Equivalent Composer command:

```bash
composer completion bash
composer completion zsh
composer completion fish
composer list
composer help require
```

## composer.json Manipulation

Directly manipulate the `composer.json` file using the `ComposerJSON` struct.

### Add/Remove Require

```go
// Add a production dependency to composer.json
err := comp.AddRequire("symfony/console", "^6.0", false)

// Add a dev dependency
err = comp.AddRequire("phpunit/phpunit", "^10.0", true)

// Remove a production dependency
err = comp.RemoveRequire("symfony/console", false)

// Remove a dev dependency
err = comp.RemoveRequire("phpunit/phpunit", true)
```

Note: `AddRequire` and `RemoveRequire` only modify the file -- they do not install or uninstall the package. Use `RequirePackage` / `Remove` for that.

### Add/Remove Scripts

```go
// Add a single-command script
err := comp.AddScript("post-install-cmd", "php artisan optimize:clear", "Clear caches after install")

// Add a multi-command script
err = comp.AddScript("test", []string{"phpunit", "phpcs"}, "Run tests and code style checks")

// Remove a script
err = comp.RemoveScript("post-install-cmd")
```

### Add Autoload Configuration

```go
// Add PSR-4 autoload
err := comp.AddAutoload("psr-4", "App\\", "src/", false)

// Add PSR-4 autoload with multiple directories
err = comp.AddAutoload("psr-4", "App\\", []string{"src/", "lib/"}, false)

// Add dev autoload
err = comp.AddAutoload("psr-4", "Tests\\", "tests/", true)

// Add classmap autoload
err = comp.AddAutoload("classmap", "", "src/Classes/", false)

// Add files autoload
err = comp.AddAutoload("files", "", "src/helpers.php", false)
```

### Set Top-Level Properties

```go
// Set project name
err := comp.SetProperty("name", "myvendor/mypackage")

// Set description
err = comp.SetProperty("description", "A fantastic PHP library")

// Set type
err = comp.SetProperty("type", "library")

// Set keywords
err = comp.SetProperty("keywords", []string{"php", "library", "awesome"})

// Set homepage
err = comp.SetProperty("homepage", "https://example.org")

// Set license
err = comp.SetProperty("license", "MIT")

// Set minimum stability
err = comp.SetProperty("minimum-stability", "beta")

// Set prefer stable
err = comp.SetProperty("prefer-stable", true)
```

### Set Config Properties

```go
// Set process timeout
err := comp.SetConfig("process-timeout", 500)

// Set vendor directory
err = comp.SetConfig("vendor-dir", "vendor")

// Disable plugins
err = comp.SetConfig("disable-plugins", true)
```

## Archive Operations

Beyond the project archive described in [Project Management](04-project-management.md), the archive functionality provides several options:

### Basic Project Archive

```go
output, err := comp.Archive("./dist")
```

### Archive with Format

```go
output, err := comp.ArchiveWithFormat("./dist", "tar")
```

### Archive with Custom Options

```go
output, err := comp.ArchiveWithOptions("./dist", map[string]string{
    "format":          "tar",
    "file":            "my-project.tar",
    "ignore-filters":  "",
})
```

### Archive a Specific Package

```go
// Archive a specific version
output, err := comp.ArchivePackage("symfony/console", "v5.4.0", "./dist")

// Archive latest version
output, err = comp.ArchivePackage("symfony/console", "", "./dist")

// With options
output, err = comp.ArchivePackageWithOptions(
    "symfony/console", "v5.4.0", "./dist",
    map[string]string{"format": "tar"},
)
```

## Normalization

### Check Normalization

Verify that `composer.json` follows the standard format:

```go
output, err := comp.CheckNormalization()
if err != nil {
    fmt.Println("composer.json needs normalization")
} else {
    fmt.Println("composer.json is properly normalized")
}
```

### Normalize composer.json

Automatically format `composer.json` to the standard layout (requires the `ergebnis/composer-normalize` plugin):

```go
output, err := comp.NormalizeComposerJson()
if err != nil {
    if strings.Contains(err.Error(), "command not found") {
        fmt.Println("Install the normalize plugin first:")
        fmt.Println("  composer global require ergebnis/composer-normalize")
    } else {
        log.Fatal(err)
    }
}
```

Equivalent Composer command:

```bash
composer normalize
```

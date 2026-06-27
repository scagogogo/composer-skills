# Package Inspection Skills

Inspect installed packages, view dependency trees, check for updates, and analyze licenses.

## Showing Packages

### Show All Installed Packages

List all packages installed in the project:

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

output, err := comp.ShowAllPackages()
fmt.Println(output)
```

CLI:

```bash
composer-skills local show --working-dir /path/to/project
```

### Show a Specific Package

Display detailed information about an installed package:

```go
output, err := comp.ShowPackage("symfony/console")
fmt.Println(output)
```

CLI:

```bash
composer-skills local show symfony/console --working-dir /path/to/project
```

The output includes version, description, authors, dependencies, and installation path.

### Show with Format

Use the generic `Run` method to specify output format:

```go
output, err := comp.Run("show", "--format=json")
// Or for a specific package:
output, err = comp.Run("show", "--format=json", "symfony/console")
```

Equivalent Composer command:

```bash
composer show --format=json
composer show --format=json symfony/console
```

## Dependency Trees

### Show Full Dependency Tree

Display the complete dependency tree for the project:

```go
output, err := comp.ShowDependencyTree("")
fmt.Println(output)
```

### Show Package Dependency Tree

Display the dependency tree for a specific package:

```go
output, err := comp.ShowDependencyTree("symfony/console")
fmt.Println(output)
```

Equivalent Composer command:

```bash
composer show --tree
composer show --tree symfony/console
```

### Reverse Dependencies (depends)

Find which packages depend on a given package:

```go
output, err := comp.ShowReverseDependencies("symfony/polyfill-mbstring")
fmt.Println("Packages depending on symfony/polyfill-mbstring:", output)
```

Equivalent Composer command:

```bash
composer depends symfony/polyfill-mbstring
```

## Why / Why-Not Analysis

### Why Is a Package Installed?

Explain why a package is installed (which dependency requires it):

```go
output, err := comp.WhyPackage("symfony/polyfill-mbstring")
fmt.Println("Installed because:", output)
```

Equivalent Composer command:

```bash
composer why symfony/polyfill-mbstring
```

### Why Not a Specific Version?

Explain why a specific version of a package cannot be installed:

```go
output, err := comp.WhyNotPackage("symfony/console", "v4.0.0")
fmt.Println("Cannot install v4.0.0 because:", output)
```

Equivalent Composer command:

```bash
composer why-not symfony/console v4.0.0
```

## Outdated Packages

### Show All Outdated Packages

```go
output, err := comp.OutdatedPackages()
fmt.Println(output)
```

### Show Only Direct Outdated Dependencies

Only show packages directly listed in `composer.json`:

```go
output, err := comp.OutdatedPackagesDirect()
fmt.Println(output)
```

### Outdated with Custom Options

Use `CheckForOutdatedPackages` for more control:

```go
// All outdated packages in JSON format
output, err := comp.CheckForOutdatedPackages(false, false, "json")

// Only direct dependencies, minor updates only
output, err = comp.CheckForOutdatedPackages(true, true, "")
```

Equivalent Composer commands:

```bash
composer outdated
composer outdated --direct
composer outdated --direct --minor-only --format=json
```

## Fund Information

### Show Funding Information

Display funding and sponsorship information for your dependencies:

```go
output, err := comp.Fund()
fmt.Println(output)
```

### Fund with JSON

Get structured funding data:

```go
fundingInfo, err := comp.FundWithJSON()
for _, info := range fundingInfo {
    if info.Funding {
        fmt.Printf("Package: %s\n", info.Name)
        for _, url := range info.URLs {
            fmt.Printf("  Sponsor: %s\n", url)
        }
    }
}
```

### Fund for a Specific Package

```go
output, err := comp.FundWithPackage("symfony/console")
fmt.Println(output)
```

### Get All Funding URLs

Extract just the URLs as a map:

```go
urls, err := comp.GetFundingURLs()
for pkg, pkgURLs := range urls {
    fmt.Printf("%s:\n", pkg)
    for _, u := range pkgURLs {
        fmt.Printf("  - %s\n", u)
    }
}
```

### Check If Any Funding Exists

```go
hasFunding, err := comp.HasFunding()
if hasFunding {
    fmt.Println("Some packages accept sponsorships. Run `composer fund` for details.")
}
```

### Fund with Custom Options

```go
output, err := comp.FundWithOptions(map[string]string{
    "format": "json",
    "no-dev": "",
})
```

## Licenses

### Show License Information

Display license information for all installed packages:

```go
output, err := comp.Licenses()
fmt.Println(output)
```

### Licenses with Format

Specify the output format:

```go
output, err := comp.LicensesWithFormat("json")
```

### Licenses with Custom Options

```go
output, err := comp.LicensesWithOptions(map[string]string{
    "format": "json",
    "no-dev": "",
})
```

### Check License Compatibility

```go
output, err := comp.CheckLicenses()
fmt.Println(output)
```

Equivalent Composer commands:

```bash
composer licenses
composer licenses --format=json
composer licenses --no-dev
composer licenses --check
```

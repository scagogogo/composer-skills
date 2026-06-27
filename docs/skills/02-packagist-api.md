# Packagist API Client Skills

The Packagist API client (`pkg/client`) provides pure Go access to the Packagist REST API. No PHP or Composer installation is required.

## Creating a Client

```go
import (
    "time"
    "github.com/scagogogo/composer-skills/pkg/client"
)

// Basic client with 30-second timeout
c := client.NewComposerClient(30 * time.Second)

// With custom base URL (e.g., for mirrors)
c = client.NewComposerClient(30 * time.Second,
    client.WithBaseURL("https://packagist.org"),
    client.WithRepoURL("https://repo.packagist.org"),
)

// With API credentials (required for package management)
c = client.NewComposerClient(30 * time.Second,
    client.WithAPICredentials("username", "api-token"),
)
```

## Package Information

### Get Package Details

Retrieves full package information including versions, maintainers, downloads, and metadata.

```go
pkg, err := c.GetPackage("symfony/console")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Name: %s\n", pkg.Package.Name)
fmt.Printf("Description: %s\n", pkg.Package.Description)
fmt.Printf("Total downloads: %d\n", pkg.Package.Downloads.Total)
fmt.Printf("Latest version: %s\n", pkg.Package.Version)
```

CLI:

```bash
composer-skills package info symfony/console
```

### Get Package Statistics

Get download statistics for a specific package.

```go
stats, err := c.GetPackageStats("symfony/console")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total downloads: %d\n", stats.PackageStats.Downloads.Total)
fmt.Printf("Monthly downloads: %d\n", stats.PackageStats.Downloads.Monthly)
fmt.Printf("Daily downloads: %d\n", stats.PackageStats.Downloads.Daily)
```

CLI:

```bash
composer-skills package stats symfony/console
```

### Get V2 Metadata

Retrieve Composer V2 metadata for a package, which includes optimized version information.

```go
data, err := c.GetPackageWithV2Metadata("symfony/console")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

CLI:

```bash
composer-skills package v2-metadata symfony/console
```

### Get Dev Versions

Get development branch versions of a package (e.g., `dev-main`, `dev-master`).

```go
data, err := c.GetPackageDevVersions("symfony/console")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

CLI:

```bash
composer-skills package dev-versions symfony/console
```

## Search

### Search by Query

Search packages using a text query with pagination support.

```go
// Basic search
results, err := c.SearchPackages("logging", 15, 1)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found %d results\n", results.Total)
for _, pkg := range results.Results {
    fmt.Printf("- %s: %s\n", pkg.Name, pkg.Description)
}

// Paginated search (page 2, 25 per page)
results, err = c.SearchPackages("http", 25, 2)
```

CLI:

```bash
composer-skills search query "logging"
composer-skills search query "http" --per-page 25 --page 2
```

### Search by Tags

Search packages by one or more tags.

```go
results, err := c.SearchPackagesByTags([]string{"logging", "psr-3"}, 15, 1)
if err != nil {
    log.Fatal(err)
}
for _, pkg := range results.Results {
    fmt.Printf("- %s\n", pkg.Name)
}
```

CLI:

```bash
composer-skills search tags --tags "logging,psr-3"
```

### Search by Type

Search packages matching a query and a specific package type.

```go
results, err := c.SearchPackagesByType("laravel", "project", 15, 1)
if err != nil {
    log.Fatal(err)
}
for _, pkg := range results.Results {
    fmt.Printf("- %s (%s)\n", pkg.Name, pkg.Type)
}
```

CLI:

```bash
composer-skills search type "laravel" --type "project"
```

## Security Advisories

### Get All Advisories

Retrieve all security advisories from Packagist.

```go
advisories, err := c.GetSecurityAdvisories()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total advisories: %d\n", len(advisories.Advisories))
```

CLI:

```bash
composer-skills security advisories
```

### Get Advisories for Specific Packages

Get security advisories for a list of package names.

```go
advisories, err := c.GetSecurityAdvisoriesForPackages([]string{
    "symfony/console",
    "monolog/monolog",
})
if err != nil {
    log.Fatal(err)
}
for pkg, advs := range advisories.Advisories {
    fmt.Printf("Package: %s\n", pkg)
    for _, adv := range advs {
        fmt.Printf("  - %s: %s\n", adv.Title, adv.Link)
    }
}
```

CLI:

```bash
composer-skills security package --packages "symfony/console,monolog/monolog"
```

### Get Advisories Since Timestamp

Get advisories updated since a specific time.

```go
since := time.Now().AddDate(0, 0, -30) // last 30 days
advisories, err := c.GetSecurityAdvisoriesSince(since)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Advisories in last 30 days: %d\n", len(advisories.Advisories))
```

CLI:

```bash
composer-skills security since --days 30
```

## Repository Operations

### Get Statistics

Get overall Packagist repository statistics.

```go
stats, err := c.GetStatistics()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total packages: %d\n", stats.Totals.Packages)
fmt.Printf("Total downloads: %d\n", stats.Totals.Downloads)
```

CLI:

```bash
composer-skills repo stats
```

### List All Packages

List all package names in the repository.

```go
list, err := c.ListPackages()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total packages: %d\n", len(list.PackageNames))
```

CLI:

```bash
composer-skills repo list
```

### List Packages by Vendor

List packages from a specific vendor/organization.

```go
list, err := c.ListPackagesByVendor("symfony")
if err != nil {
    log.Fatal(err)
}
for _, name := range list.PackageNames {
    fmt.Println(name)
}
```

CLI:

```bash
composer-skills repo list-vendor --vendor symfony
```

### List Packages by Type

List packages of a specific type (e.g., `library`, `project`, `composer-plugin`).

```go
list, err := c.ListPackagesByType("composer-plugin")
if err != nil {
    log.Fatal(err)
}
for _, name := range list.PackageNames {
    fmt.Println(name)
}
```

CLI:

```bash
composer-skills repo list-type --type composer-plugin
```

### List Packages with Additional Data

List packages with extra fields like `repository` and `type`.

```go
list, err := c.ListPackagesWithData([]string{"repository", "type"})
if err != nil {
    log.Fatal(err)
}
for name, data := range list.Packages {
    fmt.Printf("%s: %v\n", name, data)
}
```

CLI:

```bash
composer-skills repo list-with-data --fields "repository,type"
```

### List Popular Packages

Get the most popular packages ranked by downloads.

```go
popular, err := c.ListPopularPackages(100)
if err != nil {
    log.Fatal(err)
}
for _, pkg := range popular.Packages {
    fmt.Printf("%s: %d downloads\n", pkg.Name, pkg.Downloads)
}
```

CLI:

```bash
composer-skills repo popular --per-page 100
```

## Package Management

Package management operations require API credentials set via `client.WithAPICredentials()`.

### Create a Package

```go
result, err := c.CreatePackage(ctx, &domain.PackageCreateRequest{
    Repository: "https://github.com/vendor/my-package",
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created package: %s\n", result.Package.Name)
```

CLI:

```bash
composer-skills manage create --repository "https://github.com/vendor/my-package" \
    --username myuser --api-token mytoken
```

### Edit a Package

```go
result, err := c.EditPackage(ctx, "vendor/my-package", &domain.PackageEditRequest{
    Repository: "https://github.com/vendor/my-package-v2",
})
if err != nil {
    log.Fatal(err)
}
```

CLI:

```bash
composer-skills manage edit --package "vendor/my-package" \
    --repository "https://github.com/vendor/my-package-v2" \
    --username myuser --api-token mytoken
```

### Update a Package

Trigger an update for a package on Packagist.

```go
result, err := c.UpdatePackage(ctx, &domain.PackageUpdateRequest{
    Repository: "vendor/my-package",
})
if err != nil {
    log.Fatal(err)
}
```

CLI:

```bash
composer-skills manage update --package "vendor/my-package" \
    --username myuser --api-token mytoken
```

## Package Changes

Track changes to packages in the repository.

```go
// Get all changes
changes, err := c.GetPackageChanges(ctx, 0)

// Get changes since a specific timestamp
changes, err = c.GetPackageChanges(ctx, 1630000000)
if changes.Error != "" {
    fmt.Println("API warning:", changes.Error)
}
```

CLI:

```bash
composer-skills changes --since 1630000000
```

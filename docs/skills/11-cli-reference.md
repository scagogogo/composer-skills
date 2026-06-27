# CLI Reference

Complete reference for all commands available in the `composer-skills` CLI tool and the Go SDK.

## Global Options

These options apply to all commands:

| Flag | Default | Description |
|------|---------|-------------|
| `--timeout` | `60` | HTTP request timeout in seconds |
| `--base-url` | `https://packagist.org` | Base URL for Packagist API |
| `--repo-url` | `https://repo.packagist.org` | Repository URL for Composer API |
| `--username` | `""` | Packagist API username |
| `--api-token` | `""` | Packagist API token |
| `--json` | `false` | Output results as JSON |

## Package Commands

### `package info`

Get detailed information about a package from Packagist.

```bash
composer-skills package info <package-name>
```

SDK:

```go
pkg, err := c.GetPackage("symfony/console")
```

### `package stats`

Get download statistics for a package.

```bash
composer-skills package stats <package-name>
```

SDK:

```go
stats, err := c.GetPackageStats("symfony/console")
```

### `package v2-metadata`

Get package information with Composer V2 metadata format.

```bash
composer-skills package v2-metadata <package-name>
```

SDK:

```go
data, err := c.GetPackageWithV2Metadata("symfony/console")
```

### `package dev-versions`

Get development versions of a package.

```bash
composer-skills package dev-versions <package-name>
```

SDK:

```go
data, err := c.GetPackageDevVersions("symfony/console")
```

## Repository Commands

### `repo stats`

Get overall Packagist repository statistics.

```bash
composer-skills repo stats
```

SDK:

```go
stats, err := c.GetStatistics()
```

### `repo list`

List all packages in the Packagist repository.

```bash
composer-skills repo list
```

SDK:

```go
list, err := c.ListPackages()
```

### `repo list-vendor`

List packages from a specific vendor.

```bash
composer-skills repo list-vendor --vendor <vendor-name>
```

SDK:

```go
list, err := c.ListPackagesByVendor("symfony")
```

### `repo list-type`

List packages of a specific type.

```bash
composer-skills repo list-type --type <package-type>
```

SDK:

```go
list, err := c.ListPackagesByType("composer-plugin")
```

### `repo list-with-data`

List packages with additional data fields.

```bash
composer-skills repo list-with-data --fields "repository,type"
```

SDK:

```go
list, err := c.ListPackagesWithData([]string{"repository", "type"})
```

### `repo popular`

List popular packages.

```bash
composer-skills repo popular --per-page 100
```

SDK:

```go
popular, err := c.ListPopularPackages(100)
```

## Search Commands

### `search query`

Search packages by text query.

```bash
composer-skills search query <query> [--per-page 15] [--page 1]
```

SDK:

```go
results, err := c.SearchPackages("logging", 15, 1)
```

### `search tags`

Search packages by tags.

```bash
composer-skills search tags --tags "logging,psr-3" [--per-page 15] [--page 1]
```

SDK:

```go
results, err := c.SearchPackagesByTags([]string{"logging", "psr-3"}, 15, 1)
```

### `search type`

Search packages by query and type.

```bash
composer-skills search type <query> --type <package-type> [--per-page 15] [--page 1]
```

SDK:

```go
results, err := c.SearchPackagesByType("laravel", "project", 15, 1)
```

## Security Commands

### `security advisories`

Get security advisories for all packages.

```bash
composer-skills security advisories
```

SDK:

```go
advisories, err := c.GetSecurityAdvisories()
```

### `security package`

Get security advisories for specific packages.

```bash
composer-skills security package --packages "symfony/console,monolog/monolog"
```

SDK:

```go
advisories, err := c.GetSecurityAdvisoriesForPackages([]string{"symfony/console", "monolog/monolog"})
```

### `security since`

Get security advisories since a date.

```bash
composer-skills security since --days 30
```

SDK:

```go
advisories, err := c.GetSecurityAdvisoriesSince(time.Now().AddDate(0, 0, -30))
```

## Changes Commands

### `changes`

Get package changes information.

```bash
composer-skills changes --since <timestamp>
```

SDK:

```go
changes, err := c.GetPackageChanges(ctx, 1630000000)
```

## Package Management Commands

### `manage create`

Create a new package on Packagist (requires credentials).

```bash
composer-skills manage create --repository <url> --username <user> --api-token <token>
```

SDK:

```go
result, err := c.CreatePackage(ctx, &domain.PackageCreateRequest{Repository: "https://github.com/vendor/pkg"})
```

### `manage edit`

Edit an existing package on Packagist (requires credentials).

```bash
composer-skills manage edit --package <name> --repository <url> --username <user> --api-token <token>
```

SDK:

```go
result, err := c.EditPackage(ctx, "vendor/pkg", &domain.PackageEditRequest{Repository: "https://github.com/vendor/pkg-v2"})
```

### `manage update`

Update a package on Packagist (requires credentials).

```bash
composer-skills manage update --package <name> --username <user> --api-token <token>
```

SDK:

```go
result, err := c.UpdatePackage(ctx, &domain.PackageUpdateRequest{Repository: "vendor/pkg"})
```

## Local Composer Commands

These commands operate on the local Composer binary. All accept `--working-dir` to specify the project directory.

### `local install`

Install project dependencies.

```bash
composer-skills local install [--no-dev] [--optimize] [--working-dir <dir>]
```

SDK:

```go
err := comp.Install(noDev, optimize)
err := comp.InstallWithOptions(options)
```

### `local require`

Add a package dependency.

```bash
composer-skills local require <package> [--version <constraint>] [--dev] [--working-dir <dir>]
```

SDK:

```go
err := comp.RequirePackage("monolog/monolog", "^3.0", false)
err := comp.RequirePackageWithOptions("monolog/monolog", "^3.0", options)
```

### `local update`

Update dependencies.

```bash
composer-skills local update [packages...] [--no-dev] [--working-dir <dir>]
```

SDK:

```go
err := comp.Update([]string{}, noDev)
err := comp.UpdateWithOptions(packages, options)
```

### `local remove`

Remove a package dependency.

```bash
composer-skills local remove <package> [--dev] [--working-dir <dir>]
```

SDK:

```go
err := comp.Remove("monolog/monolog", false)
```

### `local audit`

Run a security audit.

```bash
composer-skills local audit [--working-dir <dir>]
```

SDK:

```go
result, err := comp.AuditWithJSON()
output, err := comp.Audit()
output, err := comp.AuditWithoutDev()
```

### `local version`

Show Composer version.

```bash
composer-skills local version
```

SDK:

```go
version, err := comp.GetVersion()
```

### `local validate`

Validate composer.json.

```bash
composer-skills local validate [--working-dir <dir>]
```

SDK:

```go
err := comp.Validate()
output, err := comp.ValidateStrict()
output, err := comp.ValidateComposerLock()
```

### `local outdated`

Show outdated packages.

```bash
composer-skills local outdated [--working-dir <dir>]
```

SDK:

```go
output, err := comp.OutdatedPackages()
output, err := comp.OutdatedPackagesDirect()
output, err := comp.CheckForOutdatedPackages(direct, minor, format)
```

### `local show`

Show package information.

```bash
composer-skills local show [package] [--working-dir <dir>]
```

SDK:

```go
output, err := comp.ShowAllPackages()
output, err := comp.ShowPackage("symfony/console")
```

### `local create-project`

Create a new project from a package.

```bash
composer-skills local create-project <package> <directory> [--version <version>] [--working-dir <dir>]
```

SDK:

```go
err := comp.CreateProject("laravel/laravel", "my-app", "")
err := comp.CreateProjectWithOptions("laravel/laravel", "my-app", "", options)
```

## Go SDK Quick Reference by Package

### pkg/client -- Packagist API Client

| Method | Description |
|--------|-------------|
| `NewComposerClient(timeout, ...options)` | Create a new API client |
| `GetPackage(name)` | Get package details |
| `GetPackageStats(name)` | Get download statistics |
| `GetPackageWithV2Metadata(name)` | Get V2 metadata |
| `GetPackageDevVersions(name)` | Get dev branch versions |
| `GetPackageChanges(ctx, since)` | Get package changes |
| `SearchPackages(query, perPage, page)` | Search by query |
| `SearchPackagesByTags(tags, perPage, page)` | Search by tags |
| `SearchPackagesByType(query, type, perPage, page)` | Search by type |
| `GetSecurityAdvisories()` | Get all advisories |
| `GetSecurityAdvisoriesForPackages(names)` | Get advisories for packages |
| `GetSecurityAdvisoriesSince(since)` | Get recent advisories |
| `GetStatistics()` | Get repository statistics |
| `ListPackages()` | List all package names |
| `ListPackagesByVendor(vendor)` | List by vendor |
| `ListPackagesByType(type)` | List by type |
| `ListPackagesWithData(fields)` | List with extra data |
| `ListPopularPackages(perPage)` | List popular packages |
| `CreatePackage(ctx, request)` | Create a package |
| `EditPackage(ctx, name, request)` | Edit a package |
| `UpdatePackage(ctx, request)` | Update a package |

### pkg/composer -- Composer CLI Wrapper

| Method | Description |
|--------|-------------|
| **Core** | |
| `New(options)` | Create Composer instance |
| `Run(args...)` | Execute arbitrary command |
| `RunWithTimeout(timeout, args...)` | Execute with timeout |
| `RunWithContext(ctx, args...)` | Execute with context |
| `SetWorkingDir(dir)` | Set working directory |
| `SetEnv(env)` | Set environment variables |
| `GetVersion()` | Get Composer version |
| **Dependencies** | |
| `Install(noDev, optimize)` | Install dependencies |
| `InstallWithOptions(options)` | Install with custom flags |
| `Update(packages, noDev)` | Update dependencies |
| `UpdateWithOptions(packages, options)` | Update with custom flags |
| `RequirePackage(name, version, dev)` | Add a package |
| `RequirePackageWithOptions(name, version, options)` | Add with custom flags |
| `Remove(name, dev)` | Remove a package |
| `BumpPackages(packages)` | Bump to latest in constraint |
| `BumpPackagesWithOptions(packages, options)` | Bump with custom flags |
| `Reinstall(name)` | Remove and re-add a package |
| `DumpAutoload(optimize)` | Regenerate autoloader |
| `DumpAutoloadWithOptions(options)` | Dump autoload with flags |
| `CheckDependencies()` | Check for dependency conflicts |
| `Suggests()` | View suggested packages |
| **Audit** | |
| `Audit()` | Run security audit |
| `AuditWithJSON()` | Audit with structured result |
| `AuditWithoutDev()` | Audit production deps only |
| `AuditWithFormat(format)` | Audit with custom format |
| `AuditWithOptions(options)` | Audit with custom flags |
| `AuditLock(lockFilePath)` | Audit a lock file |
| `HasVulnerabilities()` | Boolean check for vulns |
| `GetHighSeverityVulnerabilities()` | Get critical/high vulns |
| `GetAbandonedPackages()` | Get abandoned packages |
| **Project** | |
| `CreateProject(pkg, dir, ver)` | Create new project |
| `CreateProjectWithOptions(pkg, dir, ver, opts)` | Create with flags |
| `InitProject()` | Interactive init |
| `InitProjectWithOptions(name, desc, author, opts)` | Non-interactive init |
| `RunScript(name, args...)` | Execute a script |
| `ExecuteScript(name)` | Execute with `composer run` |
| `ListScripts()` | List defined scripts |
| `GetProjectInfo()` | Get project info |
| `ArchiveProject(dir, format)` | Archive project |
| **Package Inspection** | |
| `ShowAllPackages()` | List installed packages |
| `ShowPackage(name)` | Show package details |
| `ShowDependencyTree(name)` | Show dependency tree |
| `ShowReverseDependencies(name)` | Show what depends on package |
| `WhyPackage(name)` | Explain why installed |
| `WhyNotPackage(name, version)` | Explain why version blocked |
| `OutdatedPackages()` | Show outdated packages |
| `OutdatedPackagesDirect()` | Show outdated direct deps |
| `Search(query)` | Search packages |
| `BrowsePackage(name)` | Open package homepage |
| `BrowsePackageWithOptions(name, opts)` | Browse with flags |
| `Fund()` | Show funding info |
| `FundWithJSON()` | Get structured funding |
| `FundWithPackage(name)` | Package-specific funding |
| `GetFundingURLs()` | Get all funding URLs |
| `HasFunding()` | Check if any funding exists |
| `FundWithOptions(options)` | Fund with custom flags |
| `Licenses()` | Show license info |
| `LicensesWithFormat(format)` | Licenses with format |
| `LicensesWithOptions(options)` | Licenses with flags |
| `CheckLicenses()` | Check license compatibility |
| **Configuration** | |
| `ReadComposerJSON()` | Read composer.json |
| `WriteComposerJSON(json)` | Write composer.json |
| `AddRequire(name, ver, isDev)` | Add dependency to file |
| `RemoveRequire(name, isDev)` | Remove dependency from file |
| `AddScript(name, script, desc)` | Add script to file |
| `RemoveScript(name)` | Remove script from file |
| `AddAutoload(type, ns, paths, dev)` | Add autoload config |
| `SetConfig(key, value)` | Set config section value |
| `GetConfig(key)` | Get config section value |
| `SetProperty(prop, value)` | Set top-level property |
| `GetConfigWithGlobal(key, global)` | Get config via CLI |
| `SetConfigWithGlobal(key, val, global)` | Set config via CLI |
| `GetConfigParameter(key)` | Get any config param |
| `SetConfigParameter(key, val)` | Set any config param |
| `UnsetConfig(key)` | Delete config param |
| **Repositories** | |
| `AddRepository(name, repo)` | Add a repository |
| `RemoveRepository(name)` | Remove a repository |
| `ListRepositories()` | List repositories |
| `AddVcsRepository(name, url)` | Add VCS repo |
| `AddPathRepository(name, path, opts)` | Add path repo |
| `AddComposerRepository(name, url)` | Add Composer repo |
| `AddArtifactRepository(name, path)` | Add artifact repo |
| `AddPackagistRepository(url)` | Add Packagist |
| `DisablePackagistRepository()` | Disable Packagist |
| `EnablePackagistRepository()` | Enable Packagist |
| `AddGlobalRepository(name, repo)` | Add global repo |
| `RemoveGlobalRepository(name)` | Remove global repo |
| `ListGlobalRepositories()` | List global repos |
| `SetMinimumStability(level)` | Set min stability |
| `GetMinimumStability()` | Get min stability |
| `SetPreferStable(bool)` | Set prefer-stable |
| `GetPreferStable()` | Get prefer-stable |
| `SetPreferredInstall(val)` | Set install preference |
| `GetPreferredInstall()` | Get install preference |
| **Authentication** | |
| `GetAuthConfig()` | Get auth configuration |
| `SaveAuthConfig(config)` | Save auth configuration |
| `AddGitHubToken(domain, token)` | Add GitHub token |
| `AddGitLabToken(domain, token)` | Add GitLab token |
| `AddBitbucketToken(domain, consumer, token)` | Add Bitbucket token |
| `AddBearerToken(domain, token)` | Add Bearer token |
| `AddHTTPBasicAuth(domain, user, pass)` | Add HTTP Basic auth |
| `RemoveToken(authType, domain)` | Remove a token |
| `GetToken(authType, domain)` | Get a token |
| **Global Operations** | |
| `GlobalRequire(name, ver)` | Globally require package |
| `GlobalUpdate(packages)` | Update global packages |
| `GlobalRemove(name)` | Remove global package |
| `GlobalInstall()` | Install global deps |
| `GlobalList()` | List global packages |
| `GlobalStatus()` | Show global status |
| `GlobalDumpAutoload(optimize)` | Dump global autoload |
| `GlobalExecute(cmd, args...)` | Execute global binary |
| `GlobalHome()` | Open Composer home |
| **Platform / Diagnosis** | |
| `GetPHPVersion()` | Get PHP version |
| `GetExtensions()` | List PHP extensions |
| `HasExtension(ext)` | Check for extension |
| `CheckPlatform()` | Check platform reqs |
| `CheckPlatformWithLock()` | Check lock platform reqs |
| `IsPlatformAvailable(name, ver)` | Check specific platform |
| `CheckPlatformReqs()` | Raw platform check |
| `CheckPlatformReqsLock()` | Raw lock check |
| `Diagnose()` | Run diagnostics |
| `DiagnoseWithOptions(opts)` | Diagnostics with flags |
| `Status()` | Show local modifications |
| `StatusWithOptions(opts)` | Status with flags |
| `Check()` | Check dependency sync |
| `CheckWithOptions(opts)` | Check with flags |
| `GetEnvironmentInfo()` | Get all config info |
| `SelfUpdate()` | Update Composer |
| **Validation** | |
| `Validate()` | Validate composer.json |
| `ValidateStrict()` | Strict validation |
| `ValidateWithNoCheck()` | Format-only validation |
| `ValidateWithNoCheckPublish()` | Skip publish checks |
| `ValidateWithCheckVersion()` | Check version constraints |
| `ValidateComposerJson(strict, withDeps)` | Validate with params |
| `ValidateWithOptions(options)` | Validate with flags |
| `ValidateQuiet()` | Silent validation |
| `ValidateSchema()` | Schema-only validation |
| `ValidateComposerLock()` | Validate lock file |
| `CheckNormalization()` | Check file formatting |
| `NormalizeComposerJson()` | Format composer.json |
| `CheckForSecurityVulnerabilities()` | Audit with bool result |
| **Advanced** | |
| `CreateSatisConfig(path, name, url)` | Create Satis config |
| `AddSatisRepository(path, type, url)` | Add Satis repo |
| `BuildSatis(config, output)` | Build Satis repo |
| `InitSatis(name, homepage, dir)` | Initialize Satis |
| `UpdateSatisStability(path, level)` | Update Satis stability |
| `EnableSatisArchive(path, format)` | Enable Satis archiving |
| `AddSatisRequire(path, name, ver)` | Add Satis requirement |
| `Exec(binary, args...)` | Execute vendor binary |
| `ExecWithList()` | List available binaries |
| `ExecPHP(php, binary, args...)` | Exec with specific PHP |
| `ExecWithWorkingDir(bin, dir, args...)` | Exec in directory |
| `ExecAll(args...)` | Execute all binaries |
| `ExecCommand(cmd, args...)` | Execute any command |
| `LocalExec(cmd, args...)` | Local exec |
| `LocalExecWithOptions(cmd, opts, args...)` | Local exec with flags |
| `GenerateCompletion(shell)` | Shell completions |
| `GenerateCompletionWithOptions(shell, opts)` | Completions with flags |
| `ListCommands()` | List all commands |
| `GetCommandHelp(cmd)` | Get command help |
| `FormatVersionConstraint(ver, type)` | Format version string |
| `UpdatePackageVersion(name, ver, type)` | Update version constraint |
| `LockPackageVersion(name, ver)` | Lock to exact version |
| `GetPackageVersions(name)` | Get available versions |
| **Archive** | |
| `Archive(dest)` | Create project archive |
| `ArchiveWithFormat(dest, format)` | Archive with format |
| `ArchiveWithOptions(dest, opts)` | Archive with flags |
| `ArchivePackage(name, ver, dest)` | Archive a package |
| `ArchivePackageWithOptions(name, ver, dest, opts)` | Archive with flags |
| **Environment (package-level)** | |
| `SetEnvVariable(name, value)` | Set env variable |
| `GetEnvVariable(name)` | Get env variable |
| `SetProcessTimeout(seconds)` | Set process timeout |
| `EnableSuperuser()` | Allow root |
| `DisableSuperuser()` | Disallow root |
| `SetMemoryLimit(limit)` | Set PHP memory limit |
| `DisableInteraction()` | Disable prompts |
| `EnableInteraction()` | Enable prompts |
| `SetVendorDir(path)` | Set vendor directory |
| `SetBinDir(path)` | Set bin directory |
| `SetCaFile(path)` | Set CA cert file |
| `DisableDev()` | Skip dev dependencies |
| `EnableDev()` | Include dev dependencies |
| `SetDiscardChanges(value)` | Set change handling |
| `GetComposerPath()` | Find composer binary |

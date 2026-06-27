# Configuration Skills

Read and write `composer.json`, manage configuration, repositories, authentication, and environment variables.

## Reading and Writing composer.json

### Read composer.json

Parse `composer.json` into a structured Go type:

```go
comp, _ := composer.New(composer.DefaultOptions())
comp.SetWorkingDir("/path/to/project")

composerJSON, err := comp.ReadComposerJSON()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Name: %s\n", composerJSON.Name)
fmt.Printf("Description: %s\n", composerJSON.Description)
fmt.Printf("Type: %s\n", composerJSON.Type)
fmt.Printf("Dependencies: %v\n", composerJSON.Require)
fmt.Printf("Dev dependencies: %v\n", composerJSON.RequireDev)
```

The `ComposerJSON` struct includes all standard fields: `Name`, `Description`, `Type`, `Keywords`, `Homepage`, `License`, `Authors`, `Support`, `Require`, `RequireDev`, `Suggest`, `Autoload`, `AutoloadDev`, `Repositories`, `Config`, `Scripts`, `ScriptsDescriptions`, `Extra`, `Bin`, `Archive`, `NonFeatureBranches`, `MinimumStability`, `PreferStable`, `Replace`, `Conflict`, `Provide`.

### Write composer.json

Modify and write back:

```go
composerJSON, _ := comp.ReadComposerJSON()
composerJSON.Name = "myvendor/my-project"
composerJSON.Description = "Updated description"
err := comp.WriteComposerJSON(composerJSON)
```

## Config Get / Set

### Get a Config Value

```go
// Get a project-level config value
value, err := comp.GetConfigWithGlobal("vendor-dir", false)

// Get a global config value
value, err = comp.GetConfigWithGlobal("home", true)
```

### Set a Config Value

```go
// Set a project-level config value
err := comp.SetConfigWithGlobal("process-timeout", "500", false)

// Set a global config value
err = comp.SetConfigWithGlobal("optimize-autoloader", "true", true)
```

### Generic Config Parameter Access

```go
// Get any config parameter
name, err := comp.GetConfigParameter("name")
typ, err := comp.GetConfigParameter("type")

// Set any config parameter
err := comp.SetConfigParameter("description", "My PHP project")
err = comp.SetConfigParameter("authors.0.name", "Jane Doe")
```

### Unset a Config Parameter

```go
err := comp.UnsetConfig("repositories.old-repo")
```

### composer.json Config Section

Read and write the `config` section of `composer.json`:

```go
// Get a config value
timeout, err := comp.GetConfig("process-timeout")

// Set a config value
err := comp.SetConfig("process-timeout", 500)
err = comp.SetConfig("disable-plugins", true)
```

## Repository Management

### Add a Repository

Add a custom repository to `composer.json`:

```go
repo := composer.Repository{
    Type: composer.ComposerRepository,
    URL:  "https://composer.example.org",
}
err := comp.AddRepository("private", repo)
```

### Add a VCS Repository

```go
err := comp.AddVcsRepository("my-lib", "https://github.com/vendor/package")
```

### Add a Path Repository

For local development with multiple interdependent packages:

```go
err := comp.AddPathRepository("local", "../my-package", map[string]interface{}{
    "symlink": true,
})
```

### Add a Composer Repository

```go
err := comp.AddComposerRepository("private", "https://composer.example.org")
```

### Add an Artifact Repository

For local ZIP/TAR package archives:

```go
err := comp.AddArtifactRepository("artifacts", "./packages")
```

### Add Packagist Repository

```go
// Official Packagist
err := comp.AddPackagistRepository("https://repo.packagist.org")

// Mirror
err = comp.AddPackagistRepository("https://mirrors.aliyun.com/composer")
```

### Disable / Enable Packagist

```go
// Disable (for private-only setups)
err := comp.DisablePackagistRepository()

// Re-enable
err := comp.EnablePackagistRepository()
```

### Remove a Repository

```go
err := comp.RemoveRepository("private")
```

### List Repositories

```go
output, err := comp.ListRepositories()
fmt.Println(output)
```

### Global Repository Management

```go
// Add a global repository
repo := composer.Repository{
    Type: composer.ComposerRepository,
    URL:  "https://composer.example.org",
}
err := comp.AddGlobalRepository("global-private", repo)

// List global repositories
output, err := comp.ListGlobalRepositories()

// Remove a global repository
err = comp.RemoveGlobalRepository("global-private")
```

## Authentication

### Get Current Auth Configuration

```go
authConfig, err := comp.GetAuthConfig()
fmt.Printf("GitHub tokens: %v\n", authConfig.GitHub)
fmt.Printf("GitLab tokens: %v\n", authConfig.GitLab)
fmt.Printf("HTTP Basic: %v\n", authConfig.HTTPBasic)
```

### GitHub OAuth Token

```go
// Add a GitHub token
err := comp.AddGitHubToken("github.com", "ghp_xxxxxxxxxxxx")

// Get a GitHub token
token, err := comp.GetToken("github-oauth", "github.com")
```

### GitLab OAuth Token

```go
err := comp.AddGitLabToken("gitlab.com", "glpat-xxxxxxxxxxxx")
```

### Bitbucket OAuth Token

```go
err := comp.AddBitbucketToken("bitbucket.org", "consumer-key", "secret-key")
```

### Bearer Token

```go
err := comp.AddBearerToken("example.com", "my-bearer-token")
```

### HTTP Basic Authentication

```go
err := comp.AddHTTPBasicAuth("repo.example.com", "username", "password")
```

### Remove a Token

```go
err := comp.RemoveToken("github-oauth", "github.com")
err = comp.RemoveToken("http-basic", "repo.example.com")
```

### Save Auth Configuration

```go
config := &composer.AuthConfig{
    GitHub: map[string]string{
        "github.com": "ghp_xxxxxxxxxxxx",
    },
    HTTPBasic: map[string]string{
        "repo.example.com": "user:pass",
    },
}
err := comp.SaveAuthConfig(config)
```

## Environment Variables

Composer respects several environment variables. The SDK provides helpers for the most common ones.

### Setting Environment Variables

```go
// Set Composer home directory
composer.SetEnvVariable(composer.EnvComposerHome, "/custom/composer/home")

// Set cache directory
composer.SetEnvVariable(composer.EnvComposerCacheDir, "/custom/cache")

// Set process timeout (seconds)
composer.SetProcessTimeout(600)

// Allow running as root
composer.EnableSuperuser()

// Set PHP memory limit
composer.SetMemoryLimit("2G")

// Disable interactive prompts
composer.DisableInteraction()

// Set custom vendor directory
composer.SetVendorDir("/custom/vendor")

// Set custom bin directory
composer.SetBinDir("/custom/bin")

// Skip dev dependencies
composer.DisableDev()

// Set CA certificate file
composer.SetCaFile("/path/to/ca-bundle.crt")

// Control change handling
composer.SetDiscardChanges("stash")
```

### Getting Environment Variables

```go
home := composer.GetEnvVariable(composer.EnvComposerHome)
cache := composer.GetEnvVariable(composer.EnvComposerCacheDir)
```

### Available Environment Variables

| Constant | Variable | Description |
|----------|----------|-------------|
| `EnvComposerHome` | `COMPOSER_HOME` | Composer home directory |
| `EnvComposerCacheDir` | `COMPOSER_CACHE_DIR` | Cache directory |
| `EnvComposerProcessTimeout` | `COMPOSER_PROCESS_TIMEOUT` | Process timeout (seconds) |
| `EnvComposerAllowSuperuser` | `COMPOSER_ALLOW_SUPERUSER` | Allow root execution |
| `EnvComposerMemoryLimit` | `COMPOSER_MEMORY_LIMIT` | PHP memory limit |
| `EnvComposerDisableXdebugWarn` | `COMPOSER_DISABLE_XDEBUG_WARN` | Suppress XDebug warning |
| `EnvComposerNoInteraction` | `COMPOSER_NO_INTERACTION` | Disable interaction |
| `EnvComposerVendorDir` | `COMPOSER_VENDOR_DIR` | Custom vendor directory |
| `EnvComposerBinDir` | `COMPOSER_BIN_DIR` | Custom bin directory |
| `EnvComposerCafile` | `COMPOSER_CAFILE` | CA certificate file |
| `EnvComposerNoDev` | `COMPOSER_NO_DEV` | Skip dev dependencies |
| `EnvComposerDiscardChanges` | `COMPOSER_DISCARD_CHANGES` | Change handling strategy |
| `EnvComposerHtaccessProtect` | `COMPOSER_HTACCESS_PROTECT` | htaccess protection |
| `EnvComposerMirrorPathRepos` | `COMPOSER_MIRROR_PATH_REPOS` | Path repo mirror strategy |

### Setting Environment on the Composer Instance

You can also set environment variables on the `Composer` instance, which only affect commands run through that instance:

```go
comp.SetEnv([]string{
    "COMPOSER_HOME=/custom/home",
    "HTTP_PROXY=http://proxy.example.com:8080",
    "HTTPS_PROXY=http://proxy.example.com:8080",
})
```

## Minimum Stability and Prefer Stable

### Set Minimum Stability

Control which stability levels Composer will accept:

```go
err := comp.SetMinimumStability("beta")
```

Valid values: `stable`, `RC`, `beta`, `alpha`, `dev`.

### Get Minimum Stability

```go
stability, err := comp.GetMinimumStability()
fmt.Println("Current stability:", stability)
```

### Set Prefer Stable

```go
err := comp.SetPreferStable(true)
```

### Get Prefer Stable

```go
value, err := comp.GetPreferStable()
preferStable := value == "1"
```

### Set via composer.json Properties

```go
err := comp.SetProperty("minimum-stability", "beta")
err = comp.SetProperty("prefer-stable", true)
```

## Preferred Install

### Get Preferred Install

```go
value, err := comp.GetPreferredInstall()
fmt.Println("Preferred install:", value)
```

### Set Preferred Install

```go
err := comp.SetPreferredInstall("dist")   // "dist", "source", or "auto"
```

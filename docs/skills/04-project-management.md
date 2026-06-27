# Project Management Skills

Create, initialize, and manage PHP projects using the Composer CLI wrapper.

## Creating Projects

### Basic create-project

Create a new project from an existing package (e.g., a framework skeleton):

```go
comp, _ := composer.New(composer.DefaultOptions())

// Create a new Laravel project (latest version)
err := comp.CreateProject("laravel/laravel", "my-app", "")

// Create a specific version
err = comp.CreateProject("symfony/website-skeleton", "symfony-app", "^5.0")
```

CLI:

```bash
composer-skills local create-project laravel/laravel my-app
composer-skills local create-project symfony/website-skeleton symfony-app --version "^5.0"
```

### create-project with Options

Use `CreateProjectWithOptions` for additional flags:

```go
err := comp.CreateProjectWithOptions("laravel/laravel", "my-app", "", map[string]string{
    "no-dev":      "",
    "prefer-dist": "",
    "stability":   "dev",
})
```

For private repositories, specify the repository URL:

```go
err := comp.CreateProjectWithOptions("vendor/private-package", "my-private-app", "", map[string]string{
    "repository": "https://example.org/private-repo",
})
```

Equivalent Composer command:

```bash
composer create-project --no-dev --prefer-dist laravel/laravel my-app
```

## Initializing Projects

### Basic init

Create a `composer.json` file interactively:

```go
err := comp.InitProject()
```

This launches Composer's interactive prompt.

### init with Options

Non-interactively create a `composer.json` with all details:

```go
err := comp.InitProjectWithOptions(
    "myvendor/my-library",          // name
    "A fantastic PHP library",      // description
    "Jane Doe <jane@example.com>",  // author
    map[string]string{
        "type":           "library",
        "license":        "MIT",
        "no-interaction": "",
    },
)
```

Equivalent Composer command:

```bash
composer init \
    --name="myvendor/my-library" \
    --description="A fantastic PHP library" \
    --author="Jane Doe <jane@example.com>" \
    --type=library \
    --license=MIT \
    --no-interaction
```

## Running Scripts

### Run a Script

Execute a script defined in `composer.json`:

```go
output, err := comp.RunScript("test")
fmt.Println(output)
```

Pass additional arguments:

```go
output, err := comp.RunScript("test", "--filter=UserTest")
```

### Execute a Script (shorthand)

Use the `run` command:

```go
output, err := comp.ExecuteScript("deploy")
```

### List Available Scripts

List all scripts defined in `composer.json`:

```go
output, err := comp.ListScripts()
fmt.Println("Available scripts:", output)
```

Equivalent Composer commands:

```bash
composer run-script test
composer run-script test -- --filter=UserTest
composer run deploy
composer run-script --list
```

## Project Information

### Get Project Info

Retrieve basic project information from `composer.json`:

```go
info, err := comp.GetProjectInfo()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Name: %s\n", info.Name)
fmt.Printf("Description: %s\n", info.Description)
fmt.Printf("Type: %s\n", info.Type)
fmt.Printf("Dependencies: %d\n", len(info.Require))
```

## Archiving Projects

### Basic Archive

Create a ZIP archive of the project:

```go
err := comp.ArchiveProject("./dist", "zip")
```

### Archive with Format

Choose between `zip` and `tar`:

```go
output, err := comp.ArchiveWithFormat("./dist", "tar")
```

### Archive with Options

Full control over archive creation:

```go
output, err := comp.ArchiveWithOptions("./dist", map[string]string{
    "format":          "tar",
    "file":            "my-project.tar",
    "ignore-filters":  "",
})
```

### Archive a Specific Package

Archive a package by name and version:

```go
// Archive a specific version
output, err := comp.ArchivePackage("symfony/console", "v5.4.0", "./dist")

// Archive the latest version
output, err = comp.ArchivePackage("symfony/console", "", "./dist")
```

### Archive a Package with Options

```go
output, err := comp.ArchivePackageWithOptions(
    "symfony/console",
    "v5.4.0",
    "./dist",
    map[string]string{
        "format":          "tar",
        "ignore-filters":  "",
    },
)
```

Equivalent Composer commands:

```bash
composer archive --format=zip --dir=./dist
composer archive --format=tar --dir=./dist
composer archive symfony/console=v5.4.0 --dir=./dist
```

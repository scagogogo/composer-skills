---
layout: home

hero:
  name: "Go Composer SDK"
  text: "PHP Composer for Go"
  tagline: A comprehensive Go library for PHP Composer package manager
  image:
    src: /logo.svg
    alt: Go Composer SDK
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: View on GitHub
      link: https://github.com/scagogogo/go-composer-sdk

features:
  - icon: 🚀
    title: Complete Composer Support
    details: Full support for all standard Composer CLI commands with type-safe Go APIs
  - icon: 🛡️
    title: Type Safety
    details: Strongly typed interfaces with IDE code completion and compile-time error checking
  - icon: 🔧
    title: Comprehensive Features
    details: Dependency management, repository configuration, authentication, security auditing, and more
  - icon: 🌍
    title: Cross-Platform
    details: Native support for Windows, macOS, and Linux with platform-specific optimizations
  - icon: 📦
    title: Modular Design
    details: Well-organized code structure grouped by functionality for easy use and maintenance
  - icon: ✅
    title: Production Ready
    details: Thoroughly tested with GitHub Actions CI/CD ensuring code quality and reliability
---

## Quick Start

Install the Go Composer SDK:

```bash
go get github.com/scagogogo/go-composer-sdk
```

Create a Composer instance and start managing PHP dependencies:

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/scagogogo/go-composer-sdk/pkg/composer"
)

func main() {
    // Create a Composer instance with default options
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatalf("Failed to create Composer instance: %v", err)
    }
    
    // Check if Composer is installed
    if !comp.IsInstalled() {
        log.Fatal("Composer is not installed")
    }
    
    // Get Composer version
    version, err := comp.GetVersion()
    if err != nil {
        log.Fatalf("Failed to get Composer version: %v", err)
    }
    
    fmt.Printf("Composer version: %s\n", version)
    
    // Install dependencies
    err = comp.Install(false, false) // noDev=false, optimize=false
    if err != nil {
        log.Fatalf("Failed to install dependencies: %v", err)
    }
    
    fmt.Println("Dependencies installed successfully!")
}
```

## Why Go Composer SDK?

- **🎯 Purpose-Built**: Specifically designed for Go applications that need to manage PHP projects
- **📚 Well-Documented**: Comprehensive documentation with examples for every feature
- **🔒 Secure**: Built-in security auditing and vulnerability detection
- **⚡ Performance**: Efficient execution with proper timeout handling and context support
- **🧪 Tested**: Extensive test suite with 161+ tests ensuring reliability

## What's Included

- **Core Composer Operations**: Install, update, require, remove packages
- **Project Management**: Create projects, run scripts, validate configurations
- **Security Features**: Audit dependencies, detect vulnerabilities
- **Platform Tools**: Check PHP versions, extensions, platform requirements
- **Utility Functions**: File system operations, HTTP downloads, cross-platform support
- **Auto-Detection**: Automatically detect and install Composer if needed

## Community

- [GitHub Repository](https://github.com/scagogogo/go-composer-sdk)
- [Issue Tracker](https://github.com/scagogogo/go-composer-sdk/issues)
- [Discussions](https://github.com/scagogogo/go-composer-sdk/discussions)

## License

Go Composer SDK is released under the [MIT License](https://github.com/scagogogo/go-composer-sdk/blob/main/LICENSE).

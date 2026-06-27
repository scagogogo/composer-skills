const en = {
  nav: {
    features: 'Features',
    architecture: 'Architecture',
    quickStart: 'Quick Start',
    security: 'Security',
    coverage: 'Coverage',
    github: 'GitHub',
  },
  hero: {
    badge: 'Open Source · Go SDK · MIT License',
    tagline: 'The Missing Go SDK for the PHP Composer Ecosystem',
    subtitle:
      'Stop parsing exec.Command output by hand. One import gives you a typed, tested API for both the Packagist REST API and every Composer CLI command, plus zero-config auto-installation.',
    cta: 'Get Started',
    ctaSecondary: 'View on GitHub',
    statMethods: 'SDK Methods',
    statApi: 'API Methods',
    statCli: 'CLI Commands',
    statTests: 'Tests',
  },
  problem: {
    title: 'The Problem It Solves',
    subtitle: 'If you write Go code that touches the PHP/Composer world, you know the pain.',
    oldWay: 'The old way — fragile, untyped, no error handling',
    newWay: 'The new way — typed, tested, auto-installing',
    oldCode: `// 😩 The old way — fragile, untyped, no error handling
out, _ := exec.Command("composer", "audit").Output()
lines := strings.Split(string(out), "\\n")
// Now try to parse that...`,
    newCode: `// 😊 The new way — typed, tested, auto-installing
result, _ := comp.AuditWithJSON()
fmt.Printf("Vulnerabilities: %d\\n", result.Found)`,
    painPoints: [
      { pain: 'Raw exec.Command output parsing', solution: '234 typed Go methods with structured results' },
      { pain: 'Hand-written HTTP requests to Packagist', solution: '20 typed API methods returning Go structs' },
      { pain: '"Is Composer installed on this machine?"', solution: 'Cross-OS detector finds it anywhere' },
      { pain: '"Composer isn\'t installed, now what?"', solution: 'Auto-installer downloads Composer + PHP automatically' },
      { pain: 'Different code for different OS', solution: 'Smart defaults per platform (brew, apt, direct download)' },
    ],
    painHeader: 'Pain Point',
    solutionHeader: 'Solution',
  },
  features: {
    title: 'Key Features',
    subtitle: 'Everything you need to work with the PHP/Composer ecosystem from Go.',
    items: [
      {
        title: 'Full Composer CLI Coverage',
        description: '234 SDK methods wrapping every standard Composer command across 20 categories.',
      },
      {
        title: 'Packagist API Client',
        description: '20 methods to search, browse, and query the PHP package registry from Go (pure Go, no PHP).',
      },
      {
        title: 'Security-First',
        description: 'Audit dependencies, check vulnerabilities, validate schemas, check platform requirements.',
      },
      {
        title: 'Auto-Detection & Installation',
        description: 'Automatically finds or installs Composer (with cross-OS detection and PHP auto-install).',
      },
      {
        title: 'Cross-Platform',
        description: 'Windows, macOS, and Linux support with smart defaults.',
      },
      {
        title: 'CLI Tool',
        description: '50+ subcommands exposing all SDK capabilities from the terminal.',
      },
      {
        title: 'Structured Results',
        description: 'Type-safe return values (AuditInfo, OutdatedInfo, VersionInfo, etc.) instead of raw strings.',
      },
      {
        title: 'Convenience Methods',
        description: 'IsPackageInstalled, GetDirectDependencyNames, GetProjectSummary, and 18 more helpers.',
      },
      {
        title: 'Progressive Docs',
        description: 'From 3-line quickstart to full API reference (12 guides).',
      },
      {
        title: 'Well-Tested',
        description: '450+ tests with mock-based isolation.',
      },
    ],
  },
  architecture: {
    title: 'Architecture',
    subtitle: 'A clean three-layer architecture designed for extensibility and testability.',
    layerHeader: 'Layer',
    functionHeader: 'What it does',
    packageHeader: 'Package',
    layers: [
      { layer: 'Skills Documentation', func: 'Progressive disclosure guides (12 guides)', pkg: 'docs/skills/' },
      { layer: 'CLI Tool', func: '50+ subcommands via Cobra', pkg: 'cmd/composer-skills/' },
      { layer: 'Packagist API SDK', func: 'HTTP calls to Packagist (pure Go)', pkg: 'pkg/client, pkg/repository' },
      { layer: 'Composer CLI SDK', func: 'Executes local composer binary (234 methods)', pkg: 'pkg/composer' },
      { layer: 'Foundation', func: 'Domain models, detection, installation, utilities', pkg: 'pkg/domain, pkg/detector, pkg/installer, pkg/composerutils' },
    ],
  },
  sdkComparison: {
    title: 'Two SDKs in One',
    subtitle: 'Whether you need remote API access or local CLI control, Composer Skills has you covered.',
    packagistTitle: 'Packagist API SDK',
    composerTitle: 'Composer CLI SDK',
    fields: [
      { label: 'Package', packagist: 'pkg/client, pkg/repository', composer: 'pkg/composer' },
      { label: 'How it works', packagist: 'HTTP calls to Packagist API', composer: 'Executes local composer binary' },
      { label: 'Requires PHP?', packagist: 'No (pure Go)', composer: 'Yes (PHP 7.4+ and Composer 2.0+)' },
      { label: 'Use cases', packagist: 'Search packages, get stats, security advisories', composer: 'Install/update deps, manage projects, audit, run scripts' },
    ],
  },
  security: {
    title: 'Security-First',
    subtitle: 'Built-in security auditing and validation to keep your dependencies safe.',
    auditTitle: 'Local Audit',
    auditCode: `// Local audit with structured results
result, _ := comp.AuditWithJSON()
if result.Found > 0 {
    for _, v := range result.Advisories {
        fmt.Printf("⚠ %s: %s (%s)\\n", v.Package, v.Title, v.Severity)
    }
}`,
    remoteTitle: 'Remote Advisories',
    remoteCode: `// Remote advisories from Packagist
advisories, _ := client.GetSecurityAdvisories()`,
    validateTitle: 'Validate composer.json',
    validateCode: `// Validate composer.json before committing
result, _ := comp.ValidateStructured()`,
  },
  autoInstall: {
    title: 'Auto-Install: Zero Config',
    subtitle: 'Composer Skills handles the entire setup chain — detect → check PHP → install if missing → verify → ready.',
    code: `// That's it. If Composer is missing, it gets installed automatically.
comp, err := composer.New(composer.DefaultOptions())`,
    detectTitle: 'Detect',
    detectDesc: 'Cross-OS detection finds Composer anywhere on the system.',
    checkTitle: 'Check PHP',
    checkDesc: 'Verifies PHP is available and meets version requirements.',
    installTitle: 'Auto-Install',
    installDesc: 'Downloads and installs Composer (and PHP if needed) automatically.',
    readyTitle: 'Ready',
    readyDesc: 'Verified and ready to use — no manual configuration needed.',
  },
  coverage: {
    title: 'SDK Coverage',
    subtitle: 'Comprehensive coverage of both the Packagist REST API and Composer CLI.',
    packagistTitle: 'Packagist API (20 methods)',
    composerTitle: 'Composer CLI (234 methods across 20 categories)',
    categoryHeader: 'Category',
    methodsHeader: 'Methods',
    countHeader: 'Count',
    highlightsHeader: 'Highlights',
    packagistCategories: [
      { category: 'Package Info', methods: 'GetPackage · GetPackageStats · GetPackageWithV2Metadata · GetPackageDevVersions · GetPackageChanges' },
      { category: 'Search', methods: 'SearchPackages · SearchPackagesByTags · SearchPackagesByType' },
      { category: 'Statistics', methods: 'GetStatistics' },
      { category: 'Security', methods: 'GetSecurityAdvisories · GetSecurityAdvisoriesForPackages · GetSecurityAdvisoriesSince' },
      { category: 'Listing', methods: 'ListPackages · ListPackagesByVendor · ListPackagesByType · ListPackagesWithData · ListPopularPackages' },
      { category: 'Management', methods: 'CreatePackage · EditPackage · UpdatePackage' },
    ],
    composerCategories: [
      { category: 'Core', count: '10', highlights: 'Run, RunWithContext, RunWithTimeout, GetVersion, SelfUpdate' },
      { category: 'Dependencies', count: '16', highlights: 'Install, Update, DumpAutoload, Suggests + variants' },
      { category: 'Packages', count: '20', highlights: 'Require, Remove, Reinstall, Bump, Search, Show, Why, WhyNot' },
      { category: 'Audit', count: '10', highlights: 'Audit, AuditWithJSON, HasVulnerabilities, GetHighSeverityVulnerabilities' },
      { category: 'Project', count: '10', highlights: 'CreateProject, InitProject, RunScript, ListScripts, GetProjectInfo' },
      { category: 'Config', count: '12', highlights: 'GetConfig, SetConfig, ListConfig, ClearCache, GetComposerHome' },
      { category: 'Validation', count: '14', highlights: 'Validate, ValidateStrict, ValidateSchema, NormalizeComposerJson' },
      { category: 'Platform', count: '8', highlights: 'CheckPlatform, GetPHPVersion, GetExtensions, HasExtension' },
      { category: 'Repository', count: '18', highlights: 'AddVcsRepository, AddComposerRepository, SetMinimumStability' },
      { category: 'Global', count: '14', highlights: 'GlobalRequire, GlobalUpdate, GlobalRemove, GlobalInstall' },
      { category: 'Auth', count: '10', highlights: 'AddGitHubToken, AddGitLabToken, AddBearerToken, GetAuthConfig' },
      { category: 'Fund', count: '7', highlights: 'Fund, FundWithJSON, HasFunding, GetFundingURLs' },
      { category: 'Licenses', count: '4', highlights: 'Licenses, LicensesWithFormat, CheckLicenses' },
      { category: 'Diagnosis', count: '8', highlights: 'Diagnose, Check, Status, LocalExec' },
      { category: 'Exec', count: '8', highlights: 'Exec, ExecCommand, ExecPHP, ExecWithList' },
      { category: 'Satis', count: '8', highlights: 'InitSatis, CreateSatisConfig, BuildSatis' },
      { category: 'Version', count: '5', highlights: 'GetPackageVersions, LockPackageVersion, UpdatePackageVersion' },
      { category: 'Environment', count: '12', highlights: 'GetEnvironmentInfo, SetMemoryLimit, EnableDev, DisableInteraction' },
      { category: 'composer.json', count: '10', highlights: 'ReadComposerJSON, WriteComposerJSON, AddRequire, AddScript, AddAutoload' },
      { category: 'Archive', count: '6', highlights: 'Archive, ArchiveWithFormat, ArchivePackage' },
    ],
  },
  quickStart: {
    title: 'Quick Start',
    subtitle: 'Get up and running in minutes with just one import.',
    install: 'Install',
    installCode: 'go get github.com/scagogogo/composer-skills',
    tabPackagist: 'Packagist API (No PHP)',
    tabComposer: 'Composer CLI Wrapper',
    tabAutoInstall: 'Auto-Install',
    packagistCode: `package main

import (
    "fmt"
    "time"
    "github.com/scagogogo/composer-skills/pkg/client"
)

func main() {
    c := client.NewComposerClient(30 * time.Second)

    // Search for packages
    results, _ := c.SearchPackages("logging", 10, 1)
    fmt.Printf("Found %d packages\\n", results.Total)

    // Get package details
    pkg, _ := c.GetPackage("monolog/monolog")
    fmt.Printf("%s: %s\\n", pkg.Package.Name, pkg.Package.Description)

    // Security advisories
    advisories, _ := c.GetSecurityAdvisories()
    fmt.Printf("%d advisories\\n", len(advisories.Advisories))

    // Statistics
    stats, _ := c.GetStatistics()
    fmt.Printf("Total packages: %d\\n", stats.Packages)
}`,
    composerCode: `package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/composer-skills/pkg/composer"
)

func main() {
    comp, err := composer.New(composer.DefaultOptions())
    if err != nil {
        log.Fatal(err)
    }
    comp.SetWorkingDir("/path/to/php/project")

    // Dependency management
    comp.Install(false, true)
    comp.RequirePackage("monolog/monolog", "^3.0", false)
    comp.Update([]string{}, false)

    // Security audit (structured results)
    result, _ := comp.AuditWithJSON()
    fmt.Printf("Vulnerabilities found: %d\\n", result.Found)

    // Package inspection
    output, _ := comp.ShowDependencyTree("symfony/console")
    output, _ = comp.WhyPackage("symfony/polyfill-mbstring")
    output, _ = comp.OutdatedPackages()

    // Platform checks
    phpVer, _ := comp.GetPHPVersion()
    hasExt, _ := comp.HasExtension("mbstring")
}`,
    autoInstallCode: `package main

import (
    "fmt"
    "github.com/scagogogo/composer-skills/pkg/installer"
    "github.com/scagogogo/composer-skills/pkg/detector"
)

func main() {
    // Detect if Composer is installed
    d := detector.NewDetector()
    if d.IsInstalled() {
        path, _ := d.Detect()
        fmt.Printf("Composer found at: %s\\n", path)
        return
    }

    // Auto-install with smart OS detection
    inst := installer.NewInstaller(installer.SmartConfig())
    if err := inst.Install(); err != nil {
        fmt.Printf("Install failed: %v\\n", err)
    }
}`,
    convenienceTitle: 'Convenience Methods',
    convenienceCode: `// Quick helpers that combine multiple operations
isInstalled := comp.IsPackageInstalled("monolog/monolog")
isDev := comp.IsPackageDev("monolog/monolog")
deps := comp.GetDirectDependencyNames()
summary := comp.GetProjectSummary()
hasLock := comp.HasComposerLock()
hasVendor := comp.HasVendorDir()
abandoned := comp.GetAbandonedPackagesFromLock()
namespaces := comp.GetNamespaceMap()
scripts := comp.GetScripts()`,
  },
  useCases: {
    title: 'Use Cases',
    subtitle: 'Designed for anyone who needs to interact with the PHP/Composer ecosystem from Go.',
    items: [
      {
        title: 'CI/CD Pipelines',
        description: 'Automate composer install, run security audits, check for outdated packages.',
      },
      {
        title: 'Security Scanners',
        description: 'Query Packagist advisories, audit dependencies, check platform requirements.',
      },
      {
        title: 'Package Mirrors',
        description: 'Download package indexes, list packages, get statistics from Packagist.',
      },
      {
        title: 'Dependency Dashboards',
        description: 'Show dependency trees, check licenses, track funding, monitor outdated packages.',
      },
      {
        title: 'DevOps Automation',
        description: 'Auto-detect and install Composer, manage global packages, configure auth tokens.',
      },
      {
        title: 'Satis Builders',
        description: 'Initialize, configure, and build private Composer repositories.',
      },
    ],
  },
  footer: {
    description: 'The missing Go SDK for the PHP Composer ecosystem.',
    resources: 'Resources',
    docGettingStarted: 'Getting Started',
    docPackagist: 'Packagist API',
    docSecurity: 'Security',
    docCLI: 'CLI Reference',
    community: 'Community',
    github: 'GitHub',
    goReference: 'Go Reference',
    goReport: 'Go Report Card',
    acknowledgments: 'Built with',
    packagist: 'Packagist',
    composer: 'Composer',
    license: 'MIT License',
    copyright: '© 2024 Composer Skills. Released under the MIT License.',
  },
  tutorials: {
    title: 'Tutorials & Guides',
    subtitle: 'Step-by-step guides to help you master Composer Skills.',
    items: [
      {
        category: 'Getting Started',
        categoryColor: '#2563EB',
        title: 'Your First Packagist Query',
        description: 'Learn how to search packages, get details, and explore the Packagist API in 5 minutes.',
        difficulty: 'Beginner',
        readTime: '5 min',
      },
      {
        category: 'Composer CLI',
        categoryColor: '#0284C7',
        title: 'Managing PHP Dependencies from Go',
        description: 'Install, update, and remove Composer packages programmatically with type-safe Go methods.',
        difficulty: 'Intermediate',
        readTime: '12 min',
      },
      {
        category: 'Security',
        categoryColor: '#E11D48',
        title: 'Building a Security Audit Pipeline',
        description: 'Combine local audit, remote advisories, and validation to create a complete security workflow.',
        difficulty: 'Advanced',
        readTime: '15 min',
      },
      {
        category: 'DevOps',
        categoryColor: '#059669',
        title: 'Auto-Detect & Install in CI/CD',
        description: 'Set up zero-config Composer installation in GitHub Actions, GitLab CI, or any pipeline.',
        difficulty: 'Intermediate',
        readTime: '10 min',
      },
      {
        category: 'Integration',
        categoryColor: '#D97706',
        title: 'Satis Private Repository Builder',
        description: 'Initialize, configure, and build private Composer repositories with the Satis SDK.',
        difficulty: 'Advanced',
        readTime: '18 min',
      },
      {
        category: 'CLI Tool',
        categoryColor: '#0F172A',
        title: 'CLI Quick Reference',
        description: 'A complete guide to the 50+ CLI subcommands for daily development workflows.',
        difficulty: 'Beginner',
        readTime: '8 min',
      },
    ],
  },
  showcase: {
    title: 'Real-World Examples',
    subtitle: 'See how developers use Composer Skills in production.',
    items: [
      {
        title: 'Security Scanner Service',
        description: 'A microservice that monitors PHP packages for vulnerabilities. Uses Packagist advisory API and local audit to generate real-time security reports for hundreds of projects.',
        tags: ['Security', 'Packagist API', 'Microservice'],
        code: `// Scan all projects for vulnerabilities
for _, project := range projects {
    comp, _ := composer.New(composer.WithWorkingDir(project.Path))
    result, _ := comp.AuditWithJSON()
    if result.Found > 0 {
        alertTeam(project.Name, result.Advisories)
    }
}`,
      },
      {
        title: 'Dependency Dashboard',
        description: 'An internal dashboard that tracks outdated packages, license compliance, and dependency health across all PHP projects in the organization.',
        tags: ['Dashboard', 'Outdated', 'Licenses'],
        code: `// Gather dependency intelligence
outdated, _ := comp.OutdatedPackages()
licenses, _ := comp.LicensesWithFormat("json")
abandoned, _ := comp.GetAbandonedPackagesFromLock()
funding, _ := comp.FundWithJSON()

dashboard.Update(project, DependencyReport{
    Outdated:   outdated,
    Licenses:   licenses,
    Abandoned:  abandoned,
    HasFunding: funding,
})`,
      },
      {
        title: 'CI/CD Auto-Setup',
        description: 'A GitHub Action that automatically detects and installs Composer, runs install, audit, and validates composer.json before every deployment.',
        tags: ['CI/CD', 'Auto-Install', 'Validation'],
        code: `// Auto-setup in pipeline
inst := installer.NewInstaller(installer.SmartConfig())
if err := inst.Install(); err != nil {
    return fmt.Errorf("setup failed: %w", err)
}

comp, _ := composer.New(composer.DefaultOptions())
comp.Install(true, false)

if result, _ := comp.AuditWithJSON(); result.Found > 0 {
    return fmt.Errorf("security audit failed")
}

if _, err := comp.Validate(); err != nil {
    return fmt.Errorf("invalid composer.json")
}`,
      },
    ],
  },
}

export default en

const zh = {
  nav: {
    features: '核心特性',
    architecture: '架构',
    quickStart: '快速开始',
    security: '安全',
    coverage: '覆盖范围',
    github: 'GitHub',
  },
  hero: {
    tagline: 'PHP Composer 生态缺失的 Go SDK',
    subtitle:
      '别再手动解析 exec.Command 的输出了。一个 import 就能获得类型化、经过测试的 API，覆盖 Packagist REST API 和所有 Composer CLI 命令，还自带零配置自动安装。',
    cta: '快速开始',
    ctaSecondary: '查看 GitHub',
    statMethods: 'SDK 方法',
    statApi: 'API 方法',
    statCli: 'CLI 命令',
    statTests: '测试用例',
  },
  problem: {
    title: '它解决了什么问题',
    subtitle: '如果你写 Go 代码，并且需要跟 PHP/Composer 打交道，你一定体会过这种痛苦。',
    oldWay: '老方法 — 脆弱、无类型、没有错误处理',
    newWay: '新方法 — 类型化、经过测试、自动安装',
    oldCode: `// 😩 老方法 — 脆弱、无类型、没有错误处理
out, _ := exec.Command("composer", "audit").Output()
lines := strings.Split(string(out), "\\n")
// 然后你还得自己解析这些字符串...`,
    newCode: `// 😊 新方法 — 类型化、经过测试、自动安装
result, _ := comp.AuditWithJSON()
fmt.Printf("漏洞数: %d\\n", result.Found)`,
    painPoints: [
      { pain: '手动解析 exec.Command 输出', solution: '234 个类型化 Go 方法，返回结构化结果' },
      { pain: '手写 HTTP 请求访问 Packagist', solution: '20 个类型化 API 方法，直接返回 Go 结构体' },
      { pain: '"这台机器装了 Composer 吗？"', solution: '跨操作系统检测器，到处都能找到' },
      { pain: '"没装 Composer，怎么办？"', solution: '自动安装器，自动下载 Composer + PHP' },
      { pain: '不同操作系统要写不同代码', solution: '智能默认配置，按平台选择 brew、apt 或直接下载' },
    ],
    painHeader: '痛点',
    solutionHeader: '解决方案',
  },
  features: {
    title: '核心特性',
    subtitle: '从 Go 与 PHP/Composer 生态交互所需的一切。',
    items: [
      {
        title: '完整 Composer CLI 覆盖',
        description: '234 个 SDK 方法，20 个分类，封装所有标准 Composer 命令。',
      },
      {
        title: 'Packagist API 客户端',
        description: '20 个方法搜索、浏览和查询 PHP 包注册中心（纯 Go，无需 PHP）。',
      },
      {
        title: '安全优先',
        description: '审计依赖、检查漏洞、验证 schema、检查平台要求。',
      },
      {
        title: '自动检测与安装',
        description: '跨操作系统自动检测或安装 Composer（支持 PHP 自动安装）。',
      },
      {
        title: '跨平台',
        description: '支持 Windows、macOS 和 Linux，智能默认配置。',
      },
      {
        title: 'CLI 工具',
        description: '50+ 子命令，从终端暴露所有 SDK 能力。',
      },
      {
        title: '结构化返回值',
        description: '类型安全的返回值（AuditInfo、OutdatedInfo、VersionInfo 等），而非原始字符串。',
      },
      {
        title: '便捷方法',
        description: 'IsPackageInstalled、GetDirectDependencyNames、GetProjectSummary 等 18+ 个辅助方法。',
      },
      {
        title: '渐进式文档',
        description: '从 3 行快速入门到完整 API 参考（12 个指南）。',
      },
      {
        title: '测试完善',
        description: '450+ 测试用例，基于 Mock 隔离。',
      },
    ],
  },
  architecture: {
    title: '架构',
    subtitle: '整洁的三层架构，为可扩展性和可测试性而设计。',
    layerHeader: '层级',
    functionHeader: '功能',
    packageHeader: '包路径',
    layers: [
      { layer: 'Skills 文档层', func: '渐进式披露指南（12 个指南）', pkg: 'docs/skills/' },
      { layer: 'CLI 工具层', func: '50+ 子命令（基于 Cobra）', pkg: 'cmd/composer-skills/' },
      { layer: 'Packagist API SDK', func: 'HTTP 调用 Packagist（纯 Go）', pkg: 'pkg/client, pkg/repository' },
      { layer: 'Composer CLI SDK', func: '执行本地 composer 二进制（234 方法）', pkg: 'pkg/composer' },
      { layer: '基础层', func: '领域模型、检测、安装、工具集', pkg: 'pkg/domain, pkg/detector, pkg/installer, pkg/composerutils' },
    ],
  },
  sdkComparison: {
    title: '双 SDK 合一',
    subtitle: '无论你需要远程 API 访问还是本地 CLI 控制，Composer Skills 都能满足。',
    packagistTitle: 'Packagist API SDK',
    composerTitle: 'Composer CLI SDK',
    fields: [
      { label: '包路径', packagist: 'pkg/client, pkg/repository', composer: 'pkg/composer' },
      { label: '工作方式', packagist: 'HTTP 调用 Packagist API', composer: '执行本地 composer 二进制文件' },
      { label: '需要 PHP？', packagist: '不需要（纯 Go）', composer: '需要（PHP 7.4+ 和 Composer 2.0+）' },
      { label: '使用场景', packagist: '搜索包、获取统计、安全公告', composer: '安装/更新依赖、管理项目、审计、运行脚本' },
    ],
  },
  security: {
    title: '安全优先',
    subtitle: '内置安全审计和验证，保障依赖安全。',
    auditTitle: '本地审计',
    auditCode: `// 本地审计，结构化结果
result, _ := comp.AuditWithJSON()
if result.Found > 0 {
    for _, v := range result.Advisories {
        fmt.Printf("⚠ %s: %s (%s)\\n", v.Package, v.Title, v.Severity)
    }
}`,
    remoteTitle: '远程安全公告',
    remoteCode: `// 从 Packagist 获取远程安全公告
advisories, _ := client.GetSecurityAdvisories()`,
    validateTitle: '验证 composer.json',
    validateCode: `// 提交前验证 composer.json
result, _ := comp.ValidateStructured()`,
  },
  autoInstall: {
    title: '自动安装：零配置',
    subtitle: 'Composer Skills 处理整个安装链 — 检测 → 检查 PHP → 缺失则自动安装 → 验证 → 就绪。',
    code: `// 就这么简单。如果 Composer 不存在，会自动安装。
comp, err := composer.New(composer.DefaultOptions())`,
    detectTitle: '检测',
    detectDesc: '跨操作系统检测，在任何位置找到 Composer。',
    checkTitle: '检查 PHP',
    checkDesc: '验证 PHP 是否可用并满足版本要求。',
    installTitle: '自动安装',
    installDesc: '自动下载安装 Composer（以及 PHP，如果需要）。',
    readyTitle: '就绪',
    readyDesc: '验证通过，立即可用 — 无需手动配置。',
  },
  coverage: {
    title: 'SDK 覆盖范围',
    subtitle: '全面覆盖 Packagist REST API 和 Composer CLI。',
    packagistTitle: 'Packagist API（20 个方法）',
    composerTitle: 'Composer CLI（20 个分类共 234 个方法）',
    categoryHeader: '分类',
    methodsHeader: '方法',
    countHeader: '方法数',
    highlightsHeader: '重点方法',
    packagistCategories: [
      { category: '包信息', methods: 'GetPackage · GetPackageStats · GetPackageWithV2Metadata · GetPackageDevVersions · GetPackageChanges' },
      { category: '搜索', methods: 'SearchPackages · SearchPackagesByTags · SearchPackagesByType' },
      { category: '统计', methods: 'GetStatistics' },
      { category: '安全', methods: 'GetSecurityAdvisories · GetSecurityAdvisoriesForPackages · GetSecurityAdvisoriesSince' },
      { category: '列表', methods: 'ListPackages · ListPackagesByVendor · ListPackagesByType · ListPackagesWithData · ListPopularPackages' },
      { category: '管理', methods: 'CreatePackage · EditPackage · UpdatePackage' },
    ],
    composerCategories: [
      { category: '核心', count: '10', highlights: 'Run, RunWithContext, RunWithTimeout, GetVersion, SelfUpdate' },
      { category: '依赖管理', count: '16', highlights: 'Install, Update, DumpAutoload, Suggests + 变体' },
      { category: '包操作', count: '20', highlights: 'Require, Remove, Reinstall, Bump, Search, Show, Why, WhyNot' },
      { category: '安全审计', count: '10', highlights: 'Audit, AuditWithJSON, HasVulnerabilities, GetHighSeverityVulnerabilities' },
      { category: '项目管理', count: '10', highlights: 'CreateProject, InitProject, RunScript, ListScripts, GetProjectInfo' },
      { category: '配置', count: '12', highlights: 'GetConfig, SetConfig, ListConfig, ClearCache, GetComposerHome' },
      { category: '验证', count: '14', highlights: 'Validate, ValidateStrict, ValidateSchema, NormalizeComposerJson' },
      { category: '平台', count: '8', highlights: 'CheckPlatform, GetPHPVersion, GetExtensions, HasExtension' },
      { category: '仓库', count: '18', highlights: 'AddVcsRepository, AddComposerRepository, SetMinimumStability' },
      { category: '全局操作', count: '14', highlights: 'GlobalRequire, GlobalUpdate, GlobalRemove, GlobalInstall' },
      { category: '认证', count: '10', highlights: 'AddGitHubToken, AddGitLabToken, AddBearerToken, GetAuthConfig' },
      { category: '资金', count: '7', highlights: 'Fund, FundWithJSON, HasFunding, GetFundingURLs' },
      { category: '许可证', count: '4', highlights: 'Licenses, LicensesWithFormat, CheckLicenses' },
      { category: '诊断', count: '8', highlights: 'Diagnose, Check, Status, LocalExec' },
      { category: '执行', count: '8', highlights: 'Exec, ExecCommand, ExecPHP, ExecWithList' },
      { category: 'Satis', count: '8', highlights: 'InitSatis, CreateSatisConfig, BuildSatis' },
      { category: '版本', count: '5', highlights: 'GetPackageVersions, LockPackageVersion, UpdatePackageVersion' },
      { category: '环境', count: '12', highlights: 'GetEnvironmentInfo, SetMemoryLimit, EnableDev, DisableInteraction' },
      { category: 'composer.json', count: '10', highlights: 'ReadComposerJSON, WriteComposerJSON, AddRequire, AddScript, AddAutoload' },
      { category: '归档', count: '6', highlights: 'Archive, ArchiveWithFormat, ArchivePackage' },
    ],
  },
  quickStart: {
    title: '快速开始',
    subtitle: '只需一个 import，几分钟即可上手。',
    install: '安装',
    installCode: 'go get github.com/scagogogo/composer-skills',
    tabPackagist: 'Packagist API（无需 PHP）',
    tabComposer: 'Composer CLI 封装',
    tabAutoInstall: '自动安装',
    packagistCode: `package main

import (
    "fmt"
    "time"
    "github.com/scagogogo/composer-skills/pkg/client"
)

func main() {
    c := client.NewComposerClient(30 * time.Second)

    // 搜索包
    results, _ := c.SearchPackages("logging", 10, 1)
    fmt.Printf("找到 %d 个包\\n", results.Total)

    // 获取包详情
    pkg, _ := c.GetPackage("monolog/monolog")
    fmt.Printf("%s: %s\\n", pkg.Package.Name, pkg.Package.Description)

    // 安全公告
    advisories, _ := c.GetSecurityAdvisories()
    fmt.Printf("%d 条安全公告\\n", len(advisories.Advisories))

    // 统计信息
    stats, _ := c.GetStatistics()
    fmt.Printf("总包数: %d\\n", stats.Packages)
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

    // 依赖管理
    comp.Install(false, true)
    comp.RequirePackage("monolog/monolog", "^3.0", false)
    comp.Update([]string{}, false)

    // 安全审计（结构化结果）
    result, _ := comp.AuditWithJSON()
    fmt.Printf("发现漏洞: %d\\n", result.Found)

    // 包检查
    output, _ := comp.ShowDependencyTree("symfony/console")
    output, _ = comp.WhyPackage("symfony/polyfill-mbstring")
    output, _ = comp.OutdatedPackages()

    // 平台检查
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
    // 检测 Composer 是否已安装
    d := detector.NewDetector()
    if d.IsInstalled() {
        path, _ := d.Detect()
        fmt.Printf("Composer 已安装于: %s\\n", path)
        return
    }

    // 自动安装（智能操作系统识别）
    inst := installer.NewInstaller(installer.SmartConfig())
    if err := inst.Install(); err != nil {
        fmt.Printf("安装失败: %v\\n", err)
    }
}`,
    convenienceTitle: '便捷方法',
    convenienceCode: `// 组合多个操作的快捷辅助方法
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
    title: '使用场景',
    subtitle: '为所有需要从 Go 与 PHP/Composer 生态交互的人设计。',
    items: [
      {
        title: 'CI/CD 管线',
        description: '自动化 composer install、运行安全审计、检查过期包。',
      },
      {
        title: '安全扫描器',
        description: '查询 Packagist 安全公告、审计依赖、检查平台要求。',
      },
      {
        title: '包镜像',
        description: '下载包索引、列包、获取 Packagist 统计信息。',
      },
      {
        title: '依赖仪表盘',
        description: '展示依赖树、检查许可证、追踪资金、监控过期包。',
      },
      {
        title: 'DevOps 自动化',
        description: '自动检测并安装 Composer、管理全局包、配置认证令牌。',
      },
      {
        title: 'Satis 构建器',
        description: '初始化、配置和构建私有 Composer 仓库。',
      },
    ],
  },
  footer: {
    description: 'PHP Composer 生态缺失的 Go SDK。',
    resources: '资源',
    docGettingStarted: '入门指南',
    docPackagist: 'Packagist API',
    docSecurity: '安全',
    docCLI: 'CLI 参考',
    community: '社区',
    github: 'GitHub',
    goReference: 'Go Reference',
    goReport: 'Go Report Card',
    acknowledgments: '致谢',
    packagist: 'Packagist',
    composer: 'Composer',
    license: 'MIT 许可证',
    copyright: '© 2024 Composer Skills. 基于 MIT 许可证发布。',
  },
}

export default zh

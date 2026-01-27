# Composer 安全公告监控示例

这个示例演示了如何创建一个持续监控 Composer 包安全公告的工具。该工具可以定期运行，收集安全漏洞信息，并将结果保存到文件中以便进一步分析。

## 功能特点

- 获取 Composer 所有安全公告信息
- 筛选特定包的安全公告
- 以 JSON 格式保存安全公告数据
- 支持定期运行自动化监控

## 工作原理

1. 定义一组需要监控的 PHP 包名称
2. 从 Packagist API 获取所有安全公告
3. 将完整的安全公告数据保存到时间戳命名的 JSON 文件
4. 筛选出被监控包的安全公告，单独保存到另一个文件

## 使用方法

### 直接运行

```bash
go run main.go
```

### 构建后运行

```bash
go build -o security_monitor
./security_monitor
```

### 设置为定期任务

您可以将编译后的程序设置为 cron 任务，实现定期自动监控：

```bash
# 添加到 crontab (每天午夜运行)
0 0 * * * /path/to/security_monitor
```

## 示例输出

```
获取所有安全公告...
已保存所有公告数据到 security_data/all_advisories_20230525-120130.json
发现 symfony/symfony 有 5 个安全公告
发现 laravel/framework 有 2 个安全公告
已保存跟踪包的公告数据到 security_data/tracked_advisories_20230525-120130.json

安全监控完成。可以设置此脚本通过 cron 作业定期运行，以持续监控安全公告。
示例 cron 表达式（每天运行一次）: 0 0 * * * /path/to/security_monitor
```

## 数据存储

程序会在当前目录下创建 `security_data` 目录存储监控数据：

- `all_advisories_[时间戳].json` - 包含所有包的安全公告
- `tracked_advisories_[时间戳].json` - 仅包含被监控包的安全公告

## 自定义

您可以修改 `main.go` 文件中的以下内容来自定义监控行为：

1. `trackedPackages` 数组 - 添加或删除需要监控的包
2. `NewSecurityMonitor("security_data", trackedPackages)` - 更改数据存储目录 
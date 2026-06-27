# Composer Skills - 使用示例

本目录包含了 Composer Skills 库的使用示例，演示了如何使用该库与 Packagist API 交互以及通过 Composer CLI 管理本地 PHP 项目。

## 示例目录

示例按照复杂度和功能逐步展开，建议按照顺序查看：

### Packagist API 示例（远程操作）

1. **basic_setup** - 基本设置，展示如何初始化 Packagist API 客户端
2. **download_index** - 下载包索引，展示如何获取完整的包列表
3. **list_packages** - 列出包，展示如何获取和处理包列表
4. **get_statistics** - 获取统计数据，展示如何获取仓库统计信息
5. **security_advisories** - 安全公告，展示如何获取包的安全漏洞信息
6. **popular_packages** - 热门包，展示如何获取最受欢迎的包
7. **security_monitor** - 安全监控，展示如何构建安全监控系统

### Composer CLI 示例（本地操作）

8. **cli_basic_usage** - 基本用法，展示如何创建 Composer 实例和执行基本命令
9. **cli_package_management** - 包管理，展示安装、更新、添加、移除、搜索包
10. **cli_project_management** - 项目管理，展示创建项目、运行脚本、平台检查、依赖分析、完整性检查
11. **cli_security** - 安全审计，展示安全审计、平台要求检查、配置验证
12. **cli_inspection** - 包检查，展示查看包信息、依赖树、why分析、资金和许可证
13. **cli_configuration** - 配置管理，展示 composer.json 操作、配置命令、认证管理、仓库管理
14. **cli_global** - 全局操作，展示全局安装、更新、移除包
15. **cli_advanced** - 高级功能，展示 Satis 私有仓库、命令执行、版本约束、诊断、归档、环境变量

## 运行示例

每个示例都是独立的，可以直接在各自的目录中运行。例如：

```bash
# Packagist API 示例
cd basic_setup
go run main.go

# Composer CLI 示例（需要本地安装 PHP 和 Composer）
cd cli_basic_usage
go run 02_run_commands.go
```

## 示例特点

- 所有示例都包含详细的注释，解释每个步骤的目的和作用
- 所有必要的参数都在代码中硬编码，无需从命令行输入
- 每个示例都包含预期输出的示例，即使不运行代码也能了解结果
- 每个示例都包含错误处理，展示如何处理可能的异常情况

## 实际应用中的考虑事项

- 这些示例主要用于演示 API 的使用，实际应用中可能需要更健壮的错误处理
- 为了简化示例，有些地方使用了简化的初始化方法，实际应用中应使用适当的构造函数
- 实际应用中可能需要处理大量数据和分页，这些例子仅演示基本调用
- 在处理安全公告等关键数据时，建议实现更完善的持久化和通知机制
- Composer CLI 示例需要本地安装 PHP 和 Composer，否则会触发自动安装

## 自定义示例

如果需要自定义这些示例，可以修改以下部分：

- 修改 `ServerUrl` 可以连接到不同的 Composer 仓库
- 修改请求的包名、时间范围等参数可以获取不同的数据
- 添加代理配置可以在网络受限环境中使用
- 扩展输出处理逻辑可以实现更复杂的数据分析
- 修改 `WorkingDir` 可以操作不同的 PHP 项目

## 注意事项

- Packagist API 示例会进行实际的 API 调用，请注意避免频繁运行以免给目标服务器带来负担
- 某些 API 可能会返回大量数据，特别是包列表和安全公告，请确保有足够的内存
- 由于网络原因，API 调用可能会失败，示例中包含了基本的错误处理
- Composer CLI 示例会实际修改本地项目，建议在测试项目中运行

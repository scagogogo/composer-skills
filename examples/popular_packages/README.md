# 流行 Composer 包查询示例

这个示例展示了如何使用 Composer Crawler 库获取一组流行的 PHP Composer 包的详细信息，并将结果保存为 JSON 文件。

## 功能特点

- 获取 Composer 仓库统计数据
- 获取安全公告信息
- 批量获取多个流行包的详细信息
- 将所有获取到的数据保存为 JSON 文件

## 包含的包

该示例会获取以下流行 PHP 包的详细信息：

1. `symfony/symfony` - PHP 框架
2. `laravel/framework` - PHP 框架
3. `guzzlehttp/guzzle` - HTTP 客户端
4. `monolog/monolog` - 日志库
5. `phpunit/phpunit` - 测试框架

## 运行方式

在当前目录执行：

```bash
go run main.go
```

执行后会在当前目录生成一个名为 `popular_packages_results.json` 的文件，包含所有获取的数据。

## 示例输出

运行后的输出类似：

```
获取 Composer 仓库统计数据...
总下载量: 25000000000, 包数量: 400000, 版本数量: 3000000

获取安全公告...
获取到 120 个包的安全公告

获取流行包的信息...
  获取 symfony/symfony 的信息...
    名称: symfony/symfony
    描述: The Symfony PHP framework
    类型: library
    下载量: 150000000
    版本数: 150
    GitHub Stars: 28000

  获取 laravel/framework 的信息...
    名称: laravel/framework
    描述: The Laravel Framework.
    类型: library
    下载量: 120000000
    版本数: 80
    GitHub Stars: 25000

...

结果已保存到 popular_packages_results.json
```

## 自定义

您可以修改 `main.go` 中的 `popularPackages` 数组来获取任何您感兴趣的 Composer 包的信息。 
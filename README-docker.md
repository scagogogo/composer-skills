# 使用 Docker 运行 Composer Crawler

本文档介绍如何使用 Docker 和 Docker Compose 运行 Composer Crawler。

## 先决条件

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/) (通常随 Docker 一起安装)

## 构建和运行

### 使用 Docker

1. 构建 Docker 镜像:

```bash
docker build -t composer-skills .
```

2. 运行容器获取统计信息:

```bash
# 创建数据目录
mkdir -p data

# 运行容器获取统计信息
docker run -v $(pwd)/data:/data composer-skills -stats -output /data/statistics.json
```

3. 获取特定包信息:

```bash
docker run -v $(pwd)/data:/data composer-skills -package symfony/console -output /data/symfony-console.json
```

4. 获取安全公告:

```bash
docker run -v $(pwd)/data:/data composer-skills -advisories -output /data/advisories.json
```

### 使用 Docker Compose

使用 Docker Compose 可以更简单地运行各种预设的任务。项目包含一个 `docker-compose.yml` 文件，定义了多个服务。

#### 获取统计信息

```bash
docker-compose up composer-skills
```

结果将保存在 `./data` 目录下。

#### 获取特定包信息

```bash
docker-compose up get-package
```

这将获取 `symfony/console` 包的信息，并保存到 `./data/symfony-console.json`。

#### 获取安全公告

```bash
docker-compose up get-advisories
```

这将获取所有安全公告，并保存到 `./data/advisories.json`。

#### 运行流行包示例

```bash
docker-compose up popular-packages
```

这将运行流行包示例，并将结果保存到容器中 `/root/popular_packages_results.json`。

#### 运行安全监控示例

```bash
docker-compose up security-monitor
```

这将运行安全监控示例，结果将保存在容器中的 `/root/security_data` 目录下。

## 自定义运行

您可以通过修改 `docker-compose.yml` 文件来自定义服务的行为，或者直接使用 `docker run` 命令设置不同的参数。

例如，要获取不同包的信息:

```bash
docker run -v $(pwd)/data:/data composer-skills -package laravel/framework -output /data/laravel-framework.json
```

## 故障排除

如果遇到权限问题，可能需要调整数据目录的权限:

```bash
# 在主机上
chmod -R 777 ./data
```

或在 Dockerfile 中修改用户权限设置。 
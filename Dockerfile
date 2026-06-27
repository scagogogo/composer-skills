FROM golang:1.20-alpine AS build

WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum* ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o /composer-skills cmd/main.go

# 使用轻量级镜像
FROM alpine:latest

# 安装 CA 证书以支持 HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=build /composer-skills .

# 创建用于保存数据的目录
RUN mkdir -p /data

# 设置环境变量
ENV OUTPUT_DIR=/data

# 容器启动时运行的命令
ENTRYPOINT ["./composer-skills"]

# 默认参数，使用时可以覆盖
CMD ["--help"] 
.PHONY: build install clean

# 构建可执行文件
build:
	go build -o bin/composer-crawler ./cmd/composer-crawler

# 安装到GOPATH/bin
install:
	go install ./cmd/composer-crawler

# 清理构建产物
clean:
	rm -rf bin/

# 运行测试
test:
	go test -v ./...

# 显示帮助
help:
	@echo "使用方法:"
	@echo "  make build    - 构建可执行文件到bin目录"
	@echo "  make install  - 安装可执行文件到GOPATH/bin"
	@echo "  make clean    - 清理构建产物"
	@echo "  make test     - 运行测试"

# 默认任务
default: build 
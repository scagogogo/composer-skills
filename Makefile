.PHONY: build install clean test test-coverage test-race lint fmt vet check help docs

# 构建可执行文件
build:
	go build -o bin/composer-skills ./cmd/composer-skills

# 安装到GOPATH/bin
install:
	go install ./cmd/composer-skills

# 清理构建产物
clean:
	rm -rf bin/

# 运行测试
test:
	go test -v ./...

# 运行测试（带竞态检测）
test-race:
	go test -race ./...

# 运行测试并生成覆盖率报告
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	@echo ""
	@echo "HTML coverage report: coverage.html"
	go tool cover -html=coverage.out -o coverage.html

# 运行代码格式化
fmt:
	gofmt -w .
	goimports -w . 2>/dev/null || true

# 运行代码静态检查
vet:
	go vet ./...

# 运行所有检查（格式化 + 静态检查 + 测试）
check: fmt vet test

# 运行 lint（需要安装 golangci-lint）
lint:
	golangci-lint run ./... 2>/dev/null || echo "golangci-lint not installed, skipping"

# 显示帮助
help:
	@echo "使用方法:"
	@echo "  make build          - 构建可执行文件到bin目录"
	@echo "  make install        - 安装可执行文件到GOPATH/bin"
	@echo "  make clean          - 清理构建产物"
	@echo "  make test           - 运行测试"
	@echo "  make test-race      - 运行测试（带竞态检测）"
	@echo "  make test-coverage  - 运行测试并生成覆盖率报告"
	@echo "  make fmt            - 格式化代码"
	@echo "  make vet            - 静态检查"
	@echo "  make lint           - 运行 golangci-lint"
	@echo "  make check          - 运行所有检查（fmt + vet + test）"

# 默认任务
default: build
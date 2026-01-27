#!/bin/bash

set -e

# 颜色定义
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=============================================${NC}"
echo -e "${BLUE}     Composer Crawler - 构建脚本${NC}"
echo -e "${BLUE}=============================================${NC}"

# 创建输出目录
mkdir -p build

# 构建主程序
echo -e "${YELLOW}正在构建主程序...${NC}"
go build -o build/composer-crawler cmd/main.go
echo -e "${GREEN}主程序构建完成: build/composer-crawler${NC}"

# 构建示例
echo -e "${YELLOW}构建示例程序...${NC}"

# 流行包查询示例
echo -e "${YELLOW}构建流行包查询示例...${NC}"
go build -o build/popular-packages examples/popular_packages/main.go
echo -e "${GREEN}流行包查询示例构建完成: build/popular-packages${NC}"

# 安全监控示例
echo -e "${YELLOW}构建安全监控示例...${NC}"
go build -o build/security-monitor examples/security_monitor/main.go
echo -e "${GREEN}安全监控示例构建完成: build/security-monitor${NC}"

# 运行测试
echo -e "${YELLOW}运行测试...${NC}"
go test -v ./pkg/...
echo -e "${GREEN}测试完成${NC}"

echo -e "${BLUE}=============================================${NC}"
echo -e "${GREEN}构建完成! 生成的可执行文件在 build/ 目录下${NC}"
echo -e "${BLUE}=============================================${NC}"

echo -e "可执行以下命令运行:"
echo -e "${YELLOW}主程序:${NC} ./build/composer-crawler -help"
echo -e "${YELLOW}流行包查询:${NC} ./build/popular-packages"
echo -e "${YELLOW}安全监控:${NC} ./build/security-monitor" 
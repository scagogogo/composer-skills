#!/bin/bash

# Composer Crawler 测试覆盖率脚本
# 运行所有测试并生成覆盖率报告

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 配置参数
COVERAGE_THRESHOLD=85.0
BENCHMARK_TIME="5s"
FUZZ_TIME="10s"

echo -e "${BLUE}=============================================${NC}"
echo -e "${BLUE}    Composer Crawler 测试覆盖率报告${NC}"
echo -e "${BLUE}=============================================${NC}"

# 清理之前的覆盖率文件
echo -e "${YELLOW}清理之前的覆盖率文件...${NC}"
rm -f coverage.out coverage.html coverage_*.out

# 运行单元测试并生成覆盖率
echo -e "${YELLOW}运行单元测试...${NC}"
go test -v -coverprofile=coverage.out ./pkg/...

# 检查测试是否成功
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 单元测试通过${NC}"
else
    echo -e "${RED}✗ 单元测试失败${NC}"
    exit 1
fi

# 运行集成测试
echo -e "${YELLOW}运行集成测试...${NC}"
go test -v -coverprofile=coverage_integration.out ./cmd/...

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 集成测试通过${NC}"
else
    echo -e "${RED}✗ 集成测试失败${NC}"
    exit 1
fi

# 运行基准测试
echo -e "${YELLOW}运行基准测试...${NC}"
go test -bench=. -benchmem -benchtime=${BENCHMARK_TIME} ./pkg/client/... -run=^$ > benchmark_results.txt 2>&1

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 基准测试完成${NC}"
    echo -e "${CYAN}基准测试结果已保存到 benchmark_results.txt${NC}"
else
    echo -e "${YELLOW}⚠ 基准测试有警告，但继续执行${NC}"
fi

# 运行模糊测试（简短运行）
echo -e "${YELLOW}运行模糊测试...${NC}"
go test -fuzz=FuzzPackageInfo -fuzztime=${FUZZ_TIME} ./pkg/domain/... > fuzz_results.txt 2>&1

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 模糊测试完成${NC}"
    echo -e "${CYAN}模糊测试结果已保存到 fuzz_results.txt${NC}"
else
    echo -e "${YELLOW}⚠ 模糊测试有警告，但继续执行${NC}"
fi

# 生成详细覆盖率分析
echo -e "${YELLOW}生成详细覆盖率分析...${NC}"

# 生成函数级覆盖率报告
echo -e "${PURPLE}=== 函数级覆盖率报告 ===${NC}"
go tool cover -func=coverage.out | head -20

# 生成HTML覆盖率报告
echo -e "${YELLOW}生成HTML覆盖率报告...${NC}"
go tool cover -html=coverage.out -o coverage.html

# 获取总覆盖率
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}')
echo -e "${BLUE}总覆盖率: ${GREEN}${COVERAGE}${NC}"

# 生成覆盖率统计
echo -e "${PURPLE}=== 覆盖率统计 ===${NC}"
TOTAL_FUNCTIONS=$(go tool cover -func=coverage.out | grep -v total | wc -l)
HIGH_COVERAGE=$(go tool cover -func=coverage.out | grep -v total | awk '$3 >= 90.0 {print $0}' | wc -l)
MEDIUM_COVERAGE=$(go tool cover -func=coverage.out | grep -v total | awk '$3 >= 70.0 && $3 < 90.0 {print $0}' | wc -l)
LOW_COVERAGE=$(go tool cover -func=coverage.out | grep -v total | awk '$3 < 70.0 {print $0}' | wc -l)
ZERO_COVERAGE=$(go tool cover -func=coverage.out | grep -v total | awk '$3 == 0.0 {print $0}' | wc -l)

echo -e "${GREEN}高覆盖率函数 (≥90%): ${HIGH_COVERAGE}/${TOTAL_FUNCTIONS}${NC}"
echo -e "${YELLOW}中等覆盖率函数 (70-89%): ${MEDIUM_COVERAGE}/${TOTAL_FUNCTIONS}${NC}"
echo -e "${RED}低覆盖率函数 (<70%): ${LOW_COVERAGE}/${TOTAL_FUNCTIONS}${NC}"
echo -e "${RED}零覆盖率函数 (0%): ${ZERO_COVERAGE}/${TOTAL_FUNCTIONS}${NC}"

# 设置覆盖率门禁
COVERAGE_NUM=$(echo $COVERAGE | sed 's/%//')

if (( $(echo "$COVERAGE_NUM >= $COVERAGE_THRESHOLD" | bc -l) )); then
    echo -e "${GREEN}✓ 覆盖率达到门禁要求 (>= ${COVERAGE_THRESHOLD}%)${NC}"
else
    echo -e "${RED}✗ 覆盖率未达到门禁要求 (< ${COVERAGE_THRESHOLD}%)${NC}"
    echo -e "${YELLOW}当前覆盖率: ${COVERAGE_NUM}%${NC}"
    echo -e "${YELLOW}要求覆盖率: ${COVERAGE_THRESHOLD}%${NC}"

    # 显示需要改进的函数
    echo -e "${PURPLE}=== 需要改进的函数 (覆盖率 < 70%) ===${NC}"
    go tool cover -func=coverage.out | grep -v total | awk '$3 < 70.0 {print $0}' | head -10

    exit 1
fi

# 生成包级覆盖率报告
echo -e "${PURPLE}=== 包级覆盖率分析 ===${NC}"
echo "Package Coverage Report:" > package_coverage.txt
for pkg in $(go list ./pkg/...); do
    coverage=$(go test -coverprofile=temp.out $pkg 2>/dev/null && go tool cover -func=temp.out 2>/dev/null | grep total | awk '{print $3}' || echo "0.0%")
    printf "%-40s %s\n" "$pkg" "$coverage" | tee -a package_coverage.txt
    rm -f temp.out
done

# 生成测试统计报告
echo -e "${PURPLE}=== 测试统计报告 ===${NC}"
TEST_FILES=$(find ./pkg -name "*_test.go" | wc -l)
BENCH_FILES=$(find ./pkg -name "*_bench_test.go" | wc -l)
FUZZ_FILES=$(find ./pkg -name "*_fuzz_test.go" | wc -l)
TOTAL_TEST_FUNCTIONS=$(grep -r "^func Test" ./pkg --include="*_test.go" | wc -l)
TOTAL_BENCH_FUNCTIONS=$(grep -r "^func Benchmark" ./pkg --include="*_test.go" | wc -l)
TOTAL_FUZZ_FUNCTIONS=$(grep -r "^func Fuzz" ./pkg --include="*_test.go" | wc -l)

echo -e "${CYAN}测试文件数量: ${TEST_FILES}${NC}"
echo -e "${CYAN}基准测试文件数量: ${BENCH_FILES}${NC}"
echo -e "${CYAN}模糊测试文件数量: ${FUZZ_FILES}${NC}"
echo -e "${CYAN}单元测试函数数量: ${TOTAL_TEST_FUNCTIONS}${NC}"
echo -e "${CYAN}基准测试函数数量: ${TOTAL_BENCH_FUNCTIONS}${NC}"
echo -e "${CYAN}模糊测试函数数量: ${TOTAL_FUZZ_FUNCTIONS}${NC}"

# 生成最终报告
echo -e "${BLUE}=============================================${NC}"
echo -e "${GREEN}✓ 所有测试完成！${NC}"
echo -e "${BLUE}=============================================${NC}"

echo -e "${CYAN}生成的报告文件:${NC}"
echo -e "  📊 coverage.html - HTML覆盖率报告"
echo -e "  📈 benchmark_results.txt - 基准测试结果"
echo -e "  🔍 fuzz_results.txt - 模糊测试结果"
echo -e "  📋 package_coverage.txt - 包级覆盖率报告"

# 显示改进建议
echo -e "${PURPLE}=== 改进建议 ===${NC}"
if [ $ZERO_COVERAGE -gt 0 ]; then
    echo -e "${YELLOW}📌 发现 ${ZERO_COVERAGE} 个零覆盖率函数，建议优先添加测试${NC}"
fi

if [ $LOW_COVERAGE -gt 0 ]; then
    echo -e "${YELLOW}📌 发现 ${LOW_COVERAGE} 个低覆盖率函数，建议增加测试场景${NC}"
fi

if (( $(echo "$COVERAGE_NUM < 90.0" | bc -l) )); then
    echo -e "${YELLOW}📌 总覆盖率未达到90%，建议继续完善测试${NC}"
fi

# 显示未覆盖的函数
if [ $ZERO_COVERAGE -gt 0 ]; then
    echo -e "${YELLOW}未覆盖的函数 (前10个):${NC}"
    go tool cover -func=coverage.out | grep "0.0%" | head -10
fi

# 显示覆盖率最低的函数
echo -e "${YELLOW}覆盖率最低的函数 (前10个):${NC}"
go tool cover -func=coverage.out | grep -v "100.0%" | grep -v "0.0%" | sort -k3 -n | head -10

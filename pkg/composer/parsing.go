package composer

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// ==================== 输出解析工具 ====================

// ParseComposerShowJSON 解析 `composer show --format=json` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *PackageInfo: 包信息
//   - error: 解析错误
//
// 功能说明：
//
//	解析composer show命令的JSON输出，返回结构化的包信息。
func ParseComposerShowJSON(output string) (*PackageInfo, error) {
	return ParsePackageInfo(output)
}

// ParseComposerOutdatedJSON 解析 `composer outdated --format=json` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *OutdatedResult: 过时包信息
//   - error: 解析错误
func ParseComposerOutdatedJSON(output string) (*OutdatedResult, error) {
	return ParseOutdatedResult(output)
}

// ParseComposerAuditJSON 解析 `composer audit --format=json` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *AuditInfoResult: 审计结果
//   - error: 解析错误
func ParseComposerAuditJSON(output string) (*AuditInfoResult, error) {
	return ParseAuditInfoResult(output)
}

// ParseComposerSearchJSON 解析 `composer search --format=json` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *SearchResult: 搜索结果
//   - error: 解析错误
func ParseComposerSearchJSON(output string) (*SearchResult, error) {
	return ParseSearchResult(output)
}

// ParseDependencyTreeOutput 解析 `composer show --tree` 的文本输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []DependencyNode: 依赖树节点列表
//   - error: 解析错误
//
// 功能说明：
//
//	解析 `composer show --tree` 的文本格式输出为依赖树结构。
//	注意：如果使用 `--format=json`，请使用 ParseDependencyTreeJSON。
func ParseDependencyTreeOutput(output string) ([]DependencyNode, error) {
	var roots []DependencyNode
	var stack []*DependencyNode // 跟踪当前路径

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// 计算缩进级别
		indent := 0
		trimmed := line
		for len(trimmed) > 0 && (trimmed[0] == ' ' || trimmed[0] == '|' || trimmed[0] == '`' || trimmed[0] == '-') {
			trimmed = trimmed[1:]
			indent++
		}
		trimmed = strings.TrimSpace(trimmed)

		if trimmed == "" {
			continue
		}

		// 解析 "package/name version" 格式
		parts := strings.Fields(trimmed)
		if len(parts) < 1 {
			continue
		}

		node := DependencyNode{
			Name: parts[0],
		}
		if len(parts) >= 2 {
			node.Version = parts[1]
		}

		// 根据缩进级别确定父子关系
		level := indent / 4
		if level == 0 {
			roots = append(roots, node)
			stack = []*DependencyNode{&roots[len(roots)-1]}
		} else if level <= len(stack) {
			// 添加为当前层级的子节点
			parent := stack[level-1]
			parent.Children = append(parent.Children, node)
			// 更新栈
			if level < len(stack) {
				stack = stack[:level]
			}
			stack = append(stack, &parent.Children[len(parent.Children)-1])
		}
	}

	return roots, nil
}

// ParseDependencyTreeJSON 解析 `composer show --tree --format=json` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []DependencyNode: 依赖树节点列表
//   - error: 解析错误
func ParseDependencyTreeJSON(output string) ([]DependencyNode, error) {
	output = strings.TrimSpace(output)
	var nodes []DependencyNode
	if err := json.Unmarshal([]byte(output), &nodes); err != nil {
		return nil, fmt.Errorf("解析依赖树JSON失败: %w", err)
	}
	return nodes, nil
}

// ParseInstallOutput 解析 `composer install` 的输出，提取安装统计
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *InstallResult: 安装结果
func ParseInstallOutput(output string) *InstallResult {
	result := &InstallResult{Output: output}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 匹配 "Package operations: X installs, Y updates, Z removals"
		if strings.Contains(line, "Package operations:") {
			re := regexp.MustCompile(`(\d+)\s+install`)
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				fmt.Sscanf(matches[1], "%d", &result.PackagesInstalled)
			}
			re = regexp.MustCompile(`(\d+)\s+update`)
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				fmt.Sscanf(matches[1], "%d", &result.PackagesUpdated)
			}
			re = regexp.MustCompile(`(\d+)\s+removal`)
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				fmt.Sscanf(matches[1], "%d", &result.PackagesRemoved)
			}
		}

		// 匹配警告
		if strings.Contains(line, "Warning") || strings.Contains(line, "warning") {
			result.Warnings = append(result.Warnings, line)
		}
	}

	return result
}

// ParseUpdateOutput 解析 `composer update` 的输出，提取更新统计
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - *UpdateResult: 更新结果
func ParseUpdateOutput(output string) *UpdateResult {
	result := &UpdateResult{Output: output}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "Package operations:") {
			re := regexp.MustCompile(`(\d+)\s+update`)
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				fmt.Sscanf(matches[1], "%d", &result.PackagesUpdated)
			}
		}

		if strings.Contains(line, "Warning") || strings.Contains(line, "warning") {
			result.Warnings = append(result.Warnings, line)
		}
	}

	return result
}

// ParseRequireOutput 解析 `composer require` 的输出
//
// 参数：
//   - output: 命令输出
//   - packageName: 包名
//
// 返回值：
//   - *RequireResult: 添加包结果
func ParseRequireOutput(output string, packageName string) *RequireResult {
	result := &RequireResult{
		PackageName: packageName,
		Output:      output,
	}

	// 尝试从输出中提取安装的版本
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Installing") && strings.Contains(line, packageName) {
			// 格式: "Installing symfony/console (v5.4.0)"
			re := regexp.MustCompile(regexp.QuoteMeta(packageName) + `\s*\(([^)]+)\)`)
			if matches := re.FindStringSubmatch(line); len(matches) > 1 {
				result.Version = matches[1]
			}
		}
		if strings.Contains(line, "Warning") || strings.Contains(line, "warning") {
			result.Warnings = append(result.Warnings, line)
		}
	}

	return result
}

// ParseRemoveOutput 解析 `composer remove` 的输出
//
// 参数：
//   - output: 命令输出
//   - packageName: 包名
//
// 返回值：
//   - *RemoveResult: 移除包结果
func ParseRemoveOutput(output string, packageName string) *RemoveResult {
	result := &RemoveResult{
		PackageName: packageName,
		Output:      output,
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Warning") || strings.Contains(line, "warning") {
			result.Warnings = append(result.Warnings, line)
		}
	}

	return result
}

// ParseSelfUpdateOutput 解析 `composer self-update` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - oldVersion: 旧版本
//   - newVersion: 新版本
//   - error: 解析错误
//
// 功能说明：
//
//	解析类似 "Upgrading to 2.6.6 (from 2.6.5)..." 的输出
func ParseSelfUpdateOutput(output string) (oldVersion, newVersion string, err error) {
	output = strings.TrimSpace(output)

	// 格式1: "Upgrading to 2.6.6 (from 2.6.5)..."
	re := regexp.MustCompile(`Upgrading to (\S+)\s*\(from (\S+)\)`)
	if matches := re.FindStringSubmatch(output); len(matches) >= 3 {
		newVersion = matches[1]
		oldVersion = matches[2]
		return
	}

	// 格式2: "You are already using composer version 2.6.6."
	re2 := regexp.MustCompile(`using composer version (\S+?)[.\s]*$`)
	if matches := re2.FindStringSubmatch(output); len(matches) >= 2 {
		newVersion = strings.TrimRight(matches[1], ".")
		oldVersion = newVersion
		return
	}

	// 格式3: "Successfully updated to 2.6.6"
	re3 := regexp.MustCompile(`updated to (\S+)`)
	if matches := re3.FindStringSubmatch(output); len(matches) >= 2 {
		newVersion = matches[1]
		return
	}

	return "", "", fmt.Errorf("无法解析self-update输出: %s", output)
}

// ParseCheckPlatformReqsOutput 解析 `composer check-platform-reqs` 的文本输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []PlatformRequirement: 平台需求列表
//   - error: 解析错误
func ParseCheckPlatformReqsOutput(output string) ([]PlatformRequirement, error) {
	var reqs []PlatformRequirement
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			req := PlatformRequirement{
				Package:  fields[0],
				Version:  fields[1],
			}
			if len(fields) >= 3 {
				req.Status = fields[2]
			}
			reqs = append(reqs, req)
		}
	}
	return reqs, nil
}

// ParseFundOutput 解析 `composer fund` 的文本输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []FundInfo: 资金信息列表
func ParseFundOutput(output string) []FundInfo {
	var funds []FundInfo
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// 简单解析，格式可能为 "vendor/package - <url>"
		parts := strings.SplitN(line, "-", 2)
		if len(parts) >= 1 {
			info := FundInfo{
				Package: strings.TrimSpace(parts[0]),
			}
			if len(parts) >= 2 {
				info.URL = strings.TrimSpace(parts[1])
			}
			funds = append(funds, info)
		}
	}
	return funds
}

// ParseDiagnoseOutputAsChecks 解析 `composer diagnose` 的输出为检查项列表
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []DiagnoseCheck: 检查项列表
func ParseDiagnoseOutputAsChecks(output string) []DiagnoseCheck {
	result := ParseDiagnoseOutput(output)
	return result.Checks
}

// ParseConfigOutput 解析 `composer config` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - string: 配置值
//   - error: 解析错误
func ParseConfigOutput(output string) (string, error) {
	return strings.TrimSpace(output), nil
}

// ParseAboutOutput 解析 `composer about` 的输出
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - map[string]string: 关于信息
func ParseAboutOutput(output string) map[string]string {
	info := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return info
}

// ExtractVersionFromOutput 从任意输出中提取版本号
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - string: 版本号（如果找到）
//   - bool: 是否找到版本号
func ExtractVersionFromOutput(output string) (string, bool) {
	re := regexp.MustCompile(`(\d+\.\d+\.\d+(?:-[a-zA-Z0-9.]+)?)`)
	matches := re.FindStringSubmatch(output)
	if len(matches) >= 2 {
		return matches[1], true
	}
	return "", false
}

// ExtractPackageNamesFromOutput 从输出中提取包名列表
//
// 参数：
//   - output: 命令输出
//
// 返回值：
//   - []string: 包名列表
func ExtractPackageNamesFromOutput(output string) []string {
	re := regexp.MustCompile(`([a-z0-9]([_.-]?[a-z0-9]+)*/[a-z0-9](([_.]|-{1,2})?[a-z0-9]+)*)`)
	matches := re.FindAllString(output, -1)
	// 去重
	seen := make(map[string]bool)
	var result []string
	for _, m := range matches {
		if !seen[m] {
			seen[m] = true
			result = append(result, m)
		}
	}
	return result
}

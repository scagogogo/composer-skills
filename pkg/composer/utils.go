package composer

import "sort"

// buildOptionsArgs 将选项 map 转换为确定性的命令行参数切片。
//
// Go 中 map 的遍历顺序是随机的，如果直接遍历 options map 来构造
// composer 命令的参数，会导致：
//   - 同样的输入产生不同顺序的命令行参数（不可复现）
//   - 基于 mock 命令字符串匹配的单元测试 flaky（时而通过时而失败）
//   - 命令缓存键不稳定
//
// 因此对选项 key 排序后再生成参数，保证命令构造的确定性。
//
// 用法：
//
//	args := []string{"remove"}
//	args = append(args, buildOptionsArgs(options)...)
//	args = append(args, packageName)
//	c.Run(args...)
func buildOptionsArgs(options map[string]string) []string {
	keys := make([]string, 0, len(options))
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	args := make([]string, 0, len(keys))
	for _, key := range keys {
		value := options[key]
		if value == "" {
			args = append(args, "--"+key)
		} else {
			args = append(args, "--"+key+"="+value)
		}
	}
	return args
}

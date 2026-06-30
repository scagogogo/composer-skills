package composer

import (
	"reflect"
	"testing"
)

func TestBuildOptionsArgs_Empty(t *testing.T) {
	got := buildOptionsArgs(map[string]string{})
	if len(got) != 0 {
		t.Errorf("空 options 应返回空切片, got %v", got)
	}
}

func TestBuildOptionsArgs_FlagOnly(t *testing.T) {
	got := buildOptionsArgs(map[string]string{"dev": ""})
	want := []string{"--dev"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBuildOptionsArgs_KeyValue(t *testing.T) {
	got := buildOptionsArgs(map[string]string{"prefer-source": "dist"})
	want := []string{"--prefer-source=dist"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// TestBuildOptionsArgs_DeterministicOrder 验证多选项时输出顺序是确定性的（按 key 排序）。
// 这是修复 flaky test 的核心保证：无论 map 内部如何乱序，生成的参数顺序固定。
func TestBuildOptionsArgs_DeterministicOrder(t *testing.T) {
	options := map[string]string{
		"no-update":  "",
		"dev":        "",
		"prefer-src": "",
		"with-deps":  "",
	}
	// 期望按 key 字典序排序
	want := []string{
		"--dev",
		"--no-update",
		"--prefer-src",
		"--with-deps",
	}

	// 多次调用都应得到完全相同的结果
	for i := 0; i < 20; i++ {
		got := buildOptionsArgs(options)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("第 %d 次调用: got %v, want %v (顺序不确定导致 flaky)", i, got, want)
		}
	}
}

// TestBuildOptionsArgs_MixedFlagsAndValues 验证 flag 和 key=value 混合时仍按 key 排序
func TestBuildOptionsArgs_MixedFlagsAndValues(t *testing.T) {
	options := map[string]string{
		"optimize":     "",           // flag
		"prefer-dist":  "",           // flag
		"with-requires": "pkg:1.0",   // key=value
		"dev":          "",           // flag
	}
	want := []string{
		"--dev",
		"--optimize",
		"--prefer-dist",
		"--with-requires=pkg:1.0",
	}
	got := buildOptionsArgs(options)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/scagogogo/composer-crawler/pkg/client"
	"github.com/scagogogo/composer-crawler/pkg/domain"
)

// SecurityMonitor 表示安全监控器
type SecurityMonitor struct {
	client          *client.ComposerClient
	dataDir         string
	trackedPackages []string
}

// NewSecurityMonitor 创建一个新的安全监控器
func NewSecurityMonitor(dataDir string, packages []string) *SecurityMonitor {
	return &SecurityMonitor{
		client:          client.NewComposerClient(30 * time.Second),
		dataDir:         dataDir,
		trackedPackages: packages,
	}
}

// FetchAdvisories 获取并保存安全公告
func (m *SecurityMonitor) FetchAdvisories() error {
	// 获取所有安全公告
	fmt.Println("获取所有安全公告...")
	advisories, err := m.client.GetSecurityAdvisories()
	if err != nil {
		return fmt.Errorf("获取安全公告失败: %w", err)
	}

	// 确保数据目录存在
	if err := os.MkdirAll(m.dataDir, 0755); err != nil {
		return fmt.Errorf("创建数据目录失败: %w", err)
	}

	// 保存完整的公告数据
	timestamp := time.Now().Format("20060102-150405")
	fullDataPath := filepath.Join(m.dataDir, fmt.Sprintf("all_advisories_%s.json", timestamp))

	data, err := json.MarshalIndent(advisories, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化公告数据失败: %w", err)
	}

	if err := os.WriteFile(fullDataPath, data, 0644); err != nil {
		return fmt.Errorf("保存公告数据失败: %w", err)
	}

	fmt.Printf("已保存所有公告数据到 %s\n", fullDataPath)

	// 筛选跟踪的包的公告
	trackedAdvisories := make(map[string][]*domain.Advisory)
	for _, pkgName := range m.trackedPackages {
		if advisories, ok := advisories.Advisories[pkgName]; ok {
			trackedAdvisories[pkgName] = advisories
			fmt.Printf("发现 %s 有 %d 个安全公告\n", pkgName, len(advisories))
		}
	}

	// 如果有跟踪的包存在公告，保存单独的报告
	if len(trackedAdvisories) > 0 {
		trackedDataPath := filepath.Join(m.dataDir, fmt.Sprintf("tracked_advisories_%s.json", timestamp))

		data, err := json.MarshalIndent(trackedAdvisories, "", "  ")
		if err != nil {
			return fmt.Errorf("序列化跟踪包的公告数据失败: %w", err)
		}

		if err := os.WriteFile(trackedDataPath, data, 0644); err != nil {
			return fmt.Errorf("保存跟踪包的公告数据失败: %w", err)
		}

		fmt.Printf("已保存跟踪包的公告数据到 %s\n", trackedDataPath)
	} else {
		fmt.Println("跟踪的包中没有发现安全公告")
	}

	return nil
}

func main() {
	// 要跟踪的包列表
	trackedPackages := []string{
		"symfony/symfony",
		"laravel/framework",
		"guzzlehttp/guzzle",
		"monolog/monolog",
		"phpunit/phpunit",
	}

	// 创建安全监控器
	monitor := NewSecurityMonitor("security_data", trackedPackages)

	// 获取并保存安全公告
	if err := monitor.FetchAdvisories(); err != nil {
		log.Fatalf("监控安全公告失败: %v", err)
	}

	fmt.Println("\n安全监控完成。可以设置此脚本通过 cron 作业定期运行，以持续监控安全公告。")
	fmt.Println("示例 cron 表达式（每天运行一次）: 0 0 * * * /path/to/security_monitor")
}

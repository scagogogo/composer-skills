package composer

import (
	"errors"
	"strings"
	"testing"
)

func TestVersionCommands(t *testing.T) {
	// Reset mock outputs before test
	ClearMockOutputs()

	// Create composer with any executable path (won't be used)
	composer, err := New(Options{ExecutablePath: "/path/to/composer"})
	if err != nil {
		t.Fatalf("创建Composer实例失败: %v", err)
	}

	// 测试GetVersion命令
	t.Run("GetVersion", func(t *testing.T) {
		// 模拟版本命令输出
		ClearMockOutputs()
		SetupMockOutput("--version", "Composer version 2.5.7 2023-11-10 10:32:06", nil)

		version, err := composer.GetVersion()
		if err != nil {
			t.Errorf("GetVersion执行失败: %v", err)
		}

		if version != "2.5.7" {
			t.Errorf("版本应为'2.5.7'，实际为'%s'", version)
		}
	})

	// 测试SelfUpdate命令
	t.Run("SelfUpdate", func(t *testing.T) {
		// 模拟更新命令输出
		ClearMockOutputs()
		updateOutput := `Updating to version 2.6.0
Downloading...
Composer successfully updated to version 2.6.0.`
		SetupMockOutput("self-update", updateOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update")
		if err != nil {
			t.Errorf("执行self-update命令失败: %v", err)
		}

		if !strings.Contains(output, "Composer successfully updated") {
			t.Errorf("更新输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate预览
	t.Run("SelfUpdatePreview", func(t *testing.T) {
		// 模拟预览更新命令输出
		ClearMockOutputs()
		previewOutput := `The latest version is 2.6.0
Updates to Composer 2.6.0:
- Fixed issue with package resolution
- Added new features for plugin developers
- Improved performance of install command`
		SetupMockOutput("self-update --dry-run", previewOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update", "--dry-run")
		if err != nil {
			t.Errorf("执行self-update --dry-run命令失败: %v", err)
		}

		if !strings.Contains(output, "The latest version is 2.6.0") {
			t.Errorf("更新预览输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate到快照版本
	t.Run("SelfUpdateSnapshot", func(t *testing.T) {
		// 模拟更新到快照版本的输出
		ClearMockOutputs()
		snapshotOutput := `Updating to snapshot version 2.6.0-alpha3
Downloading...
Composer successfully updated to version 2.6.0-alpha3.`
		SetupMockOutput("self-update --snapshot", snapshotOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update", "--snapshot")
		if err != nil {
			t.Errorf("执行self-update --snapshot命令失败: %v", err)
		}

		if !strings.Contains(output, "Composer successfully updated to version 2.6.0-alpha3") {
			t.Errorf("快照更新输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate回滚
	t.Run("SelfUpdateRollback", func(t *testing.T) {
		// 模拟回滚命令输出
		ClearMockOutputs()
		rollbackOutput := `Rolling back to version 2.5.5
Downloading...
Composer successfully rolled back to version 2.5.5.`
		SetupMockOutput("self-update --rollback", rollbackOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update", "--rollback")
		if err != nil {
			t.Errorf("执行self-update --rollback命令失败: %v", err)
		}

		if !strings.Contains(output, "Composer successfully rolled back") {
			t.Errorf("回滚输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate到特定版本
	t.Run("SelfUpdateToVersion", func(t *testing.T) {
		// 模拟更新到特定版本的输出
		ClearMockOutputs()
		versionOutput := `Updating to version 2.4.4
Downloading...
Composer successfully updated to version 2.4.4.`
		SetupMockOutput("self-update 2.4.4", versionOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update", "2.4.4")
		if err != nil {
			t.Errorf("执行self-update 2.4.4命令失败: %v", err)
		}

		if !strings.Contains(output, "Composer successfully updated to version 2.4.4") {
			t.Errorf("更新到特定版本的输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate到特定通道
	t.Run("SelfUpdateToChannel", func(t *testing.T) {
		// 模拟更新到特定通道的输出
		ClearMockOutputs()
		channelOutput := `Updating to stable channel
Downloading...
Composer successfully updated to version 2.5.7 (stable channel).`
		SetupMockOutput("self-update --stable", channelOutput, nil)

		// 直接使用Run方法
		output, err := composer.Run("self-update", "--stable")
		if err != nil {
			t.Errorf("执行self-update --stable命令失败: %v", err)
		}

		if !strings.Contains(output, "Composer successfully updated") || !strings.Contains(output, "stable channel") {
			t.Errorf("更新到特定通道的输出不符合预期: %s", output)
		}
	})

	// 测试SelfUpdate错误情况
	t.Run("SelfUpdateError", func(t *testing.T) {
		// 模拟更新失败的情况
		ClearMockOutputs()
		errorMsg := "Failed to download new version"
		SetupMockOutput("self-update", errorMsg, errors.New("update failed"))

		// 直接使用Run方法
		_, err := composer.Run("self-update")
		if err == nil {
			t.Errorf("更新失败应返回错误")
		}
	})
}

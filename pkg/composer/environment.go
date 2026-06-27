package composer

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// EnvironmentVariable 表示 Composer 环境变量
type EnvironmentVariable string

// 定义常用的 Composer 环境变量
const (
	// COMPOSER_HOME 设置 Composer 主目录
	EnvComposerHome EnvironmentVariable = "COMPOSER_HOME"
	// COMPOSER_CACHE_DIR 设置 Composer 缓存目录
	EnvComposerCacheDir EnvironmentVariable = "COMPOSER_CACHE_DIR"
	// COMPOSER_PROCESS_TIMEOUT 设置 Composer 进程超时时间
	EnvComposerProcessTimeout EnvironmentVariable = "COMPOSER_PROCESS_TIMEOUT"
	// COMPOSER_ALLOW_SUPERUSER 允许以 root 身份运行 Composer
	EnvComposerAllowSuperuser EnvironmentVariable = "COMPOSER_ALLOW_SUPERUSER"
	// COMPOSER_MEMORY_LIMIT 设置 PHP 内存限制
	EnvComposerMemoryLimit EnvironmentVariable = "COMPOSER_MEMORY_LIMIT"
	// COMPOSER_DISABLE_XDEBUG_WARN 禁用 XDebug 警告
	EnvComposerDisableXdebugWarn EnvironmentVariable = "COMPOSER_DISABLE_XDEBUG_WARN"
	// COMPOSER_NO_INTERACTION 禁用交互提示
	EnvComposerNoInteraction EnvironmentVariable = "COMPOSER_NO_INTERACTION"
	// COMPOSER_VENDOR_DIR 设置 vendor 目录位置
	EnvComposerVendorDir EnvironmentVariable = "COMPOSER_VENDOR_DIR"
	// COMPOSER_BIN_DIR 设置 bin 目录位置
	EnvComposerBinDir EnvironmentVariable = "COMPOSER_BIN_DIR"
	// COMPOSER_CAFILE 设置 CA 证书文件
	EnvComposerCafile EnvironmentVariable = "COMPOSER_CAFILE"
	// COMPOSER_NO_DEV 不安装开发依赖
	EnvComposerNoDev EnvironmentVariable = "COMPOSER_NO_DEV"
	// COMPOSER_DISCARD_CHANGES 控制修改的处理
	EnvComposerDiscardChanges EnvironmentVariable = "COMPOSER_DISCARD_CHANGES"
	// COMPOSER_HTACCESS_PROTECT 控制 htaccess 保护
	EnvComposerHtaccessProtect EnvironmentVariable = "COMPOSER_HTACCESS_PROTECT"
	// COMPOSER_MIRROR_PATH_REPOS 控制路径仓库的镜像策略
	EnvComposerMirrorPathRepos EnvironmentVariable = "COMPOSER_MIRROR_PATH_REPOS"
)

// SetEnvVariable 设置 Composer 环境变量
func SetEnvVariable(name EnvironmentVariable, value string) error {
	return os.Setenv(string(name), value)
}

// GetEnvVariable 获取 Composer 环境变量值
func GetEnvVariable(name EnvironmentVariable) string {
	return os.Getenv(string(name))
}

// SetProcessTimeout 设置 Composer 进程超时时间（秒）
func SetProcessTimeout(seconds int) error {
	return os.Setenv(string(EnvComposerProcessTimeout), strconv.Itoa(seconds))
}

// EnableSuperuser 允许以 root 身份运行 Composer
func EnableSuperuser() error {
	return os.Setenv(string(EnvComposerAllowSuperuser), "1")
}

// DisableSuperuser 禁止以 root 身份运行 Composer
func DisableSuperuser() error {
	return os.Unsetenv(string(EnvComposerAllowSuperuser))
}

// SetMemoryLimit 设置 PHP 内存限制
func SetMemoryLimit(limit string) error {
	return os.Setenv(string(EnvComposerMemoryLimit), limit)
}

// DisableInteraction 禁用交互提示
func DisableInteraction() error {
	return os.Setenv(string(EnvComposerNoInteraction), "1")
}

// EnableInteraction 启用交互提示
func EnableInteraction() error {
	return os.Unsetenv(string(EnvComposerNoInteraction))
}

// SetVendorDir 设置 vendor 目录位置
func SetVendorDir(path string) error {
	return os.Setenv(string(EnvComposerVendorDir), path)
}

// SetBinDir 设置 bin 目录位置
func SetBinDir(path string) error {
	return os.Setenv(string(EnvComposerBinDir), path)
}

// SetCaFile 设置 CA 证书文件
func SetCaFile(path string) error {
	return os.Setenv(string(EnvComposerCafile), path)
}

// DisableDev 禁用开发依赖
func DisableDev() error {
	return os.Setenv(string(EnvComposerNoDev), "1")
}

// EnableDev 启用开发依赖
func EnableDev() error {
	return os.Unsetenv(string(EnvComposerNoDev))
}

// SetDiscardChanges 设置是否丢弃更改
func SetDiscardChanges(value string) error {
	return os.Setenv(string(EnvComposerDiscardChanges), value)
}

// GetComposerPath 获取 Composer 可执行文件路径
func GetComposerPath() (string, error) {
	path, err := exec.LookPath("composer")
	if err != nil {
		path, err = exec.LookPath("composer.phar")
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// GetEnvironmentInfo 获取 Composer 环境信息
func (c *Composer) GetEnvironmentInfo() (map[string]string, error) {
	output, err := c.Run("config", "--list")
	if err != nil {
		return nil, err
	}

	// 解析输出
	info := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			info[key] = value
		}
	}

	return info, nil
}

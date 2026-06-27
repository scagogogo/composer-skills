package composer

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/scagogogo/composer-skills/pkg/detector"
	"github.com/scagogogo/composer-skills/pkg/installer"
)

// 常见错误
var (
	ErrComposerNotFound     = errors.New("未找到composer可执行文件")
	ErrComposerInstallation = errors.New("安装composer失败")
	ErrCommandExecution     = errors.New("执行composer命令失败")
	ErrUnsupportedOS        = errors.New("不支持的操作系统")
	ErrInitFailed           = errors.New("初始化Composer失败")
	ErrSelfUpdateFailed     = errors.New("composer自更新失败")
	ErrRequirePackageFailed = errors.New("安装包失败")
	ErrInstallFailed        = errors.New("composer install失败")
	ErrUpdateFailed         = errors.New("composer update失败")
	ErrRemoveFailed         = errors.New("删除包失败")
	ErrValidateFailed       = errors.New("composer validate失败")
	ErrDumpAutoloadFailed   = errors.New("composer dump-autoload失败")
	ErrCreateProjectFailed  = errors.New("创建项目失败")
	ErrRunScriptFailed      = errors.New("运行脚本失败")
	ErrShowPackageFailed    = errors.New("查看包失败")
	ErrSearchFailed         = errors.New("搜索失败")
	ErrInitProjectFailed    = errors.New("初始化项目失败")
	ErrGetComposerHome      = errors.New("获取Composer主目录失败")
)

// mockCommandOutput 存储不同命令的预定义输出，用于测试目的
var mockCommandOutput = map[string]struct {
	output string
	err    error
}{}

// mockMutexSimple 保护mockCommandOutput的并发访问
var mockMutexSimple sync.RWMutex

// testMode 标识是否处于测试模式，测试模式下会跳过某些验证
var testMode = false

// SetupMockOutput 为特定命令设置模拟输出
//
// 参数：
//   - command: 需要模拟的命令字符串
//   - output: 模拟的命令输出
//   - err: 模拟的错误信息
//
// 功能说明：
//
//	该方法用于测试目的，可以为特定的Composer命令设置预定义的输出和错误信息。
//	这样在单元测试中可以模拟命令执行而不需要实际执行Composer命令。
//
// 用法示例：
//
//	composer.SetupMockOutput("composer require", "Package operations: 1 install, 0 updates, 0 removals", nil)
func SetupMockOutput(command string, output string, err error) {
	mockMutexSimple.Lock()
	defer mockMutexSimple.Unlock()
	testMode = true // 启用测试模式
	mockCommandOutput[command] = struct {
		output string
		err    error
	}{
		output: output,
		err:    err,
	}
}

// ClearMockOutputs 清除所有模拟输出
//
// 功能说明：
//
//	该方法用于清除之前通过SetupMockOutput设置的所有模拟输出。
//	通常在测试前后使用，以确保测试环境的干净。
//
// 用法示例：
//
//	composer.ClearMockOutputs()
func ClearMockOutputs() {
	mockMutexSimple.Lock()
	defer mockMutexSimple.Unlock()
	testMode = true // 保持测试模式
	mockCommandOutput = map[string]struct {
		output string
		err    error
	}{}
}

// getMockOutput 如果存在模拟输出，则返回命令的模拟输出
func getMockOutput(args ...string) (string, error, bool) {
	if len(args) == 0 {
		return "", nil, false
	}

	mockMutexSimple.RLock()
	defer mockMutexSimple.RUnlock()

	// Try to match the exact command
	cmdStr := args[0]
	for i := 1; i < len(args); i++ {
		cmdStr += " " + args[i]
	}

	if mock, ok := mockCommandOutput[cmdStr]; ok {
		return mock.output, mock.err, true
	}

	// If no exact match, try to match just the first argument (the command)
	if mock, ok := mockCommandOutput[args[0]]; ok {
		return mock.output, mock.err, true
	}

	return "", nil, false
}

// Composer 封装PHP Composer的功能
type Composer struct {
	// composer可执行文件的路径
	executablePath string
	// 工作目录
	workingDir string
	// 是否自动安装
	autoInstall bool
	// 自定义安装器
	installer *installer.Installer
	// 自定义检测器
	detector *detector.Detector
	// 环境变量
	env []string
	// 默认超时时间
	defaultTimeout time.Duration
}

// Options 用于自定义Composer实例的选项
type Options struct {
	// composer可执行文件的路径
	ExecutablePath string
	// 工作目录
	WorkingDir string
	// 是否在未找到composer时自动安装
	AutoInstall bool
	// 自定义安装器
	Installer *installer.Installer
	// 自定义检测器
	Detector *detector.Detector
	// 环境变量
	Env []string
	// 默认超时时间
	DefaultTimeout time.Duration
}

// DefaultOptions 返回默认选项
//
// 返回值：
//   - Options: 包含默认配置的选项结构体
//
// 功能说明：
//
//	该方法返回Composer实例的默认配置选项。
//	默认配置包括：
//	- 使用当前目录作为工作目录
//	- 自动安装Composer（如果未找到）
//	- 10分钟的默认超时时间
//
// 用法示例：
//
//	options := composer.DefaultOptions()
//	comp, err := composer.New(options)
func DefaultOptions() Options {
	return Options{
		WorkingDir:     "",
		AutoInstall:    true,
		DefaultTimeout: 10 * time.Minute,
	}
}

// New 创建一个新的Composer实例
//
// 参数：
//   - options: 自定义Composer实例的选项
//
// 返回值：
//   - *Composer: 创建的Composer实例
//   - error: 如果创建过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法会创建一个新的Composer实例。如果未指定可执行文件路径，它会尝试
//	检测系统中已安装的Composer。如果未找到且autoInstall设置为true，则会
//	尝试自动安装Composer。
//
// 用法示例：
//
//	options := composer.DefaultOptions()
//	options.WorkingDir = "/path/to/project"
//	comp, err := composer.New(options)
//	if err != nil {
//	    log.Fatalf("初始化Composer失败: %v", err)
//	}
func New(options Options) (*Composer, error) {
	c := &Composer{
		executablePath: options.ExecutablePath,
		workingDir:     options.WorkingDir,
		autoInstall:    options.AutoInstall,
		installer:      options.Installer,
		detector:       options.Detector,
		env:            options.Env,
		defaultTimeout: options.DefaultTimeout,
	}

	// 如果未提供安装器，使用默认安装器
	if c.installer == nil {
		c.installer = installer.NewInstaller(installer.DefaultConfig())
	}

	// 如果未提供检测器，使用默认检测器
	if c.detector == nil {
		c.detector = detector.NewDetector()
	}

	// 如果未指定可执行文件路径，则尝试检测
	if c.executablePath == "" {
		execPath, err := c.detector.Detect()
		if err != nil {
			if c.autoInstall {
				// 尝试安装composer
				if err := c.installer.Install(); err != nil {
					return nil, fmt.Errorf("%w: %v", ErrComposerInstallation, err)
				}

				// 重新尝试检测
				execPath, err = c.detector.Detect()
				if err != nil {
					return nil, fmt.Errorf("%w: %v", ErrComposerNotFound, err)
				}
			} else {
				return nil, fmt.Errorf("%w: %v", ErrComposerNotFound, err)
			}
		}
		c.executablePath = execPath
	} else {
		// 如果指定了可执行文件路径，验证文件是否存在（测试模式下跳过）
		if !testMode {
			if _, err := os.Stat(c.executablePath); os.IsNotExist(err) {
				return nil, fmt.Errorf("%w: 指定的可执行文件不存在: %s", ErrComposerNotFound, c.executablePath)
			}
		}
	}

	return c, nil
}

// SetWorkingDir 设置composer命令的工作目录
//
// 参数：
//   - dir: 要设置的工作目录路径
//
// 功能说明：
//
//	该方法用于设置Composer操作的工作目录。所有Composer命令将在此目录下执行。
//	适用于在不同的项目目录之间切换时使用。
//
// 用法示例：
//
//	comp.SetWorkingDir("/path/to/php/project")
func (c *Composer) SetWorkingDir(dir string) {
	c.workingDir = dir
}

// SetEnv 设置环境变量
//
// 参数：
//   - env: 环境变量数组，格式为["KEY=VALUE", ...]
//
// 功能说明：
//
//	该方法用于设置执行Composer命令时的环境变量。
//	可以用来配置HTTP代理、身份验证信息或其他影响Composer行为的环境变量。
//
// 用法示例：
//
//	// 设置HTTP代理
//	comp.SetEnv([]string{
//	    "HTTP_PROXY=http://proxy.example.com:8080",
//	    "HTTPS_PROXY=http://proxy.example.com:8080",
//	    "COMPOSER_HOME=/custom/composer/home"
//	})
func (c *Composer) SetEnv(env []string) {
	c.env = env
}

// Run 执行composer命令并返回输出
//
// 参数：
//   - args: 命令参数，第一个参数是composer子命令
//
// 返回值：
//   - string: 命令的标准输出
//   - error: 如果命令执行失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法是SDK的核心方法，用于执行任意Composer命令。
//	它会调用系统上的Composer可执行文件，并传递指定的参数。
//	默认使用10分钟的超时时间。
//
// 用法示例：
//
//	// 执行"composer show"命令
//	output, err := comp.Run("show")
//	if err != nil {
//	    log.Fatalf("执行命令失败: %v", err)
//	}
//	fmt.Println(output)
//
//	// 执行带参数的命令
//	output, err = comp.Run("require", "symfony/console", "--dev")
func (c *Composer) Run(args ...string) (string, error) {
	return c.RunWithTimeout(c.defaultTimeout, args...)
}

// RunWithTimeout 在指定超时时间内执行composer命令
//
// 参数：
//   - timeout: 命令执行的最大超时时间
//   - args: 命令参数，第一个参数是composer子命令
//
// 返回值：
//   - string: 命令的标准输出
//   - error: 如果命令执行失败或超时，则返回相应的错误信息
//
// 功能说明：
//
//	该方法与Run类似，但允许指定自定义的超时时间。
//	对于可能长时间运行的命令特别有用。
//
// 用法示例：
//
//	// 执行可能需要很长时间的安装命令，设置30分钟超时
//	output, err := comp.RunWithTimeout(30*time.Minute, "install")
//	if err != nil {
//	    log.Fatalf("安装超时或失败: %v", err)
//	}
func (c *Composer) RunWithTimeout(timeout time.Duration, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return c.RunWithContext(ctx, args...)
}

// RunWithContext 在指定上下文中执行composer命令
//
// 参数：
//   - ctx: 上下文，可用于取消或设置超时
//   - args: 命令参数，第一个参数是composer子命令
//
// 返回值：
//   - string: 命令的标准输出
//   - error: 如果命令执行失败，则返回相应的错误信息
//
// 功能说明：
//
//	该方法提供对命令执行的最大控制权，允许使用自定义上下文。
//	上下文可以用于超时控制、取消执行或传递其他值。
//
// 用法示例：
//
//	// 创建可以手动取消的上下文
//	ctx, cancel := context.WithCancel(context.Background())
//
//	// 在另一个goroutine中根据条件取消
//	go func() {
//	    time.Sleep(5 * time.Second)
//	    cancel()
//	}()
//
//	// 执行命令
//	output, err := comp.RunWithContext(ctx, "update")
//	if err != nil {
//	    if errors.Is(err, context.Canceled) {
//	        fmt.Println("命令被取消")
//	    } else {
//	        log.Fatalf("执行命令失败: %v", err)
//	    }
//	}
func (c *Composer) RunWithContext(ctx context.Context, args ...string) (string, error) {
	// 检查是否有模拟输出
	if output, err, ok := getMockOutput(args...); ok {
		return output, err
	}

	// 创建命令
	cmd := exec.CommandContext(ctx, c.executablePath, args...)

	// 设置工作目录
	if c.workingDir != "" {
		cmd.Dir = c.workingDir
	}

	// 设置环境变量
	if len(c.env) > 0 {
		cmd.Env = c.env
	}

	// 执行命令并获取输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("%w: %v, output: %s", ErrCommandExecution, err, string(out))
	}

	return string(out), nil
}

// GetExecutablePath 获取composer可执行文件的路径
//
// 返回值：
//   - string: composer可执行文件的完整路径
//
// 功能说明：
//
//	该方法返回Composer实例使用的可执行文件路径。
//
// 用法示例：
//
//	execPath := comp.GetExecutablePath()
//	fmt.Printf("使用的Composer可执行文件: %s\n", execPath)
func (c *Composer) GetExecutablePath() string {
	return c.executablePath
}

// GetWorkingDir 获取当前的工作目录
//
// 返回值：
//   - string: 当前设置的工作目录路径
//
// 功能说明：
//
//	该方法返回Composer实例当前使用的工作目录。
//
// 用法示例：
//
//	workDir := comp.GetWorkingDir()
//	fmt.Printf("当前工作目录: %s\n", workDir)
func (c *Composer) GetWorkingDir() string {
	return c.workingDir
}

// GetEnv 获取当前设置的环境变量
//
// 返回值：
//   - []string: 当前设置的环境变量数组
//
// 功能说明：
//
//	该方法返回Composer实例当前使用的环境变量。
//
// 用法示例：
//
//	env := comp.GetEnv()
//	fmt.Println("当前环境变量:")
//	for _, e := range env {
//	    fmt.Println(e)
//	}
func (c *Composer) GetEnv() []string {
	return c.env
}

// IsInstalled 检查Composer是否已安装
//
// 返回值：
//   - bool: 如果Composer已安装，则返回true；否则返回false
//
// 功能说明：
//
//	该方法检查当前Composer实例是否已经有指向有效Composer可执行文件的路径。
//	它仅检查路径是否存在，不会验证可执行文件是否可正常工作。
//
// 用法示例：
//
//	if comp.IsInstalled() {
//	    fmt.Println("Composer已安装")
//	} else {
//	    fmt.Println("Composer未安装")
//	}
func (c *Composer) IsInstalled() bool {
	return c.executablePath != ""
}

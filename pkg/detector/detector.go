package detector

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// 常见错误
var (
	ErrExecutableNotFound = errors.New("未找到composer可执行文件")
)

// Detector 是用于检测系统中Composer安装情况的工具
type Detector struct {
	// 可能的composer可执行文件路径
	possiblePaths []string
}

// NewDetector 创建一个新的Composer检测器
func NewDetector() *Detector {
	return &Detector{
		possiblePaths: defaultPossiblePaths(),
	}
}

// SetPossiblePaths 设置可能的Composer路径
func (d *Detector) SetPossiblePaths(paths []string) {
	d.possiblePaths = paths
}

// AddPossiblePath 添加可能的Composer路径
func (d *Detector) AddPossiblePath(path string) {
	d.possiblePaths = append(d.possiblePaths, path)
}

// Detect 尝试在系统中检测Composer可执行文件
// 返回Composer可执行文件的完整路径，如果未找到则返回错误
func (d *Detector) Detect() (string, error) {
	// 首先检查环境变量中是否指定了composer
	if envPath := os.Getenv("COMPOSER_PATH"); envPath != "" {
		if isExecutable(envPath) {
			return envPath, nil
		}
	}

	// 检查 COMPOSER_HOME 环境变量下的 composer
	if home := os.Getenv("COMPOSER_HOME"); home != "" {
		homeBin := home + "/vendor/bin/composer"
		if isExecutable(homeBin) {
			return homeBin, nil
		}
	}

	// 然后检查默认可能的路径
	for _, path := range d.possiblePaths {
		if isExecutable(path) {
			return path, nil
		}
	}

	// 尝试使用which/where命令查找
	executable := "composer"
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", executable)
	} else {
		cmd = exec.Command("which", executable)
	}

	output, err := cmd.Output()
	if err == nil {
		path := strings.TrimSpace(string(output))
		if isExecutable(path) {
			return path, nil
		}
	}

	// 尝试查找 composer.phar
	pharPaths := []string{"composer.phar", "./composer.phar"}
	if runtime.GOOS != "windows" {
		if home, err := os.UserHomeDir(); err == nil {
			pharPaths = append(pharPaths, home+"/composer.phar")
		}
	}
	for _, pharPath := range pharPaths {
		if fileExists(pharPath) {
			return pharPath, nil
		}
	}

	return "", ErrExecutableNotFound
}

// DetectVerbose attempts to detect Composer and returns detailed information
type DetectionResult struct {
	// Path is the detected Composer executable path
	Path string
	// Method describes how Composer was detected (e.g., "env:COMPOSER_PATH", "which", "default_path")
	Method string
	// IsPhar indicates if the detected path is a .phar file
	IsPhar bool
}

// DetectVerbose detects Composer and returns detailed detection information
func (d *Detector) DetectVerbose() (*DetectionResult, error) {
	// Check COMPOSER_PATH environment variable
	if envPath := os.Getenv("COMPOSER_PATH"); envPath != "" {
		if isExecutable(envPath) {
			return &DetectionResult{
				Path:   envPath,
				Method: "env:COMPOSER_PATH",
				IsPhar: strings.HasSuffix(envPath, ".phar"),
			}, nil
		}
	}

	// Check COMPOSER_HOME
	if home := os.Getenv("COMPOSER_HOME"); home != "" {
		homeBin := home + "/vendor/bin/composer"
		if isExecutable(homeBin) {
			return &DetectionResult{
				Path:   homeBin,
				Method: "env:COMPOSER_HOME/vendor/bin",
				IsPhar: false,
			}, nil
		}
	}

	// Check default paths
	for _, path := range d.possiblePaths {
		if isExecutable(path) {
			return &DetectionResult{
				Path:   path,
				Method: "default_path",
				IsPhar: strings.HasSuffix(path, ".phar"),
			}, nil
		}
	}

	// Try which/where
	executable := "composer"
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", executable)
	} else {
		cmd = exec.Command("which", executable)
	}

	output, err := cmd.Output()
	if err == nil {
		path := strings.TrimSpace(string(output))
		if isExecutable(path) {
			return &DetectionResult{
				Path:   path,
				Method: "which/where",
				IsPhar: strings.HasSuffix(path, ".phar"),
			}, nil
		}
	}

	return nil, ErrExecutableNotFound
}

// IsInstalled 检查Composer是否已安装
func (d *Detector) IsInstalled() bool {
	_, err := d.Detect()
	return err == nil
}

// 检查路径是否指向一个可执行文件
func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	// 检查是否是普通文件
	if info.Mode().IsRegular() {
		// 在Unix系统上，检查是否有执行权限
		if runtime.GOOS != "windows" {
			return (info.Mode().Perm() & 0111) != 0
		}
		return true // 在Windows上无法直接检查执行权限，假定文件可执行
	}

	return false
}

// fileExists checks if a file exists at the given path
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

// 根据不同的操作系统返回可能的composer路径
func defaultPossiblePaths() []string {
	var paths []string

	// 添加平台特定的路径
	paths = append(paths, getPlatformSpecificPaths()...)

	// 添加当前目录中的通用路径
	paths = append(paths, "./composer", "./composer.phar")

	return paths
}

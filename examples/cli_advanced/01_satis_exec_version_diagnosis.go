package cli_advanced

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01Satis 演示如何使用Satis私有仓库
func Example01Satis() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 示例1：初始化Satis配置
	fmt.Println("1. 初始化Satis配置...")
	err = c.InitSatis("my-packages", "https://packages.example.com", "/path/to/satis")
	if err != nil {
		log.Printf("初始化Satis失败: %v", err)
	} else {
		fmt.Println("Satis初始化成功")
	}

	// 示例2：创建Satis配置
	fmt.Println("\n2. 创建Satis配置...")
	err = c.CreateSatisConfig("/path/to/satis/satis.json", "my-packages", "https://packages.example.com")
	if err != nil {
		log.Printf("创建Satis配置失败: %v", err)
	} else {
		fmt.Println("Satis配置创建成功")
	}

	// 示例3：添加Satis仓库
	fmt.Println("\n3. 添加Satis仓库...")
	err = c.AddSatisRepository("/path/to/satis/satis.json", "vcs", "https://github.com/myorg/myrepo")
	if err != nil {
		log.Printf("添加Satis仓库失败: %v", err)
	} else {
		fmt.Println("Satis仓库已添加")
	}

	// 示例4：添加Satis需求
	fmt.Println("\n4. 添加Satis需求...")
	err = c.AddSatisRequire("/path/to/satis/satis.json", "myorg/mypackage", "*")
	if err != nil {
		log.Printf("添加Satis需求失败: %v", err)
	} else {
		fmt.Println("Satis需求已添加")
	}

	// 示例5：构建Satis
	fmt.Println("\n5. 构建Satis...")
	output, err := c.BuildSatis("/path/to/satis/satis.json", "/path/to/satis/web")
	if err != nil {
		log.Printf("构建Satis失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例6：启用Satis归档
	fmt.Println("\n6. 启用Satis归档...")
	err = c.EnableSatisArchive("/path/to/satis/satis.json", "zip")
	if err != nil {
		log.Printf("启用Satis归档失败: %v", err)
	} else {
		fmt.Println("Satis归档已启用")
	}

	// 示例7：更新Satis稳定性
	fmt.Println("\n7. 更新Satis稳定性...")
	err = c.UpdateSatisStability("/path/to/satis/satis.json", "stable")
	if err != nil {
		log.Printf("更新Satis稳定性失败: %v", err)
	} else {
		fmt.Println("Satis稳定性已更新")
	}
}

// Example02Exec 演示如何执行命令
func Example02Exec() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：执行二进制文件
	fmt.Println("1. 执行二进制文件...")
	output, err := c.Exec("phpunit", "--version")
	if err != nil {
		log.Printf("执行命令失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：执行自定义命令
	fmt.Println("\n2. 执行自定义命令...")
	output, err = c.ExecCommand("php", "-v")
	if err != nil {
		log.Printf("执行命令失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例3：使用指定PHP执行
	fmt.Println("\n3. 使用指定PHP执行...")
	output, err = c.ExecPHP("/usr/bin/php8.1", "phpunit", "--version")
	if err != nil {
		log.Printf("执行命令失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例4：获取可执行文件列表
	fmt.Println("\n4. 获取可执行文件列表...")
	binaries, err := c.ExecWithList()
	if err != nil {
		log.Printf("获取可执行文件列表失败: %v", err)
	} else {
		for _, bin := range binaries {
			fmt.Printf("  - %s\n", bin)
		}
	}

	// 示例5：在指定目录执行
	fmt.Println("\n5. 在指定目录执行...")
	output, err = c.ExecWithWorkingDir("phpunit", "/path/to/project", "--version")
	if err != nil {
		log.Printf("执行命令失败: %v", err)
	} else {
		fmt.Println(output)
	}
}

// Example03VersionConstraints 演示版本约束操作
func Example03VersionConstraints() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：获取包的版本信息
	fmt.Println("1. 获取包的版本信息...")
	output, err := c.GetPackageVersions("symfony/console")
	if err != nil {
		log.Printf("获取版本信息失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：锁定包的版本
	fmt.Println("\n2. 锁定包的版本...")
	err = c.LockPackageVersion("symfony/console", "v6.0.0")
	if err != nil {
		log.Printf("锁定版本失败: %v", err)
	} else {
		fmt.Println("版本已锁定")
	}

	// 示例3：更新包的版本约束
	fmt.Println("\n3. 更新包的版本约束...")
	err = c.UpdatePackageVersion("symfony/console", "^6.0", composer.CaretVersion)
	if err != nil {
		log.Printf("更新版本约束失败: %v", err)
	} else {
		fmt.Println("版本约束已更新")
	}

	// 示例4：格式化版本约束
	fmt.Println("\n4. 格式化版本约束...")
	constraint := composer.FormatVersionConstraint("1.2.3", composer.CaretVersion)
	fmt.Printf("Caret版本约束: %s\n", constraint) // ^1.2.3

	constraint = composer.FormatVersionConstraint("1.2.3", composer.TildeVersion)
	fmt.Printf("Tilde版本约束: %s\n", constraint) // ~1.2.3
}

// Example04Diagnosis 演示诊断功能
func Example04Diagnosis() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：运行诊断
	fmt.Println("1. 运行诊断...")
	output, err := c.Diagnose()
	if err != nil {
		log.Printf("诊断失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：检查环境
	fmt.Println("\n2. 检查环境...")
	output, err = c.Check()
	if err != nil {
		log.Printf("检查失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例3：查看状态
	fmt.Println("\n3. 查看状态...")
	output, err = c.Status()
	if err != nil {
		log.Printf("获取状态失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例4：获取环境信息
	fmt.Println("\n4. 获取环境信息...")
	envInfo, err := c.GetEnvironmentInfo()
	if err != nil {
		log.Printf("获取环境信息失败: %v", err)
	} else {
		for key, value := range envInfo {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	// 示例5：获取项目信息
	fmt.Println("\n5. 获取项目信息...")
	projectInfo, err := c.GetProjectInfo()
	if err != nil {
		log.Printf("获取项目信息失败: %v", err)
	} else {
		fmt.Printf("项目名称: %s\n", projectInfo.Name)
		fmt.Printf("项目描述: %s\n", projectInfo.Description)
	}
}

// Example05Archive 演示归档功能
func Example05Archive() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：归档项目
	fmt.Println("1. 归档项目...")
	output, err := c.Archive("./dist")
	if err != nil {
		log.Printf("归档失败: %v", err)
	} else {
		fmt.Printf("归档结果: %s\n", output)
	}

	// 示例2：以指定格式归档
	fmt.Println("\n2. 以指定格式归档...")
	output, err = c.ArchiveWithFormat("./dist", "zip")
	if err != nil {
		log.Printf("归档失败: %v", err)
	} else {
		fmt.Printf("归档结果: %s\n", output)
	}

	// 示例3：归档特定包
	fmt.Println("\n3. 归档特定包...")
	output, err = c.ArchivePackage("symfony/console", "v6.0.0", "./dist")
	if err != nil {
		log.Printf("归档包失败: %v", err)
	} else {
		fmt.Printf("归档结果: %s\n", output)
	}
}

// Example06Environment 演示环境变量操作
func Example06Environment() {
	// 示例1：设置Composer内存限制
	fmt.Println("1. 设置Composer内存限制...")
	err := composer.SetMemoryLimit("2G")
	if err != nil {
		fmt.Printf("设置内存限制失败: %v\n", err)
	} else {
		fmt.Println("内存限制已设置为2G")
	}

	// 示例2：设置进程超时
	fmt.Println("\n2. 设置进程超时...")
	err = composer.SetProcessTimeout(600)
	if err != nil {
		fmt.Printf("设置进程超时失败: %v\n", err)
	} else {
		fmt.Println("进程超时已设置为600秒")
	}

	// 示例3：设置Vendor目录
	fmt.Println("\n3. 设置Vendor目录...")
	err = composer.SetVendorDir("vendor-custom")
	if err != nil {
		fmt.Printf("设置Vendor目录失败: %v\n", err)
	} else {
		fmt.Println("Vendor目录已设置为vendor-custom")
	}

	// 示例4：设置Bin目录
	fmt.Println("\n4. 设置Bin目录...")
	err = composer.SetBinDir("bin-custom")
	if err != nil {
		fmt.Printf("设置Bin目录失败: %v\n", err)
	} else {
		fmt.Println("Bin目录已设置为bin-custom")
	}

	// 示例5：获取Composer路径
	fmt.Println("\n5. 获取Composer路径...")
	path, err := composer.GetComposerPath()
	if err != nil {
		fmt.Printf("获取Composer路径失败: %v\n", err)
	} else {
		fmt.Printf("Composer路径: %s\n", path)
	}

	// 示例6：启用/禁用交互模式
	fmt.Println("\n6. 启用/禁用交互模式...")
	err = composer.DisableInteraction()
	if err != nil {
		fmt.Printf("禁用交互模式失败: %v\n", err)
	}
	err = composer.EnableInteraction()
	if err != nil {
		fmt.Printf("启用交互模式失败: %v\n", err)
	}
}

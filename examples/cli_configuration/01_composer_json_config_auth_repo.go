package cli_configuration

import (
	"fmt"
	"log"

	"github.com/scagogogo/composer-skills/pkg/composer"
)

// Example01ComposerJson 演示如何操作composer.json
func Example01ComposerJson() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：读取composer.json
	fmt.Println("1. 读取composer.json...")
	composerJSON, err := c.ReadComposerJSON()
	if err != nil {
		log.Printf("读取composer.json失败: %v", err)
	} else {
		fmt.Printf("项目名称: %s\n", composerJSON.Name)
		fmt.Printf("项目描述: %s\n", composerJSON.Description)
		fmt.Printf("项目类型: %s\n", composerJSON.Type)
	}

	// 示例2：设置项目属性
	fmt.Println("\n2. 设置项目属性...")
	err = c.SetProperty("name", "myvendor/mypackage")
	if err != nil {
		log.Printf("设置项目名称失败: %v", err)
	}
	err = c.SetProperty("description", "一个很棒的PHP库")
	if err != nil {
		log.Printf("设置项目描述失败: %v", err)
	}
	err = c.SetProperty("keywords", []string{"php", "library"})
	if err != nil {
		log.Printf("设置关键词失败: %v", err)
	}

	// 示例3：添加/移除依赖
	fmt.Println("\n3. 添加/移除依赖...")
	err = c.AddRequire("symfony/console", "^6.0", false)
	if err != nil {
		log.Printf("添加依赖失败: %v", err)
	}
	err = c.AddRequire("phpunit/phpunit", "^10.0", true)
	if err != nil {
		log.Printf("添加开发依赖失败: %v", err)
	}
	err = c.RemoveRequire("old/package", false)
	if err != nil {
		log.Printf("移除依赖失败: %v", err)
	}

	// 示例4：添加/移除脚本
	fmt.Println("\n4. 添加/移除脚本...")
	err = c.AddScript("test", "phpunit", "运行测试")
	if err != nil {
		log.Printf("添加脚本失败: %v", err)
	}
	err = c.AddScript("post-install-cmd", []string{"php artisan optimize:clear"}, "安装后清理缓存")
	if err != nil {
		log.Printf("添加脚本失败: %v", err)
	}
	err = c.RemoveScript("old-script")
	if err != nil {
		log.Printf("移除脚本失败: %v", err)
	}

	// 示例5：添加自动加载配置
	fmt.Println("\n5. 添加自动加载配置...")
	err = c.AddAutoload("psr-4", "App\\", "src/", false)
	if err != nil {
		log.Printf("添加自动加载失败: %v", err)
	}
	err = c.AddAutoload("psr-4", "Tests\\", "tests/", true)
	if err != nil {
		log.Printf("添加开发自动加载失败: %v", err)
	}

	// 示例6：设置/获取配置
	fmt.Println("\n6. 设置/获取配置...")
	err = c.SetConfig("process-timeout", 500)
	if err != nil {
		log.Printf("设置配置失败: %v", err)
	}
	timeout, err := c.GetConfig("process-timeout")
	if err != nil {
		log.Printf("获取配置失败: %v", err)
	} else {
		fmt.Printf("进程超时时间: %v\n", timeout)
	}
}

// Example02Config 演示如何使用composer config命令
func Example02Config() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：获取配置值
	fmt.Println("1. 获取配置值...")
	value, err := c.GetConfigWithGlobal("vendor-dir", false)
	if err != nil {
		log.Printf("获取配置失败: %v", err)
	} else {
		fmt.Printf("vendor-dir: %s\n", value)
	}

	// 示例2：设置配置值
	fmt.Println("\n2. 设置配置值...")
	err = c.SetConfigWithGlobal("vendor-dir", "vendor", false)
	if err != nil {
		log.Printf("设置配置失败: %v", err)
	}

	// 示例3：获取全局配置
	fmt.Println("\n3. 获取全局配置...")
	value, err = c.GetConfigWithGlobal("bin-dir", true)
	if err != nil {
		log.Printf("获取全局配置失败: %v", err)
	} else {
		fmt.Printf("全局 bin-dir: %s\n", value)
	}

	// 示例4：获取Composer主目录
	fmt.Println("\n4. 获取Composer主目录...")
	home, err := c.GetComposerHome()
	if err != nil {
		log.Printf("获取Composer主目录失败: %v", err)
	} else {
		fmt.Printf("Composer主目录: %s\n", home)
	}

	// 示例5：清除缓存
	fmt.Println("\n5. 清除缓存...")
	err = c.ClearCache()
	if err != nil {
		log.Printf("清除缓存失败: %v", err)
	} else {
		fmt.Println("缓存已清除")
	}
}

// Example03Auth 演示如何管理认证配置
func Example03Auth() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	// 示例1：获取认证配置
	fmt.Println("1. 获取认证配置...")
	authConfig, err := c.GetAuthConfig()
	if err != nil {
		log.Printf("获取认证配置失败: %v", err)
	} else {
		fmt.Printf("GitHub令牌数: %d\n", len(authConfig.GitHub))
		fmt.Printf("GitLab令牌数: %d\n", len(authConfig.GitLab))
	}

	// 示例2：添加GitHub令牌
	fmt.Println("\n2. 添加GitHub令牌...")
	err = c.AddGitHubToken("github.com", "your-github-token")
	if err != nil {
		log.Printf("添加GitHub令牌失败: %v", err)
	} else {
		fmt.Println("GitHub令牌已添加")
	}

	// 示例3：添加GitLab令牌
	fmt.Println("\n3. 添加GitLab令牌...")
	err = c.AddGitLabToken("gitlab.com", "your-gitlab-token")
	if err != nil {
		log.Printf("添加GitLab令牌失败: %v", err)
	} else {
		fmt.Println("GitLab令牌已添加")
	}

	// 示例4：添加Bearer令牌
	fmt.Println("\n4. 添加Bearer令牌...")
	err = c.AddBearerToken("example.com", "your-bearer-token")
	if err != nil {
		log.Printf("添加Bearer令牌失败: %v", err)
	} else {
		fmt.Println("Bearer令牌已添加")
	}

	// 示例5：添加HTTP Basic认证
	fmt.Println("\n5. 添加HTTP Basic认证...")
	err = c.AddHTTPBasicAuth("example.com", "username", "password")
	if err != nil {
		log.Printf("添加HTTP Basic认证失败: %v", err)
	} else {
		fmt.Println("HTTP Basic认证已添加")
	}

	// 示例6：获取特定令牌
	fmt.Println("\n6. 获取特定令牌...")
	token, err := c.GetToken("github-oauth", "github.com")
	if err != nil {
		log.Printf("获取令牌失败: %v", err)
	} else {
		fmt.Printf("GitHub令牌: %s\n", token)
	}

	// 示例7：移除令牌
	fmt.Println("\n7. 移除令牌...")
	err = c.RemoveToken("github-oauth", "github.com")
	if err != nil {
		log.Printf("移除令牌失败: %v", err)
	} else {
		fmt.Println("令牌已移除")
	}
}

// Example04Repository 演示如何管理仓库
func Example04Repository() {
	c, err := composer.New(composer.DefaultOptions())
	if err != nil {
		log.Fatalf("无法创建Composer实例: %v", err)
	}

	c.SetWorkingDir("/path/to/project")

	// 示例1：列出仓库
	fmt.Println("1. 列出仓库...")
	output, err := c.ListRepositories()
	if err != nil {
		log.Printf("列出仓库失败: %v", err)
	} else {
		fmt.Println(output)
	}

	// 示例2：添加VCS仓库
	fmt.Println("\n2. 添加VCS仓库...")
	err = c.AddVcsRepository("my-vcs", "https://github.com/myorg/myrepo")
	if err != nil {
		log.Printf("添加VCS仓库失败: %v", err)
	} else {
		fmt.Println("VCS仓库已添加")
	}

	// 示例3：添加Composer仓库
	fmt.Println("\n3. 添加Composer仓库...")
	err = c.AddComposerRepository("private-repo", "https://packages.example.com")
	if err != nil {
		log.Printf("添加Composer仓库失败: %v", err)
	} else {
		fmt.Println("Composer仓库已添加")
	}

	// 示例4：添加Path仓库
	fmt.Println("\n4. 添加Path仓库...")
	err = c.AddPathRepository("local-lib", "../my-lib", nil)
	if err != nil {
		log.Printf("添加Path仓库失败: %v", err)
	} else {
		fmt.Println("Path仓库已添加")
	}

	// 示例5：添加Artifact仓库
	fmt.Println("\n5. 添加Artifact仓库...")
	err = c.AddArtifactRepository("my-artifacts", "/path/to/artifacts")
	if err != nil {
		log.Printf("添加Artifact仓库失败: %v", err)
	} else {
		fmt.Println("Artifact仓库已添加")
	}

	// 示例6：移除仓库
	fmt.Println("\n6. 移除仓库...")
	err = c.RemoveRepository("my-vcs")
	if err != nil {
		log.Printf("移除仓库失败: %v", err)
	} else {
		fmt.Println("仓库已移除")
	}

	// 示例7：禁用/启用Packagist
	fmt.Println("\n7. 禁用/启用Packagist...")
	err = c.DisablePackagistRepository()
	if err != nil {
		log.Printf("禁用Packagist失败: %v", err)
	} else {
		fmt.Println("Packagist已禁用")
	}
	err = c.EnablePackagistRepository()
	if err != nil {
		log.Printf("启用Packagist失败: %v", err)
	} else {
		fmt.Println("Packagist已启用")
	}

	// 示例8：设置最小稳定性
	fmt.Println("\n8. 设置最小稳定性...")
	err = c.SetMinimumStability("stable")
	if err != nil {
		log.Printf("设置最小稳定性失败: %v", err)
	} else {
		fmt.Println("最小稳定性已设置为stable")
	}

	// 示例9：设置优先稳定版本
	fmt.Println("\n9. 设置优先稳定版本...")
	err = c.SetPreferStable(true)
	if err != nil {
		log.Printf("设置优先稳定版本失败: %v", err)
	} else {
		fmt.Println("已启用优先稳定版本")
	}
}

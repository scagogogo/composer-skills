package composer

// About 显示Composer的简短信息
//
// 返回值：
//   - string: Composer的简短描述信息
//   - error: 如果获取信息过程中发生错误，则返回相应的错误信息
//
// 功能说明：
//
//	该方法显示Composer本身的简短信息。
//	相当于执行`composer about`命令。
//
// 用法示例：
//
//	output, err := comp.About()
//	if err != nil {
//	    log.Fatalf("获取Composer信息失败: %v", err)
//	}
//	fmt.Println(output)
func (c *Composer) About() (string, error) {
	return c.Run("about")
}

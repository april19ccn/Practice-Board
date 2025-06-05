package main

import "fmt"

func main() {
	var name string
	var age int

	// 示例输入可能为：
	// "Alice 25"   -> 完整
	// "Bob"        -> 只有名字
	// "30 Charlie" -> 数字在前
	// n, err := fmt.Scanf("%s %d", &name, &age) // 这里传的是指针

	// "Age:25 Name:Alice"
	n, err := fmt.Scanf("Age:%d Name:%s", &age, &name)

	if err != nil {
		// 部分成功时：err 包含错误细节，n 是成功解析的字段数
		fmt.Printf("Parsed %d field(s): Name=%s\nError: %v\n", n, name, err)
	} else {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}

// 1. 渐进式解析
// 即使输入不完整（如只有名字），也能获取有效部分（name），而非完全失败。

// 2. 灵活空格处理
// 格式化字符串中的空格自动匹配任意数量空白字符（包括Tab、换行等），例如：
// "Alice 25"
// "Alice  25"
// "Alice\n25"

// 3. 复杂格式支持
// 可通过格式化字符串处理固定模式，例如带前缀的输入：
// 输入："Age:25 Name:Alice"
// fmt.Scanf("Age:%d Name:%s", &age, &name)

// fmt.Scanf 的核心价值在于通过格式化字符串显式定义输入结构，在解析混合类型或不规则数据时提供“部分成功”的能力，而非全有或全无。适用于日志解析、命令行工具等需要容忍脏数据的场景。

// 声明一个greeting包去收集相关函数
package greetings

import "fmt"

// Hello returns a greeting for the named person.
// 实现一个名称为Hello的函数
// 首字母大写的函数会被导出，意味着可以在不同的包中被调用
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// :=表达式会根据=右侧的内容自动推导变量类型，类似于下面这种方式的简写
	//var message string
	//message = fmt.Sprintf("Hello, %v. Welcome!!", name)
	//通过Sprintf创建一个格式化字符串
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

package main

import "fmt"

func main() {
	//未初始化的整形变量，会被初始化为0
	var x int
	fmt.Println(x)

	//简写的形式，y和z都是bool型变量，只需要在最后一个变量后面声明类型
	var y, z bool
	fmt.Println(y, z)

	//在变量声明的时候赋值以对其进行初始化，可以省略类型，会根据其初始值推断类型
	var a, b, c = 1, 'g', true
	//变量a,b的类型会被自动推导
	fmt.Println(a, b, c)

	//声明短变量
	i, j, k := true, false, "hello!"
	fmt.Println(i, j, k)
}

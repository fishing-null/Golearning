package main

import "fmt"

func main() {
	//声明两个变量
	var a int = 42
	b := 45
	//通过&取地址运算符获取这两个元素的地址
	pa, pb := &a, &b
	//pa pb是变量a b在内存中存储的地址
	fmt.Println(pa, pb)

	//使用*解引用来访问地址对应的元素
	*pa = 666
	//引用a并对其进行赋值
	fmt.Println(a)
	//go的指针不能进行运算

	//直接声明一个指针型的变量
	var pc *int
	c := 65
	pc = &c
	fmt.Println(pc, *pc)
}

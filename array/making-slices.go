package main

import "fmt"

func main() {
	//通过内置函数make，可以创建切片
	//切片a的容量为10，长度为0
	a := make([]int, 0, 10)
	//make函数会分配一个元素为零值（对于数组来说是nil）的数组并返回一个引用了他的切片。
	fmt.Println(a, len(a), cap(a))
	//向切片中追加元素,当元素个数超过底层数组长度的时候,会进行扩容
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 11)
	a = append(a, 12)
	fmt.Println(a, len(a), cap(a))

}

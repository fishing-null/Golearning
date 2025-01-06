package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// 当多个形参类型相同时，除了最后一个类型，其他的可以省略类型
func sub(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//这里return没有直接返回的返回值，会返回裸返回值
	return
}

//func split(sum int) (int, int) {
//	x := sum * 4 / 9
//	y := sum - x
//	//这里return没有直接返回的返回值，会返回裸返回值
//	return x, y
//}

func main() {
	//调用函数
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//裸返回值
	c, d := split(56)
	fmt.Println(c, d)
}

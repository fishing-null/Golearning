package main

import "fmt"

func outter() func(int) int {
	//正常的函数调用中 在outter调用完毕，sum会被销毁
	sum := 0
	return func(x int) int {
		//由于闭包的特性，sum的生命周期被延长
		sum += x
		fmt.Println("inner", sum)
		return sum
	}
}

func main() {
	f := outter()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

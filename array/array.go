package main

import "fmt"

func main() {
	//数组的长度不能改变
	var a [10]string

	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	//通过短语句声明一个数组
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

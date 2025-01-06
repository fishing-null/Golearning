package main

import "fmt"

func main() {
	fmt.Println("counting")
	//因为defer的特性，元素会被逆序打印
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	//先执行的代码，等执行完毕后，才执行上面的代码
	fmt.Println("hello")

}

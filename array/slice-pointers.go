package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	slice := primes[1:4]
	slice[2] = 10000

	fmt.Println(slice)
	//primes中下标为2的元素已经被改变了
	fmt.Println(primes)
	fmt.Println(primes[2])
}

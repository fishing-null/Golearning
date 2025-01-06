package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	count := 1
	//初始化语句中没有任何内容
	for ; count < 1000; fmt.Println("count", count) {
		count++
	}

	//for是Go中的while
	total := 0
	for total < 1000 {
		total++
		fmt.Println("total", total)
	}
}

package main

import "fmt"

var (
	pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
)

func main() {
	//range遍历
	//range + 数组可以对一个数组进行迭代 每次迭代返回两个参数 第一个是下标 第二个是对应位置的元素
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//只会拿到值
	for _, v := range pow {
		fmt.Printf("value:%d\n", v)
	}

	//只会拿到下标
	for i, _ := range pow {
		fmt.Printf("index:%d\n", i)
	}
}

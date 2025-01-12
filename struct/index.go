package main

import "fmt"

// 寻找x在切片s中的下标，不存在则返回-1
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	//Index方法可以在不同类型的切片上使用
	intArr := []int{15, 30, 45, 60}
	fmt.Println(Index(intArr, 60))

	strArr := []string{"foo", "bar", "baz"}
	fmt.Println(Index(strArr, "ajw"))
}

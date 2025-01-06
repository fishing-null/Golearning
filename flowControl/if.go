package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	//对数字开方，如果是负数则返回虚数
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//通过短表达式进行变量声明
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

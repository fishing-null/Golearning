package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//将float64类型的f转换位uint64并赋值给z
	//短运算符会自动推导变量类型
	z := f
	//通过var带有必须进行显示的类型转换，否则会报错
	//这种写法是不可行的
	//var z uint = f
	fmt.Println(x, y, z)
	fmt.Printf("%T", f)
}

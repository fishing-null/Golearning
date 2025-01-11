package main

import (
	"fmt"
)

type Pair struct {
	X, Y float64
}

// 传址
func (v *Pair) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 传值
func (v Pair) _Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v1 := Pair{3, 4}
	v2 := Pair{3, 4}
	fmt.Println("v1", v1, "v2", v2)
	//通过Scale方法能够修改v1中的内容
	v1.Scale(5)
	//_Scale方法不能修改v2中的内容
	v2._Scale(5)
	fmt.Println("v1", v1, "v2", v2)
}

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// 传址
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//声明一个Vertex类型的变量v
	var v Vertex
	v = Vertex{3, 4}
	//声明指针 指向变量v的地址
	pv := &v
	//通过v和pv调用Scale方法，结果是一致的，指针会自动重定向
	fmt.Println("v:", v, "pv", pv)
	//v首先扩大十倍
	pv.Scale(10)
	//然后再扩大八倍
	v.Scale(8)
	fmt.Println("v:", v, "pv", pv)

}

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var (
	// key为string型 value为Vertex自定义类型
	// 只是声明了变量，未对其进行初始化
	m map[string]Vertex
)

func main() {
	//通过make函数返回给定类型的映射
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	//常用的映射操作

	//在映射中插入或修改元素
	n := make(map[int]int)
	n['1'] = 1
	fmt.Println(n['1'])

	//获取映射中的一个元素
	elem := n['1']
	fmt.Println(elem)

	//删除映射中的元素
	delete(n, '1')
	fmt.Println(n['1'])

	//通过双赋值判断某个映射是否存在
	a, ok := n['1']
	//如果不存在 a 会是映射元素的零值，ok为false
	fmt.Println("双映射：", a, ok)
}

package main

import "fmt"

type Person struct {
	//大写声明，才能被外部访问
	Name        string
	Age, Height int
}

func main() {
	p := Person{Name: "Bob", Age: 24, Height: 184}
	fmt.Println(p)
	//通过.来访问结构体字段
	fmt.Println(p.Name, p.Age, p.Height)
}

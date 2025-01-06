package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	//会对表达式进行判断 t.Hour() < 12为真则会执行case内的内容
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}

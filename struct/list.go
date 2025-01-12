package main

// 一个可以存储任意类型的单链表
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

}

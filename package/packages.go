package main

//在这里导入路径
//包名是路径的最后一个元素名
import (
	"fmt"
	"math/rand"
)

// 程序从main包开始运行，不导入的话，程序就无法运行。
func main() {
	fmt.Println("我最喜欢的数字是 ", rand.Intn(100))
}

package main

import "fmt"

func main() {
	var i int
	var f float64
	var s string
	var b bool
	//int，float64,string,bool对应的零值分别是0，0，"",false
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

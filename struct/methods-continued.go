package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (p MyFloat) Abs() float64 {
	if p < 0 {
		return float64(-p)
	}
	return float64(p)
}

func main() {
	var a MyFloat = -math.Sqrt2
	fmt.Println(a.Abs())
}

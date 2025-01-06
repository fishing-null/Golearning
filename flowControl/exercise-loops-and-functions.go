package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	prev := 0.0
	for z := x; z-prev != 0; {
		prev = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return prev
}

func main() {
	fmt.Println(Sqrt(2))
}

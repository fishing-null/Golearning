package main

import "fmt"

const Pi float64 = 3.14

func main() {
	const World string = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

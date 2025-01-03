package main

import (
	"fmt"

	"example.com/greetings"
)

// Go的可执行代码必须放在main内部
func main() {
	// Get a greetings message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

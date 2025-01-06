package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go的运行环境")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

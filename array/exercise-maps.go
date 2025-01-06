package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var ret = make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		ret[v] = ret[v] + 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}

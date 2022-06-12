package main

import (
	"fmt"
	"strings"
)

var result map[string]int

func wordCount(s string) map[string]int {
	words := strings.Fields(s)

	result = make(map[string]int)

	for i, v := range words {
		result[v] = i
	}

	return result
}

func main() {
	fmt.Println(wordCount("hi hello da no"))
}

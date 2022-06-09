package main

import "fmt"

func main() {
	value := 1

	for value < 10 {
		value += value
	}

	fmt.Println(value)
}

package main

import "fmt"

func main() {
	value := 0

	for i := 0; i < 10; i++ {
		value += i
	}

	fmt.Println(value)
}

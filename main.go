package main

import "fmt"

func main() {
	numbers := [6]int{1, 2, 3, 4, 5, 6}

	var sliceOne []int = numbers[1:4]
	fmt.Println(sliceOne)

	numbers[2] = 10
	fmt.Println(sliceOne)
}

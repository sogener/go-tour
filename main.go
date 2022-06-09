package main

import "fmt"

func main() {
	var phrases [2]string

	phrases[0] = "Hello golang"
	phrases[1] = "From ide"

	fmt.Println(phrases[0], phrases[1])

	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)
}

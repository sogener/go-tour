package main

import "fmt"

type Numbers struct {
	numberOne, numberTwo int
}

func sumNumbers(n Numbers) int {
	return n.numberOne + n.numberTwo
}

func main() {
	n := Numbers{numberTwo: 2, numberOne: 1}

	fmt.Println(sumNumbers(n))
}

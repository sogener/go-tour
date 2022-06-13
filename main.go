package main

import "fmt"

type Numbers struct {
	numberOne, numberTwo int
}

func (n Numbers) sumNumbers() int {
	return n.numberOne + n.numberTwo
}

func main() {
	n := Numbers{numberTwo: 2, numberOne: 1}

	fmt.Println(n.sumNumbers())
}

package main

import "fmt"

func sumNumbers(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

func main() {
	divideNumbers := func(numberOne, numberTwo int) int {
		return numberOne / numberTwo
	}

	dividedNumber := divideNumbers(4, 2)
	fmt.Println(dividedNumber)

	result := sumNumbers(dividedNumber, 10)

	fmt.Println(result)
}

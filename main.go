package main

import "fmt"

func main() {
	test1()

	defer fmt.Println("hi")
}

func test1() {
	defer fmt.Println("yopta")

	fmt.Println("hello1")
}

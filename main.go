package main

import "fmt"

type MyString string

func (s MyString) randomMessage() string {
	return "hello123"
}

func main() {
	message := MyString("hello")

	fmt.Println(message)
	fmt.Println(message.randomMessage())
}

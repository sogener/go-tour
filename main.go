package main

import "fmt"

func main() {
	names := [4]string{
		"Gleb",
		"Jhon",
		"Alexander",
		"Mike",
	}

	sliceOne := names[0:1]
	sliceSecond := names[1:3]

	// output: [Gleb] [Jhon Alexander]
	fmt.Println(sliceOne, sliceSecond)

	sliceSecond[1] = "HELLO123"

	// output: [Jhon HELLO123]
	fmt.Println(sliceSecond)

	// output: [Gleb Jhon HELLO123 Mike]
	// Изменился изначальный массив
	fmt.Println(names)

}

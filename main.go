package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"hello1": {40.11, 2.6},
		"hello2": {10, 20},
	}

	fmt.Println(m)
}

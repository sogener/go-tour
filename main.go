package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]int)

	m["hello1"] = 40
	fmt.Println(m)

	m["hello1"] = 20
	fmt.Println(m)

	delete(m, "hello1")

	v, ok := m["hello1"]
	fmt.Println("The value: ", v, "Present?", ok)
}

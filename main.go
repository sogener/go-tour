package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map[key] => Value
// key => some random STRING key
// value = Vertex{Lat, Long}

// Создаём глобальную переменную и её описание (интерфейс)
var m map[string]Vertex

func main() {
	// Создаём эту переменную
	m = make(map[string]Vertex)
	// Добавляем значения
	m["Hello123"] = Vertex{Lat: 40.2222, Long: -66.231}

	fmt.Println(m["Hello123"])
}

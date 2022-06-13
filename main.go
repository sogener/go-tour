package main

import "fmt"

type Sentence struct {
	firstWord, secondWord string
}

func (s *Sentence) appendWords() string {
	return s.firstWord + " " + s.secondWord
}

func main() {
	sentence := Sentence{firstWord: "hello", secondWord: "yopta"}

	fmt.Println(sentence.appendWords())
}

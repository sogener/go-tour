package main

import "fmt"

type Sentence struct {
	firstWord, secondWord string
}

func (s *Sentence) appendWords() string {
	return s.firstWord + " " + s.secondWord
}

func changeWordsInSentence(s *Sentence, firstWord, secondWord string) {
	s.firstWord = firstWord
	s.secondWord = secondWord
}

func main() {
	sentence := Sentence{firstWord: "hello", secondWord: "yopta"}
	fmt.Println(sentence.appendWords())

	changeWordsInSentence(&sentence, "poka", "cat")
	fmt.Println(sentence.appendWords())
}

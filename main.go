package main

import (
	"fmt"

	"github.com/changduckGO/Practice-Go/mydict"
)

func main() {
	dictionary := mydict.Dictioary{}
	baseWord := "Hello"
	def := "Greeting"

	dictionary.Add(baseWord, def)
	dictionary.Search(baseWord)
	// dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}

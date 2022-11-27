package main

import (
	"fmt"

	"github.com/changduckGO/Practice-Go/mydict"
)

func main() {
	dictionary := mydict.Dictioary{}
	word := "Hello"
	def := "Greeting"

	errAdd := dictionary.Add(word, def)
	if errAdd != nil {
		fmt.Println(errAdd)
	}

	defSearch, errSearch := dictionary.Search(word)
	if errSearch != nil {
		fmt.Println("Not Found", word)
	} else {
		fmt.Println("Found", word, "Definition:", defSearch)
	}

}

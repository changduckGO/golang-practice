package main

import (
	"fmt"

	"github.com/changduckGO/Practice-Go/mydict"
)

func main() {
	dictionary := mydict.Dictioary{"FIRST": "FIRST WORLD"}
	definition, err := dictionary.Search("SECOND")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition)
}

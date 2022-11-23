package main

import (
	"fmt"

	"github.com/changduckGO/Practice-Go/accounts"
)

func main() {
	account := accounts.NewAccount("harry")
	fmt.Println(account)
}

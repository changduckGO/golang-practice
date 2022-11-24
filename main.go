package main

import (
	"fmt"

	"github.com/changduckGO/Practice-Go/accounts"
)

func main() {
	account := accounts.NewAccount("harry")
	account.Deposit(100)

	fmt.Print(account.Balance())
}

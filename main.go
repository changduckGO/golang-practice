package main

import (
	"fmt"
	"log"

	"github.com/changduckGO/Practice-Go/accounts"
)

func main() {
	account := accounts.NewAccount("harry")
	account.Deposit(100)
	fmt.Print(account.Balance())

	err := account.Withdraw(120)
	if err != nil {
		log.Fatalln(err) // Prinln() 후 프로그램 종료
	}

	fmt.Println(account.Balance())
}

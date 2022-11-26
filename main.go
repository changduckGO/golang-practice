package main

import (
	"fmt"
	"log"

	"github.com/changduckGO/Practice-Go/accounts"
)

func main() {
	account := accounts.NewAccount("harry")
	account.Deposit(100)

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err) // Prinln() 후 프로그램 종료
	}

	account.ChangeOwner("SCD")

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}

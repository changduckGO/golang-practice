// accounts는 컴파일하지 않을 것이므로 main() 사용하지 않음
package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit X amount on my account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Witdraw X amount on my account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// Balance of my account
func (a Account) Balance() int {
	return a.balance
}

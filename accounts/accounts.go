// accounts는 컴파일하지 않을 것이므로 main() 사용하지 않음
package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit X amount on my account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of my account
func (a Account) Balance() int {
	return a.balance
}

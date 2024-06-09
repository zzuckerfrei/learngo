package accounts

import (
	"errors"
	"fmt"
)

// Account struct
// Upper case for public export
type Account struct {
	owner   string
	balance int
}

// name of error variable should be like 'err000'
var errNoMoney = errors.New("you cant't withdraw. not enough money")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// (a *Account) -> receiver
// *Account -> this means 'don't make copy.'
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// Golang doesn't have 'Exception' like try-catch, try-except..
// you need to handle all erros by yourself
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
	fmt.Println("this is", a.owner, "'s account from now on")
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// String -> default method
// edit String method like this
func (a Account) String() string {
	return fmt.Sprint("---\n", a.Owner(), "'s account.\nHas: ", a.Balance())
}

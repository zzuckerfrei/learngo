package main

import (
	"fmt"

	accounts "github.com/zzuckerfrei/learngo/account"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(100)

	fmt.Println(account.Balance())

	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	account.ChangeOwner("zzuckerfrei")
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(account)
}

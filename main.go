package main

import (
	"fmt"

	"github.com/canu0205/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("canu")
	account.Deposit(10)
	if err := account.Withdraw(20); err != nil {
		fmt.Println(err)
	}
	account.ChangeOwner("canu2")
	fmt.Println(account.Balance(), account.Owner())
}

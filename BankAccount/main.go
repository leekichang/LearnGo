package main

import (
	"fmt"

	"github.com/leekichang/LearnGo/BankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("kichang")
	fmt.Println(account.Balance())
	account.Deposit(1000)
	fmt.Println(account.Balance())
	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

	account.ChangeOwner("Yujin")
	fmt.Println(account.Owner())
	fmt.Println(account)
	account.String()
}

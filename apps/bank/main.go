// Bank Account Manager
// Create a class called Account which will be an abstract class for three other classes
// called CheckingAccount, SavingsAccount and BusinessAccount. Manage credits and debits
// from these accounts through an ATM style program.
package main

import (
	"fmt"
	"strconv"

	"bank/accounts"
)

func deposit(account accounts.Account, amount int64) {
	fmt.Printf("Depositing %s to account\n", strconv.Itoa(int(amount)))
	account.Deposit(amount)
	displayBalance(account)
}

func withdraw(account accounts.Account, amount int64) {
	fmt.Printf("Withdrawing %s from account\n", strconv.Itoa(int(amount)))
	account.Withdraw(amount)
	displayBalance(account)
}

func displayBalance(account accounts.Account) {
	fmt.Printf("Account balance %s\n", strconv.Itoa(int(account.Balance())))
}

func main() {
	checking := accounts.NewCheckingAccount(100)
	savings := accounts.NewSavingsAccount(0)

	//ATM
	fmt.Println("Checking Account")
	displayBalance(checking)
	deposit(checking, 75)
	withdraw(checking, 50)

	fmt.Println("\nSavings Account")
	displayBalance(savings)
	deposit(savings, 10000)
	deposit(savings, 2000)
}

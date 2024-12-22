package BankAcc

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.Balance += amount
		fmt.Printf("Deposited %.2f successfully.\n", amount)
	} else {
		fmt.Println("Deposit amount must be positive.")
	}
}

func (ba *BankAccount) Withdraw(amount float64) {
	if amount > 0 {
		if ba.Balance >= amount {
			ba.Balance -= amount
			fmt.Printf("Withdrew %.2f successfully.\n", amount)
		} else {
			fmt.Println("Insufficient balance.")
		}
	} else {
		fmt.Println("Withdrawal amount must be positive.")
	}
}

func (ba *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", ba.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}

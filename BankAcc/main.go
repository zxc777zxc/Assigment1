package BankAcc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	account := &BankAccount{
		AccountNumber: "1234567890",
		HolderName:    "Adolf",
		Balance:       0,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu:\n1. Deposit\n2. Withdraw\n3. Get balance\n4. Exit")
		fmt.Print("Enter your choice: ")
		choiceInput, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceInput))
		if err != nil {
			fmt.Println("Invalid input, please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter deposit amount: ")
			amountInput, _ := reader.ReadString('\n')
			amount, err := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Deposit(amount)

		case 2:
			fmt.Print("Enter withdrawal amount: ")
			amountInput, _ := reader.ReadString('\n')
			amount, err := strconv.ParseFloat(strings.TrimSpace(amountInput), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Withdraw(amount)

		case 3:
			account.GetBalance()

		case 4:
			fmt.Println("Exiting the program. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

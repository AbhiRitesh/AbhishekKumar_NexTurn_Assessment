package main

import (
	"errors"
	"fmt"
)

// Account struct for a bank account
type Account struct {
	ID                int
	Name              string
	Balance           float64
	History           []string
}

// menu options
const (
	DepositOption   = 1
	WithdrawOption  = 2
	ViewBalance     = 3
	TransactionHist = 4
	ExitOption      = 5
)

// Deposeit function
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	a.History = append(a.History, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw function
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount should be greater than zero")
	}
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount
	a.History = append(a.History, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// menu options
func DisplayMenu() {
	fmt.Println("\nBank Transaction System")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. View Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
}

func (a *Account) TransactionHistory() {
	if len(a.History) == 0 {
		fmt.Println("No transactions.")
		return
	}

	fmt.Println("Transaction History:")
	for _, transaction := range a.History {
		fmt.Println(transaction)
	}
}

func main() {
	account := Account{
		ID:   1,
		Name: "John Doe",
		Balance: 0.0,
		History: []string{},
	}

	for {
		DisplayMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case DepositOption:
			var amount float64
			fmt.Print("Enter amount: ")
			fmt.Scanln(&amount)
			if err := account.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case WithdrawOption:
			var amount float64
			fmt.Print("Enter amount: ")
			fmt.Scanln(&amount)
			if err := account.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdraw successful.")
			}
		case ViewBalance:
			fmt.Printf("Current balance: %.2f\n", account.Balance)
		case TransactionHist:
			account.TransactionHistory()
		case ExitOption:
			fmt.Println("Exiting. Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

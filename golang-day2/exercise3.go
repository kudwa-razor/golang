package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (a *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	a.balance += amount
	fmt.Printf("Deposited Rs.%d, New Balance: Rs.%d\n", amount, a.balance)
	a.mutex.Unlock()
}

func (a *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mutex.Lock()
	if a.balance-amount >= 200 {
		a.balance -= amount
		fmt.Printf("Withdrawn Rs.%d, New Balance: Rs.%d\n", amount, a.balance)
	} else {
		fmt.Printf("Failed to withdraw Rs.%d, Minimum balance of Rs.200 required. Current balance: Rs.%d\n", amount, a.balance)
	}
	a.mutex.Unlock()
}

func main() {
	account := BankAccount{balance: 500}
	var wg sync.WaitGroup

	for {
		var choice, amount int
		fmt.Println("Choose an operation:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		if choice == 3 {
			break
		}

		fmt.Print("Enter amount: ")
		fmt.Scan(&amount)

		if choice == 1 {
			wg.Add(1)
			go account.Deposit(amount, &wg)
		} else if choice == 2 {
			wg.Add(1)
			go account.Withdraw(amount, &wg)
		} else {
			fmt.Println("Invalid choice, please try again.")
		}
	}

	wg.Wait()
	fmt.Printf("Final Balance: Rs.%d\n", account.balance)
}

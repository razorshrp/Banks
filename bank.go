package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	var amount float64
	fmt.Println("Welcome to Go Bank!!!")
	for i := 0; i <= i; i++ {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			} else {
				accountBalance += amount
				fmt.Println("Your Final Balance is: ", accountBalance)
			}
		} else if choice == 3 {
			fmt.Print("Enter Withdraw Amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			if amount > accountBalance {
				fmt.Println("Withdraw Amount is Greater than Account Balance")
				continue
			} else {
				accountBalance -= amount
				fmt.Println("Your Final Balance is: ", accountBalance)
			}
		} else {
			fmt.Println("Goodbye!!")
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank")
}

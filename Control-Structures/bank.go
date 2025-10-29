package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		var depositAmmount float64
		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmmount)
		accountBalance += depositAmmount
		fmt.Println("Balance has been updated! New balance:", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Withdraw amount: ")
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Balance has been updated! New balance:", accountBalance)
	}
}

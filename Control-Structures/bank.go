package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmmount)

			if depositAmmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}

			accountBalance += depositAmmount
			fmt.Println("Balance has been updated! New balance:", accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount! You can't withdraw more than you have. Current balance:", accountBalance)
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance has been updated! New balance:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go bank!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmmount float64
		// 	fmt.Print("Your deposit: ")
		// 	fmt.Scan(&depositAmmount)

		// 	if depositAmmount <= 0 {
		// 		fmt.Println("Invalid amount! Must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmmount
		// 	fmt.Println("Balance has been updated! New balance:", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawalAmount float64
		// 	fmt.Print("Withdrawal amount: ")
		// 	fmt.Scan(&withdrawalAmount)

		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Invalid amount! Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawalAmount > accountBalance {
		// 		fmt.Println("Invalid amount! You can't withdraw more than you have. Current balance:", accountBalance)
		// 		continue
		// 	}
		// 	accountBalance -= withdrawalAmount
		// 	fmt.Println("Balance has been updated! New balance:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}

}

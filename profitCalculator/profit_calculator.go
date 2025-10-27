package main

import (
	"fmt"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter a revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("Enter a tax rate: ")
	// fmt.Scan(&taxRate)

	revenue := getUserInput("Enter a revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter a tax rate: ")

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(consoleMessage string) float64 {
	var userInput float64
	fmt.Print(consoleMessage)
	fmt.Scan(&userInput)
	return userInput
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

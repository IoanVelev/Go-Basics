package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter a revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter a tax rate: ")
	fmt.Scan(&taxRate)

	earnings_Before_Tax := revenue - expenses
	earnings_After_Tax := earnings_Before_Tax * taxRate / 100
	ratio := earnings_Before_Tax / earnings_After_Tax

	fmt.Print(earnings_Before_Tax)
	fmt.Print(earnings_After_Tax)
	fmt.Print(ratio)
	fmt.Print("This is profit calculator")
}

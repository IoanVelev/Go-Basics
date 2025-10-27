package main

import (
	"fmt"
)

var taxRate float64

func main() {
	var revenue float64
	var expenses float64

	// fmt.Print("Enter a revenue: ")
	// fmt.Scan(&revenue)
	revenue = fetchSelectedRevenue(revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter a tax rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateValues(revenue, expenses)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func fetchSelectedRevenue(revenue float64) float64 {
	fmt.Print("Enter a revenue: ")
	fmt.Scan(&revenue)
	return revenue
}

func calculateValues(revenue, expenses float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

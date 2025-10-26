package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment amount: ")
	printText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	printText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	printText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureFinalValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureFinalValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formatedFV := fmt.Sprintf("Future value: %.0f\n", futureValue)
	formatedFFV := fmt.Sprintf("Future final value (adjusted for inflation): %.0f\n", futureFinalValue)

	// fmt.Printf("Future value: %.0f\n", futureValue)
	// fmt.Printf("Future final value (adjusted for inflation): %.0f", futureFinalValue)

	fmt.Print(formatedFV, formatedFFV)
}

func printText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	ffv := fv / math.Pow(1+inflationRate/100, years)

	return fv, ffv
}

package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureFinalValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value: %v\n", futureValue)
	fmt.Printf("Future final value (adjusted for inflation): %v", futureFinalValue)
}

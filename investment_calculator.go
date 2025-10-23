package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	futureValue := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	futureFinalValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureFinalValue)
}

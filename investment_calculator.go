package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	futureValue := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}

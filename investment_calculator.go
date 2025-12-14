package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the No. of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the exptected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value = %.2f", futureValue)
	fmt.Printf("Future Real Value = %.2f", futureRealValue)

}

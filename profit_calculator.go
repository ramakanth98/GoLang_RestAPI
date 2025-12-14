package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formattedEBT := fmt.Sprintf("EBT = %.2f \n", ebt)
	fmt.Print(formattedEBT)
	fmt.Printf("Profit = %.2f \n", profit)
	fmt.Printf("Ratio of EBT and Profit = %.2f \n", ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

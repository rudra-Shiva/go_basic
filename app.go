package main

import (
	"fmt"
)

func main() {
	fmt.Print("Profit Calculator\n")

	var earning float64
	var expenses float64
	var taxRate float64
	fmt.Print("Earnings Value :")
	fmt.Scan(&earning)

	fmt.Print("Expenses :")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate :")
	fmt.Scan(&taxRate)
	fmt.Print("Earning beforer  Tax: ", earning-expenses, "\n")
	profit := (earning - expenses) * (1 - taxRate/100)
	fmt.Printf("Net Profit : %.2f", profit)
	fmt.Printf("\nRatio : %.2f", (earning-expenses)/profit)

	formattedFv := fmt.Sprintf("\nRatio: %.1f\n", (earning-expenses)/profit)
	formattedFvs := fmt.Sprintf(`Other way
	 formatted String Ratio: %.1f`, (earning-expenses)/profit)
	fmt.Print(formattedFv, " and ", formattedFvs)

}

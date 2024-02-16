package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello World\n")
	const inflamationRate float64 = 2.5
	var investmentValue float64
	var years float64
	var expectedReturnRate float64
	fmt.Print("Investment Value :")
	fmt.Scan(&investmentValue)

	fmt.Print("Years :")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate :")
	fmt.Scan(&expectedReturnRate)
	futurevalue := investmentValue * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futurevalue / math.Pow(1+inflamationRate/100, years)
	fmt.Println(futurevalue)
	fmt.Println(futureRealValue)
}

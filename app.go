package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello World\n")
	var investmentValue float64 = 10000
	years := 10
	expectedReturnRate := 5.5
	futurevalue := investmentValue * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futurevalue)
}

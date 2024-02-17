package main

import (
	"fmt"
	"log"
	"math"
)

// Constant also define here or outside main function
const value int = 5
const value1 string = "Hai String"

func main() {
	fmt.Print("Function\n")
	const inflamationRate float64 = 2.5
	var investmentValue float64
	var years float64
	var expectedReturnRate float64
	// fmt.Print("Investment Value :")
	// formattedTexts("Investment Value :")
	// fmt.Scan(&investmentValue)
	investmentValue = getUserInpur("Revenu: ")
	fmt.Print(investmentValue, "\n")
	fmt.Print("Years :")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate :")
	fmt.Scan(&expectedReturnRate)

	futurevalue := investmentValue * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futurevalue / math.Pow(1+inflamationRate/100, years)

	fmt.Println(futurevalue)
	fmt.Println(futureRealValue)

	ss := withParametersAndSingleRetrunType("ha", "Atul")
	fmt.Print("Function with returm type: ", ss)

}

func formattedTexts(text string) {
	fmt.Print(text)
}

func withParameter(text int, text2 string) {
	fmt.Print(text, " ", text2)
}
func withParameters(text, text2 string) {
	fmt.Print(text, " ", text2)
}

// with return type
func withParametersAndSingleRetrunType(text, text2 string) string {
	fmt.Print(text, " ", text2, "\n")
	return text
}

// with return type
func withParametersAndRetrunType(text, text2 string) (string, string) {
	fmt.Print(text, " ", text2)
	return text, text2
}

// with return type
func withParametersAndRetrunTypeVarabile(text, text2 string) (fs string, fsv string) {
	fmt.Print(text, " ", text2)
	fs = text
	fsv = text2
	return
}

// function for declaring variable
func getUserInpur(prompt string) float64 {
	var userInput float64
	fmt.Print(prompt)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatal("scan failed:", err)
	}
	return userInput
}

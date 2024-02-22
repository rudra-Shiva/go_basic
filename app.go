package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers Program in Go") // Output: 8
	age := 32
	agePointer := &age
	fmt.Print("Age: ", age, "\n") // Output: Age: 32
	adultYears := getAdultYears(*agePointer)
	fmt.Println("Adult age: ", adultYears)
	adultYear := getAdultYear(agePointer)
	fmt.Println("Adult age: ", adultYear)
	editAgeToAdultYears(agePointer)
	fmt.Println(age) // Output: 0
}

func getAdultYears(age int) int {
	return age - 18
}

// passing pointer for calcuating address
func getAdultYear(age *int) int {
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}

package main

import (
	"fmt"
)

func main() {
	fmt.Print("Recursion World\n")
	result := factorial(4)
	fmt.Println(result)
	r1 := recurFactorial(5)
	fmt.Println(r1)
}

// factorial with normal methods
func factorial(number int) int {
	if number <= 0 {
		return 0
	}
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

// using recursion
func recurFactorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * recurFactorial(number-1)
}

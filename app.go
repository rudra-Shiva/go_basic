package main

import "fmt"

func main() {
	fmt.Print("Variadic Function for  summing numbers:\n")
	nums := []int{1, 2, 3, 4}
	sums := sumup(nums...) // use the ... to unpack nums into arguments of Sum function
	fmt.Printf("The sum is %d\n", sums)

	// calling a variadic function with no arguments
	fmt.Println("\nCalling a variadic function with arguments:")
	sumss := sum(1, nums...)
	fmt.Println(sumss)
	summm := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(summm)

}

//normal  way: one of type function uses for variadic function
func sumup(numbers ...int) int {
	total := 0
	fmt.Printf("The number of arguments passed is %d.\n", len(numbers))
	for _, value := range numbers {
		total += value
	}
	return total

}

//other way function define to variadic

func sum(startingNumber int, number ...int) int {
	sum := 0
	for _, val := range number {
		sum += val
	}
	return sum
}

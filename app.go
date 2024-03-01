package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	number := []int{1, 2, 3, 4, 5}
	moreNumber := []int{1, 3, 4, 5, 6, 7, 8, 9}
	extraNUmber := []int{4, 5, 6, 7, 8, 9, 10}
	dNumber := doubledNumber(&number, func(number int) int {
		return number * 2
	}) // anonymous function using like that in Go lang
	fmt.Printf("%v\n", dNumber) // [2 4 6 8 10]
	tripleNumber := doubledNumber(&number, triple)
	fmt.Println(tripleNumber)
	trnasForm1 := getTransformerFunction(&moreNumber)
	trnasForm2 := getTransformerFunction(&extraNUmber)
	tr := doubledNumber(&number, trnasForm1)
	tr2 := doubledNumber(&extraNUmber, trnasForm2)
	fmt.Println(tr)  //
	fmt.Println(tr2) // <nil>

}
func doubledNumber(number *[]int, transform transformFn) []int {
	doubledNumbers := []int{}
	for _, value := range *number {
		doubledNumbers = append(doubledNumbers, transform(value))
	}
	return doubledNumbers
}

// func doubledNumber(number *[]int, transform func(int) int) []int {
// 	doubledNumbers := []int{}
// 	for _, value := range *number {
// 		doubledNumbers = append(doubledNumbers, transform(value))
// 	}
// 	return doubledNumbers
// }

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(number *[]int) transformFn {

	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

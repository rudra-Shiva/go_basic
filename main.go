package main

import "fmt"

type str string

func (text str) customType() {
	fmt.Println(text)
}

func main() {
	fmt.Println("GO Lang")
	var t1 str = "Hello World!"
	t1.customType()
}

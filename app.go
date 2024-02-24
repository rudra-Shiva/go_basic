package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
)

func main() {
	fmt.Print("Structs Projects\n")
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err) // terminate the program if there is an error while reading from input
	}
	note, err := note.New(title, content) // create a new instance of struct "Note" with given values for Title and Content else{ else { // otherwise continue with the rest of the else {
	note.Display()
	note.Save() // display the note data on console
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter the title of your note: ")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	Content, err := getUserInput("Enter the content of your note: \n")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return strings.TrimSpace(title), Content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	// return input[:len(input)-1] // remove trailing newline character
	return input, nil
}

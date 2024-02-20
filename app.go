package main

import (
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64) {
	os.WriteFile("balance.txt", []byte(fmt.Sprint(balance)), os.ModePerm)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, err
	}
	balance := string(data)
	balances, _ := strconv.ParseFloat(balance, 64)
	return float64(balances), nil
}

// This function is used to deposit money into the account of a user. It takes in an amount as a parameter and adds it to the current balance stored in the file 'balance
func depositMoneyIntoAccount(amount float64) {
	current_balance, _ := readBalanceFromFile()
	writeBalanceToFile(current_balance + amount)
}

func withdrawMoneyFromAccount(amount float64) bool {
	if current_balance, _ := readBalanceFromFile(); current_balance < amount {
		return false // Not enough money in account
	} else {
		depositMoneyIntoAccount(-amount)
		return true
	}
}

// Main function of the program
func main() {
	accountBalance, _ := readBalanceFromFile()
	// for i := 0; i < 2; i++ {
	for {
		fmt.Println("Welcome to Rudra Bank")
		fmt.Println("What do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		// wantCheckBalance := choice == 1
		// if wantCheckBalance {
		// 	fmt.Println("Balance Now: ", accountBalance)
		// }
		switch choice {
		case 1:
			fmt.Println("Balance Now: ", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, please enter a positive number.")
				// return  //basically it use for final result and end of program
				continue // it use for continue next iteration of loop and not the function
			} else {
				accountBalance += depositAmount
				writeBalanceToFile(accountBalance)
				fmt.Printf("Deposited Amount %v successfully.\n", depositAmount)
				fmt.Println("New Balance: ", accountBalance)
			}
			// accountBalance += depositAmount
			// fmt.Printf("Deposited Amount %v Successfully\n", depositAmount)
			// fmt.Println("New Balance is : ", accountBalance)
		case 3:

			fmt.Print("WithDraw Money: ")
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 || withdrawMoney > accountBalance {
				fmt.Println("Invalid amount or not enough balance in your account.")
				// return
				continue
			}
			if withdrawMoney > accountBalance {
				fmt.Println("Sorry! You don't have enough amount in your account.")

				// return
				continue
			}
			// accountBalance = accountBalance - withdrawMoney
			withdrawMoneyFromAccount(withdrawMoney)
			accountBalance -= withdrawMoney
			fmt.Printf("Successfully Withdrawn Amount Rs.%v \n", withdrawMoney)
			fmt.Println("Remaining Balance is : ", accountBalance)
		default:
			fmt.Println("Good bye!!")
			fmt.Println("Thanks for choosing our Bank")
			// break //it use for  break out from loop and not the function
			return
		}

	}

}

package main

import (
	"fmt"

	"example.com/main/fileops"
)

// Main function of the program
func main() {
	accountBalance, _ := fileops.ReadBalanceFromFile()
	// for i := 0; i < 2; i++ {
	for {
		presentOptions()
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
				fileops.WriteBalanceToFile(accountBalance)
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
			fileops.WithdrawMoneyFromAccount(withdrawMoney)
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

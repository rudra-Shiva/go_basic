package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	// for i := 0; i < 2; i++ {
	for { //enfinity  loop
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
		if choice == 1 {
			fmt.Println("Balance Now: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, please enter a positive number.")
				// return  //basically it use for final result and end of program
				continue // it use for continue next iteration of loop and not the function
			} else {
				accountBalance += depositAmount
				fmt.Printf("Deposited Amount %v successfully.\n", depositAmount)
				fmt.Println("New Balance: ", accountBalance)
			}
			// accountBalance += depositAmount
			// fmt.Printf("Deposited Amount %v Successfully\n", depositAmount)
			// fmt.Println("New Balance is : ", accountBalance)
		} else if choice == 3 {
			fmt.Print("WithDraw Mone: ")
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
			accountBalance -= withdrawMoney
			fmt.Printf("Successfully Withdrawn Amount Rs.%v \n", withdrawMoney)
			fmt.Println("Remaining Balance is : ", accountBalance)

		} else {
			fmt.Println("Good bye!!")
			// return
			break //it use for  break out from loop and not the function
		}

	}
	fmt.Println("Thanks for choosing our Bank")

}

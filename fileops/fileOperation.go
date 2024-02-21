package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(balance float64) {
	os.WriteFile("balance.txt", []byte(fmt.Sprint(balance)), os.ModePerm)
}

func ReadBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, err
	}
	balance := string(data)
	balances, _ := strconv.ParseFloat(balance, 64)
	return float64(balances), nil
}

// This function is used to deposit money into the account of a user. It takes in an amount as a parameter and adds it to the current balance stored in the file 'balance
func DepositMoneyIntoAccount(amount float64) {
	current_balance, _ := ReadBalanceFromFile()
	WriteBalanceToFile(current_balance + amount)
}

func WithdrawMoneyFromAccount(amount float64) bool {
	if current_balance, _ := ReadBalanceFromFile(); current_balance < amount {
		return false // Not enough money in account
	} else {
		DepositMoneyIntoAccount(-amount)
		return true
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const storageFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(storageFile)

	if err != nil {
		return 0, errors.New("could not find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("could not parse balance file")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(storageFile, []byte(balanceText), os.FileMode(0644))
}

func main() {
	var choice int
	balance, err := getBalanceFromFile()

	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome to the Bank, what would you like to do?")
	for {
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is $%v\n", balance)
		case 2:
			var deposit float64
			fmt.Println("How much would you like to deposit?")
			fmt.Scan(&deposit)
			balance += deposit
			fmt.Printf("Your balance is $%v\n", balance)
			writeBalanceToFile(balance)
		case 3:
			var withdraw float64
			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("You do not have enough money to withdraw that amount")
			} else {
				balance -= withdraw
				fmt.Printf("Your balance is $%v\n", balance)
				writeBalanceToFile(balance)
			}
		default:
			fmt.Println("Thank you for using the Bank")
			return
		}
	}
}

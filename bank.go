package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const storageFile = "balance.txt"

func main() {
	var choice int
	balance, err := fileops.GetFloatFromFile(storageFile)

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
			fileops.WriteFloatToFile(storageFile, balance)
		case 3:
			var withdraw float64
			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("You do not have enough money to withdraw that amount")
			} else {
				balance -= withdraw
				fmt.Printf("Your balance is $%v\n", balance)
				fileops.WriteFloatToFile(storageFile, balance)
			}
		default:
			fmt.Println("Thank you for using the Bank")
			return
		}
	}
}

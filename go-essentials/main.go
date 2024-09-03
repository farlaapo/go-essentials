package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "balance.txt"

func readFileBalance() (float64, error) {
	date, err := os.ReadFile(accountBalance)
	if err != nil {
		return 1000, errors.New("filed to ReadFile")
	}

	balanceTxt := string(date)
	balance, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil {
		return 1000, errors.New("filed to to  ParseFloat")
	}
	return balance, nil

}

func writeFileBalance(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalance, []byte(balanceTxt), 0644)

}

func main() {
	var accountBalance, err = readFileBalance()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")

	}

	fmt.Println("Welcom to GO Bank! ")

	for {
		fmt.Println("What You Want To Do: ")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var Choice int
		fmt.Print("your Choice:  ")
		fmt.Scan(&Choice)

		// fmt.Println("your Choice", Choice)

		switch Choice {
		case 1:
			fmt.Println("Your Balance is  :", accountBalance)
		case 2:
			fmt.Println("Your Deposit is :")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("your AccountBalance", accountBalance)
			writeFileBalance(accountBalance)
		case 3:
			fmt.Println("Your Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount")
				return
			}
			accountBalance -= withdrawAmount
			fmt.Println("your AccountBalance", accountBalance)
			writeFileBalance(accountBalance)
		default:
			fmt.Println("Good Bye! ")
			return

		}
	}

}

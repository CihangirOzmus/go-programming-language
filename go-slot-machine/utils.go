package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to My Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, lets play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		_, err := fmt.Scan(&bet)
		if err != nil {
			return 0 // if not uint!
		}

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}

	return bet
}

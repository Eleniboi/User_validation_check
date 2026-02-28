package main

import "fmt"

func main() {
	var balance int = 10000
	var name string

	fmt.Print("what is your name: ")
	fmt.Scanln(&name)
	fmt.Println(name, "WELCOME TO SAM's ATM!!")
	for {
		fmt.Println(name, "'s", "current balance: ", balance)
		fmt.Print("Enter amount to withdraw (0 to exit): ")

		var withdraw int

		_, err := fmt.Scanln(&withdraw)

		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if withdraw == 0 {
			fmt.Println(name, "Thank you for using our ATM.")
			break
		}
		if withdraw < 0 {
			fmt.Println("Invalid amount")
			continue
		}
		if withdraw > balance {
			fmt.Println("Insufficient balance")
			continue
		}
		balance = balance - withdraw
		fmt.Println("withdrawal successful.")
	}
}

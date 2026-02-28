package main

import "fmt"

/*this program will ask user for a
number and print all numbers from 1 up to that users
number, if the user enters a non numeric character or
a 0 value print error and stop! */

func main() {

	var i int
	fmt.Print("Enter a number: ")

	_, err := fmt.Scanln(&i)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if i < 1 {
		fmt.Println("Invalid Number Enter a Valid Number: ")
		return
	}
	for p := 3; p <= i; p += 2 {
		/*if p%2 == 0 */ {
			fmt.Println(p)
		}
	}

}

package main

import "fmt"

func main() {

	var i int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&i)

	if err != nil {
		fmt.Println("Invalid number: ")
		return
	}
	if i < 1 {
		fmt.Println("Invalid number, Enter number greater than 0")
		return
	}
	for p := 1; p <= i; p++ {
		if p == 7 {
			break
		}
		fmt.Println(p)
	}
}

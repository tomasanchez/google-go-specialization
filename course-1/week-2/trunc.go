package main

import "fmt"

// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated
// version of the floating point number that was entered. Truncation is the process of removing the digits to the right
// of the decimal place.

func main() {

	var enteredNumber float32

	fmt.Println("Enter a floating point number: ")

	_, err := fmt.Scanln(&enteredNumber)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Truncated number: ")
	fmt.Println(int(enteredNumber))

}

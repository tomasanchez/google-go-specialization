package main

import (
	"math/rand"
	"time"
)

// Write two goroutines which have a race condition when executed concurrently.
// Explain what the race condition is and how it can occur.

// changeNumber changes the value of a number
func changeNumber(luckyNumber *int) {
	*luckyNumber = rand.Int()
}

// printLuckNumber prints the value of a number
func printLuckNumber(luckyNumber *int) {
	println("In this routine is:", *luckyNumber)
}

// main is a demonstration of race condition
func main() {

	// let us assume that the lucky number is 7
	luckyNumber := 7

	println("Lucky number may be 7 or other")

	// One routine changes the value of the lucky number
	go changeNumber(&luckyNumber)

	// Another routine prints the value of the lucky number
	go printLuckNumber(&luckyNumber)

	// The race condition is on the value of the lucky number.

	// Wait a second
	time.Sleep(1 * time.Second)
	// The value is different.
	println("However the value is:", luckyNumber)
}

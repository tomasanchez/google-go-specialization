package main

import "fmt"

// Write a Bubble Sort program in Go.
// The program should prompt the user to type in a sequence of up to 10 integers.
//The program should print the integers out on one line, in sorted order, from least to greatest.
//Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int) {

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}

}

func main() {

	// Input 10 integers
	var input [10]int
	fmt.Println("Enter 10 integers: ")
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
	fmt.Println("Input: ", input)
	BubbleSort(input[:])
	fmt.Println("Bubble Sorted: ", input)

}

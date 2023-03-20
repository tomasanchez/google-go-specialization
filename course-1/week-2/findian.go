package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
// The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’,
// and contains the character ‘a’.
//
// The program should print “Not Found!” otherwise.
// The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

func main() {

	errorMsg := "Not Found!"
	successMsg := "Found!"
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter some characters: ")
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(errorMsg)
		return
	}
	line = strings.TrimSuffix(strings.TrimSpace(strings.ToLower(line)), "\r\n")

	if strings.HasPrefix(line, "i") &&
		strings.HasSuffix(line, "n") &&
		strings.Contains(line, "a") {
		fmt.Println(successMsg)
	} else {
		fmt.Println(errorMsg)
	}

}

package main

import (
	"fmt"
	"strings"
)

// Write a program which allows the user to get information about a predefined set of animals.

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
// Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
// The first string is the name of an animal, either “cow”, “bird”, or “snake”.
// The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”
// . Your program should process each request by printing out the requested data.

// Your program should call the appropriate method when the user makes a request.

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func PromptUser() (string, string) {

	var animalName, animalInfo string

	fmt.Print("> ")

	_, err := fmt.Scan(&animalName, &animalInfo)
	if err != nil {
		return "", ""
	}

	return strings.ToLower(animalName), strings.ToLower(animalInfo)
}

func ExecuteRequest(animalInfo string, animal Animal) {
	switch animalInfo {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid animal info")
	}
}

func main() {

	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		animalName, animalInfo := PromptUser()

		animal, ok := animals[animalName]

		if !ok {
			fmt.Println("Invalid animal name")
			continue
		}

		ExecuteRequest(animalInfo, animal)
	}

}

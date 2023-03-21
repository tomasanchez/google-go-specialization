package main

import (
	"fmt"
	"strings"
)

// Write a program which allows the user to get information about a predefined set of animals.

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
// Your program should continue in this loop forever.
//
// Every command from the user must be either a “newanimal” command or a “query” command.
//
// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”.
// The second string is an arbitrary string which will be the name of the new animal.
// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
//
// Each “query” command must be a single line containing 3 strings.
// The first string is “query”.
// The second string is the name of the animal.
// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
//
// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
// which take no arguments and return no values.
//
// Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak()

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func PromptUser() (string, string, string) {

	var command, animalName, arg2 string

	fmt.Print("> ")

	_, err := fmt.Scan(&command, &animalName, &arg2)
	if err != nil {
		return "", "", ""
	}

	return strings.ToLower(command), strings.ToLower(animalName), strings.ToLower(arg2)
}

func ExecuteQuery(method string) func(Animal) {

	switch method {

	case "eat":
		return func(a Animal) {
			a.Eat()
		}

	case "move":
		return func(a Animal) {
			a.Move()
		}

	case "speak":
		return func(a Animal) {
			a.Speak()
		}

	default:
		return func(a Animal) {
			fmt.Println("Invalid animal method")
		}
	}
}

func CreateAnimal(animal string) (Animal, error) {
	switch animal {
	case "cow":
		return Cow{}, nil
	case "bird":
		return Bird{}, nil
	case "snake":
		return Snake{}, nil
	default:
		return nil, fmt.Errorf("invalid animal type")
	}
}

func main() {

	animals := map[string]Animal{}

	for {
		command, animalName, arg2 := PromptUser()

		switch command {

		case "newanimal":
			// Create a new animal of the requested type
			animal, err := CreateAnimal(arg2)
			if err != nil {
				fmt.Println(err)
				continue
			}
			animals[animalName] = animal
			fmt.Println("Created it!")

		case "query":
			animal, ok := animals[animalName]

			if !ok {
				fmt.Println("Animal not found")
				continue
			}

			method := ExecuteQuery(arg2)
			method(animal)
		default:
			fmt.Println("Invalid command")
		}

	}

}

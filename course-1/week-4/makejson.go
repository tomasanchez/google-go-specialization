package main

import (
	"encoding/json"
	"fmt"
)

//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
//Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

func main() {

	var name string
	var address string

	// get user input for name
	fmt.Println("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get user input for address
	fmt.Println("Enter your address: ")
	_, err = fmt.Scan(&address)
	if err != nil {
		fmt.Println(err)
		return
	}

	//create a map
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	//convert map to json
	marshalled, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
		return
	}

	//print json
	fmt.Println(string(marshalled))

}

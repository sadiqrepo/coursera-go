package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	var jsonList = make(map[string]string)
	fmt.Print("Please enter a name:\n")
	in := bufio.NewReader(os.Stdin)
	name, err1 := in.ReadString('\n')

	if err1 != nil {
		fmt.Println("An error occured while reading entered name. Please try again", err1)
		return
	}

	name = strings.TrimSpace(name)
	jsonList["name"] = name

	fmt.Print("Please enter an address:\n")
	ad := bufio.NewReader(os.Stdin)
	address, err2 := ad.ReadString('\n')

	if err2 != nil {
		fmt.Println("An error occured while reading entered address. Please try again", err2)
		return
	}

	address = strings.TrimSpace(address)
	jsonList["address"] = address

	json_data, err3 := json.Marshal(jsonList)

	if err3 != nil {

		fmt.Println("An error occurred while marshalling into json object. Please try again", err3)
	}

	fmt.Println(string(json_data))

}


/**
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.


*/
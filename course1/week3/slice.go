package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){

	var intList []int = make([]int, 3)
	var input string
	
	for{
		fmt.Println("Enter an Integer or X to exit: ")
		_, err := fmt.Scan(&input)

		if err != nil {
            fmt.Println("Please provide a valid input")
            continue
        }

		if input == "X" || input == "x"{
			fmt.Println("Program Exited...")
			break

		}

		addList, err := strconv.Atoi(input)
		intList = append(intList, addList)
		sort.Ints(intList)
		fmt.Println(intList)
	}
}
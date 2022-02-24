package main

import(
	"fmt"
)

func main(){
	var num float64
	fmt.Println("Enter a floating point number: ")
	fmt.Scanln(&num)
	fmt.Printf("%.0f\n", num)
}

/*

Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered. 
Truncation is the process of removing the digits to the right of the decimal place.

func main() {
	var scannedFloat float32

	fmt.Print("Enter a floating point number: ")
	if _, err := fmt.Scan(&scannedFloat); err == nil {
		fmt.Printf("Truncated integer: %d\n", int(scannedFloat))
	} else {
		fmt.Println("Error scanning float")
	}
}


package main

import (
	"fmt"
	"math"
)

func main() {
	var inputNum float64
	fmt.Println("Enter a float number (ex: 1.12):")
	for {
		_, err := fmt.Scan(&inputNum)
		if err != nil {
			fmt.Println("Input is not a valid number. Try again.")
		} else {
			fmt.Printf("The number you entered '%v' is truncated to '%v'", inputNum, math.Trunc(inputNum))
			break
		}
	}
}
*/
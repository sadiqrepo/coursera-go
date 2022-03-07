package main


import (
	"fmt"
)




func BubbleSort(input []int){
	for endIndex := len(input)-1; endIndex > 0; endIndex--{
		for startIndex := 0; startIndex < endIndex; startIndex++{
			if input[startIndex] > input[startIndex+1] {
				Swap(input, startIndex, startIndex+1)
			}			
		}
	}
	
}

func Swap(slice []int, startValue int, nextValue int){
	if startValue == nextValue {
		return
	}
		temp := slice[startValue]
		slice[startValue] = slice[nextValue]
		slice[nextValue] = temp

}


func DisplayInput(input []int){

	
	for _, value := range input{
		fmt.Print(" ",value)
	}
	fmt.Println()
}


func main(){

	input := make([]int, 10)
	for i := 0; i < 10; i++{
		fmt.Print("Enter i[",i, "]: ")
		fmt.Scan(&input[i])
	}

	fmt.Println("Scanned input values are:")
	DisplayInput(input)

	

	fmt.Println("Sorted input values:")
	BubbleSort(input)
	DisplayInput(input)
	

}


/**

Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

*/
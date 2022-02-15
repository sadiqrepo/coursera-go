package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	filePath := ""
	fmt.Println("Please enter the full path to text file: ")
	fmt.Scanln(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("An error occurred while reading file. Please try again", err)
	}

	defer file.Close()

	var userList []name

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := string(sc.Text())
		firstName := strings.Split(line, " ")[0]
		lastName := strings.Split(line, " ")[1]

		uList := name{firstName, lastName}

		userList = append(userList, uList)

	}

	if err := sc.Err(); err != nil {
		log.Fatal("Error in Reading file. Please try again.", err)
	}

	for _, uList := range userList {

		fmt.Print(uList.fname)
		fmt.Print(" ")
		fmt.Println(uList.lname)
	}

}

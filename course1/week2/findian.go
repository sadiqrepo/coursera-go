package main

//import "fmt"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    
    fmt.Print("Please enter a string:\n")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
   	line = strings.TrimSpace(line)
	line = strings.ToLower(line)
	if err != nil{
		fmt.Println("Please enter a valid string!!")
	}else if strings.Contains(line,"a") && (line[0]=='i') && (line[len(line)-1]=='n' ){
        fmt.Println("Found!")
    }else{
        fmt.Println("Not Found!")
    }
}

  /**
Write a program which prompts the user to enter a string. 
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. 
The program should print “Found!” if the entered string starts with the character ‘i’, 
ends with the character ‘n’, and contains the character ‘a’. 
The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
 “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
  The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”. 
  */
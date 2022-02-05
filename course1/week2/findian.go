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

  
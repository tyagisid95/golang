package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome :=
		"Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating of our pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Print("Thankyou for the rating:", input);

	/*Now i wanted to add +1 in the rating
	  - Firs we need to convert string to int or float
	  -using package "strconv"
	  - to remove space use  strings.TrimSpace(input)
	**/
	// numRating, err  :=strconv.ParseFloat(input,64)
	numRating, err  :=strconv.ParseFloat(strings.TrimSpace(input),64)

	//nil signifies not empty something exist
	if err !=nil{
		 // strconv.ParseFloat(input,64)
    fmt.Println(err)
	// strconv.ParseFloat: parsing "3\n": invalid syntax
	}else{
		//Trim
		fmt.Println("Rating is:",numRating+1)
	}
	
}

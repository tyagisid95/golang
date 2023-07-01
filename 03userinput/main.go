package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to  user input"
	fmt.Println(welcome)
	fmt.Print("Enter the rating of our pizza:")

	/* for input using bufio package*/
	reader := bufio.NewReader(os.Stdin)

	// comma ok || err errr
    // to store value after reading using comma ok variable ,
	// firs var store value of reader and after comma(,) read error
	input, _ := reader.ReadString('\n') //\n => till get new line
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)

}

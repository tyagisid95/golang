package main

import (
	"fmt"
)

// var ptr *int
/*
if you declare var as a pointer but not initialize , it will contain<nil>
*/

func main() {
	fmt.Println("Welecome to a class on pointers")
	//  var ptr *int
	//  fmt.Println("If You declare var as a pointer but not initialize , it will contain<nil> ", ptr)

	//case2
     myNumber :=23
	 var ptr = &myNumber  // referecing myNumber variable

	 fmt.Println("Will return memory address", ptr)

	 fmt.Println("Will return actual value", *ptr)

	 *ptr = *ptr +3
	 fmt.Println("Pointer value will be override",ptr)
}

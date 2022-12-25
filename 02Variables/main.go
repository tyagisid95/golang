package main

import "fmt"

func main()  {
	var username string = "siddhartha"
	fmt.Println(username)
	fmt.Printf("type of username : %T \n",username)

	var isLoogedIn bool= false
	fmt.Printf("isloogedin type: %T \n",isLoogedIn)

	// var somevalue int= 255
	// fmt.Printf("\n type of somevalue: %T",somevalue)

	// uint8	8 bits/1 byte	0 to 255
    // uint16	16 bits/2 byte	0 to 65535
    // uint32	32 bits/4 byte	0 to 4294967295
    // uint64	64 bits/8 byte	0 to 18446744073709551615   
	var somevalue uint32= 256
	fmt.Println(somevalue)
	fmt.Printf("type of somevalue: %T \n",somevalue)

	// float32	32 bits	-3.4e+38 to 3.4e+38.
    // float64	64 bits	-1.7e+308 to +1.7e+308.

	// no var
	txt3 := "World 1"
	fmt.Println(txt3)

}
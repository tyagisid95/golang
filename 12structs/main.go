package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs is GOLANG")
	/*
		-No Inheritance in Golang
		-No super or parent

	*/
	siddhartha := User{"siddhartha", "sidtyagi@gmail.com", true, 18}
	fmt.Println(siddhartha)
    fmt.Printf("User details with key %+v ", siddhartha)
	fmt.Printf("Name is %v & Email is %v ", siddhartha.Name, siddhartha.Email)
}

// how to define struct
// Noted all are caps keys and type struct =>User
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

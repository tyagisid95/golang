package main

import (
	"fmt"
)

func main() {
	fmt.Println("Methods in golang")
	/*
		-No Inheritance in Golang
		-No super or parent

	*/
	siddhartha := User{"siddhartha", "sidtyagi@gmail.com", true, 18}
	fmt.Println(siddhartha)
    fmt.Printf("User details with key %+v ", siddhartha)
	fmt.Printf("Name is %v & Email is %v \n", siddhartha.Name, siddhartha.Email)
     siddhartha.getStatus()
	 siddhartha.newEmail()
	 fmt.Printf("Name is %v & Email is %v \n", siddhartha.Name, siddhartha.Email)
}

// how to define struct
// Noted all are caps keys and type struct =>User
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User)getStatus(){
  fmt.Println("User status :",u.Status)
}
/* it will not override user email*/
// func (u User)newEmail(){
// 	u.Email ="siddhartha@gmail.com"
// 	fmt.Println("New email is :",u.Email)
// }

/* if you want to override existing email with new email
-User Pointer
*/
func (u *User)newEmail(){
	u.Email ="siddhartha@gmail.com"
	fmt.Println("New email is :",u.Email)
}
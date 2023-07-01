package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentDate := time.Now()
	fmt.Println("Present Date is: ", presentDate)

	//date form 01-02-2006
	fmt.Println(presentDate.Format("01-02-2006"))
	//for day  have to add Monday
	fmt.Println(presentDate.Format("01-02-2006 Monday"))
    //for current time "15:04:05"
	fmt.Println(presentDate.Format("01-02-2006 15:04:05 Monday"))

	// 
	createdDate := time.Date(2023, time.July,01,23,0,0,0,time.UTC);
	fmt.Println("Created Date Is :",createdDate.Format("01-02-2006 Monday"))
}

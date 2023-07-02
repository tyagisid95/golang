package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the my Array")
    //syntax varname [length mandatory] variable type
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "orange"
	fruitList[3] = "kiwi"

	fmt.Println("Fruitlist is :", fruitList)// [Apple orange   kiwi] it will print space on index 2 if not putting value
	fmt.Println("Fruitlist length is :", len(fruitList))

	//case2 declare and initializing 
     fmt.Println(`-----------------------------Case2---declare and initializing ----------`)
	var veggiesList = [3]string{"potato,tomato,beans"}
	fmt.Println("Veggies list is: ",veggiesList)
	fmt.Println("Veggies list length is: ",len(veggiesList))

}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(`Welcome to learn type slices, dont need to define length as we dfined array`)
	//syntax

	var fruitList = []string{}
	fmt.Printf("Type of fruitList is :%T ", fruitList)

	//append method expends automatically memory for us
	fruitList = append(fruitList, "apple", "mango", "kiwi")
	fmt.Println(fruitList)
	//range not inclusive [x:n] [startIndex:endIndex]

	//  fruitList = append(fruitList[1:])
	//  fmt.Println("start with index 1",fruitList) //[mango kiwi]

	//  fruitList = append(fruitList[:2])
	//  fmt.Println("end with index 2",fruitList)

	fruitList = append(fruitList[1:2])
	fmt.Println("start with 1 and end with 2 ", fruitList) //[mango]

	//Another wayt to using make()

	highscores := make([]int, 4)
	//default allocation
	highscores[0] = 834
	highscores[1] = 334
	highscores[2] = 534
	highscores[3] = 634
	//  highscores[4]=934// out of bound due to lenght 4

	fmt.Println(highscores)
	//dynamically allocation
	highscores = append(highscores, 900, 800, 1000) // allocate memory on run time
	fmt.Println(highscores)

	//for sorting array

	sort.Ints(highscores)
	fmt.Println(highscores)

	//check whether is sorted or not

	fmt.Println("isSorted:", sort.IntsAreSorted(highscores)) //true

	//How to remove a value from slices based on index

	var courses = []string{"Reactjs", "Javascript", "Swift", "Python", "Ruby"}
	fmt.Println("Courses are :", courses)

	const index = 2
	fmt.Println("Courses on index 2 is:",courses)

	//print till index =>2 [Reactjs, Javascript]
	//start with index+1 =>3 [Python, Ruby]
	//then append
    courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated courses are :", courses)

}

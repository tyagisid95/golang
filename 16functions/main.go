package main

import "fmt"

func main() {
	fmt.Println("Functions in golangs")
	greeter()

	/* Function inside function is not allowed */
	// func greeterTwo(){
	// 	fmt.Println("Another method")
	// }
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Sum of twon mubers is : ", result)

	proAdd := proAdder(3,5,6,7)
	fmt.Println("Pro adder : ",proAdd)
}

func greeter() {
	fmt.Println("Namaste")
}
func greeterTwo() {
	fmt.Println("Another method")
}

/*
 argument type  is must
return type of function is must
*/
func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) int{
	total :=0

	for _,v:= range values{
		 total = total +v
	}
	return total
}



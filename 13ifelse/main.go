package main

import "fmt"

func main() {
	fmt.Println("If else in GOLANG")

	loginCount := 10
	var result string
	//case 1
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		fmt.Println("Exactly 10")
	}
	fmt.Print(result)
 
    //cas2
	if 9%2==0{
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	//case3
	//lets suppose storing data from webrequest and checking those data
	if num :=3; num<10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is greater than 10")
	}
}

package main

import "fmt"

func main(){
	defer fmt.Println("Defer in golang")
	fmt.Println("Hello")
	defer fmt.Println("World")
	/*
	If you use keyword as defer it will excute at last
	If there are multiple statment marked as defer , will follo statck pattern (LIFO)
	*/ getDefer()
}

func getDefer(){
  for i :=0 ; i< 5; i++{
	defer fmt.Println("vlaue is :", i)
  }
}

/*Output
Hello
4 3 2 1 0
World
Defer in golang
*/
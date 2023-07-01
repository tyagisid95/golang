package main

import ("fmt")

func main() {

	fmt.Println(`
	/*  ---------Memory Management
	GC -> Garbage collection automatically happen
		  -Out of scope  or nil


 There are two type to create memory
 1) new()
   - Allocate memory but not initilalize
   -you will get a memory address
   -zeroed  storage(Imp for interview)

 majority using make()
 
 2 make()
  -Allocate memory and initialze
  -you will get a memory address
  -non zeroed storage (Imp for interview)
*/`)
	
}

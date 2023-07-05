package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	/* file creation happened with os package*/
	/* file read operation happened with ioutil package*/
	/*io.WriteString helps to write content ina file*/
    content := "This content will go in a file."
	file, err := os.Create("./sampleFile.text")

	// if err !=nil{
	// 	panic(err)
	// }
	checkErr(err)
	length, err := io.WriteString(file, content)
	// if err !=nil{
	// 	panic(err)
	// }
	checkErr(err)
	fmt.Println("Length of file is : ",length)

	//recommended mark as defer cand close so that next operation can be happened later on
	defer file.Close()
	readFile("./sampleFile.text")
}

func readFile(filename string){
	dataByte,err :=ioutil.ReadFile(filename)
	// if err !=nil{
	// 	panic(err)
	// }
	checkErr(err)
	fmt.Println("File data in bytes:", dataByte)
	fmt.Println("File data in string:", string(dataByte))
}

func checkErr(err error){
  if err !=nil{
	panic(err)
  }
}
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Handling webrequest in golangs")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is type of : %T", resp)

	// caller's responsiblity to close the connections
	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	// fmt.Println("Databytes are :", databytes)
	responseBody := string(databytes)
	fmt.Printf("\nResponse is in string: %+v", responseBody)
}

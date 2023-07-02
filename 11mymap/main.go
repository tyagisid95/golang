package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")
    // map[key]value
	languages := make(map[string]string)
	languages["JS"]="Javscript"
	languages["PY"]="Python"
	languages["RB"]="Ruby"
	fmt.Println("Languages are : ",languages)

	fmt.Println("JS stands for : ",languages["JS"])

	//delete from map
	delete(languages,"RB")
	fmt.Println("Updated languages are : ",languages)

	//loops in GO LANG
	 for key,value := range languages {
		fmt.Printf("For key %v, & and for value %v\n", key,value);
		
	 }

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switc case in golangs")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 //6=>0 to 5 range not inclusive
	fmt.Println("Dice number is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1  and you can open")
	case 2:
		fmt.Println("You can move to spot 2")
	case 3:
		fmt.Println("You can move to spot 3")
	case 4:
		fmt.Println("You can move to spot 4")
	case 5:
		fmt.Println("You can move to spot 5")
	case 6:
		fmt.Println("You can move to spot 6 and roll dice again")
	default:
		fmt.Println("What was that!")

	}
}

package main

import "fmt"

func main() {
	fmt.Printf("Welcome to the Band Name Generator\n")
	fmt.Printf("What's the name of the city you grew up in?\n")
	var city string
	fmt.Scanf("%s", &city)

	var pet string
	fmt.Printf("What's your pet's name?\n")
	fmt.Scanf("%s", &pet)
	bandName := fmt.Sprintf("%s %s", city, pet)
	fmt.Printf("Your band name could be %s", bandName)
}

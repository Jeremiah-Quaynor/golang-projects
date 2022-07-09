package main

import "fmt"

func main() {
	var firstName string = "Jeremiah"
	lastName := "Quaynor"
	var curYear int = 2022
	birthDate := 2000
	// runes can be used to store special characters
	var emojis rune = '🔥'
	nextAge := curYear - birthDate + 1

	fmt.Println(firstName, lastName, string(emojis))
	fmt.Println("Your current age is ", curYear-birthDate)
	fmt.Printf("Your next age is %v", nextAge)
}

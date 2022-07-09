package main

import "fmt"


func main() {
	const PiNumber = float32(3.14)
	radius := float32(5)
	var circum float32 = PiNumber * 2 * radius 

	fmt.Printf("the circumference is: %v ", circum)
}
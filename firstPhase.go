package main
import "fmt"

func main() {
	var firstName string = "Jeremiah"
	lastName := "Quaynor"
	var curYear int = 2022
	birthDate := 2000
	fmt.Println(firstName, lastName)
	fmt.Println("Your current age is ", curYear-birthDate)
	fmt.Println("Your next age is ", curYear-birthDate+1)

}
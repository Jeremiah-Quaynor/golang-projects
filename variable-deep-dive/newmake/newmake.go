// package variabledeepdive
package main

import "fmt"


func main () {
	// new & make 

	number := new(int)
	fmt.Print(number)
	fmt.Print(*number)

	anotherNumber := 0
	numberAddress := &anotherNumber

	fmt.Println(numberAddress)


	// hobbies := []string{"Sports", "Reading"}

	hobbies := make([]string, 2, 10)
	moreHobbies := new([]string)
	// evenMoreHobbies := []string{} //make([]string)

	fmt.Println(*moreHobbies)

	*moreHobbies = append(*moreHobbies, "Sports")

	// aMap := make(map[string]int, 5) // making a map 

	hobbies[0] = "Sports"
	hobbies[1] = "Cooking 2"

	fmt.Println(hobbies)

	hobbies = append(hobbies, "Cooking", "Dancing")


	fmt.Println(hobbies)

}
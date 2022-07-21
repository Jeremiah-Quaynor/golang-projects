package main

import "fmt"


var userName string = "Jerry"

func main() {
	shouldContinue := true
	userName := "Jeremiah" //shadow the higher-level variable 

	if shouldContinue {
		userAge := 34

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}
printData()
}

func printData () {

	fmt.Println(userName)

}
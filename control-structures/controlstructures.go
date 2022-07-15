package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge, err :=getUserAge()
	if err != nil {
		fmt.Print("Invalid Input!")
		return
	}
	//taking user input 
	// var userAge int
	// fmt.Print("Please enter your age: ")
	// fmt.Scan(&userAge)

	if (userAge >= 30 && userAge <= 50) || (userAge >= 80) {
		fmt.Println("You are eligible for our senior jobs.")
	}else if userAge >=18 {
		fmt.Println("Welcome the the club!")
	}else {
		fmt.Println("You are way too young!")
	}
	fmt.Println("Goodbye!")

}

func getUserAge()(int, error) {
	reader := bufio.NewReader(os.Stdin)
	
	// error handling: 
	fmt.Print("Please enter your year of birth: ")
	userYearInput, _:= reader.ReadString('\n')
	userYearInput = strings.Replace(userYearInput, "\n","", -1)
	userYear,err := strconv.ParseInt(userYearInput, 0, 64)

	return int(userYear), err 
} 
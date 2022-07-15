package main

import (
	"bufio"
	"reflect"
	"strconv"
	"strings"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice := getUserChoice()

	if selectedChoice == 0 {
		fmt.Println("invalid choice, exiting!")
		return
	}

	if selectedChoice == 1 {
		calculateSumUpToNumber()

	} else if selectedChoice == 2 {
		calculateFactorial()

	} else if selectedChoice == 3 {

		calculateSumManually()
	} else {

		calculateSumList()
	}

}

func getUserChoice() (int) {
	var userInput int
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers to the number X")
	fmt.Println("1) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("3) Sum up a list of numbers")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&userInput)
	// userInput, err := reader.ReadString('\n')

	// if err != nil {
	// 	return "", err
	// }
	// userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == 1 ||
		userInput == 2 ||
		userInput == 3 ||
		userInput == 4 {
		return userInput

	} else {
		return 0
		// return "", errors.New("INVALID INPUT")
	}
}

func calculateSumUpToNumber() {
	var chosenNumber int
	fmt.Println("Please enter your number: ")
	fmt.Scan(&chosenNumber)

	// chosenNumber, err := getInputNumber()
	// if err != nil {
	// 	fmt.Println("Invalid input")
	// }
	sum := 0
	for i := 0; i < chosenNumber; i++ {
		sum += i
	}

	fmt.Printf("Result: %v", sum)

}

func calculateFactorial() {
	var chosenNumber int
	fmt.Print("Please enter your number: ")
	fmt.Scan(&chosenNumber)
	fact := 1
	for i:=1; i<= chosenNumber; i++ {
		fact = fact* i 
	}

	fmt.Printf("fact of %v is: %v", chosenNumber, fact )

}

func calculateSumManually() {
	//Has buags
	isEnteringNumbers := true
	sum := 0
	fmt.Println("Enter Zero to quit")

	for isEnteringNumbers {
		var num int
		fmt.Scan(&num)
		if  reflect.TypeOf(num) != reflect.TypeOf(5) {
			fmt.Println("Invalid input")
			isEnteringNumbers = false
		}
		sum = sum + num;	
	}
	fmt.Printf("The sum of your values is: %v", sum)
}

func calculateSumList() {
	fmt.Println("Please enter a comma-separated list of numbers.")

	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return 
	}

	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")
	sum := 0

	for index, value:= range inputNumbers{
		number,_ := strconv.ParseInt(value,0,64)
		sum = sum + int(number)
		fmt.Printf("Index: %v, Value: %v \n",index,value)
	}
	fmt.Printf("Result: %v", sum)
}

func getInputNumber() int {
	fmt.Println("Please make your choice: ")

	var chosenNumber int
	fmt.Scan(&chosenNumber)
	// inputNumber, err := reader.ReadString('\n')

	// if err !=nil {
	// 	return  0, err
	// }

	// inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	// chosenNumber, err :=strconv.ParseInt(inputNumber, 0, 64)

	// if err !=nil {
	// 	return 0, err
	// }

	return chosenNumber
}

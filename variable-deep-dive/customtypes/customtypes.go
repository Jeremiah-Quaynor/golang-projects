package main

import "fmt"

type person struct {
	name string
	age  int
}

type personData map[string]person //creating custom data type


type customNumber int

func (number customNumber) pow (power int)  customNumber {
	result := number
	for i := 1; i < power; i++ {
		result = result * number
	}
	return result
}

func main() {

	var people personData = personData{
		"P1" : {"jerry", 23},
	}

	fmt.Println(people)


	var dummyNumber customNumber = 4

	poweredNumber := dummyNumber.pow(4)
	fmt.Print(poweredNumber)

}
package main

import (
	"fmt"
)







func main () {
	detials := make(map[int]int)

	detials[833435] = 8712343
	detials[23438] = 678675
	detials[67887] = 3468345


	welcomeMessage()
	accNum, pin := getDetails()
	fmt.Print(accNum,pin)
	isValid :=verifyDetails(accNum,pin)
	fmt.Print(isValid)
	performTransaction()
	terminateSession()

}

func welcomeMessage() {
	fmt.Println("############################")
	fmt.Println("Welcome To Your Friendly ATM")
	fmt.Println("-------------------")
}

func getDetails () (int, int) {
	var bankAcc int 
	var pin int 

	fmt.Print("Please Your Account Number: ")
	fmt.Scan(&bankAcc)

	fmt.Print("Please Your Account Pin: ")
	fmt.Scan(&pin)

	if bankAcc == 0 || pin == 0 {
	fmt.Print("Account number or Pin invalid")
	return 0,0
	}
	return bankAcc,pin
}

func verifyDetails (acccNum, pin) bool {

	return true
}

func performTransaction () {

}


func terminateSession () {

}

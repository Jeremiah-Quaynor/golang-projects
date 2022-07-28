package main

import (
	"fmt"
)

// var reader = bufio.NewReader(os.Stdin)

func main () {
	welcomeMessage()
	getDetails()
	verifyDetails()
	performTransaction()
	terminateSession()

}

func welcomeMessage() {
	fmt.Println("############################")
	fmt.Println("Welcome To Your Friendly ATM")
	fmt.Println("-------------------")
}

func getDetails () (int, int) {
	var bankAcc, pin int 

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

func verifyDetails () {

}

func performTransaction () {

}


func terminateSession () {

}

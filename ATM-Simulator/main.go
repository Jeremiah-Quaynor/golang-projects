package main

import (
	"fmt"
)

func main() {
	welcomeMessage()
	accNum, pin := getDetails()
	fmt.Println("Authenticating account...")
	isValid := verifyDetails(accNum, pin)
	if isValid {
		performTransaction()
	}else {
		fmt.Println("You have entered a wrong account number or pin code please try again")
	}
	terminateSession()

}

func welcomeMessage() {
	fmt.Println("############################")
	fmt.Println("Welcome To Your Friendly ATM")
	fmt.Println("-------------------")
}

func getDetails() (int, int) {
	var bankAcc int
	var pin int

	fmt.Print("Please Your Account Number: ")
	fmt.Scan(&bankAcc)

	fmt.Print("Please Your Account Pin: ")
	fmt.Scan(&pin)

	if bankAcc == 0 || pin == 0 {
		fmt.Print("Account number or Pin invalid")
		return 0, 0
	}
	return bankAcc, pin
}

func verifyDetails(acccNum int, pin int) bool {
	var BankDetials = map[int]int{
		123456: 1234,
		4111111111111111: 8923,
		5496198584584769: 8232,
		2223000048400011: 0233,
		2223520043560014: 8932,
		378282246310005: 9003,
	}
	var state bool = false
	for bankNumber, pincode := range BankDetials {
		if bankNumber == acccNum && pincode == pin {
			state = state ||true
		}
	}

	return state
}

func performTransaction() {
	var request int
	var amount int
	fmt.Println("Dear Sir/Madam, welcome...")
	fmt.Println("")
	fmt.Println("1)Check Balance		2)Withdrawal\n3)Savings		4)Bank Statement")
	fmt.Scan(&request)
	switch request {
	case 1:
		fmt.Print("CHECK BALANCE")
		fmt.Print("Your balance is $400.00")
	case 2:
		fmt.Print("WITHDRAWAL")
		fmt.Print("Please enter an amount")
		fmt.Scan(&amount)
		fmt.Printf("You have withdrew $%v from your account", &amount)
	case 3:
		fmt.Print("WELCOME TO YOUR SAVINGS")
	case 4:
		fmt.Print("BANK STATEMENT")
	default:
		fmt.Print("GOING TO SLEEP NOW...")
	}

}

func terminateSession() {
	fmt.Print("THANK YOU FOR BANKING WITH US")
}

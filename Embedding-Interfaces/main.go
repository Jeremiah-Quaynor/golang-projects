package main

import (
	"bufio"
	"fmt"
	"os"

)

type Storer interface {
	Store(fileName string)
}
type Prompter interface {
	PromptForInput() 
	Storer

}

type PromptStorer interface {
	Prompter

}

type userInputData struct {
	input string

}

//embedding structs
type user struct {
	firstName string
	lastName string 
	*userInputData
}

func newUser (first string, last string) *user{
	return &user{
		firstName: first,
		lastName: last,
		userInputData: &userInputData{},
	}
}

func newUserInputData () *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForInput() {
	fmt.Print("Your input please: ")

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Fetching user input failed!")
		return
	}

	usr.input = userInput
}

func (usr *userInputData ) Store(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Creating file failed!")

	}

	defer file.Close() //might want to wrap into anonymous fn for error handling

	file.WriteString(usr.input)

}

func main () {
	data := newUserInputData()
	jay := newUser("Jeremiah", "Quaynor")

	jay.PromptForInput()
	jay.Store("jay.txt")

	// data.PromptForInput()
	// data.Store("user1.txt")

	handleUserInput(data)

}

func handleUserInput(container Prompter) {
	fmt.Println("Ready store your data!")
	container.PromptForInput()
	container.Store("user2.txt")
	
}
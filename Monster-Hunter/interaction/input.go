package interaction

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "os/user"
	// "strings"
)

// var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	// for {
	// 	playerChoice, _ := getPlayerInput()

	// 	if playerChoice == "1" {
	// 		return "ATTACK"
	// 	}else if playerChoice == "2" {
	// 		return "HEAL"
	// 	}else if playerChoice == "3" && specialAttackIsAvailable {
	// 		return "SPEACIAL ATTACK"
	// 	}
	// 	fmt.Println("Fetching the user input fialed. Please try again.")
	// }
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == 1 {
			return "ATTACK"
		}else if playerChoice == 2 {
			return "HEAL"
		}else if playerChoice == 3 && specialAttackIsAvailable {
			return "SPEACIAL ATTACK"
		}
		fmt.Println("Fetching the user input fialed. Please try again.")
	}
}

func getPlayerInput() (int, error) {
	fmt.Print("Your Choice here: ")
	var userInput int
	var err error 
	fmt.Scan(&userInput, &err)
	// userInput, err := reader.ReadString('\n')

	if err != nil {
		return  0 ,err
	}
	// userInput  = strings.Replace(userInput, "\n", "", -1)

	// return userInput, nil
	return userInput, nil
}
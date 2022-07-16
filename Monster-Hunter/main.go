package main

import (
	"github.com/jeremiah-quaynor/Monster-Hunter/interaction"
)

var currentRound = 0

func main() {

	startGame()

	winner := ""// Player || Monster || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame () {
	interaction.PrintGreeting()
}

func executeRound() (string) {
	currentRound++ 
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvaibleActions(isSpecialRound)

	return ""
}

func endGame () {

}
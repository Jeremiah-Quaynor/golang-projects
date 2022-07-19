package main

import (
	"github.com/jeremiah-quaynor/Monster-Hunter/interaction"
	"github.com/jeremiah-quaynor/Monster-Hunter/actions"
)

var currentRound = 0

func main() {

	startGame()

	winner := ""// Player || Monster || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}



func startGame () {
	interaction.PrintGreeting()
}

func executeRound() (string) {
	currentRound++ 
	isSpecialRound := currentRound%3 == 0
	interaction.ShowAvaibleActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealValue int 
	var monsterAttackDmg int


	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	}	else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	}else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action : userChoice,
		PlayerHealth: playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue: playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)

	if playerHealth <= 0 {
		return "Monster"

	}else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame (winner string) {
	interaction.DeclareWinner(winner)
}
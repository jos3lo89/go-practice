package ch2

import (
	"fmt"
	"math/rand"
	"time"
)

// rockPaperScissorDemo runs a simple rock, paper and scissors game
func RockPaperScissor() {
	choices := []string{"rock", "paper", "scissors"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var playerChoice string

	fmt.Print("enter rock, paper or scissors: ")
	fmt.Scan(&playerChoice)

	if playerChoice != "rock" && playerChoice != "paper" && playerChoice != "scissors" {
		fmt.Println("invalid option")
		return
	}

	computerChoice := choices[r.Intn(len(choices))]
	fmt.Println("computer choice:", computerChoice)

	if playerChoice == computerChoice {
		fmt.Println("it's a tie")
	} else if (playerChoice == "rock" && computerChoice == "scissors") || (playerChoice == "paper" && computerChoice == "rock") || (playerChoice == "scissors" && computerChoice == "paper") {
		fmt.Print("you win")
	} else {
		fmt.Print("you lose")
	}

}

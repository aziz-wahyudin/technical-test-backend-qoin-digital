package main

import (
	"fmt"
	"math/rand"
	"time"
)

func judiDadu(numGamblers, numDice int) int {
	rand.Seed(time.Now().UnixNano())

	// initialize the gamblers' scores and dice
	gamblerDice := make([]int, numGamblers)
	gamblerScores := make([]int, numGamblers)
	for i := 0; i < numGamblers; i++ {
		gamblerDice[i] = numDice
	}

	// play the game until only one gambler has dice
	for {
		// each gambler rolls their dice
		for i := 0; i < numGamblers; i++ {
			for j := 0; j < gamblerDice[i]; j++ {
				roll := rand.Intn(6) + 1
				if roll == 6 {
					gamblerScores[i] += 6
					gamblerDice[i]--
				} else if roll == 1 {
					gamblerDice[i]--
					if i == numGamblers-1 {
						gamblerDice[0]++
					} else {
						gamblerDice[i+1]++
					}
				}
			}
		}

		// check if the game is over
		remaininggamblers := 0
		for i := 0; i < numGamblers; i++ {
			if gamblerDice[i] > 0 {
				remaininggamblers++
			}
		}
		if remaininggamblers == 1 {
			break
		}
	}

	// find the gambler with the highest score
	maxScore := -1
	winner := 0
	for i := 0; i < numGamblers; i++ {
		if gamblerScores[i] > maxScore {
			maxScore = gamblerScores[i]
			winner = i + 1
		}
	}

	// return the winner's gambler number
	return winner
}

func main() {
	fmt.Println("Aww judi, menjanjikan kemenanagan~ ... and the winner is player number:", judiDadu(3, 4))
}

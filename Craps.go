package main

import "fmt"
import "math/rand"

func main() {
	fmt.Println("From PL Team #1")

	var gameStatus, sum, myPoint int

	sum = rollDice() // First roll of the dice
	//fmt.Println(sum)

	switch sum {
	case 7:
		fallthrough
	case 11:
		gameStatus = 1 // Win on first roll
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 12:
		gameStatus = 2 // Lose on first roll
	default:
		gameStatus = 0 // Remember point
		myPoint = sum
		fmt.Printf("Point is: %d\n", myPoint)
	}

	for gameStatus == 0 { // Keep rolling
		sum = rollDice()
		//fmt.Println(sum)
		switch sum {
		case myPoint: // Win by making point
			gameStatus = 1
		case 7: // Lose by rolling 7
			gameStatus = 2
		}
	}

	if gameStatus == 1 {
		fmt.Println("Player wins")
	} else {
		fmt.Println("Player loses")
	}
}

func rollDice() int {
	var die1, die2, workSum int
	die1 = rand.Intn(6)+1
	die2 = rand.Intn(6)+1
	workSum = die1 + die2

	fmt.Printf("Player rolled %d, %d: %d\n", die1, die2, workSum)

	return workSum
}

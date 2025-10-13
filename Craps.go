package main

import "fmt"

func main() {
	fmt.Println("From PL Team #1")

	var gameStatus, sum, myPoint int

	sum = rollDice()

	switch sum {
	case 7:
		fallthrough
	case 11:
		gameStatus = 1
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 12:
		gameStatus = 2
	default:
		gameStatus = 0
		myPoint = sum
		fmt.Printf("Point is: %d\n", myPoint)
		
	}

	// While loop commented to avoid infinite loop
	//while (gameStatus == 0) {}

	if gameStatus == 1 {
		fmt.Println("Player wins")
	} else {
		fmt.Println("Player loses")
	}
}

func rollDice() int {
	var die1, die2, workSum int
	// Placeholder values
	die1 = -1
	die2 = -1
	workSum = die1 + die2

	return workSum
}

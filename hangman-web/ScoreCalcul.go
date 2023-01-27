package main

import "fmt"

func ScoreCalcul(difficulty string, live int, elapsed float64) int {

	var score int

	//Function to calcul the score of the player from the game given in parameters
	if difficulty == "easy" {
		score += 3
		fmt.Printf("+3")
	}
	if difficulty == "medium" {
		score += 5
		fmt.Printf("+5")
	}
	if difficulty == "hard" {
		score += 7
		fmt.Printf("+7")
	}
	if live >= 9 {
		score += 8
		fmt.Printf("+8")
	}
	if live >= 5 && live < 9 {
		score += 5
		fmt.Printf("+5")
	}
	if live > 0 && live < 5 {
		score += 2
		fmt.Printf("+2")
	}
	if elapsed <= 10 {
		score += 16
		fmt.Printf("+16")
	}
	if elapsed > 10 && elapsed <= 20 {
		score += 10
		fmt.Printf("+10")
	}
	if elapsed > 20 && elapsed <= 30 {
		score += 5
		fmt.Printf("+5")
	}
	if elapsed > 30 {
		score += 2
		fmt.Printf("+2")
	}

	return score
}

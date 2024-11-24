package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Symbols
var symbols = []string{"🍒", "🍋", "⭐", "💎", "🎰"}

func performSpin() []string {
	reels := make([]string, 3) // slice of strings with 3 elements
	for i := 0; i < 3; i++ {
		reels[i] = symbols[rand.Intn(len(symbols))]
	}
	return reels
} 

// Check if the combination is winning
func isWinningCombination(spinResult []string) bool {
	// if all reels are the same symbol, then it's a winning combination
	return spinResult[0] == spinResult[1] && spinResult[1] == spinResult[2]
}

func main() {
	// Initialize randomness
	rand.NewSource(time.Now().UnixNano())	
 
	// Perform a spin and store the result
	spinResult := performSpin()

	// Print the spin result
	fmt.Println("Spin Result:", spinResult)

	// Check if the combination is winning
	if isWinningCombination(spinResult) {
		fmt.Println("Win! 🎉")
	} else {
		fmt.Println("Lose! 😢")
	}
}

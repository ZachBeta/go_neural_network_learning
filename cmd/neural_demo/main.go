package main

import (
	"fmt"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

func main() {
	// Set random seed for reproducibility
	neural.SetRandomSeed(time.Now().UnixNano())

	// Create a new neural network
	// 9 inputs (board state), 9 outputs (move probabilities)
	network := neural.NewNetwork(9, 9)

	// Print network structure
	fmt.Println("Neural Network Structure:")
	neural.PrintNetwork(network)

	// Create a new game board
	board := game.NewBoard()

	// Print initial board
	fmt.Println("\nInitial Board:")
	fmt.Println(board)

	// Convert board to neural network input
	input := neural.BoardToInput(board)
	fmt.Println("\nNeural Network Input:")
	for i, val := range input {
		fmt.Printf("Input[%d] = %.2f\n", i, val)
	}

	// Perform forward pass
	output := network.Forward(input)
	fmt.Println("\nNeural Network Output:")
	neural.PrintOutput(output)

	// Convert output to move probabilities
	probabilities := neural.OutputToMoveProbabilities(output)
	fmt.Println("\nMove Probabilities:")
	neural.PrintMoveProbabilities(probabilities)

	// Select best move
	bestMove := neural.SelectBestMove(probabilities)
	row, col := neural.MoveIndexToRowCol(bestMove)
	fmt.Printf("\nBest Move: (%d,%d)\n", row, col)

	// Make the move
	board.MakeMove(row, col)
	fmt.Println("\nBoard After Move:")
	fmt.Println(board)

	// Simulate a few more moves
	for i := 0; i < 3; i++ {
		// Switch player
		board.SwitchPlayer()

		// Convert board to neural network input
		input = neural.BoardToInput(board)

		// Perform forward pass
		output = network.Forward(input)

		// Convert output to move probabilities
		probabilities = neural.OutputToMoveProbabilities(output)

		// Select best move
		bestMove = neural.SelectBestMove(probabilities)
		row, col = neural.MoveIndexToRowCol(bestMove)

		// Make the move
		board.MakeMove(row, col)
		fmt.Printf("\nBoard After Move %d:\n", i+1)
		fmt.Println(board)
	}

	// Check for winner
	board.CheckWinner()
	fmt.Printf("\nGame Status: %v\n", board.GetStatus())
}

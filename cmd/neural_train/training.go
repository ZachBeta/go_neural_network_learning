package main

import (
	"math/rand"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

// playGameWithVisualization plays a complete game and returns the game record
func playGameWithVisualization(network *neural.Network, epsilon float64, displayDelay time.Duration) GameRecord {
	// Create a new game board
	board := game.NewBoard()

	// Create a game record
	record := GameRecord{
		States: make([]GameState, 0),
		Winner: "",
	}

	// Play until the game is over
	moveNum := 0
	for board.GetStatus() == game.InProgress {
		// Get the current player
		currentPlayer := board.GetCurrentPlayer()
		playerStr := "X"
		if currentPlayer == game.O {
			playerStr = "O"
		}

		// Convert board to neural network input
		input := neural.BoardToInput(board)

		// Get move probabilities
		output := network.Forward(input)
		probabilities := neural.OutputToMoveProbabilities(output)

		// Select a move
		var move int
		if rand.Float64() < epsilon {
			// Exploration: choose a random valid move
			move = selectRandomValidMove(board)
		} else {
			// Exploitation: use network's prediction
			move = neural.SelectBestMove(probabilities)
		}

		// Convert move index to row and column
		row, col := neural.MoveIndexToRowCol(move)

		// Make the move
		board.MakeMove(row, col)

		// Create a game state
		state := GameState{
			Board:         board,
			Move:          move,
			Player:        playerStr,
			Probabilities: probabilities,
		}

		// Add the state to the record
		record.States = append(record.States, state)

		// Display the game state
		ClearScreen()
		DisplayGameProgress(0, 1, moveNum, playerStr, epsilon)
		DisplayBoard(board)
		DisplayMoveProbabilities(probabilities, move)
		DisplayStrategyInfo(board, move)

		// Wait for the specified delay
		time.Sleep(displayDelay)

		// Check for winner
		board.CheckWinner()

		// Increment move number
		moveNum++
	}

	// Set the winner
	if board.GetStatus() == game.Won {
		// Switch back to the previous player since they won
		board.SwitchPlayer()
		if board.GetCurrentPlayer() == game.X {
			record.Winner = "X"
		} else {
			record.Winner = "O"
		}
	} else {
		record.Winner = "Draw"
	}

	// Set the result for each state
	for i := range record.States {
		state := &record.States[i]
		if record.Winner == "Draw" {
			state.Result = 0.0
		} else if state.Player == record.Winner {
			state.Result = 1.0
		} else {
			state.Result = -1.0
		}
	}

	return record
}

// selectRandomValidMove selects a random valid move
func selectRandomValidMove(board *game.Board) int {
	// Get all valid moves
	validMoves := make([]int, 0)
	for i := 0; i < 9; i++ {
		row, col := neural.MoveIndexToRowCol(i)
		if board.Get(row, col) == game.Empty {
			validMoves = append(validMoves, i)
		}
	}

	// Select a random move
	if len(validMoves) > 0 {
		return validMoves[rand.Intn(len(validMoves))]
	}

	// This should never happen if the game is not over
	return 0
}

// updateNetworkWeights updates the network weights based on a batch of game states
func updateNetworkWeights(network *neural.Network, batch []GameState, learningRate float64) {
	// For each state in the batch
	for _, state := range batch {
		// Convert board to neural network input
		input := neural.BoardToInput(state.Board)

		// Get the target output
		target := make([]float64, 9)
		target[state.Move] = state.Result

		// Update the network weights
		// This is a simplified version of backpropagation
		// In a real implementation, we would use a proper backpropagation algorithm
		output := network.Forward(input)
		for i := range output {
			// Calculate the error
			error := target[i] - output[i]

			// Update the weights
			// This is a very simplified version
			// In a real implementation, we would update the weights properly
			// using the error and the derivative of the activation function
			for j := range input {
				// Update the weight
				// This is a very simplified version
				// In a real implementation, we would update the weights properly
				// using the error and the derivative of the activation function
				network.GetOutputLayer().GetNeuron(i).Weights[j] += learningRate * error * input[j]
			}

			// Update the bias
			// This is a very simplified version
			// In a real implementation, we would update the bias properly
			// using the error and the derivative of the activation function
			network.GetOutputLayer().GetNeuron(i).Bias += learningRate * error
		}
	}
}

// displayTrainingProgress displays the training progress
func displayTrainingProgress(gameNum, totalGames int, record GameRecord, epsilon float64) {
	// Count wins and draws
	xWins := 0
	oWins := 0
	draws := 0

	if record.Winner == "X" {
		xWins = 1
	} else if record.Winner == "O" {
		oWins = 1
	} else {
		draws = 1
	}

	// Display training progress
	DisplayTrainingProgress(gameNum, totalGames, xWins, oWins, draws, epsilon)
}

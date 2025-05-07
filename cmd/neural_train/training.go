package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/display"
	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/logger"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

var (
	// Logger for detailed gameplay information
	gameLogger *logger.Logger
)

func init() {
	// Initialize the logger
	var err error
	gameLogger, err = logger.NewLogger("training", logger.INFO)
	if err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	gameLogger.Info("Training session started")
}

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

		// Log the game state to file
		logGameState(board, move, moveNum, playerStr, epsilon, probabilities)

		// Wait for the specified delay (very short for training)
		time.Sleep(time.Millisecond)

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

	// Log game result
	logGameResult(record)

	return record
}

// logGameState logs detailed game state information to file
func logGameState(board *game.Board, move, moveNum int, player string, epsilon float64, probabilities []float64) {
	row, col := neural.MoveIndexToRowCol(move)
	gameLogger.Info("\nMove %d - Player %s (ε=%.3f)", moveNum+1, player, epsilon)
	gameLogger.Info("Selected move: (%d,%d)", row, col)
	gameLogger.Info("Board state:\n%s", board.String())

	// Log move probabilities with more detail
	gameLogger.Info("Move probabilities:")
	for i, prob := range probabilities {
		r, c := neural.MoveIndexToRowCol(i)
		if i == move {
			gameLogger.Info("  Position (%d,%d): %.2f%% [SELECTED]", r, c, prob*100)
		} else {
			gameLogger.Info("  Position (%d,%d): %.2f%%", r, c, prob*100)
		}
	}

	// Log strategic analysis
	gameLogger.Info("\nStrategic Analysis:")
	if isForkCreation(board, move) {
		gameLogger.Info("  ✓ Fork Creation detected")
	}
	if isForkBlocking(board, move) {
		gameLogger.Info("  ✓ Fork Blocking detected")
	}
	if isWinningMove(board, move) {
		gameLogger.Info("  ✓ Winning Move detected")
	}
	if isBlockingMove(board, move) {
		gameLogger.Info("  ✓ Blocking Move detected")
	}

	// Log board evaluation
	gameLogger.Info("\nBoard Evaluation:")
	gameLogger.Info("  Empty cells: %d", countEmptyCells(board))
	gameLogger.Info("  Potential winning lines: %d", countPotentialWinningLines(board))
	gameLogger.Info("==========================================\n")
}

// logGameResult logs the game result to file with more detail
func logGameResult(record GameRecord) {
	gameLogger.Info("\nGame Summary")
	gameLogger.Info("============")
	gameLogger.Info("Winner: %s", record.Winner)
	gameLogger.Info("Total moves: %d", len(record.States))

	// Log move statistics
	xMoves := 0
	oMoves := 0
	for _, state := range record.States {
		if state.Player == "X" {
			xMoves++
		} else {
			oMoves++
		}
	}
	gameLogger.Info("X moves: %d", xMoves)
	gameLogger.Info("O moves: %d", oMoves)

	// Log strategic move counts
	forkCreations := 0
	forkBlocks := 0
	winningMoves := 0
	blockingMoves := 0

	for _, state := range record.States {
		if isForkCreation(state.Board, state.Move) {
			forkCreations++
		}
		if isForkBlocking(state.Board, state.Move) {
			forkBlocks++
		}
		if isWinningMove(state.Board, state.Move) {
			winningMoves++
		}
		if isBlockingMove(state.Board, state.Move) {
			blockingMoves++
		}
	}

	gameLogger.Info("\nStrategic Move Statistics:")
	gameLogger.Info("  Fork Creations: %d", forkCreations)
	gameLogger.Info("  Fork Blocks: %d", forkBlocks)
	gameLogger.Info("  Winning Moves: %d", winningMoves)
	gameLogger.Info("  Blocking Moves: %d", blockingMoves)
	gameLogger.Info("==========================================\n")
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
	// Use the display package to show progress
	display.DisplayProgressBar(gameNum, totalGames)
}

// Helper functions for board evaluation
func countEmptyCells(board *game.Board) int {
	count := 0
	for i := 0; i < 9; i++ {
		row, col := neural.MoveIndexToRowCol(i)
		if board.Get(row, col) == game.Empty {
			count++
		}
	}
	return count
}

func countPotentialWinningLines(board *game.Board) int {
	// This is a simplified version - you might want to implement a more sophisticated check
	count := 0
	// Check rows
	for i := 0; i < 3; i++ {
		if board.Get(i, 0) == game.Empty || board.Get(i, 1) == game.Empty || board.Get(i, 2) == game.Empty {
			count++
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if board.Get(0, i) == game.Empty || board.Get(1, i) == game.Empty || board.Get(2, i) == game.Empty {
			count++
		}
	}
	// Check diagonals
	if board.Get(0, 0) == game.Empty || board.Get(1, 1) == game.Empty || board.Get(2, 2) == game.Empty {
		count++
	}
	if board.Get(0, 2) == game.Empty || board.Get(1, 1) == game.Empty || board.Get(2, 0) == game.Empty {
		count++
	}
	return count
}

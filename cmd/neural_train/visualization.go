package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DisplayBoard displays the current board state
func DisplayBoard(board *game.Board) {
	fmt.Println("\nCurrent Board:")
	fmt.Print(board.String())
}

// DisplayMoveProbabilities displays the move probabilities
func DisplayMoveProbabilities(probabilities []float64, selectedMove int) {
	fmt.Println("\nMove Probabilities:")
	for i, prob := range probabilities {
		row, col := neural.MoveIndexToRowCol(i)
		if i == selectedMove {
			fmt.Printf("Position (%d,%d): %.2f%% [SELECTED]\n", row, col, prob*100)
		} else {
			fmt.Printf("Position (%d,%d): %.2f%%\n", row, col, prob*100)
		}
	}
}

// DisplayGameProgress displays the progress of the current game
func DisplayGameProgress(gameNum, totalGames int, moveNum int, currentPlayer string, epsilon float64) {
	fmt.Printf("\nGame %d/%d (Move %d) - Player: %s (Epsilon: %.2f)\n",
		gameNum+1, totalGames, moveNum+1, currentPlayer, epsilon)
}

// DisplayGameResult displays the result of a game
func DisplayGameResult(record GameRecord) {
	fmt.Println("\nGame Result:")
	if record.Winner == "X" {
		fmt.Println("Player X wins!")
	} else if record.Winner == "O" {
		fmt.Println("Player O wins!")
	} else {
		fmt.Println("It's a draw!")
	}
}

// DisplayTrainingProgress displays the overall training progress
func DisplayTrainingProgress(gameNum, totalGames int, xWins, oWins, draws int, epsilon float64) {
	winRate := float64(xWins+oWins) / float64(gameNum+1) * 100
	fmt.Printf("\nTraining Progress: %d/%d games (%.1f%% complete)\n",
		gameNum+1, totalGames, float64(gameNum+1)/float64(totalGames)*100)
	fmt.Printf("X Wins: %d, O Wins: %d, Draws: %d (Win Rate: %.1f%%)\n",
		xWins, oWins, draws, winRate)
	fmt.Printf("Current Epsilon: %.2f\n", epsilon)
}

// DisplayStrategyInfo displays information about the strategy used
func DisplayStrategyInfo(board *game.Board, move int) {
	// Check for fork creation
	if isForkCreation(board, move) {
		fmt.Println("\n[STRATEGY] Fork Creation Detected!")
	}

	// Check for fork blocking
	if isForkBlocking(board, move) {
		fmt.Println("\n[STRATEGY] Fork Blocking Detected!")
	}

	// Check for winning move
	if isWinningMove(board, move) {
		fmt.Println("\n[STRATEGY] Winning Move Detected!")
	}

	// Check for blocking opponent's winning move
	if isBlockingMove(board, move) {
		fmt.Println("\n[STRATEGY] Blocking Move Detected!")
	}
}

// VisualizeGame visualizes a complete game with the given delay
func VisualizeGame(record GameRecord, displayDelay time.Duration) {
	for i, state := range record.States {
		ClearScreen()

		// Display game progress
		fmt.Printf("\nGame Replay - Move %d/%d\n", i+1, len(record.States))

		// Display board
		DisplayBoard(state.Board)

		// Display move probabilities
		DisplayMoveProbabilities(state.Probabilities, state.Move)

		// Display strategy info
		DisplayStrategyInfo(state.Board, state.Move)

		// Wait for the specified delay
		time.Sleep(displayDelay)
	}

	// Display final result
	ClearScreen()
	DisplayBoard(record.States[len(record.States)-1].Board)
	DisplayGameResult(record)
}

// isForkCreation checks if a move creates a fork
func isForkCreation(board *game.Board, move int) bool {
	// Implementation will be added later
	return false
}

// isForkBlocking checks if a move blocks a fork
func isForkBlocking(board *game.Board, move int) bool {
	// Implementation will be added later
	return false
}

// isWinningMove checks if a move is a winning move
func isWinningMove(board *game.Board, move int) bool {
	// Implementation will be added later
	return false
}

// isBlockingMove checks if a move blocks the opponent's winning move
func isBlockingMove(board *game.Board, move int) bool {
	// Implementation will be added later
	return false
}

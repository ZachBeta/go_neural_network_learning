package main

import (
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
	gameLogger.Info("\nCurrent Board:")
	gameLogger.Info("%s", board.String())
}

// DisplayMoveProbabilities displays the move probabilities
func DisplayMoveProbabilities(probabilities []float64, selectedMove int) {
	gameLogger.Info("\nMove Probabilities:")
	for i, prob := range probabilities {
		row, col := neural.MoveIndexToRowCol(i)
		if i == selectedMove {
			gameLogger.Info("Position (%d,%d): %.2f%% [SELECTED]", row, col, prob*100)
		} else {
			gameLogger.Info("Position (%d,%d): %.2f%%", row, col, prob*100)
		}
	}
}

// DisplayGameProgress displays the progress of the current game
func DisplayGameProgress(gameNum, totalGames int, moveNum int, currentPlayer string, epsilon float64) {
	gameLogger.Info("\nGame %d/%d (Move %d) - Player: %s (Epsilon: %.2f)",
		gameNum+1, totalGames, moveNum+1, currentPlayer, epsilon)
}

// DisplayGameResult displays the result of a game
func DisplayGameResult(record GameRecord) {
	gameLogger.Info("\nGame Result:")
	if record.Winner == "X" {
		gameLogger.Info("Player X wins!")
	} else if record.Winner == "O" {
		gameLogger.Info("Player O wins!")
	} else {
		gameLogger.Info("It's a draw!")
	}
}

// DisplayTrainingProgress displays the overall training progress
func DisplayTrainingProgress(gameNum, totalGames int, xWins, oWins, draws int, epsilon float64) {
	winRate := float64(xWins+oWins) / float64(gameNum+1) * 100
	gameLogger.Info("\nTraining Progress: %d/%d games (%.1f%% complete)",
		gameNum+1, totalGames, float64(gameNum+1)/float64(totalGames)*100)
	gameLogger.Info("X Wins: %d, O Wins: %d, Draws: %d (Win Rate: %.1f%%)",
		xWins, oWins, draws, winRate)
	gameLogger.Info("Current Epsilon: %.2f", epsilon)
}

// DisplayStrategyInfo displays information about the strategy used
func DisplayStrategyInfo(board *game.Board, move int) {
	// Check for fork creation
	if isForkCreation(board, move) {
		gameLogger.Info("\n[STRATEGY] Fork Creation Detected!")
	}

	// Check for fork blocking
	if isForkBlocking(board, move) {
		gameLogger.Info("\n[STRATEGY] Fork Blocking Detected!")
	}

	// Check for winning move
	if isWinningMove(board, move) {
		gameLogger.Info("\n[STRATEGY] Winning Move Detected!")
	}

	// Check for blocking opponent's winning move
	if isBlockingMove(board, move) {
		gameLogger.Info("\n[STRATEGY] Blocking Move Detected!")
	}
}

// VisualizeGame visualizes a complete game with the given delay
func VisualizeGame(record GameRecord, displayDelay time.Duration) {
	for i, state := range record.States {
		ClearScreen()

		// Display game progress
		gameLogger.Info("\nGame Replay - Move %d/%d", i+1, len(record.States))

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
	// Make a temporary board to test the move
	tempBoard := board.Clone()
	row, col := neural.MoveIndexToRowCol(move)
	tempBoard.MakeMove(row, col)

	// Count winning lines after the move
	winningLines := 0
	for i := 0; i < 9; i++ {
		testRow, testCol := neural.MoveIndexToRowCol(i)
		if tempBoard.Get(testRow, testCol) == game.Empty {
			// Try this move
			testBoard := tempBoard.Clone()
			testBoard.MakeMove(testRow, testCol)
			testBoard.CheckWinner()
			if testBoard.GetStatus() == game.Won {
				winningLines++
			}
		}
	}

	// A fork creates at least 2 winning lines
	return winningLines >= 2
}

// isForkBlocking checks if a move blocks a fork
func isForkBlocking(board *game.Board, move int) bool {
	// Make a temporary board to test the move
	tempBoard := board.Clone()
	row, col := neural.MoveIndexToRowCol(move)
	tempBoard.MakeMove(row, col)

	// Count opponent's winning lines after our move
	winningLines := 0
	for i := 0; i < 9; i++ {
		testRow, testCol := neural.MoveIndexToRowCol(i)
		if tempBoard.Get(testRow, testCol) == game.Empty {
			// Try opponent's move
			testBoard := tempBoard.Clone()
			testBoard.SwitchPlayer()
			testBoard.MakeMove(testRow, testCol)
			testBoard.CheckWinner()
			if testBoard.GetStatus() == game.Won {
				winningLines++
			}
		}
	}

	// A fork block prevents opponent from having 2 winning lines
	return winningLines < 2
}

// isWinningMove checks if a move is a winning move
func isWinningMove(board *game.Board, move int) bool {
	// Make a temporary board to test the move
	tempBoard := board.Clone()
	row, col := neural.MoveIndexToRowCol(move)
	tempBoard.MakeMove(row, col)
	tempBoard.CheckWinner()

	// Check if this move wins
	return tempBoard.GetStatus() == game.Won
}

// isBlockingMove checks if a move blocks the opponent's winning move
func isBlockingMove(board *game.Board, move int) bool {
	// Make a temporary board to test the move
	tempBoard := board.Clone()
	row, col := neural.MoveIndexToRowCol(move)
	tempBoard.MakeMove(row, col)

	// Switch to opponent's turn
	tempBoard.SwitchPlayer()

	// Check if opponent can win after our move
	for i := 0; i < 9; i++ {
		testRow, testCol := neural.MoveIndexToRowCol(i)
		if tempBoard.Get(testRow, testCol) == game.Empty {
			// Try opponent's move
			testBoard := tempBoard.Clone()
			testBoard.MakeMove(testRow, testCol)
			testBoard.CheckWinner()
			if testBoard.GetStatus() == game.Won {
				return false // Opponent can still win
			}
		}
	}

	// If we get here, opponent cannot win after our move
	return true
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ZachBeta/go_neural_network_learning/internal/utils"
	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
)

// playerToString converts a Cell value to a string representation
func playerToString(player game.Cell) string {
	switch player {
	case game.X:
		return "X"
	case game.O:
		return "O"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("Enter 'q' to quit at any time.")

	// Set log level to INFO for a cleaner output
	utils.SetLogLevel(utils.INFO)

	// Start a new game
	board := game.NewBoard()

	// Game loop
	for {
		// Display the board
		fmt.Println("\nCurrent board:")
		fmt.Print(board.String())

		// Check if the game is over
		if board.GetStatus() != game.InProgress {
			fmt.Println("Game over!")
			if board.GetStatus() == game.Won {
				// Switch back to the previous player since they won
				board.SwitchPlayer()
				fmt.Printf("Player %s wins!\n", playerToString(board.GetCurrentPlayer()))
			} else {
				fmt.Println("It's a draw!")
			}
			break
		}

		// Get player input
		fmt.Printf("Player %s's turn. Enter row and column (0-2) separated by space: ",
			playerToString(board.GetCurrentPlayer()))

		// Read input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		utils.HandleError(err, false)
		if err != nil {
			continue
		}

		// Check for quit command
		input = strings.TrimSpace(input)
		if input == "q" {
			fmt.Println("Quitting game.")
			break
		}

		// Parse input
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter row and column (0-2) separated by space.")
			continue
		}

		// Convert to integers
		row, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid row. Please enter a number between 0 and 2.")
			continue
		}

		col, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid column. Please enter a number between 0 and 2.")
			continue
		}

		// Validate row and column
		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid position. Row and column must be between 0 and 2.")
			continue
		}

		// Make the move
		if !board.MakeMove(row, col) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		// Check for winner
		board.CheckWinner()
	}
}

package neural

import (
	"math"

	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
)

// BoardToInput converts a game board to a neural network input vector
// The input vector is a flattened representation of the board
// X = 1.0, O = -1.0, Empty = 0.0
func BoardToInput(board *game.Board) []float64 {
	input := make([]float64, 9)

	for i := 0; i < 9; i++ {
		row := i / 3
		col := i % 3

		switch board.Get(row, col) {
		case game.X:
			input[i] = 1.0
		case game.O:
			input[i] = -1.0
		case game.Empty:
			input[i] = 0.0
		}
	}

	return input
}

// OutputToMoveProbabilities converts neural network output to move probabilities
// The output is normalized to sum to 1.0 using softmax
func OutputToMoveProbabilities(output []float64) []float64 {
	// Apply softmax to convert to probabilities
	probabilities := make([]float64, len(output))

	// Find the maximum value for numerical stability
	maxVal := output[0]
	for _, val := range output {
		if val > maxVal {
			maxVal = val
		}
	}

	// Calculate sum of exponentials
	sum := 0.0
	for i, val := range output {
		probabilities[i] = math.Exp(val - maxVal)
		sum += probabilities[i]
	}

	// Normalize to get probabilities
	if sum > 0 {
		for i := range probabilities {
			probabilities[i] /= sum
		}
	}

	return probabilities
}

// SelectBestMove selects the best move based on move probabilities
// It returns the index of the best move (0-8)
func SelectBestMove(probabilities []float64) int {
	bestMove := 0
	bestProb := probabilities[0]

	for i, prob := range probabilities {
		if prob > bestProb {
			bestProb = prob
			bestMove = i
		}
	}

	return bestMove
}

// MoveIndexToRowCol converts a move index (0-8) to row and column coordinates
func MoveIndexToRowCol(moveIndex int) (row, col int) {
	row = moveIndex / 3
	col = moveIndex % 3
	return row, col
}

// RowColToMoveIndex converts row and column coordinates to a move index (0-8)
func RowColToMoveIndex(row, col int) int {
	return row*3 + col
}

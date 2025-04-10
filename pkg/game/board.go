package game

import (
	"github.com/ZachBeta/go_neural_network_learning/internal/utils"
)

// Cell represents a single cell in the board
type Cell int

const (
	Empty Cell = iota
	X
	O
)

// GameStatus represents the current status of the game
type GameStatus int

const (
	InProgress GameStatus = iota
	Won
	Draw
)

// Board represents the game board using a flat array
type Board struct {
	cells         [9]Cell
	currentPlayer Cell
	status        GameStatus
}

// NewBoard creates a new empty board
func NewBoard() *Board {
	utils.Info("Creating new game board")
	return &Board{
		cells:         [9]Cell{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		currentPlayer: X,
		status:        InProgress,
	}
}

// Get returns the cell value at the given position
func (b *Board) Get(row, col int) Cell {
	if !isValidPosition(row, col) {
		utils.Debug("Invalid position requested: row=%d, col=%d", row, col)
		return Empty
	}
	return b.cells[row*3+col]
}

// Set sets the cell value at the given position
func (b *Board) Set(row, col int, value Cell) bool {
	if !isValidPosition(row, col) {
		utils.Debug("Attempted to set invalid position: row=%d, col=%d", row, col)
		return false
	}
	b.cells[row*3+col] = value
	utils.Debug("Set position (%d,%d) to %v", row, col, value)
	return true
}

// isValidPosition checks if the given position is valid
func isValidPosition(row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3
}

// String returns a string representation of the board
func (b *Board) String() string {
	var result string
	for i := 0; i < 9; i++ {
		switch b.cells[i] {
		case Empty:
			result += "."
		case X:
			result += "X"
		case O:
			result += "O"
		}
		if (i+1)%3 == 0 {
			result += "\n"
		}
	}
	return result
}

// SwitchPlayer switches the current player
func (b *Board) SwitchPlayer() {
	if b.currentPlayer == X {
		b.currentPlayer = O
	} else {
		b.currentPlayer = X
	}
	utils.Debug("Switched current player to %v", b.currentPlayer)
}

// UpdateStatus updates the game status
func (b *Board) UpdateStatus(status GameStatus) {
	oldStatus := b.status
	b.status = status
	if oldStatus != status {
		utils.Info("Game status changed from %v to %v", oldStatus, status)
	}
}

// GetCurrentPlayer returns the current player
func (b *Board) GetCurrentPlayer() Cell {
	return b.currentPlayer
}

// GetStatus returns the current game status
func (b *Board) GetStatus() GameStatus {
	return b.status
}

// MakeMove attempts to make a move at the specified position
func (b *Board) MakeMove(row, col int) bool {
	if b.status != InProgress {
		utils.Debug("Move rejected: game is not in progress (status=%v)", b.status)
		return false
	}
	if !isValidPosition(row, col) || b.Get(row, col) != Empty {
		utils.Debug("Move rejected: invalid position or position already occupied (row=%d, col=%d)", row, col)
		return false
	}
	utils.Info("Player %v making move at position (%d,%d)", b.currentPlayer, row, col)
	b.Set(row, col, b.currentPlayer)
	b.SwitchPlayer()
	return true
}

// CheckWinner checks for a winner and updates the game status
func (b *Board) CheckWinner() {
	// Check rows
	for i := 0; i < 3; i++ {
		if b.Get(i, 0) != Empty && b.Get(i, 0) == b.Get(i, 1) && b.Get(i, 1) == b.Get(i, 2) {
			utils.Info("Winner detected: horizontal win in row %d", i)
			b.UpdateStatus(Won)
			return
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if b.Get(0, j) != Empty && b.Get(0, j) == b.Get(1, j) && b.Get(1, j) == b.Get(2, j) {
			utils.Info("Winner detected: vertical win in column %d", j)
			b.UpdateStatus(Won)
			return
		}
	}

	// Check diagonals
	if b.Get(0, 0) != Empty && b.Get(0, 0) == b.Get(1, 1) && b.Get(1, 1) == b.Get(2, 2) {
		utils.Info("Winner detected: diagonal win (top-left to bottom-right)")
		b.UpdateStatus(Won)
		return
	}
	if b.Get(0, 2) != Empty && b.Get(0, 2) == b.Get(1, 1) && b.Get(1, 1) == b.Get(2, 0) {
		utils.Info("Winner detected: diagonal win (top-right to bottom-left)")
		b.UpdateStatus(Won)
		return
	}

	// Check for draw
	isFull := true
	for i := 0; i < 9; i++ {
		if b.cells[i] == Empty {
			isFull = false
			break
		}
	}
	if isFull {
		utils.Info("Game ended in a draw")
		b.UpdateStatus(Draw)
	}
}

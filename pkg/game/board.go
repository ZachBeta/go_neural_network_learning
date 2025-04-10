package game

// Cell represents a single cell in the board
type Cell int

const (
	Empty Cell = iota
	X
	O
)

// Board represents the game board using a flat array
type Board struct {
	cells [9]Cell
}

// NewBoard creates a new empty board
func NewBoard() *Board {
	return &Board{
		cells: [9]Cell{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
	}
}

// Get returns the cell value at the given position
func (b *Board) Get(row, col int) Cell {
	if !isValidPosition(row, col) {
		return Empty
	}
	return b.cells[row*3+col]
}

// Set sets the cell value at the given position
func (b *Board) Set(row, col int, value Cell) bool {
	if !isValidPosition(row, col) {
		return false
	}
	b.cells[row*3+col] = value
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

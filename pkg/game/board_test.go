package game

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	for i := 0; i < 9; i++ {
		if board.cells[i] != Empty {
			t.Errorf("Expected empty cell at position %d, got %v", i, board.cells[i])
		}
	}
}

func TestGetSet(t *testing.T) {
	board := NewBoard()

	// Test valid positions
	if !board.Set(0, 0, X) {
		t.Error("Failed to set valid position")
	}
	if board.Get(0, 0) != X {
		t.Error("Failed to get correct value")
	}

	// Test invalid positions
	if board.Set(3, 0, X) {
		t.Error("Should not set invalid position")
	}
	if board.Get(3, 0) != Empty {
		t.Error("Should return Empty for invalid position")
	}
}

func TestString(t *testing.T) {
	board := NewBoard()
	board.Set(0, 0, X)
	board.Set(1, 1, O)
	board.Set(2, 2, X)

	expected := "X..\n.O.\n..X\n"
	if board.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, board.String())
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()

	// Test valid move
	if !board.MakeMove(0, 0) {
		t.Error("Failed to make valid move")
	}
	if board.Get(0, 0) != X {
		t.Error("Board not updated correctly after valid move")
	}
	if board.GetCurrentPlayer() != O {
		t.Error("Player not switched after valid move")
	}

	// Test invalid move
	if board.MakeMove(0, 0) {
		t.Error("Should not make move at occupied position")
	}
	if board.Get(0, 0) != X {
		t.Error("Board should not change after invalid move")
	}
}

func TestCheckWinner(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Board)
		expected GameStatus
	}{
		{
			name: "horizontal win - top row",
			setup: func(b *Board) {
				b.Set(0, 0, X)
				b.Set(0, 1, X)
				b.Set(0, 2, X)
			},
			expected: Won,
		},
		{
			name: "horizontal win - middle row",
			setup: func(b *Board) {
				b.Set(1, 0, O)
				b.Set(1, 1, O)
				b.Set(1, 2, O)
			},
			expected: Won,
		},
		{
			name: "horizontal win - bottom row",
			setup: func(b *Board) {
				b.Set(2, 0, X)
				b.Set(2, 1, X)
				b.Set(2, 2, X)
			},
			expected: Won,
		},
		{
			name: "vertical win - left column",
			setup: func(b *Board) {
				b.Set(0, 0, O)
				b.Set(1, 0, O)
				b.Set(2, 0, O)
			},
			expected: Won,
		},
		{
			name: "vertical win - middle column",
			setup: func(b *Board) {
				b.Set(0, 1, X)
				b.Set(1, 1, X)
				b.Set(2, 1, X)
			},
			expected: Won,
		},
		{
			name: "vertical win - right column",
			setup: func(b *Board) {
				b.Set(0, 2, O)
				b.Set(1, 2, O)
				b.Set(2, 2, O)
			},
			expected: Won,
		},
		{
			name: "diagonal win - top-left to bottom-right",
			setup: func(b *Board) {
				b.Set(0, 0, X)
				b.Set(1, 1, X)
				b.Set(2, 2, X)
			},
			expected: Won,
		},
		{
			name: "diagonal win - top-right to bottom-left",
			setup: func(b *Board) {
				b.Set(0, 2, O)
				b.Set(1, 1, O)
				b.Set(2, 0, O)
			},
			expected: Won,
		},
		{
			name: "draw - full board",
			setup: func(b *Board) {
				b.Set(0, 0, X)
				b.Set(0, 1, O)
				b.Set(0, 2, X)
				b.Set(1, 0, O)
				b.Set(1, 1, X)
				b.Set(1, 2, O)
				b.Set(2, 0, O)
				b.Set(2, 1, X)
				b.Set(2, 2, O)
			},
			expected: Draw,
		},
		{
			name:     "no win - empty board",
			setup:    func(b *Board) {},
			expected: InProgress,
		},
		{
			name: "no win - partial board",
			setup: func(b *Board) {
				b.Set(0, 0, X)
				b.Set(1, 1, O)
			},
			expected: InProgress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := NewBoard()
			tt.setup(board)
			board.CheckWinner()
			if board.GetStatus() != tt.expected {
				t.Errorf("expected status %v, got %v", tt.expected, board.GetStatus())
			}
		})
	}
}

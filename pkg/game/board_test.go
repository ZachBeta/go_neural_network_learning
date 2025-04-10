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
	// This is a placeholder test that will be implemented
	// when we create the move implementation
	t.Run("should make valid move", func(t *testing.T) {
		// TODO: Implement test
	})
}

func TestCheckWinner(t *testing.T) {
	// This is a placeholder test that will be implemented
	// when we create the win condition checking
	t.Run("should detect horizontal win", func(t *testing.T) {
		// TODO: Implement test
	})
}

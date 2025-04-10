# Neural Network Package Tutorial

This tutorial explains the implementation of a simple neural network for the Tic-tac-toe game. The network is designed to learn and play the game by understanding board positions and making strategic moves.

## Overview

The neural network package implements a single-layer perceptron, which is the simplest form of a neural network. It's perfect for learning Tic-tac-toe because:
- The game has a small, fixed input size (9 board positions)
- The output is a set of probabilities for each possible move
- The game rules are deterministic and well-defined

## Components

### 1. Activation Function (`activation.go`)

The activation function determines how a neuron processes its input. We use the Sigmoid function:

```go
type ActivationFunction interface {
    Activate(x float64) float64
    Derivative(x float64) float64
    Name() string
}
```

The Sigmoid function maps any input to a value between 0 and 1:
- f(x) = 1 / (1 + e^(-x))
- Perfect for producing move probabilities
- Smooth gradient for learning

### 2. Neuron (`neuron.go`)

A neuron is the basic processing unit:

```go
type Neuron struct {
    Weights    []float64
    Bias       float64
    Activation ActivationFunction
}
```

Key features:
- Weights: Connection strengths to inputs
- Bias: Threshold for activation
- Forward pass: 
  1. Calculate weighted sum: sum = bias + Σ(weight_i * input_i)
  2. Apply activation function
- Xavier/Glorot initialization for better training

### 3. Layer (`layer.go`)

A layer contains multiple neurons that process inputs in parallel:

```go
type Layer struct {
    Neurons []*Neuron
    Output  []float64
}
```

Features:
- Multiple neurons process the same input
- Each neuron can learn different patterns
- Output is stored for backpropagation

### 4. Network (`network.go`)

The network combines layers to process inputs:

```go
type Network struct {
    InputLayer  *Layer
    OutputLayer *Layer
}
```

Current implementation:
- Single-layer perceptron
- Input layer (pass-through)
- Output layer with sigmoid activation
- Produces move probabilities

## Usage Example

```go
// Create a network for Tic-tac-toe
network := neural.NewNetwork(9, 9)  // 9 inputs (board), 9 outputs (moves)

// Convert board state to input using the game integration
gameBoard := game.NewBoard()  // Create a new game board
gameBoard.Set(0, 0, game.X)   // Set X at position (0,0)
gameBoard.Set(2, 0, game.O)   // Set O at position (2,0)
input := neural.BoardToInput(gameBoard)  // Convert to neural network input

// Get move probabilities
probabilities := network.Forward(input)

// Convert probabilities to move
moveIndex := neural.SelectBestMove(probabilities)
row, col := neural.MoveIndexToRowCol(moveIndex)
```

## Board State Handling

The neural network package provides functions to convert between game board states and neural network inputs:

```go
// In neural/game_integration.go

// BoardToInput converts a game board to a neural network input vector
// X = 1.0, O = -1.0, Empty = 0.0
func BoardToInput(board *game.Board) []float64 {
    // Implementation details...
}

// OutputToMoveProbabilities converts neural network output to move probabilities
func OutputToMoveProbabilities(output []float64) []float64 {
    // Implementation details...
}

// SelectBestMove selects the best move based on move probabilities
func SelectBestMove(probabilities []float64) int {
    // Implementation details...
}

// MoveIndexToRowCol converts a move index (0-8) to row and column coordinates
func MoveIndexToRowCol(moveIndex int) (row, col int) {
    // Implementation details...
}

// RowColToMoveIndex converts row and column coordinates to a move index (0-8)
func RowColToMoveIndex(row, col int) int {
    // Implementation details...
}
```

Benefits of this approach:
1. Direct integration with the game package
2. Type safety through the game.Board type
3. Consistent mapping between game state and neural network input
4. Helper functions for move selection and conversion
5. Clear documentation of the expected values

## Training Process

The network learns through:
1. Forward pass: Process board state to get move probabilities
2. Move selection: Choose move based on probabilities
3. Game outcome: Get reward (win/loss/draw)
4. Backpropagation: Update weights based on outcome
5. Repeat: Play many games to improve

## Next Steps

1. Implement backpropagation for learning
2. Add hidden layers for more complex patterns
3. Implement experience replay
4. Add exploration vs exploitation
5. Create training pipeline

## Testing

The package includes comprehensive tests:
- Activation function behavior
- Neuron forward pass
- Layer processing
- Network output
- Move probability conversion

Run tests with:
```bash
go test ./pkg/neural/...
```

## Best Practices

1. Always initialize the network with proper sizes
2. Use the sigmoid activation for move probabilities
3. Normalize board state inputs (-1, 0, 1)
4. Handle invalid moves gracefully
5. Log training progress

## Common Patterns

### Board State Encoding
```go
// Using the game integration
gameBoard := game.NewBoard()
// Set up the board state
gameBoard.Set(0, 0, game.X)
gameBoard.Set(1, 1, game.O)
// Convert to neural network input
input := neural.BoardToInput(gameBoard)
```

### Move Selection
```go
// Get move probabilities
probs := network.Forward(input)

// Convert to move
moveIndex := neural.SelectBestMove(probs)
row, col := neural.MoveIndexToRowCol(moveIndex)

// Make the move
gameBoard.Set(row, col, game.X)  // or game.O depending on the current player
```

## Troubleshooting

Common issues and solutions:
1. Network produces same move repeatedly
   - Check initialization
   - Verify input normalization
   - Adjust learning rate

2. Poor learning performance
   - Increase training games
   - Adjust network architecture
   - Check reward function

3. Invalid moves
   - Validate network output
   - Implement fallback strategy
   - Log problematic states 
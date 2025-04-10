# Neural Network Package

This package implements a neural network from scratch in Go, starting with a single-layer perceptron and evolving to more complex architectures. The implementation is designed to learn and play Tic-Tac-Toe.

## Package Structure

- `activation.go`: Contains activation function interfaces and implementations
- `neuron.go`: Implements the basic neuron structure
- `layer.go`: Implements a layer of neurons
- `network.go`: Implements a simple neural network
- `game_integration.go`: Contains functions to convert between game states and neural network inputs/outputs
- `utils.go`: Contains utility functions for the neural network

## Components

### Activation Functions
- Sigmoid: Maps any input to a value between 0 and 1
- (Future) Tanh: Maps any input to a value between -1 and 1
- (Future) ReLU: Returns the input if positive, 0 otherwise

### Neuron
The basic building block of the neural network, consisting of:
- Weights: A slice of float64 values, one for each input
- Bias: A single float64 value to shift the activation function
- Activation function: A function to transform the weighted sum

### Layer
A collection of neurons that process the same inputs and produce a vector of outputs.

### Network
A simple neural network with input and output layers.

## Usage

```go
// Create a new neural network
network := neural.NewNetwork(9, 9) // 9 inputs (board state), 9 outputs (move probabilities)

// Convert game board to neural network input
input := neural.BoardToInput(board)

// Get move probabilities from the network
output := network.Forward(input)
moveProbs := neural.OutputToMoveProbabilities(output)

// Select the best move
bestMove := neural.SelectBestMove(moveProbs)
```

## Future Enhancements
- Multi-layer neural network
- Advanced activation functions
- Convolutional layers for pattern recognition
- Recurrent layers for sequential decision making 
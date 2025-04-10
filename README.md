# Go Neural Network Learning

A tutorial project implementing neural networks from scratch in Go, using Tic-Tac-Toe as a learning problem. This project demonstrates the implementation of both single-layer perceptrons and multi-layer neural networks, leveraging Go's concurrency features for optimized training on Apple M1 hardware.

## Project Structure

```
.
├── cmd/
│   └── tictactoe/     # Main application entry point
├── pkg/
│   ├── game/          # Game logic and board representation
│   ├── network/       # Neural network implementation
│   └── training/      # Training logic and batch processing
└── internal/
    └── utils/         # Internal utilities and helpers
```

## Features

- Single-layer perceptron implementation
- Multi-layer neural network evolution
- Concurrent training optimization
- Tic-Tac-Toe game implementation
- CLI interface for interaction
- Training visualization and metrics

## Requirements

- Go 1.21 or later
- Apple M1/M2 hardware (for optimal performance)

## Getting Started

1. Clone the repository:
   ```bash
   git clone git@github.com:ZachBeta/go_neural_network_learning.git
   cd go_neural_network_learning
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run cmd/tictactoe/main.go
   ```

## Development

This project follows a phase-based development approach. See `PHASES.md` for detailed information about the implementation phases and progress.

## License

MIT License

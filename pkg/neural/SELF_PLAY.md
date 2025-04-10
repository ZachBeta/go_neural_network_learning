# Self-Play Training Tutorial

This tutorial explains how to implement self-play training for the Tic-tac-toe neural network, inspired by AlphaZero's approach.

## Overview

Self-play training allows the neural network to learn by playing against itself, discovering strategies through exploration and reinforcement learning. This approach has several advantages:

1. **Continuous Learning**: The network can train 24/7 without human intervention
2. **Strategy Discovery**: Can discover novel strategies through exploration
3. **Balanced Training**: Both players (X and O) learn equally
4. **Scalable**: Can generate vast amounts of training data

## Implementation

### 1. Game State Recording

```go
// GameState represents a single state in a game
type GameState struct {
    Board      *game.Board
    Move       int
    Result     float64  // 1.0 for win, -1.0 for loss, 0.0 for draw
    Player     string   // "X" or "O"
}

// GameRecord represents a complete game
type GameRecord struct {
    States []GameState
    Winner string
}
```

### 2. Move Selection with Exploration

```go
// SelectMove chooses a move using epsilon-greedy exploration
func SelectMove(network *Network, board *game.Board, epsilon float64) int {
    if rand.Float64() < epsilon {
        // Exploration: choose random valid move
        return selectRandomValidMove(board)
    }
    
    // Exploitation: use network's prediction
    input := BoardToInput(board)
    probabilities := network.Forward(input)
    return SelectBestMove(probabilities)
}
```

### 3. Training Loop

```go
// TrainSelfPlay runs the self-play training loop
func TrainSelfPlay(network *Network, numGames int) {
    for game := 0; game < numGames; game++ {
        // Play a complete game
        record := playGame(network)
        
        // Convert game record to training data
        trainingData := convertToTrainingData(record)
        
        // Update network weights
        network.Update(trainingData)
        
        // Log progress
        logProgress(game, record)
    }
}
```

### 4. Experience Replay

```go
// ExperienceBuffer stores game states for batch training
type ExperienceBuffer struct {
    states  []GameState
    maxSize int
}

// Add adds a new game state to the buffer
func (b *ExperienceBuffer) Add(state GameState) {
    if len(b.states) >= b.maxSize {
        // Remove oldest state
        b.states = b.states[1:]
    }
    b.states = append(b.states, state)
}

// Sample returns a random batch of states
func (b *ExperienceBuffer) Sample(batchSize int) []GameState {
    // Implementation details...
}
```

## Training Process

1. **Initial Phase (Random Play)**
   - High exploration rate (epsilon = 0.9)
   - Mostly random moves
   - Building initial experience buffer

2. **Development Phase**
   - Decreasing exploration rate
   - Network starts recognizing patterns
   - Basic strategies emerge

3. **Refinement Phase**
   - Low exploration rate (epsilon = 0.1)
   - Strong strategy development
   - Pattern recognition

## Monitoring Progress

### 1. Win/Loss Statistics
```go
type TrainingStats struct {
    GamesPlayed    int
    XWins          int
    OWins          int
    Draws          int
    AverageGameLength float64
}
```

### 2. Move Distribution Analysis
```go
// AnalyzeMoveDistribution tracks move preferences
type MoveDistribution struct {
    Moves          [9]int
    TotalMoves     int
    WinningMoves   [9]int
    TotalWins      int
}
```

### 3. Strategy Development Tracking
```go
// Track emerging strategies
type StrategyTracker struct {
    ForkPatterns    int
    BlockingMoves   int
    WinningPatterns int
}
```

## Example Training Session

```go
func main() {
    // Initialize network
    network := NewNetwork(9, 9)
    
    // Create experience buffer
    buffer := NewExperienceBuffer(10000)
    
    // Training parameters
    numGames := 1000
    batchSize := 32
    learningRate := 0.01
    
    // Start training
    for epoch := 0; epoch < numGames; epoch++ {
        // Calculate exploration rate
        epsilon := math.Max(0.1, 1.0 - float64(epoch)/float64(numGames))
        
        // Play game and collect experience
        record := playGame(network, epsilon)
        buffer.Add(record)
        
        // Sample batch and update network
        batch := buffer.Sample(batchSize)
        network.Update(batch, learningRate)
        
        // Log progress
        logProgress(epoch, record)
    }
}
```

## Best Practices

1. **Exploration vs Exploitation**
   - Start with high exploration (90%)
   - Gradually decrease to 10%
   - Use temperature parameter for move selection

2. **Experience Replay**
   - Maintain large buffer of past games
   - Sample randomly for batch training
   - Prioritize recent experiences

3. **Network Updates**
   - Use small batch sizes initially
   - Increase batch size as training progresses
   - Implement learning rate decay

4. **Monitoring**
   - Track win/loss ratios
   - Monitor move distributions
   - Log emerging strategies
   - Visualize network behavior

## Common Issues and Solutions

1. **Network Stagnation**
   - Increase exploration rate
   - Add noise to move selection
   - Implement curriculum learning

2. **Poor Strategy Development**
   - Increase experience buffer size
   - Adjust reward function
   - Implement self-play temperature

3. **Training Instability**
   - Reduce learning rate
   - Implement gradient clipping
   - Use experience replay

## Next Steps

1. Implement advanced strategies recognition
2. Add visualization tools
3. Optimize training parameters
4. Implement parallel self-play
5. Add strategy analysis tools 
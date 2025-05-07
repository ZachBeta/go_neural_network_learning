package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/display"
	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

// TrainingStats tracks training statistics
type TrainingStats struct {
	XWins         int
	OWins         int
	Draws         int
	MoveCounts    [9]int
	ForkCreates   int
	ForkBlocks    int
	WinningMoves  int
	BlockingMoves int
	LastSaveTime  time.Time
}

// TrainingParams holds the parameters for self-play training
type TrainingParams struct {
	NumGames      int
	BatchSize     int
	LearningRate  float64
	EpsilonStart  float64
	EpsilonEnd    float64
	EpsilonDecay  float64
	DisplayDelay  time.Duration
	SaveInterval  int
	MaxBufferSize int
	LogInterval   int
}

// GameState represents a single state in a game
type GameState struct {
	Board         *game.Board
	Move          int
	Result        float64 // 1.0 for win, -1.0 for loss, 0.0 for draw
	Player        string  // "X" or "O"
	Probabilities []float64
}

// GameRecord represents a complete game
type GameRecord struct {
	States []GameState
	Winner string
}

// ExperienceBuffer stores game states for batch training
type ExperienceBuffer struct {
	states  []GameState
	maxSize int
}

// NewExperienceBuffer creates a new experience buffer
func NewExperienceBuffer(maxSize int) *ExperienceBuffer {
	return &ExperienceBuffer{
		states:  make([]GameState, 0, maxSize),
		maxSize: maxSize,
	}
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
	if len(b.states) < batchSize {
		return b.states
	}

	// Create a copy of the buffer to avoid modifying the original
	buffer := make([]GameState, len(b.states))
	copy(buffer, b.states)

	// Shuffle the buffer
	for i := len(buffer) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		buffer[i], buffer[j] = buffer[j], buffer[i]
	}

	// Return the first batchSize elements
	return buffer[:batchSize]
}

// Size returns the current size of the buffer
func (b *ExperienceBuffer) Size() int {
	return len(b.states)
}

// DefaultTrainingParams returns the default training parameters
func DefaultTrainingParams() TrainingParams {
	return TrainingParams{
		NumGames:      1000,
		BatchSize:     32,
		LearningRate:  0.01,
		EpsilonStart:  0.9,
		EpsilonEnd:    0.1,
		EpsilonDecay:  0.995,
		DisplayDelay:  500 * time.Millisecond,
		SaveInterval:  100,
		MaxBufferSize: 10000,
		LogInterval:   10, // Log every 10 games
	}
}

func main() {
	// Set random seed for reproducibility
	rand.Seed(time.Now().UnixNano())

	// Get training parameters
	params := DefaultTrainingParams()

	// Create a new neural network
	network := neural.NewNetwork(9, 9)

	// Create experience buffer
	buffer := NewExperienceBuffer(params.MaxBufferSize)

	// Initialize training stats
	stats := &TrainingStats{
		LastSaveTime: time.Now(),
	}

	// Create a channel for handling interrupts
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Print initial information
	fmt.Println("Neural Network Self-Play Training")
	fmt.Println("Detailed logs will be written to logs/training_*.log")
	fmt.Println("Press Ctrl+C to stop training")
	fmt.Println()

	// Start a goroutine to handle user input
	go handleUserInput(interrupt)

	// Training loop
	for gameNum := 0; gameNum < params.NumGames; gameNum++ {
		// Calculate exploration rate
		epsilon := math.Max(params.EpsilonEnd, params.EpsilonStart*math.Pow(params.EpsilonDecay, float64(gameNum)))

		// Play a game and collect experience
		record := playGameWithVisualization(network, epsilon, params.DisplayDelay)

		// Update statistics
		updateStats(stats, record)

		// Add game states to buffer
		for _, state := range record.States {
			buffer.Add(state)
		}

		// Sample batch and update network
		if buffer.Size() >= params.BatchSize {
			batch := buffer.Sample(params.BatchSize)
			updateNetworkWeights(network, batch, params.LearningRate)
		}

		// Display progress
		displayTrainingProgress(gameNum, params.NumGames, record, epsilon)

		// Log detailed statistics periodically
		if gameNum > 0 && gameNum%params.LogInterval == 0 {
			// Move to next line after progress bar
			fmt.Println()
			logDetailedStats(gameNum, stats, epsilon)
		}

		// Display statistics every 100 games
		if gameNum > 0 && gameNum%100 == 0 {
			display.DisplayStatistics(gameNum, params.NumGames, stats.XWins, stats.OWins, stats.Draws, epsilon)
		}

		// Save network periodically
		if gameNum > 0 && gameNum%params.SaveInterval == 0 {
			saveNetwork(network, gameNum)
			stats.LastSaveTime = time.Now()
		}

		// Check for interrupt
		select {
		case <-interrupt:
			fmt.Println("\nTraining interrupted. Saving network...")
			saveNetwork(network, gameNum)
			logDetailedStats(gameNum, stats, epsilon)
			return
		default:
			// Continue training
		}
	}

	// Training completed
	fmt.Println("\nTraining completed!")
	logDetailedStats(params.NumGames-1, stats, params.EpsilonEnd)

	// Save the final network
	saveNetwork(network, params.NumGames-1)
}

// handleUserInput handles user input during training
func handleUserInput(interrupt chan<- os.Signal) {
	// Implementation will be added later
}

// updateStats updates the training statistics
func updateStats(stats *TrainingStats, record GameRecord) {
	// Update win/draw counts
	if record.Winner == "X" {
		stats.XWins++
	} else if record.Winner == "O" {
		stats.OWins++
	} else {
		stats.Draws++
	}

	// Update move counts and strategy detection
	for _, state := range record.States {
		stats.MoveCounts[state.Move]++

		// Check for strategic moves
		if isForkCreation(state.Board, state.Move) {
			stats.ForkCreates++
		}
		if isForkBlocking(state.Board, state.Move) {
			stats.ForkBlocks++
		}
		if isWinningMove(state.Board, state.Move) {
			stats.WinningMoves++
		}
		if isBlockingMove(state.Board, state.Move) {
			stats.BlockingMoves++
		}
	}
}

// logDetailedStats logs detailed training statistics
func logDetailedStats(gameNum int, stats *TrainingStats, epsilon float64) {
	// Calculate percentages
	xWinRate := float64(stats.XWins) / float64(gameNum+1) * 100
	oWinRate := float64(stats.OWins) / float64(gameNum+1) * 100
	drawRate := float64(stats.Draws) / float64(gameNum+1) * 100

	// Log detailed statistics
	gameLogger.Info("\n=== Detailed Training Statistics (Game %d) ===", gameNum+1)
	gameLogger.Info("Win Rates:")
	gameLogger.Info("  X: %.1f%% (%d wins)", xWinRate, stats.XWins)
	gameLogger.Info("  O: %.1f%% (%d wins)", oWinRate, stats.OWins)
	gameLogger.Info("  Draws: %.1f%% (%d draws)", drawRate, stats.Draws)

	// Log move distribution
	gameLogger.Info("\nMove Distribution:")
	for i, count := range stats.MoveCounts {
		row, col := neural.MoveIndexToRowCol(i)
		percentage := float64(count) / float64(gameNum+1) * 100
		gameLogger.Info("  Position (%d,%d): %.1f%% (%d moves)", row, col, percentage, count)
	}

	// Log strategic moves
	gameLogger.Info("\nStrategic Moves:")
	gameLogger.Info("  Fork Creations: %d (%.1f%% of games)",
		stats.ForkCreates, float64(stats.ForkCreates)/float64(gameNum+1)*100)
	gameLogger.Info("  Fork Blocks: %d (%.1f%% of games)",
		stats.ForkBlocks, float64(stats.ForkBlocks)/float64(gameNum+1)*100)
	gameLogger.Info("  Winning Moves: %d (%.1f%% of games)",
		stats.WinningMoves, float64(stats.WinningMoves)/float64(gameNum+1)*100)
	gameLogger.Info("  Blocking Moves: %d (%.1f%% of games)",
		stats.BlockingMoves, float64(stats.BlockingMoves)/float64(gameNum+1)*100)

	// Log training parameters
	gameLogger.Info("\nTraining Parameters:")
	gameLogger.Info("  Epsilon: %.3f", epsilon)
	gameLogger.Info("  Time since last save: %v", time.Since(stats.LastSaveTime))
	gameLogger.Info("==========================================\n")
}

// saveNetwork saves the network to a file
func saveNetwork(network *neural.Network, gameNum int) {
	// Implementation will be added later
}

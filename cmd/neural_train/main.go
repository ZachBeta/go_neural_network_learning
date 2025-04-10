package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZachBeta/go_neural_network_learning/pkg/game"
	"github.com/ZachBeta/go_neural_network_learning/pkg/neural"
)

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

	// Create a channel for handling interrupts
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Print initial information
	fmt.Println("Neural Network Self-Play Training")
	fmt.Println("Press Ctrl+C to stop training")
	fmt.Println("Press 'p' to pause/resume")
	fmt.Println("Press 's' to save the current network")
	fmt.Println("Press 'h' to display help")
	fmt.Println()

	// Start a goroutine to handle user input
	go handleUserInput(interrupt)

	// Training loop
	for gameNum := 0; gameNum < params.NumGames; gameNum++ {
		// Calculate exploration rate
		epsilon := math.Max(params.EpsilonEnd, params.EpsilonStart*math.Pow(params.EpsilonDecay, float64(gameNum)))

		// Play a game and collect experience
		record := playGame(network, epsilon, params.DisplayDelay)

		// Add game states to buffer
		for _, state := range record.States {
			buffer.Add(state)
		}

		// Sample batch and update network
		if buffer.Size() >= params.BatchSize {
			batch := buffer.Sample(params.BatchSize)
			updateNetwork(network, batch, params.LearningRate)
		}

		// Display progress
		displayProgress(gameNum, params.NumGames, record, epsilon)

		// Save network periodically
		if gameNum > 0 && gameNum%params.SaveInterval == 0 {
			saveNetwork(network, gameNum)
		}

		// Check for interrupt
		select {
		case <-interrupt:
			fmt.Println("\nTraining interrupted. Saving network...")
			saveNetwork(network, gameNum)
			return
		default:
			// Continue training
		}
	}

	fmt.Println("\nTraining completed!")
}

// handleUserInput handles user input during training
func handleUserInput(interrupt chan<- os.Signal) {
	// Implementation will be added later
}

// playGame plays a complete game and returns the game record
func playGame(network *neural.Network, epsilon float64, displayDelay time.Duration) GameRecord {
	// Implementation will be added later
	return GameRecord{}
}

// updateNetwork updates the network weights based on a batch of game states
func updateNetwork(network *neural.Network, batch []GameState, learningRate float64) {
	// Implementation will be added later
}

// displayProgress displays the training progress
func displayProgress(gameNum, totalGames int, record GameRecord, epsilon float64) {
	// Implementation will be added later
}

// saveNetwork saves the network to a file
func saveNetwork(network *neural.Network, gameNum int) {
	// Implementation will be added later
}

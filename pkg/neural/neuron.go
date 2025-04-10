package neural

import (
	"math"
	"math/rand"
)

// Neuron represents a single neuron in a neural network
type Neuron struct {
	// Weights are the connection strengths to the inputs
	Weights []float64

	// Bias is the threshold for activation
	Bias float64

	// Activation is the activation function used by the neuron
	Activation ActivationFunction
}

// NewNeuron creates a new neuron with the specified number of inputs and activation function
func NewNeuron(inputSize int, activation ActivationFunction) *Neuron {
	if activation == nil {
		activation = &Sigmoid{}
	}

	neuron := &Neuron{
		Weights:    make([]float64, inputSize),
		Bias:       0.0,
		Activation: activation,
	}

	// Initialize weights with Xavier/Glorot initialization
	neuron.InitializeWeights()

	return neuron
}

// InitializeWeights initializes the weights using Xavier/Glorot initialization
func (n *Neuron) InitializeWeights() {
	// Xavier/Glorot initialization
	// Scale factor depends on the number of inputs
	scale := math.Sqrt(2.0 / float64(len(n.Weights)))

	for i := range n.Weights {
		// Initialize with random values scaled by the scale factor
		n.Weights[i] = rand.NormFloat64() * scale
	}

	// Initialize bias to a small random value
	n.Bias = rand.NormFloat64() * 0.01
}

// Forward performs a forward pass through the neuron
// It calculates the weighted sum of inputs and applies the activation function
func (n *Neuron) Forward(input []float64) float64 {
	if len(input) != len(n.Weights) {
		// Handle error: input size doesn't match weight size
		return 0.0
	}

	// Calculate weighted sum
	sum := n.Bias
	for i, weight := range n.Weights {
		sum += weight * input[i]
	}

	// Apply activation function
	return n.Activation.Activate(sum)
}

// GetWeights returns a copy of the neuron's weights
func (n *Neuron) GetWeights() []float64 {
	weights := make([]float64, len(n.Weights))
	copy(weights, n.Weights)
	return weights
}

// SetWeights sets the neuron's weights
func (n *Neuron) SetWeights(weights []float64) {
	if len(weights) != len(n.Weights) {
		// Handle error: weight size doesn't match
		return
	}

	copy(n.Weights, weights)
}

// GetBias returns the neuron's bias
func (n *Neuron) GetBias() float64 {
	return n.Bias
}

// SetBias sets the neuron's bias
func (n *Neuron) SetBias(bias float64) {
	n.Bias = bias
}

package neural

import (
	"math"
	"math/rand"

	"github.com/ZachBeta/go_neural_network_learning/pkg/logger"
)

// SetRandomSeed sets the random seed for reproducibility
func SetRandomSeed(seed int64) {
	rand.Seed(seed)
}

// PrintWeights prints the weights of a neuron
func PrintWeights(neuron *Neuron) {
	logger.Info("Weights: [")
	for i, weight := range neuron.Weights {
		if i > 0 {
			logger.Info(", ")
		}
		logger.Info("%.4f", weight)
	}
	logger.Info("], Bias: %.4f", neuron.Bias)
}

// PrintLayer prints the weights of all neurons in a layer
func PrintLayer(layer *Layer) {
	logger.Info("Layer with %d neurons:", len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		logger.Info("Neuron %d: ", i)
		PrintWeights(neuron)
	}
}

// PrintNetwork prints the structure of a network
func PrintNetwork(network *Network) {
	logger.Info("Network Structure:")
	logger.Info("Output Layer: %d neurons", network.OutputLayer.GetNeuronCount())
	PrintLayer(network.OutputLayer)
}

// PrintOutput prints the output of a layer
func PrintOutput(output []float64) {
	logger.Info("Output: [")
	for i, val := range output {
		if i > 0 {
			logger.Info(", ")
		}
		logger.Info("%.4f", val)
	}
	logger.Info("]")
}

// PrintMoveProbabilities prints the move probabilities
func PrintMoveProbabilities(probabilities []float64) {
	logger.Info("Move Probabilities:")
	for i, prob := range probabilities {
		row, col := MoveIndexToRowCol(i)
		logger.Info("(%d,%d): %.4f", row, col, prob)
	}
}

// CalculateMSE calculates the mean squared error between predicted and target values
func CalculateMSE(predicted, target []float64) float64 {
	if len(predicted) != len(target) {
		return math.NaN()
	}

	sum := 0.0
	for i, pred := range predicted {
		diff := pred - target[i]
		sum += diff * diff
	}

	return sum / float64(len(predicted))
}

// CalculateCrossEntropy calculates the cross-entropy loss between predicted and target values
func CalculateCrossEntropy(predicted, target []float64) float64 {
	if len(predicted) != len(target) {
		return math.NaN()
	}

	sum := 0.0
	for i, pred := range predicted {
		// Avoid log(0)
		if pred < 1e-10 {
			pred = 1e-10
		}
		if pred > 1.0-1e-10 {
			pred = 1.0 - 1e-10
		}

		sum += target[i] * math.Log(pred)
	}

	return -sum
}

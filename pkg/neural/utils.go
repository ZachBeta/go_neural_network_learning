package neural

import (
	"fmt"
	"math"
	"math/rand"
)

// SetRandomSeed sets the random seed for reproducibility
func SetRandomSeed(seed int64) {
	rand.Seed(seed)
}

// PrintWeights prints the weights of a neuron
func PrintWeights(neuron *Neuron) {
	fmt.Printf("Weights: [")
	for i, weight := range neuron.Weights {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%.4f", weight)
	}
	fmt.Printf("], Bias: %.4f\n", neuron.Bias)
}

// PrintLayer prints the weights of all neurons in a layer
func PrintLayer(layer *Layer) {
	fmt.Printf("Layer with %d neurons:\n", len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		fmt.Printf("Neuron %d: ", i)
		PrintWeights(neuron)
	}
}

// PrintNetwork prints the structure of a network
func PrintNetwork(network *Network) {
	fmt.Println("Network Structure:")
	fmt.Printf("Output Layer: %d neurons\n", network.OutputLayer.GetNeuronCount())
	PrintLayer(network.OutputLayer)
}

// PrintOutput prints the output of a layer
func PrintOutput(output []float64) {
	fmt.Printf("Output: [")
	for i, val := range output {
		if i > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%.4f", val)
	}
	fmt.Printf("]\n")
}

// PrintMoveProbabilities prints the move probabilities
func PrintMoveProbabilities(probabilities []float64) {
	fmt.Println("Move Probabilities:")
	for i, prob := range probabilities {
		row, col := MoveIndexToRowCol(i)
		fmt.Printf("(%d,%d): %.4f\n", row, col, prob)
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

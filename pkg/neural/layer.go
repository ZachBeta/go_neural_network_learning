package neural

// Layer represents a layer of neurons in a neural network
type Layer struct {
	// Neurons is a slice of neurons in the layer
	Neurons []*Neuron

	// Output is the output of the layer after a forward pass
	Output []float64
}

// NewLayer creates a new layer with the specified number of neurons and inputs
func NewLayer(neuronCount, inputSize int, activation ActivationFunction) *Layer {
	layer := &Layer{
		Neurons: make([]*Neuron, neuronCount),
		Output:  make([]float64, neuronCount),
	}

	// Create neurons
	for i := 0; i < neuronCount; i++ {
		layer.Neurons[i] = NewNeuron(inputSize, activation)
	}

	return layer
}

// Forward performs a forward pass through the layer
// It processes the input through all neurons in the layer
func (l *Layer) Forward(input []float64) []float64 {
	// Process input through all neurons
	for i, neuron := range l.Neurons {
		l.Output[i] = neuron.Forward(input)
	}

	return l.Output
}

// GetNeurons returns a copy of the layer's neurons
func (l *Layer) GetNeurons() []*Neuron {
	neurons := make([]*Neuron, len(l.Neurons))
	copy(neurons, l.Neurons)
	return neurons
}

// GetOutput returns a copy of the layer's output
func (l *Layer) GetOutput() []float64 {
	output := make([]float64, len(l.Output))
	copy(output, l.Output)
	return output
}

// GetNeuronCount returns the number of neurons in the layer
func (l *Layer) GetNeuronCount() int {
	return len(l.Neurons)
}

// GetNeuron returns a specific neuron in the layer
func (l *Layer) GetNeuron(index int) *Neuron {
	if index < 0 || index >= len(l.Neurons) {
		return nil
	}
	return l.Neurons[index]
}

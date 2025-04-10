package neural

// Network represents a simple neural network with input and output layers
type Network struct {
	// InputLayer is the input layer of the network
	InputLayer *Layer

	// OutputLayer is the output layer of the network
	OutputLayer *Layer
}

// NewNetwork creates a new neural network with the specified input and output sizes
func NewNetwork(inputSize, outputSize int) *Network {
	// Create a simple network with just input and output layers
	// For a single-layer perceptron, the input layer is just a pass-through
	// and the output layer is the actual layer of neurons

	network := &Network{}

	// Create output layer with sigmoid activation
	network.OutputLayer = NewLayer(outputSize, inputSize, &Sigmoid{})

	return network
}

// Forward performs a forward pass through the network
// It processes the input through all layers in the network
func (n *Network) Forward(input []float64) []float64 {
	// For a single-layer perceptron, we just pass the input to the output layer
	return n.OutputLayer.Forward(input)
}

// GetOutputLayer returns the output layer of the network
func (n *Network) GetOutputLayer() *Layer {
	return n.OutputLayer
}

// GetOutput returns the output of the network after a forward pass
func (n *Network) GetOutput() []float64 {
	return n.OutputLayer.GetOutput()
}

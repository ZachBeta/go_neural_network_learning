package neural

import (
	"math"
)

// ActivationFunction defines the interface for activation functions
type ActivationFunction interface {
	// Activate applies the activation function to the input
	Activate(x float64) float64

	// Derivative returns the derivative of the activation function at the given input
	Derivative(x float64) float64

	// Name returns the name of the activation function
	Name() string
}

// Sigmoid implements the sigmoid activation function
// f(x) = 1 / (1 + e^(-x))
type Sigmoid struct{}

// Activate applies the sigmoid function to the input
func (s *Sigmoid) Activate(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Derivative returns the derivative of the sigmoid function at the given input
// f'(x) = f(x) * (1 - f(x))
func (s *Sigmoid) Derivative(x float64) float64 {
	sig := s.Activate(x)
	return sig * (1.0 - sig)
}

// Name returns the name of the activation function
func (s *Sigmoid) Name() string {
	return "sigmoid"
}

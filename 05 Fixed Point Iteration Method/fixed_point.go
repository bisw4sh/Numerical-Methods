package main

import (
	"fmt"
	"math"
)

// Define the function for which we want to find the fixed point.
func f(x float64) float64 {
	// You should replace this function with the one you want to find the fixed point for.
	// For this example, we'll use a simple function: f(x) = x^2 - 2.
	return x*x - 2
}

// FixedPointIteration finds the fixed point of a given function using the fixed-point iteration method.
func FixedPointIteration(initialGuess float64, tolerance float64, maxIterations int) (float64, error) {
	x := initialGuess

	for i := 0; i < maxIterations; i++ {
		// Calculate the next approximation using the fixed-point iteration formula.
		nextX := f(x)

		// Check for convergence by comparing the absolute difference with the tolerance.
		if math.Abs(nextX-x) < tolerance {
			return nextX, nil
		}

		x = nextX
	}

	return 0, fmt.Errorf("fixed point not found after %d iterations", maxIterations)
}

func main() {
	initialGuess := 1.0       // Initial guess for the fixed point.
	tolerance := 1e-6        // Tolerance for convergence.
	maxIterations := 100     // Maximum number of iterations.

	result, err := FixedPointIteration(initialGuess, tolerance, maxIterations)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Fixed Point: %f\n", result)
	}
}
package main

import (
	"fmt"
)

// ForwardDifferenceInterpolation estimates the value of a function at a given point using forward difference interpolation.
func ForwardDifferenceInterpolation(x float64, dataX []float64, dataY []float64) float64 {
	n := len(dataX)
	if n != len(dataY) || n == 0 {
		panic("Invalid input data")
	}

	// Initialize the result
	result := 0.0

	// Calculate the forward differences
	for i := 0; i < n; i++ {
		term := dataY[i]
		for j := 0; j < i; j++ {
			term *= (x - dataX[j]) / (dataX[i] - dataX[j])
		}
		result += term
	}

	return result
}

func main() {
	// Example data points
	dataX := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	dataY := []float64{2.0, 3.0, 6.0, 11.0, 18.0}

	// Point at which you want to interpolate the value
	interpolationPoint := 2.5

	// Perform forward difference interpolation
	interpolatedValue := ForwardDifferenceInterpolation(interpolationPoint, dataX, dataY)

	fmt.Printf("Interpolated value at x=%.2f: %.2f\n", interpolationPoint, interpolatedValue)
}
package main

import (
	"fmt"
	"math"
)

// Define the function to be integrated
func f(x float64) float64 {
	return math.Sin(x)
}

// Trapezoidal rule for numerical integration
func trapezoidalRule(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2.0
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return h * sum
}

func main() {
	// Define the interval [a, b] and the number of subintervals n
	a := 0.0
	b := math.Pi
	n := 1000

	// Calculate the integral using the Trapezoidal rule
	result := trapezoidalRule(a, b, n)

	fmt.Printf("Result of numerical integration: %f\n", result)
}
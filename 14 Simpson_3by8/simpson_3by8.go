package main

import (
	"fmt"
	"math"
)

// Define the function to be integrated
func f(x float64) float64 {
	return math.Sin(x) // Change this to your desired function
}

func main() {
	// Define the integration interval [a, b]
	a := 0.0
	b := math.Pi

	// Number of subintervals (must be a multiple of 3)
	n := 6

	// Calculate the step size (h)
	h := (b - a) / float64(n)

	// Initialize the result variable
	result := f(a) + f(b)

	// Calculate the sum of the interior points
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%3 == 0 {
			result += 2 * f(x)
		} else {
			result += 3 * f(x)
		}
	}

	// Apply the Simpson's 3/8 Rule formula
	result = (3 * h / 8) * result

	fmt.Printf("Approximate integral value: %f\n", result)
}
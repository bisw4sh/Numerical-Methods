package main

import (
	"fmt"
	"math"
)

func main() {
	// Define the function for which you want to find the root.
	// Here, we will find the square root of a number using the bisection method.
	function := func(x float64) float64 {
		return x*x - 2 // Change this function to the one you want to find the root of
	}

	// Define the interval [a, b] where you want to search for the root.
	a := 0.0
	b := 2.0

	// Define the tolerance level for the root approximation.
	epsilon := 1e-6

	// Perform the bisection method to find the root.
	root, err := bisection(function, a, b, epsilon)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the approximate root.
	fmt.Printf("Approximate Root: %v\n", root)
}

func bisection(f func(float64) float64, a, b, epsilon float64) (float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, fmt.Errorf("bisection method is not applicable to the given interval")
	}

	var root float64
	for {
		root = (a + b) / 2
		if math.Abs(f(root)) < epsilon {
			break
		}
		if f(a)*f(root) < 0 {
			b = root
		} else {
			a = root
		}
	}

	return root, nil
}
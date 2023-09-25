package main

import (
    "fmt"
    "math"
)

// Define the function for which you want to find the root.
func f(x float64) float64 {
    // For example, let's find the square root of 2.
    return x*x - 2
}

// Define the derivative of the function.
func df(x float64) float64 {
    return 2*x
}

// Implement the Newton-Raphson method with progress printing.
func newtonRaphson(initialGuess, tolerance float64, maxIterations int) (float64, error) {
    x := initialGuess

    for i := 0; i < maxIterations; i++ {
        fValue := f(x)
        if math.Abs(fValue) < tolerance {
            fmt.Printf("Converged after %d iterations\n", i)
            return x, nil // Found a root within the tolerance
        }

        dfValue := df(x)
        if dfValue == 0 {
            return 0, fmt.Errorf("Derivative is zero, cannot continue")
        }

        x = x - fValue/dfValue
        fmt.Printf("Iteration %d: x = %f\n", i, x)
    }

    return 0, fmt.Errorf("Max iterations reached")
}

func main() {
    initialGuess := 1.0
    tolerance := 1e-6
    maxIterations := 100

    root, err := newtonRaphson(initialGuess, tolerance, maxIterations)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Approximate root: %f\n", root)
    }
}
package main

import (
    "fmt"
    "math"
)

// Define the function for which you want to find the root.
func f(x float64) float64 {
    // Example function: f(x) = x^2 - 4
    return x*x - 4
}

// SecantMethod finds the root of a function using the secant method.
func SecantMethod(f func(float64) float64, x0, x1, tol float64, maxIterations int) (float64, error) {
    for i := 0; i < maxIterations; i++ {
        f0 := f(x0)
        f1 := f(x1)

        if math.Abs(f1) < tol {
            return x1, nil // Found a root within tolerance
        }

        // Update the estimate using the secant formula
        x2 := x1 - f1*(x1-x0)/(f1-f0)

        x0 = x1
        x1 = x2
    }

    return 0, fmt.Errorf("failed to converge after %d iterations", maxIterations)
}

func main() {
    // Initial guesses for the root and tolerance
    x0 := 1.0
    x1 := 2.0
    tol := 1e-6
    maxIterations := 100

    root, err := SecantMethod(f, x0, x1, tol, maxIterations)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Approximate root: %.6f\n", root)
    }
}
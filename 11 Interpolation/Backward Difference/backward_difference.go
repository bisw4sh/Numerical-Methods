package main

import (
    "fmt"
)

func backwardDifferenceInterpolation(x []float64, y []float64, value float64) float64 {
    n := len(x)
    result := y[n-1]

    for i := 1; i < n; i++ {
        term := 1.0
        for j := 0; j < i; j++ {
            term *= (value - x[n-j-1]) / (x[n-j-1] - x[n-i-1])
        }
        result += term * y[n-i-1]
    }

    return result
}

func main() {
    // Example data points
    x := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    y := []float64{1.0, 4.0, 9.0, 16.0, 25.0}

    // Value to interpolate
    value := 2.5

    // Perform backward difference interpolation
    interpolatedValue := backwardDifferenceInterpolation(x, y, value)

    fmt.Printf("Interpolated value at %.2f: %.2f\n", value, interpolatedValue)
}
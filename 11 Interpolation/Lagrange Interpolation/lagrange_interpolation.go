package main

import (
    "fmt"
)

// lagrangeInterpolation calculates the Lagrange interpolation polynomial for a given set of points.
func lagrangeInterpolation(x []float64, y []float64, point float64) float64 {
    result := 0.0
    n := len(x)

    for i := 0; i < n; i++ {
        term := y[i]
        for j := 0; j < n; j++ {
            if i != j {
                term *= (point - x[j]) / (x[i] - x[j])
            }
        }
        result += term
    }

    return result
}

func main() {
    // Sample data points
    x := []float64{1.0, 2.0, 3.0, 4.0}
    y := []float64{2.0, 3.0, 5.0, 7.0}

    // Point at which to interpolate
    point := 2.5

    // Calculate the interpolated value at the given point
    interpolatedValue := lagrangeInterpolation(x, y, point)

    fmt.Printf("Interpolated value at x=%.2f is %.2f\n", point, interpolatedValue)
}
package main

import (
    "fmt"
    "math"
)

// Define the function to be integrated.
func f(x float64) float64 {
    return math.Sin(x) // You can replace this with your own function.
}

// Simpson's 1/3 rule for numerical integration.
func simpsons13(a, b float64, n int) float64 {
    h := (b - a) / float64(n)
    sum := f(a) + f(b)

    for i := 1; i < n; i++ {
        x := a + float64(i)*h
        if i%2 == 0 {
            sum += 2 * f(x)
        } else {
            sum += 4 * f(x)
        }
    }

    return (h / 3) * sum
}

func main() {
    a := 0.0 // Lower limit of integration
    b := math.Pi // Upper limit of integration
    n := 1000 // Number of subintervals

    result := simpsons13(a, b, n)
    fmt.Printf("Approximate integral: %f\n", result)
}
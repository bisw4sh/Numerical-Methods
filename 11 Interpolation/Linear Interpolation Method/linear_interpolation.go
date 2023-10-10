package main

import (
    "fmt"
)

// Linear interpolation function
func lerp(start, end, t float64) float64 {
    return start + (end-start)*t
}

func main() {
    // Define the start and end values
    startValue := 10.0
    endValue := 20.0

    // Define a time (t) value between 0 and 1 to interpolate between the start and end values
    t := 0.5

    // Perform linear interpolation
    interpolatedValue := lerp(startValue, endValue, t)

    fmt.Printf("Interpolated value at t=%.2f: %.2f\n", t, interpolatedValue)
}
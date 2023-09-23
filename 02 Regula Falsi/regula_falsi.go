package main

import (
	"fmt"
	"math"
)

// Define the function for which you want to find the root.
func f(x float64) float64 {
	return x*x - 4
}

// RegulaFalsi finds the root of the function f(x) using the Regula Falsi method.
func RegulaFalsi(f func(float64) float64, a, b, e float64) float64 {
	// Check if the initial values bracket the root.
	if f(a)*f(b) > 0 {
		fmt.Println("Given initial values do not bracket the root.")
		return math.NaN()
	}

	// Initialize variables.
	c := a - (f(a)*(b-a))/(f(b)-f(a))
	fc := f(c)

	// Iterate until the desired accuracy is reached.
	for math.Abs(fc) > e {
		// Print the current iteration.
		fmt.Printf("a: %f\tb: %f\tc: %f\tf(c): %f\n", a, b, c, fc)

		// Update the initial values and the functional value.
		if f(a)*fc < 0 {
			b = c
		} else {
			a = c
		}
		c = a - (f(a)*(b-a))/(f(b)-f(a))
		fc = f(c)
	}

	// Return the root.
	return c
}

func main() {
	// Input the initial values and tolerance.
	a := 0.0
	b := 3.0
	e := 0.0001

	// Find the root using the Regula Falsi method.
	root := RegulaFalsi(f, a, b, e)

	// Print the root.
	fmt.Printf("Root is: %f\n", root)
}
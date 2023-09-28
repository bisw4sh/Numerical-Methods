package main

import (
	"fmt"
)

func gaussElimination(A [][]float64, b []float64) []float64 {
	n := len(b)
	x := make([]float64, n)

	// Forward elimination
	for k := 0; k < n-1; k++ {
		for i := k + 1; i < n; i++ {
			factor := A[i][k] / A[k][k]
			for j := k; j < n; j++ {
				A[i][j] -= factor * A[k][j]
			}
			b[i] -= factor * b[k]
		}
	}

	// Backward substitution
	for i := n - 1; i >= 0; i-- {
		x[i] = b[i]
		for j := i + 1; j < n; j++ {
			x[i] -= A[i][j] * x[j]
		}
		x[i] /= A[i][i]
	}

	return x
}

func main() {
	A := [][]float64{
		{2, 1, -1},
		{-3, -1, 2},
		{-2, 1, 2},
	}
	b := []float64{8, -11, -3}

	x := gaussElimination(A, b)

	fmt.Println("Solution:")
	for i, val := range x {
		fmt.Printf("x%d = %f\n", i+1, val)
	}
}
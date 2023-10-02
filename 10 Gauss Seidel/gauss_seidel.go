package main

import (
    "fmt"
    "math"
)

func gaussSeidelIteration(A [][]float64, b []float64, initialGuess []float64, tolerance float64, maxIterations int) []float64 {
    n := len(b)
    x := make([]float64, n)
    copy(x, initialGuess)

    for k := 0; k < maxIterations; k++ {
        for i := 0; i < n; i++ {
            sum := 0.0
            for j := 0; j < n; j++ {
                if i != j {
                    sum += A[i][j] * x[j]
                }
            }
            x[i] = (b[i] - sum) / A[i][i]
        }

        // Check for convergence
        converged := true
        for i := 0; i < n; i++ {
            if math.Abs(x[i]-initialGuess[i]) > tolerance {
                converged = false
                break
            }
        }

        if converged {
            fmt.Printf("Converged after %d iterations\n", k+1)
            return x
        }

        copy(initialGuess, x)
    }

    fmt.Println("Did not converge within the maximum number of iterations")
    return x
}

func main() {
    A := [][]float64{
        {4, 1, 2},
        {3, 5, 1},
        {1, 1, 3},
    }

    b := []float64{4, 7, 3}
    initialGuess := []float64{0, 0, 0}
    tolerance := 1e-6
    maxIterations := 100

    solution := gaussSeidelIteration(A, b, initialGuess, tolerance, maxIterations)

    fmt.Println("Solution:", solution)
}
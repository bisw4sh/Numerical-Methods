package main

import (
    "fmt"
)

func gaussJordanElimination(matrix [][]float64, b []float64) []float64 {
    n := len(matrix)

    // Create an augmented matrix [matrix|b]
    augmentedMatrix := make([][]float64, n)
    for i := 0; i < n; i++ {
        augmentedMatrix[i] = make([]float64, n+1)
        for j := 0; j < n; j++ {
            augmentedMatrix[i][j] = matrix[i][j]
        }
        augmentedMatrix[i][n] = b[i]
    }

    // Perform row operations to reduce the matrix to row echelon form
    for i := 0; i < n; i++ {
        // Find the pivot row (row with the largest absolute value in the current column)
        maxRowIndex := i
        for j := i + 1; j < n; j++ {
            if abs(augmentedMatrix[j][i]) > abs(augmentedMatrix[maxRowIndex][i]) {
                maxRowIndex = j
            }
        }

        // Swap the current row with the pivot row
        augmentedMatrix[i], augmentedMatrix[maxRowIndex] = augmentedMatrix[maxRowIndex], augmentedMatrix[i]

        // Make the diagonal element of the current row equal to 1
        pivot := augmentedMatrix[i][i]
        for j := i; j <= n; j++ {
            augmentedMatrix[i][j] /= pivot
        }

        // Eliminate other rows
        for j := 0; j < n; j++ {
            if j != i {
                factor := augmentedMatrix[j][i]
                for k := i; k <= n; k++ {
                    augmentedMatrix[j][k] -= factor * augmentedMatrix[i][k]
                }
            }
        }
    }

    // Extract the solutions from the augmented matrix
    solutions := make([]float64, n)
    for i := 0; i < n; i++ {
        solutions[i] = augmentedMatrix[i][n]
    }

    return solutions
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    // Example usage:
    matrix := [][]float64{{2, -1, 1}, {3, 2, 4}, {1, 3, 3}}
    b := []float64{5, 23, 21}

    solutions := gaussJordanElimination(matrix, b)

    fmt.Println("Solution:")
    for i, sol := range solutions {
        fmt.Printf("x%d = %f\n", i+1, sol)
    }
}
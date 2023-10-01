package main

import (
	"fmt"
)

func printMatrix(mat [][]float64) {
	rows := len(mat)
	cols := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%.2f\t", mat[i][j])
		}
		fmt.Println()
	}
}

func swapRows(mat [][]float64, i, j int) {
	mat[i], mat[j] = mat[j], mat[i]
}

func scaleRow(mat [][]float64, row int, scalar float64) {
	for i := range mat[row] {
		mat[row][i] *= scalar
	}
}

func addScaledRow(mat [][]float64, sourceRow, targetRow int, scalar float64) {
	for i := range mat[targetRow] {
		mat[targetRow][i] += scalar * mat[sourceRow][i]
	}
}

func gaussJordanInverse(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	// Create an augmented matrix [matrix | I]
	// where I is the identity matrix of the same size.
	identity := make([][]float64, rows)
	for i := range identity {
		identity[i] = make([]float64, cols)
		identity[i][i] = 1.0
	}
	augmentedMatrix := make([][]float64, rows)
	for i := range matrix {
		augmentedMatrix[i] = append(matrix[i], identity[i]...)
	}

	// Perform Gauss-Jordan elimination to get the inverse.
	for col := 0; col < cols; col++ {
		// Find the pivot row.
		pivotRow := -1
		for i := col; i < rows; i++ {
			if augmentedMatrix[i][col] != 0 {
				pivotRow = i
				break
			}
		}
		if pivotRow == -1 {
			fmt.Println("Matrix is not invertible.")
			return nil
		}

		// Swap the pivot row with the current row.
		swapRows(augmentedMatrix, col, pivotRow)

		// Scale the pivot row to make the pivot element 1.
		pivotElement := augmentedMatrix[col][col]
		scaleRow(augmentedMatrix, col, 1/pivotElement)

		// Eliminate other rows.
		for i := 0; i < rows; i++ {
			if i != col {
				addScaledRow(augmentedMatrix, col, i, -augmentedMatrix[i][col])
			}
		}
	}

	// Extract the inverse from the augmented matrix.
	inverse := make([][]float64, rows)
	for i := range inverse {
		inverse[i] = augmentedMatrix[i][cols:]
	}

	return inverse
}

func main() {
	// Example usage
	matrix := [][]float64{
		{2, 1},
		{5, 3},
	}

	fmt.Println("Original Matrix:")
	printMatrix(matrix)

	inverse := gaussJordanInverse(matrix)
	if inverse != nil {
		fmt.Println("Inverse Matrix:")
		printMatrix(inverse)
	}
}
package main

import "math"

func calcDiagonal(matrix [][]int) int {
	diagonalLeftToRight := 0
	diagonalRightToLeft := 0
	size := len(matrix)

	// Menghitung jumlah diagonal kiri ke kanan dan kanan ke kiri
	for i := 0; i < size; i++ {
		diagonalLeftToRight += matrix[i][i]
		diagonalRightToLeft += matrix[i][size-i-1]
	}

	// Menghitung selisih mutlak antara jumlah diagonal
	diff := int(math.Abs(float64(diagonalLeftToRight - diagonalRightToLeft)))

	return diff
}

package goutil

import "testing"

func TestPrintMatrix(t *testing.T) {
	matrix := make([][]int, 5)
	for i := 0; i < 5; i++ {
		matrix[i] = make([]int, 7)
		for j := 0; j < 7; j++ {
			matrix[i][j] = i * j
		}
	}
	PrintMatrix(matrix)
}

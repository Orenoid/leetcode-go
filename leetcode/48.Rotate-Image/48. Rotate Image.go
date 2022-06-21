package main

func rotate(matrix [][]int) {
	for i := range matrix {
		left, right := 0, len(matrix[0])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
	n := len(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])-i-1; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
}

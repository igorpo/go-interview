package arrays

// Given an NxN matrix, rotate it by 90 degress clockwise in place

// Key here is to understand that M[i][j] -> M[n-j-1][i] in rotated matrix
func RotateMatrix90CW(matrix [][]int) {
	n := len(matrix[0])

	// cycle through the layers
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

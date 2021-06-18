package array

func rotateMap(matrix [][]int) {
	var tmp int
	l := len(matrix)
	n := l - 1
	for i := 0; i <= n/2; i++ {
		for j := 0; j < l/2; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = tmp
		}
	}
}

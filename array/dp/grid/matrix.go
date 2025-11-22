package grid

/*
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

https://leetcode.com/problems/01-matrix/description/
*/

func updateMatrixRecursive(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])

	// Initialize the result matrix
	resultMat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		resultMat[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				resultMat[i][j] = 0
			} else {
				resultMat[i][j] = getNearestZero(i, j, rows, cols, mat)
			}
		}
	}

	return resultMat
}

func getNearestZero(i, j, rows, cols int, mat [][]int) int {
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return 100000
	}

	if mat[i][j] == 0 {
		return 0
	}

	return min(
		1+getNearestZero(i-1, j, rows, cols, mat),
		1+getNearestZero(i+1, j, rows, cols, mat),
		1+getNearestZero(i, j-1, rows, cols, mat),
		1+getNearestZero(i, j+1, rows, cols, mat))
}

func updateMatrixDP(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])

	resultMat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		resultMat[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				resultMat[i][j] = 0
			} else {
				resultMat[i][j] = -1
			}
		}
	}

	// Compute distances for all cells
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			getNearestZeroDP(i, j, rows, cols, resultMat)
		}
	}

	return resultMat
}

func getNearestZeroDP(i, j, rows, cols int, mat [][]int) int {
	// If out of bounds, return a large value (infinite distance, not reachable)
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return 1e9
	}

	// If already computed, return the memoized value
	if mat[i][j] != -1 {
		return mat[i][j]
	}

	mat[i][j] = min(
		1+getNearestZeroDP(i-1, j, rows, cols, mat),
		1+getNearestZeroDP(i+1, j, rows, cols, mat),
		1+getNearestZeroDP(i, j-1, rows, cols, mat),
		1+getNearestZeroDP(i, j+1, rows, cols, mat))

	return mat[i][j]
}

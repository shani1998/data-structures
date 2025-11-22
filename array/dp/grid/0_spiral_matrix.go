package grid

/*
Given an m x n matrix, return all elements of the matrix in spiral order.
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

observations: in one go.. we need top rint outer peripheral of the metrics
 right --> btm --> left --> top
*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	startR, startC := 0, 0
	endR, endC := rows-1, cols-1

	result := make([]int, 0, rows*cols)

	for startR <= endR && startC <= endC {

		// left → right
		for c := startC; c <= endC; c++ {
			result = append(result, matrix[startR][c])
		}

		// top → bottom
		for r := startR + 1; r <= endR; r++ {
			result = append(result, matrix[r][endC])
		}

		// right → left (only if more than one row remains)
		if startR < endR {
			for c := endC - 1; c >= startC; c-- { //
				result = append(result, matrix[endR][c])
			}
		}

		// bottom → top (only if more than one column remains)
		if startC < endC {
			for r := endR - 1; r > startR; r-- {
				result = append(result, matrix[r][startC])
			}
		}

		startR++
		startC++
		endR--
		endC--
	}

	return result
}

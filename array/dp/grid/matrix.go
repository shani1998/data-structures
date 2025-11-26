package grid

/*
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

https://leetcode.com/problems/01-matrix/description/
*/

// When we find the distaince for each cell using DFS or BFS, we explore all four directions
// until we find a cell with 0. The distance is then returned.
// Time complexity: O((N*M)^2) because for each cell (N*M), in worst case we may have to explore all cells (N*M) to find a 0.
// Space complexity: O(N*M) for the visited matrix.
func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {

			// For each cell create a fresh visited matrix
			visited := make([][]bool, rows)
			for r, _ := range visited {
				visited[r] = make([]bool, cols)
			}

			mat[i][j] = getMinDistanceDfs(i, j, rows, cols, mat, visited)
		}
	}

	return mat
}

func getMinDistanceDfs(i, j, rows, cols int, mat [][]int, visited [][]bool) int {
	// Out of bounds
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return 1e9
	}

	// If this path already visited => stop (avoid infinite loop)
	if visited[i][j] {
		return 1e9
	}

	// If cell is zero => distance = 0
	if mat[i][j] == 0 {
		return 0
	}

	visited[i][j] = true

	// Explore 4 directions
	up := 1 + getMinDistanceDfs(i-1, j, rows, cols, mat, visited)
	down := 1 + getMinDistanceDfs(i+1, j, rows, cols, mat, visited)
	left := 1 + getMinDistanceDfs(i, j-1, rows, cols, mat, visited)
	right := 1 + getMinDistanceDfs(i, j+1, rows, cols, mat, visited)

	visited[i][j] = false // IMPORTANT — unmark for next paths

	return min(up, down,left, right)
}

func getMinDistanceBfs(i, j, rows, cols int, mat [][]int) int {
	visited := make([][]bool, rows)
	for a := range visited {
		visited[a] = make([]bool, cols)
	}

	queue := [][]int{{i, j}}
	visited[i][j] = true
	distance := 0 	// Distance from (i,j)

	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) > 0 {
		size := len(queue)

		// BFS LEVEL LOOP
		for k := 0; k < size; k++ {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]

			// If this cell is zero, return distance
			if mat[r][c] == 0 {
				return distance
			}

			// explore neighbors
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] {
					visited[nr][nc] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}

		// Increase distance AFTER exploring one full BFS level
		distance++
	}

	return distance
}

// Optimized BFS from all 0s simultaneously
// Time complexity: O(N*M) because each cell is processed at most once.
// Space complexity: O(N*M) for the queue and distance matrix.
func updateMatrix1(mat [][]int) [][]int {
    rows, cols := len(mat), len(mat[0])
    queue := make([][2]int, 0)

    // Initialize distances
    distance := make([][]int, rows)
    for i := 0; i < rows; i++ {
        distance[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            if mat[i][j] == 0 {
                distance[i][j] = 0
                queue = append(queue, [2]int{i, j}) // push all zeros
            } else {
                distance[i][j] = 1e9 // infinity
            }
        }
    }

    dirs := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}

    // BFS from all zeros
    for front := 0; front < len(queue); front++ {
        r, c := queue[front][0], queue[front][1]

        for _, d := range dirs {
            nr, nc := r + d[0], c + d[1]

            if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
                if distance[nr][nc] > distance[r][c] + 1 {
				     // We discovered a shorter route to (nr,nc).
                    distance[nr][nc] = distance[r][c] + 1
                    queue = append(queue, [2]int{nr, nc})
                }
            }
        }
    }

    return distance
}
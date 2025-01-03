package questions

/*
Input: M[][] = {{‘1’, ‘1’, ‘0’, ‘0’, ‘0’},
				{‘0’, ‘1’, ‘0’, ‘0’, ‘1’},
				{‘1’, ‘0’, ‘0’, ‘1’, ‘1’},
				{‘0’, ‘0’, ‘0’, ‘0’, ‘0’},
				{‘1’, ‘0’, ‘1’, ‘1’, ‘0’}}

Output: 4
*/

func dfs(grid [][]byte, visited [][]bool, i, j, rows, cols int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfs(grid, visited, i+1, j, rows, cols)
	dfs(grid, visited, i, j+1, rows, cols)
	dfs(grid, visited, i-1, j, rows, cols)
	dfs(grid, visited, i, j-1, rows, cols)
}

func bfs(grid [][]byte, visited [][]bool, i, j, rows, cols int) {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{{i, j}}
	visited[i][j] = true

	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, d := range directions {
			newX, newY := x+d[0], y+d[1]
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && !visited[newX][newY] && grid[newX][newY] == '1' {
				visited[newX][newY] = true
				queue = append(queue, []int{newX, newY})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	// Initialize visited array
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var noOfIslands int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				dfs(grid, visited, i, j, rows, cols)
				noOfIslands++
			}
		}
	}
	return noOfIslands
}

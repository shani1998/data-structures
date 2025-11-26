package questions

/*
Input: M[][] = {{‘1’, ‘1’, ‘0’, ‘0’, ‘0’},
				{‘0’, ‘1’, ‘0’, ‘0’, ‘1’},
				{‘1’, ‘0’, ‘0’, ‘1’, ‘1’},
				{‘0’, ‘0’, ‘0’, ‘0’, ‘0’},
				{‘1’, ‘0’, ‘1’, ‘1’, ‘0’}}

Output: 4
*/

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

func dfs(grid [][]byte, visited [][]bool, i, j, rows, cols int) {
    // boundary checks
    if i < 0 || i >= rows || j < 0 || j >= cols {
        return
    }

    // if already visited or water
    if visited[i][j] || grid[i][j] == '0' {
        return
    }

    visited[i][j] = true

    // explore neighbors (4 directions)
    dfs(grid, visited, i+1, j, rows, cols)
    dfs(grid, visited, i-1, j, rows, cols)
    dfs(grid, visited, i, j+1, rows, cols)
    dfs(grid, visited, i, j-1, rows, cols)
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    rows := len(grid)
    cols := len(grid[0])

    visited := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]bool, cols)
    }

    count := 0

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                dfs(grid, visited, i, j, rows, cols)
                count++
            }
        }
    }

    return count
}
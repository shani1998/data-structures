package grid

/*
https://leetcode.com/problems/rotting-oranges/description/

You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
*/

func orangesRotting(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    queue := [][]int{}
    fresh := 0

    // 1. Add all rotten oranges to queue
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 2 {
                queue = append(queue, []int{r, c})
            } else if grid[r][c] == 1 {
                fresh++
            }
        }
    }

    // no fresh oranges → 0 minutes
    if fresh == 0 {
        return 0
    }

    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    minutes := -1

    // 2. BFS level by level
    for len(queue) > 0 {
        size := len(queue)
        minutes++

        for i := 0; i < size; i++ {
            cell := queue[0]
            queue = queue[1:]

            r, c := cell[0], cell[1]

            for _, d := range dirs {
                nr, nc := r+d[0], c+d[1]

                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
                    grid[nr][nc] = 2
                    fresh--
                    queue = append(queue, []int{nr, nc})
                }
            }
        }
    }

    if fresh > 0 {
        return -1
    }

    return minutes
}
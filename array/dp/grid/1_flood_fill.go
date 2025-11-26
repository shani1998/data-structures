package grid


// Flood Fill Algorithm using BFS time complexity O(N*M), space complexity O(N*M)
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    rows, cols := len(image), len(image[0])
    originalColor := image[sr][sc]

    if originalColor == color {
        return image
    }

    queue := [][]int{{sr, sc}}
    image[sr][sc] = color

    dirs := [][]int{{1,0},{0,1},{-1,0},{0,-1}}

    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        r, c := cell[0], cell[1]

        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]

            if nr >= 0 && nr < rows && nc >= 0 && nc < cols && image[nr][nc] == originalColor {
                image[nr][nc] = color       // recolor
                queue = append(queue, []int{nr, nc})
            }
        }
    }

    return image
}

// Flood Fill Algorithm using DFS time complexity O(N*M), space complexity O(N*M)
func floodFillDFS(image [][]int, sr int, sc int, color int) [][]int {
    totalRows, totalCols := len(image), len(image[0])
    originalColor := image[sr][sc]

    // If the starting pixel already has the new color, return immediately
    if originalColor == color {
        return image
    }    

    floodFillRecurse(image, sr, sc, color, originalColor, totalRows, totalCols)
    return image
}

func floodFillRecurse(image [][]int, sr int, sc int, color int, originalColor int, totalRows int, totalCols int) {
	// Base conditions: Out of bounds or different color
	if sr < 0 || sr >= totalRows || sc < 0 || sc >= totalCols || image[sr][sc] != originalColor {
		return
	}
	// Change color
	image[sr][sc] = color
	
	// Recursive calls in all 4 directions
	floodFillRecurse(image, sr+1, sc, color, originalColor, totalRows, totalCols) // Down
	floodFillRecurse(image, sr-1, sc, color, originalColor, totalRows, totalCols) // Up
	floodFillRecurse(image, sr, sc+1, color, originalColor, totalRows, totalCols) // Right
	floodFillRecurse(image, sr, sc-1, color, originalColor, totalRows, totalCols) // Left
}

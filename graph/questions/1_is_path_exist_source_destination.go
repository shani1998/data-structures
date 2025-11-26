package questions

/*
Given a graph represented by n nodes and edges, you need to check if there is a valid path from source to destination.
Ex:
Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
Output: true

Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
Output: false
*/
// time:  O(V + E), space: O(V + E)
func validPath(n int, edges [][]int, source int, destination int) bool {
	adjList := make(map[int][]int)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		adjList[x] = append(adjList[x], y)
		adjList[y] = append(adjList[y], x)
	}

	visited := make([]bool, n)

	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}
		visited[node] = true
		for _, neighbor := range adjList[node] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}
		return false
	}

	return dfs(source)
}

package graph

import "fmt"

func (g *Graph) DFSUtil(node *Node, visited map[*Node]bool) {
	visited[node] = true
	fmt.Printf("%s ", node.Data)
	for _, neighbour := range g.Edges[*node] {
		if _, isVisited := visited[neighbour]; !isVisited {
			g.DFSUtil(neighbour, visited)
		}
	}
}

func (g *Graph) DFS() {
	g.lock.Lock()
	defer g.lock.Unlock()
	visited := make(map[*Node]bool)
	if len(g.Nodes) > 0 {
		g.DFSUtil(g.Nodes[0], visited)
	}
}

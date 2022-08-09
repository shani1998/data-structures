package graph

import (
	"fmt"
	"github.com/shani1998/data-structures/queue"
)

func (g *Graph) BFSUtil(node *Node, visited map[*Node]bool) {
	queue := queue.NewQueue(len(g.Nodes))
	queue.Push(node)
	visited[node] = true
	fmt.Printf("%s ", node.Data)

	for queue.Length() > 0 {
		qSize := queue.Length()

		for i := 0; i < qSize; i++ {
			item, _ := queue.Pop()
			//traverse all neighbors that are not visited
			for _, neighbour := range g.Edges[*item.(*Node)] {
				if _, isVisited := visited[neighbour]; !isVisited {
					visited[neighbour] = true
					fmt.Printf("%s ", neighbour.Data)
					queue.Push(neighbour)
				}
			}
		}
	}
}

func (g *Graph) BFS() {
	g.lock.Lock()
	defer g.lock.Unlock()
	visited := make(map[*Node]bool)
	if len(g.Nodes) > 0 {
		// use node-0 as starting point
		g.BFSUtil(g.Nodes[0], visited)
	}
}

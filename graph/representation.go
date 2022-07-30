package graph

import "sync"

/*
A graph is a representation of a network structure. There are tons of graph real world examples, the Internet and the social graph being the classic ones.
It’s basically a set of Nodes connected by Edges, A graph can be represented in mainly two ways. They are:
Adjacency List & Adjacency Matrix.
*/

// Node a vertex in a graph
type Node struct {
	Data any
}

// Graph Adjacency list representation of a graph.
type Graph struct {
	// vertices of a graph
	Nodes []*Node
	// a vertex connected by Edges
	Edges map[Node][]*Node
	// to synchronize the operations
	lock sync.RWMutex
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(n *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if n == nil {
		return
	}
	g.Nodes = append(g.Nodes, n)
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}
	// it's an undirected graph, which means that adding an
	// edge from A to B also adds an edge from B to A.
	g.Edges[*n1] = append(g.Edges[*n1], n2)
	g.Edges[*n2] = append(g.Edges[*n2], n1)
}

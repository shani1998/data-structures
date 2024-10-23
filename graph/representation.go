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

func AdjacencyMatrix() {
	// A graph can be represented in mainly two ways. They are:
	// Adjacency List & Adjacency Matrix.
	// Adjacency Matrix is a 2D array of size V x V where V is the number of vertices in a graph.
	// Let the 2D array be adj[][], a slot adj[i][j] = 1 indicates that there is an edge from vertex i to vertex j.
	// Adjacency matrix for undirected graph is always symmetric.
	// Adjacency Matrix is also used to represent weighted graphs.
	// If adj[i][j] = w, then there is an edge from vertex i to vertex j with weight w.
}

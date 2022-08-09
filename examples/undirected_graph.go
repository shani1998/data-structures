package main

import (
	"fmt"
	"github.com/shani1998/data-structures/graph"
)

func main() {
	var g graph.Graph
	/*
	       A ---- D
	    /    \
	   B      C
	    \    /
	      E------F
	*/

	nA := &graph.Node{Data: "A"}
	nB := &graph.Node{Data: "B"}
	nC := &graph.Node{Data: "C"}
	nD := &graph.Node{Data: "D"}
	nE := &graph.Node{Data: "E"}
	nF := &graph.Node{Data: "F"}
	g.AddNode(nA)
	g.AddNode(nB)
	g.AddNode(nC)
	g.AddNode(nD)
	g.AddNode(nE)
	g.AddNode(nF)

	g.AddEdge(nA, nB)
	g.AddEdge(nA, nC)
	g.AddEdge(nB, nE)
	g.AddEdge(nC, nE)
	g.AddEdge(nE, nF)
	g.AddEdge(nD, nA)

	var s string
	for i := 0; i < len(g.Nodes); i++ {
		s += fmt.Sprintf("%s", g.Nodes[i].Data) + " -> "
		near := g.Edges[*g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += fmt.Sprintf("%s", near[j].Data) + " "
		}
		s += "\n"
	}
	fmt.Println(s)

	fmt.Println("---DFS Traversal ---")
	g.DFS()
	fmt.Println("\n---BFS Traversal ---")
	g.BFS()
}

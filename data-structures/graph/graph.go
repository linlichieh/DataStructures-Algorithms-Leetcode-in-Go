package graph

import "fmt"

type Graph struct {
	numVertices   int
	adjacencyList map[int][]int
}

func (g *Graph) AddVertex(v int) {
	if len(g.adjacencyList) == 0 {
		g.adjacencyList = make(map[int][]int)
	}
	if _, ok := g.adjacencyList[v]; !ok {
		g.adjacencyList[v] = make([]int, 0)
		g.numVertices++
	}
}

func (g *Graph) AddEdge(u int, v int) error {
	if _, ok := g.adjacencyList[u]; !ok {
		return fmt.Errorf("vertex '%d' doesn't exist", u)
	}
	if _, ok := g.adjacencyList[v]; !ok {
		return fmt.Errorf("vertex '%d' doesn't exist", v)
	}
	if !inList(g.adjacencyList[u], v) {
		g.adjacencyList[u] = append(g.adjacencyList[u], v)
	}
	if !inList(g.adjacencyList[v], u) {
		g.adjacencyList[v] = append(g.adjacencyList[v], u)
	}
	return nil
}

func inList(s []int, vertex int) bool {
	for _, v := range s {
		if v == vertex {
			return true
		}
	}
	return false
}

func (g *Graph) GetNeighbors(v int) []int {
	return g.adjacencyList[v]
}

func (g *Graph) IsAdjacent(u, v int) bool {
	for _, neighbor := range g.adjacencyList[u] {
		if neighbor == v {
			return true
		}
	}
	return false
}

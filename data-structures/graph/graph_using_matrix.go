package graph

type GraphMatrix struct {
	numVertices     int
	adjacencyMatrix [][]int
}

func NewGraphMatrix(num int) *GraphMatrix {
	matrix := make([][]int, num)
	for i := 0; i < num; i++ {
		matrix[i] = make([]int, num)

	}
	return &GraphMatrix{numVertices: num, adjacencyMatrix: matrix}
}

func (g *GraphMatrix) AddEdge(u int, v int) {
	g.adjacencyMatrix[u][v] = 1
	g.adjacencyMatrix[v][u] = 1
}

func (g *GraphMatrix) GetNeighbors(v int) []int {
	neighbors := []int{}
	for i, isAdjacent := range g.adjacencyMatrix[v] {
		if isAdjacent == 1 {
			neighbors = append(neighbors, i)
		}
	}
	return neighbors
}

func (g *GraphMatrix) IsAdjacent(u, v int) bool {
	return g.adjacencyMatrix[u][v] == 1
}

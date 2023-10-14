package leetcode207

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]int, numCourses)

	// Store prerequisite courses by parent course
	// {0: [1,2], 1: [3,4], 2: [], 3: [4], 4: []}
	graph := make(map[int][]int, numCourses)
	for _, pair := range prerequisites {
		graph[pair[0]] = append(graph[pair[0]], pair[1])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, graph, visited) {
			return false
		}
	}
	return true
}

func dfs(node int, graph map[int][]int, visited []int) bool {
	// cycle detected
	if visited[node] == 1 {
		return false
	}
	// visited
	if visited[node] == 2 {
		return true
	}

	// Make it as visiting
	visited[node] = 1
	for _, neighbour := range graph[node] {
		if !dfs(neighbour, graph, visited) {
			return false
		}
	}
	// NOTE: This is key to reduce repeated work
	visited[node] = 2

	return true
}

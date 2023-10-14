package leetcode207

func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	graph := make(map[int][]int, numCourses)

	for _, pair := range prerequisites {
		// NOTE: It's different from DFS building the graph
		graph[pair[1]] = append(graph[pair[1]], pair[0])
		// Counting the indegree for each course
		indegree[pair[0]]++
	}

	// Add course with 0 indegree into the queue
	queue := []int{}
	for i, degree := range indegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}

	counter := 0
	bfs(&counter, queue, graph, indegree)

	return counter == numCourses
}

func bfs(counter *int, queue []int, graph map[int][]int, indegree []int) {
	if len(queue) == 0 {
		return
	}

	*counter++
	node := queue[0]
	queue = queue[1:]
	for _, nextNode := range graph[node] {
		// Decrease indegree by 1
		indegree[nextNode]--
		// Whenever the indegree of node is 0, add it into the queue
		if indegree[nextNode] == 0 {
			queue = append(queue, nextNode)
		}
	}
	bfs(counter, queue, graph, indegree)
}

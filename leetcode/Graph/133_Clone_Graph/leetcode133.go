package leetcode133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNodes := make(map[int]*Node)
	return dfs(node, newNodes)
}

func dfs(node *Node, newNodes map[int]*Node) *Node {
	// If the node is already in the list, it indicates that it has already been processed.
	if newNodes[node.Val] != nil {
		return newNodes[node.Val]
	}

	// Clone the node
	clone := &Node{Val: node.Val}
	newNodes[node.Val] = clone

	// Clone it's neighbors
	for _, neighbor := range node.Neighbors {
		// Call itself
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, newNodes))
	}
	return clone
}

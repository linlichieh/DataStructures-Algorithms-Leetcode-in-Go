package tree

func (t *BST) BFS() (list []int) {
	var popNode *node
	queue := []*node{t.root}
	for len(queue) > 0 {
		// Pop a node from queue
		popNode = queue[0]
		queue = queue[1:]

		// Add node that is just poped into the result list
		list = append(list, popNode.val)

		// Add the left and right nodes to the queue if they exist
		if popNode.left != nil {
			queue = append(queue, popNode.left)
		}
		if popNode.right != nil {
			queue = append(queue, popNode.right)
		}
	}

	return
}

func (t *BST) BFSRecur(queue []*node, list []int) []int {
	if len(queue) == 0 {
		return list
	}

	popNode := queue[0]
	queue = queue[1:]

	list = append(list, popNode.val)

	if popNode.left != nil {
		queue = append(queue, popNode.left)
	}
	if popNode.right != nil {
		queue = append(queue, popNode.right)
	}

	return t.BFSRecur(queue, list)
}

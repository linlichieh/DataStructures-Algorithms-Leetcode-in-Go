package leetcode450

// Example:
//     100
//        \
//         125 <- remove
//        /   \
//      115    150
//            /   \
//         135     175
//        /   \
//     126     140

// Non-recursive implementation
func nonRecurDeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	// Store the parent node for later use
	var parentNode *TreeNode
	currentNode := root

	for currentNode != nil && currentNode.Val != key {
		parentNode = currentNode

		// Find the node in left or right tree
		if key < currentNode.Val {
			// left
			currentNode = currentNode.Left
		} else {
			// right
			currentNode = currentNode.Right
		}
	}

	// Can't find the value that matches any node in the tree
	if currentNode == nil {
		return root
	}

	// Node has been found. The solution depends on the number of children of the node to be removed
	switch {
	case currentNode.Left == nil && currentNode.Right == nil: // case 0: no child
		// It means root is the only node in the tree and will be removed
		if parentNode == nil {
			return nil
		}
		if parentNode.Left == currentNode {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	case currentNode.Left != nil && currentNode.Right == nil: // case 1: 1 child
		// case 2: 1 child
		if parentNode == nil {
			return currentNode.Left
		}
		if currentNode == parentNode.Right {
			parentNode.Right = currentNode.Left
		} else {
			parentNode.Left = currentNode.Left
		}
	case currentNode.Left == nil && currentNode.Right != nil: // case 1: 1 child
		// case 2: 1 child
		if parentNode == nil {
			return currentNode.Right
		}
		if currentNode == parentNode.Left {
			parentNode.Left = currentNode.Right
		} else {
			parentNode.Right = currentNode.Right
		}
	default: // case 3: 2 children
		// Find the minimum in right tree
		leftmostParent := currentNode
		leftmostNode := currentNode.Right
		for leftmostNode.Left != nil {
			leftmostParent = leftmostNode
			leftmostNode = leftmostNode.Left
		}

		// Swap currentNode with leftmost
		if leftmostParent != currentNode {
			// leftmostNode is the farest left node, parent linkes ot its right node if existed
			leftmostParent.Left = leftmostNode.Right
			leftmostNode.Right = currentNode.Right
		}
		leftmostNode.Left = currentNode.Left

		// If the node to be removed is the root, return the new one.
		if parentNode == nil {
			return leftmostNode
		}

		if parentNode.Left == currentNode {
			parentNode.Left = leftmostNode
		} else {
			parentNode.Right = leftmostNode
		}
	}
	return root
}

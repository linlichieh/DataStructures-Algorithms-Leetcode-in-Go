package tree

type node struct {
	left  *node
	right *node
	val   int
}

type BST struct {
	root *node
}

func (t *BST) Insert(val int) {
	if t.root == nil {
		t.root = &node{val: val}
		return
	}

	currentNode := t.root
	for {
		if val < currentNode.val {
			if currentNode.left == nil {
				currentNode.left = &node{val: val}
				break
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = &node{val: val}
				break
			}
			currentNode = currentNode.right
		}
	}
}

func (t *BST) Lookup(val int) bool {
	if t.root == nil {
		return false
	}

	currentNode := t.root
	for currentNode != nil {
		if val < currentNode.val {
			// left
			currentNode = currentNode.left
		} else if val > currentNode.val {
			// right
			currentNode = currentNode.right
		} else {
			// val enquals to currentNode.val
			return true
		}
	}
	return false
}

// NOTE
//   - The `Remove` implementation can be achieved through both recursive and non-recursive methods.
//     Please refer to these solutions in `leetcode/Binary-Search-Tree/450_Delete_Node_in_a_BST/`
// func (t *BST) Remove(val int)

package tree

// inorder
func (t *BST) DFSInOrder() []int {
	var list []int
	t.traverseInOrder(t.root, &list)
	return list
}

func (t *BST) traverseInOrder(n *node, list *[]int) {
	if n.left != nil {
		t.traverseInOrder(n.left, list)
	}

	*list = append(*list, n.val)

	if n.right != nil {
		t.traverseInOrder(n.right, list)
	}
}

// preorder
func (t *BST) DFSPreOrder() []int {
	var list []int
	t.traversePreOrder(t.root, &list)
	return list
}

func (t *BST) traversePreOrder(n *node, list *[]int) {
	*list = append(*list, n.val)

	if n.left != nil {
		t.traversePreOrder(n.left, list)
	}

	if n.right != nil {
		t.traversePreOrder(n.right, list)
	}
}

// postorder
func (t *BST) DFSPostOrder() []int {
	var list []int
	t.traversePostOrder(t.root, &list)
	return list
}

func (t *BST) traversePostOrder(n *node, list *[]int) {
	if n.left != nil {
		t.traversePostOrder(n.left, list)
	}

	if n.right != nil {
		t.traversePostOrder(n.right, list)
	}

	*list = append(*list, n.val)
}

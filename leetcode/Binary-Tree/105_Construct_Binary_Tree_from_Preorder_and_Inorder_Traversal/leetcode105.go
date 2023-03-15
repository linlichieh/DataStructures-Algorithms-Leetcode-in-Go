package leetcode105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int)
	for k, v := range inorder {
		inMap[v] = k
	}
	return recBuildTree(preorder, 0, inorder, &inMap, 0, len(inorder)-1)
}

func recBuildTree(preorder []int, preCurr int, inorder []int, inMap *map[int]int, inStart int, inEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	// get the index of inorder by current preorder element
	inMid := (*inMap)[preorder[preCurr]]

	// calculate the length where curr on the right side of mid in preorder should start
	leftSubtreeLen := inMid - inStart

	root := &TreeNode{
		Val:   preorder[preCurr],
		Left:  recBuildTree(preorder, preCurr+1, inorder, inMap, inStart, inMid-1),
		Right: recBuildTree(preorder, preCurr+1+leftSubtreeLen, inorder, inMap, inMid+1, inEnd),
	}

	return root
}

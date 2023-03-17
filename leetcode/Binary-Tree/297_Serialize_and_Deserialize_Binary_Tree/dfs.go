package leetcode297

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	dfsIdx      int
	bfsListIdx  int
	bfsQueueIdx int
}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	return fmt.Sprintf("[%s]", this.dfsSer(root))
}

func (this *Codec) dfsSer(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%s,%s", node.Val, this.dfsSer(node.Left), this.dfsSer(node.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	// remove the open and close square
	data = data[1 : len(data)-1]
	list := strings.Split(data, ",")
	return this.dfsDes(list)
}

func (this *Codec) dfsDes(list []string) *TreeNode {
	if len(list) == this.dfsIdx {
		return nil
	}
	// get the val from the list
	val := list[this.dfsIdx]
	this.dfsIdx++
	// if it's null, stop continuing
	if val == "null" {
		return nil
	}

	intVal, _ := strconv.Atoi(val)
	node := &TreeNode{Val: intVal}
	node.Left = this.dfsDes(list)
	node.Right = this.dfsDes(list)
	return node
}

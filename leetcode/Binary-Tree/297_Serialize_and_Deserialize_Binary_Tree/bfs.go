package leetcode297

import (
	"fmt"
	"strconv"
	"strings"
)

// Serializes a tree to a single string.
func (this *Codec) bfsSerialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	str := this.bfsSer([]*TreeNode{root})

	// remove the last comma
	str = str[:len(str)-1]
	return fmt.Sprintf("[%s]", str)
}

func (this *Codec) bfsSer(queue []*TreeNode) string {
	if len(queue) == 0 {
		return ""
	}
	// Pop the first node from the queue
	node := queue[0]
	queue = queue[1:]
	if node == nil {
		// Append `null` into the result
		return fmt.Sprintf("null,%s", this.bfsSer(queue))
	}
	// Append the left and right node of current node into the queue
	queue = append(queue, node.Left)
	queue = append(queue, node.Right)
	// Append the value into the result
	return fmt.Sprintf("%d,%s", node.Val, this.bfsSer(queue))
}

// Deserializes your encoded data to tree.
func (this *Codec) bfsDeserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}

	// Extract the value from the array and make it a slice
	data = data[1 : len(data)-1]
	list := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: rootVal}
	this.bfsDes(list[1:], 0, []*TreeNode{root}, 0)
	return root
}

func (this *Codec) bfsDes(list []string, bfsListIdx int, queue []*TreeNode, bfsQueueIdx int) {
	if this.bfsListIdx == len(list) {
		return
	}

	// pop the first value from the queue
	node := queue[this.bfsQueueIdx]
	this.bfsQueueIdx++

	// add the left node for the current node if it's not null
	if list[this.bfsListIdx] != "null" {
		leftVal, _ := strconv.Atoi(list[this.bfsListIdx])
		node.Left = &TreeNode{Val: leftVal}
		queue = append(queue, node.Left)
	}
	this.bfsListIdx++

	// Check the length of list before poping again
	if len(list) > this.bfsListIdx {
		// add the right node for the current node if it's not null
		if list[this.bfsListIdx] != "null" {
			rightVal, _ := strconv.Atoi(list[this.bfsListIdx])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		this.bfsListIdx++
	}
	this.bfsDes(list, this.bfsListIdx, queue, this.bfsQueueIdx)
}

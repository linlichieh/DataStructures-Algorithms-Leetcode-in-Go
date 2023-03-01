package leetcode206

// recursive solution
// Demonstration: [1 -> 2 -> 3 -> 4]
//
// 1st loop:
//
//	head = 1
//	newHead := recursion(2)
//	2.Next = 1			|
//	1.Next = nil		|
//	return newHead		|
//						|
//					 2nd loop:
//						head = 2
//						newHead := recursion(3)
//						3.Next = 2			|
//						2.Next = nil		|
//						return newHead		|
//											|
//										3nd loop:
//											head = 3
//											newHead := recursion(4)
//											4.Next = 3			|
//											3.Next = nil		|
//											return newHead		|
//																|
//															return 4
func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	curr := head
	nodes := []*ListNode{}

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	mid := len(nodes) / 2

	return nodes[mid]
}

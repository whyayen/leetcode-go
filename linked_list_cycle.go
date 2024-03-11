package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	curr := head

	for curr != nil {
		_, isExist := visited[curr]
		if isExist {
			return true
		} else {
			visited[curr] = true
		}

		curr = curr.Next
	}

	return false
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* 最直覺的做法：用 map 儲存訪問過的節點
 * 更好的做法：(Floyd's Cycle Detection Algorithm) 用 fast, slow 指標，龜兔賽跑演算法，當 fast 追到 slow 的時候就是有 cycle
 * 被提出證明的算法
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

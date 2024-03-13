package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var previous *ListNode

	for curr != nil {
		originalNext := curr.Next
		curr.Next = previous
		previous = curr
		curr = originalNext
	}

	return previous
}

// recursive version:
// func reverseList(head *ListNode) *ListNode {
// 		return recursiveReverse(head, nil)
// }

//func recursiveReverse(curr *ListNode, previous *ListNode) *ListNode {
//	if curr == nil {
//		return previous
//	}
//
//	originalNext := curr.Next
//	curr.Next = previous
//	return recursiveReverse(originalNext, curr)
//}

func main() {
	n5 := &ListNode{
		Val: 5,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	result := reverseList(n1)
	fmt.Printf("%+v", result)
	for result != nil {
		fmt.Printf("%+v\n", result.Val)
		result = result.Next
	}
}

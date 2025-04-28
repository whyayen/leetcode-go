package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var previousHead *ListNode
	var currentHead *ListNode

	for list1 != nil || list2 != nil {
		previousHead = currentHead
		if (list1 != nil && list2 != nil && (list1.Val <= list2.Val)) || list2 == nil {
			currentHead = list1
			list1 = list1.Next
		} else {
			currentHead = list2
			list2 = list2.Next
		}

		if previousHead != nil {
			previousHead.Next = currentHead
		}

		if head == nil {
			head = currentHead
		}
	}

	return head
}

func main() {
	var firstHead *ListNode
	var secondHead *ListNode
	var previousFirstHead *ListNode
	var previousSecondHead *ListNode

	for i := 1; i <= 10; i++ {
		currentFirstHead := ListNode{
			Val: i,
		}
		if previousFirstHead != nil {
			previousFirstHead.Next = &currentFirstHead
		}
		currentSecondHead := ListNode{
			Val: i * 2,
		}
		if previousSecondHead != nil {
			previousSecondHead.Next = &currentSecondHead
		}

		if i == 1 {
			firstHead = &currentFirstHead
			secondHead = &currentSecondHead
		}

		previousFirstHead = &currentFirstHead
		previousSecondHead = &currentSecondHead
	}

	result := mergeTwoLists(secondHead, firstHead)

	for result != nil {
		fmt.Printf("%+v\n", result.Val)
		result = result.Next
	}
}

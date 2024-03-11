package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depthCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := depthCheck(root.Left)
	r := depthCheck(root.Right)

	fmt.Printf("%+v / %+v / %+v\n", root.Val, l, r)

	if l == -1 || r == -1 || math.Abs(float64(l-r)) > 1 {
		return -1
	}

	return max(l, r) + 1
}

func isBalanced(root *TreeNode) bool {
	h := depthCheck(root)
	return !(h == -1)
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	l1 := &TreeNode{
		Val: 2,
	}
	r1 := &TreeNode{
		Val: 5,
	}
	root.Left = l1
	root.Right = r1

	l2l := &TreeNode{
		Val: 3,
	}
	l1.Left = l2l
	l3l := &TreeNode{
		Val: 4,
	}
	l2l.Left = l3l

	r2r := &TreeNode{
		Val: 6,
	}
	r1.Right = r2r
	r3r := &TreeNode{
		Val: 7,
	}
	r2r.Right = r3r
	isBalanced(root)
}

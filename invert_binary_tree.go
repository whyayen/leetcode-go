package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func swap(current *TreeNode) {
	l, r := current.Left, current.Right
	current.Right = l
	current.Left = r
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	nodesShouldSwap := []*TreeNode{root}

	for len(nodesShouldSwap) != 0 {
		node := nodesShouldSwap[0]
		swap(node)

		if node.Left != nil {
			nodesShouldSwap = append(nodesShouldSwap, node.Left)
		}

		if node.Right != nil {
			nodesShouldSwap = append(nodesShouldSwap, node.Right)
		}

		nodesShouldSwap = nodesShouldSwap[1:]
	}

	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
	}
	l1 := &TreeNode{
		Val: 2,
	}
	root.Left = l1
	r1 := &TreeNode{
		Val: 7,
	}
	root.Right = r1

	l1l := &TreeNode{
		Val: 1,
	}
	l1r := &TreeNode{
		Val: 3,
	}
	l1.Left = l1l
	l1.Right = l1r

	r1l := &TreeNode{
		Val: 6,
	}
	r1r := &TreeNode{
		Val: 9,
	}
	r1.Left = r1l
	r1.Right = r1r

	result := invertTree(root)

	nodesShouldSwap := []*TreeNode{result}
	for len(nodesShouldSwap) != 0 {
		node := nodesShouldSwap[0]
		fmt.Printf("%+v \n", node.Val)

		nodesShouldSwap = nodesShouldSwap[1:]

		if node.Left != nil && node.Right != nil {
			nodesShouldSwap = append(nodesShouldSwap, node.Left, node.Right)
		}
	}
}

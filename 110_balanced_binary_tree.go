package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	var depth func(*TreeNode) int
	isBalanced := true

	depth = func(node *TreeNode) int {
		d := 0

		if node != nil {
			l := depth(node.Left)
			r := depth(node.Right)

			if int(math.Abs(float64(l)-float64(r))) > 1 {
				isBalanced = false
			}

			d = max(l, r) + 1
		}

		return d
	}
	depth(root)

	return isBalanced
}

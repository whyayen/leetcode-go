package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	depth, longest := depthOfNode(root, 0)
	return max(longest, depth-1)
}

func depthOfNode(root *TreeNode, longest int) (int, int) {
	depth, newLongest := 0, 0

	if root != nil {
		l, ll := depthOfNode(root.Left, longest)
		r, rl := depthOfNode(root.Right, longest)

		depth, newLongest = max(l, r)+1, max(longest, l+r, ll, rl)
	}

	return depth, newLongest
}

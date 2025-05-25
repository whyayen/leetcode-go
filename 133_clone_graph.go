package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if cloned, ok := visited[node.Val]; ok {
			return cloned
		}

		clone := &Node{
			Val: node.Val,
		}
		visited[node.Val] = clone

		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}

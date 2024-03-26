package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	var firstNode *Node
	queue := []*Node{node}
	newNodes := map[int]*Node{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		newCurr, currIsExist := newNodes[curr.Val]

		if !currIsExist {
			newCurr = &Node{
				Val: curr.Val,
			}
			newNodes[curr.Val] = newCurr
		}

		if firstNode == nil {
			firstNode = newCurr
		}

		for _, neighbor := range curr.Neighbors {
			newNeighbor, neighborIsExist := newNodes[neighbor.Val]

			if !neighborIsExist {
				newNeighbor := &Node{
					Val: neighbor.Val,
				}
				newNodes[neighbor.Val] = newNeighbor
				newCurr.Neighbors = append(newCurr.Neighbors, newNeighbor)
				queue = append(queue, neighbor)
			} else {
				newCurr.Neighbors = append(newCurr.Neighbors, newNeighbor)
			}
		}
	}

	return firstNode
}

func main() {
	first := &Node{
		Val: 1,
	}
	second := &Node{
		Val: 2,
	}
	third := &Node{
		Val: 3,
	}
	fourth := &Node{
		Val: 4,
	}
	first.Neighbors = []*Node{second, third}
	second.Neighbors = []*Node{first, fourth}
	third.Neighbors = []*Node{first, fourth}
	fourth.Neighbors = []*Node{second, third}

	result := cloneGraph(first)
	fmt.Printf("%+v / %+v, %+v, %+v\n", first, result, result.Neighbors[0], result.Neighbors[1])
}

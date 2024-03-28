package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, pre := range prerequisites {
		x := pre[0]
		y := pre[1]
		graph[x] = append(graph[x], y)
	}

	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, visited, i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited []int, node int) bool {
	if visited[node] == 1 {
		return true // Cycle detected
	}

	if visited[node] == -1 {
		return false // visited & no cycle detected
	}

	visited[node] = 1

	for _, neighbor := range graph[node] {
		if hasCycle(graph, visited, neighbor) {
			return true // cycle detected
		}
	}

	visited[node] = -1
	return false
}

func main() {
	result := canFinish(4, [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	})
	fmt.Printf("%+v\n", result)
}

package main

import (
	"fmt"
)

func updateMatrix(mat [][]int) [][]int {
	var queue [][]int
	m, n := len(mat), len(mat[0])
	maxDistance := m * n

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = maxDistance
			}
		}

	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			r, c := cell[0]+dir[0], cell[1]+dir[1]
			if r >= 0 && c >= 0 && r < m && c < n && mat[r][c] > mat[cell[0]][cell[1]]+1 {
				queue = append(queue, []int{r, c})
				mat[r][c] = mat[cell[0]][cell[1]] + 1
			}
		}
	}

	return mat
}

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	result := updateMatrix(mat)
	fmt.Printf("%+v\n", result)
}

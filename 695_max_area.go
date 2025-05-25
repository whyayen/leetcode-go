package main

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// visit 函數用來走訪整個 graph
	// 如果 x, y 超過 graph 大小或者遇到 0 就直接回傳 0
	// 否則就是把 grid[x][y] 設為 0 表示走過了，且 res 設為 1 （表示有 1 個 island）並往四個方向走，把回傳值加總
	// 遍歷每一格比對當前 maxArea 取較大的值
	var visit func(int, int) int
	visit = func(x int, y int) int {
		if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 0 {
			return 0
		}

		grid[x][y] = 0
		res := 1

		res += visit(x-1, y)
		res += visit(x, y-1)
		res += visit(x+1, y)
		res += visit(x, y+1)

		return res
	}

	startX, startY := 0, 0
	maxArea := 0
	for startX < rows && startY < cols {
		maxArea = max(maxArea, visit(startX, startY))

		if startY+1 < cols {
			startY++
		} else {
			startY = 0
			startX++
		}
	}

	return maxArea
}
